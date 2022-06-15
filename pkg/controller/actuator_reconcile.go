// Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controller

import (
	"bytes"
	"context"
	"fmt"

	ciliumv1alpha1 "github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium/v1alpha1"
	"github.com/gardener/gardener-extension-networking-cilium/pkg/charts"
	"github.com/gardener/gardener-extension-networking-cilium/pkg/cilium"

	extensionscontroller "github.com/gardener/gardener/extensions/pkg/controller"
	"github.com/gardener/gardener/extensions/pkg/webhook"
	extensionswebhookshoot "github.com/gardener/gardener/extensions/pkg/webhook/shoot"
	"github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	gardenerkubernetes "github.com/gardener/gardener/pkg/client/kubernetes"
	"github.com/gardener/gardener/pkg/extensions"
	"github.com/gardener/gardener/pkg/utils/chart"
	"github.com/gardener/gardener/pkg/utils/flow"
	kutil "github.com/gardener/gardener/pkg/utils/kubernetes"
	"github.com/gardener/gardener/pkg/utils/managedresources"
	"github.com/gardener/gardener/pkg/utils/managedresources/builder"
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// CiliumConfigSecretName is the name of the secret used for the managed resource of networking cilium
	CiliumConfigSecretName = "extension-networking-cilium-config"
	// ShootWebhooksResourceName is the name of the managed resource for the gardener networking extension cilium webhooks
	ShootWebhooksResourceName = "extension-cilium-shoot-webhooks"
)

func withLocalObjectRefs(refs ...string) []corev1.LocalObjectReference {
	var localObjectRefs []corev1.LocalObjectReference
	for _, ref := range refs {
		localObjectRefs = append(localObjectRefs, corev1.LocalObjectReference{Name: ref})
	}
	return localObjectRefs
}

func ciliumSecret(cl client.Client, ciliumConfig []byte, namespace string) (*builder.Secret, []corev1.LocalObjectReference) {
	return builder.NewSecret(cl).
		WithKeyValues(map[string][]byte{charts.CiliumConfigKey: ciliumConfig}).
		WithNamespacedName(namespace, CiliumConfigSecretName), withLocalObjectRefs(CiliumConfigSecretName)
}

func applyMonitoringConfig(ctx context.Context, seedClient client.Client, chartApplier gardenerkubernetes.ChartApplier, network *extensionsv1alpha1.Network, deleteChart bool) error {
	ciliumControlPlaneMonitoringChart := &chart.Chart{
		Name: cilium.MonitoringName,
		Path: cilium.CiliumMonitoringChartPath,
		Objects: []*chart.Object{
			{
				Type: &corev1.ConfigMap{},
				Name: cilium.MonitoringName,
			},
		},
	}

	if deleteChart {
		return client.IgnoreNotFound(ciliumControlPlaneMonitoringChart.Delete(ctx, seedClient, network.Namespace))
	}

	return ciliumControlPlaneMonitoringChart.Apply(ctx, chartApplier, network.Namespace, nil, "", "", nil)
}

// Reconcile implements Network.Actuator.
func (a *actuator) Reconcile(ctx context.Context, network *extensionsv1alpha1.Network, cluster *extensionscontroller.Cluster) error {
	var (
		networkConfig *ciliumv1alpha1.NetworkConfig
		err           error
	)

	if network.Spec.ProviderConfig != nil {
		networkConfig, err = CiliumNetworkConfigFromNetworkResource(network)
		if err != nil {
			return err
		}
	}

	if cluster.Shoot.Spec.Kubernetes.KubeProxy != nil && cluster.Shoot.Spec.Kubernetes.KubeProxy.Enabled != nil && *cluster.Shoot.Spec.Kubernetes.KubeProxy.Enabled && cluster.Shoot.Spec.Kubernetes.KubeProxy.Mode != nil && *cluster.Shoot.Spec.Kubernetes.KubeProxy.Mode == "IPVS" {
		if cluster.Shoot.Annotations[v1beta1constants.AnnotationNodeLocalDNS] == "true" {
			return field.Forbidden(field.NewPath("spec", "kubernetes", "kubeProxy", "mode"), "Running kube-proxy with IPVS mode is forbidden in conjunction with node local dns enabled")
		}
	}

	if len(a.shootWebhooks) > 0 {
		if err := ReconcileShootWebhooks(ctx, a.client, network.Namespace, cilium.Name, a.webhookServerPort, a.shootWebhooks, cluster); err != nil {
			return fmt.Errorf("could not reconcile shoot webhooks: %w", err)
		}
	}

	// Create shoot chart renderer
	chartRenderer, err := a.chartRendererFactory.NewChartRendererForShoot(cluster.Shoot.Spec.Kubernetes.Version)
	if err != nil {
		return fmt.Errorf("could not create chart renderer for shoot '%s': %w", network.Namespace, err)
	}

	ciliumChart, err := charts.RenderCiliumChart(chartRenderer, networkConfig, network, cluster)
	if err != nil {
		return err
	}

	secret, secretRefs := ciliumSecret(a.client, ciliumChart, network.Namespace)
	if err := secret.Reconcile(ctx); err != nil {
		return err
	}

	if err := builder.
		NewManagedResource(a.client).
		WithNamespacedName(network.Namespace, CiliumConfigSecretName).
		WithSecretRefs(secretRefs).
		WithInjectedLabels(map[string]string{constants.ShootNoCleanup: "true"}).
		Reconcile(ctx); err != nil {
		return err
	}

	if err := applyMonitoringConfig(ctx, a.client, a.chartApplier, network, false); err != nil {
		return err
	}

	return a.updateProviderStatus(ctx, network, networkConfig)
}

func marshalWebhooks(webhooks []admissionregistrationv1.MutatingWebhook, name string) ([]byte, error) {
	var (
		buf     = new(bytes.Buffer)
		encoder = json.NewYAMLSerializer(json.DefaultMetaFactory, nil, nil)

		apiVersion, kind                            = admissionregistrationv1.SchemeGroupVersion.WithKind("MutatingWebhookConfiguration").ToAPIVersionAndKind()
		mutatingWebhookConfiguration runtime.Object = &admissionregistrationv1.MutatingWebhookConfiguration{
			TypeMeta: metav1.TypeMeta{
				APIVersion: apiVersion,
				Kind:       kind,
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: webhook.NamePrefix + name + webhook.NameSuffixShoot,
			},
			Webhooks: webhooks,
		}
	)

	if err := encoder.Encode(mutatingWebhookConfiguration, buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// ReconcileShootWebhooks deploys the shoot webhook configuration and a managed resource that contains the MutatingWebhookConfiguration.
func ReconcileShootWebhooks(ctx context.Context, c client.Client, namespace, providerName string, serverPort int, shootWebhooks []admissionregistrationv1.MutatingWebhook, cluster *extensionscontroller.Cluster) error {
	if err := extensionswebhookshoot.EnsureNetworkPolicy(ctx, c, namespace, providerName, serverPort); err != nil {
		return fmt.Errorf("could not create or update network policy for shoot webhooks in namespace '%s': %w", namespace, err)
	}

	if cluster.Shoot == nil {
		return fmt.Errorf("no shoot found in cluster resource")
	}

	webhookConfiguration, err := marshalWebhooks(shootWebhooks, providerName)
	if err != nil {
		return err
	}
	data := map[string][]byte{"mutatingwebhookconfiguration.yaml": webhookConfiguration}

	if err := managedresources.Create(ctx, c, namespace, ShootWebhooksResourceName, false, "", data, nil, nil, nil); err != nil {
		return fmt.Errorf("could not create or update managed resource '%s/%s' containing shoot webhooks: %w", namespace, ShootWebhooksResourceName, err)
	}

	return nil
}

// ReconcileShootWebhooksForAllNamespaces reconciles the shoot webhooks in all shoot namespaces of the given
// provider type. This is necessary in case the webhook port is changed (otherwise, the network policy would only be
// updated again as part of the ControlPlane reconciliation which might only happen in the next 24h).
func ReconcileShootWebhooksForAllNamespaces(ctx context.Context, c client.Client, providerName, providerType string, port int, shootWebhooks []admissionregistrationv1.MutatingWebhook) error {
	namespaceList := &corev1.NamespaceList{}
	if err := c.List(ctx, namespaceList, client.MatchingLabels{
		v1beta1constants.GardenRole:         v1beta1constants.GardenRoleShoot,
		v1beta1constants.LabelShootProvider: providerType,
	}); err != nil {
		return err
	}

	fns := make([]flow.TaskFn, 0, len(namespaceList.Items))

	for _, namespace := range namespaceList.Items {
		var (
			networkPolicy     = extensionswebhookshoot.GetNetworkPolicyMeta(namespace.Name, providerName)
			namespaceName     = namespace.Name
			networkPolicyName = networkPolicy.Name
		)

		fns = append(fns, func(ctx context.Context) error {
			if err := c.Get(ctx, kutil.Key(namespaceName, networkPolicyName), &networkingv1.NetworkPolicy{}); err != nil {
				if !apierrors.IsNotFound(err) {
					return err
				}
				return nil
			}

			cluster, err := extensions.GetCluster(ctx, c, namespaceName)
			if err != nil {
				return err
			}

			return ReconcileShootWebhooks(ctx, c, namespaceName, providerName, port, shootWebhooks, cluster)
		})
	}

	return flow.Parallel(fns...)(ctx)
}
