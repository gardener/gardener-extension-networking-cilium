// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"context"
	"fmt"

	extensionsconfig "github.com/gardener/gardener/extensions/pkg/apis/config"
	extensionscontroller "github.com/gardener/gardener/extensions/pkg/controller"
	"github.com/gardener/gardener/extensions/pkg/util"
	"github.com/gardener/gardener/extensions/pkg/webhook"
	extensionshootwebhook "github.com/gardener/gardener/extensions/pkg/webhook/shoot"
	v1beta1helper "github.com/gardener/gardener/pkg/apis/core/v1beta1/helper"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	gardenerkubernetes "github.com/gardener/gardener/pkg/client/kubernetes"
	"github.com/gardener/gardener/pkg/utils"
	"github.com/gardener/gardener/pkg/utils/chart"
	"github.com/gardener/gardener/pkg/utils/managedresources"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/gardener/gardener-extension-networking-cilium/charts"
	ciliumv1alpha1 "github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium/v1alpha1"
	chartspkg "github.com/gardener/gardener-extension-networking-cilium/pkg/charts"
	"github.com/gardener/gardener-extension-networking-cilium/pkg/cilium"
)

const (
	// CiliumConfigManagedResourceName is the name of the managed resource of networking cilium
	CiliumConfigManagedResourceName = "extension-networking-cilium-config"
	// ShootWebhooksResourceName is the name of the managed resource for the gardener networking extension cilium webhooks
	ShootWebhooksResourceName = "extension-cilium-shoot-webhooks"
)

func applyMonitoringConfig(ctx context.Context, seedClient client.Client, chartApplier gardenerkubernetes.ChartApplier, network *extensionsv1alpha1.Network, deleteChart bool) error {
	ciliumControlPlaneMonitoringChart := &chart.Chart{
		Name:       cilium.MonitoringName,
		EmbeddedFS: charts.InternalChart,
		Path:       cilium.CiliumMonitoringChartPath,
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
func (a *actuator) Reconcile(ctx context.Context, _ logr.Logger, network *extensionsv1alpha1.Network, cluster *extensionscontroller.Cluster) error {
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

	if networkConfig != nil {
		if networkConfig.Overlay != nil && networkConfig.Overlay.Enabled {
			if networkConfig.TunnelMode == nil || networkConfig.TunnelMode != nil && *networkConfig.TunnelMode == ciliumv1alpha1.Disabled {
				// use vxlan as default overlay network
				networkConfig.TunnelMode = (*ciliumv1alpha1.TunnelMode)(pointer.String(string(ciliumv1alpha1.VXLan)))
			}
			networkConfig.IPv4NativeRoutingCIDREnabled = pointer.Bool(false)
		}
		if networkConfig.Overlay != nil && !networkConfig.Overlay.Enabled {
			networkConfig.TunnelMode = (*ciliumv1alpha1.TunnelMode)(pointer.String(string(ciliumv1alpha1.Disabled)))
			networkConfig.IPv4NativeRoutingCIDREnabled = pointer.Bool(true)
		}
	}

	if cluster.Shoot.Spec.Kubernetes.KubeProxy != nil &&
		pointer.BoolDeref(cluster.Shoot.Spec.Kubernetes.KubeProxy.Enabled, false) &&
		cluster.Shoot.Spec.Kubernetes.KubeProxy.Mode != nil &&
		*cluster.Shoot.Spec.Kubernetes.KubeProxy.Mode == "IPVS" &&
		v1beta1helper.IsNodeLocalDNSEnabled(cluster.Shoot.Spec.SystemComponents) {
		return field.Forbidden(field.NewPath("spec", "kubernetes", "kubeProxy", "mode"), "Running kube-proxy with IPVS mode is forbidden in conjunction with node local dns enabled")
	}

	if a.atomicShootWebhookConfig != nil {
		value := a.atomicShootWebhookConfig.Load()
		webhookConfig, ok := value.(*webhook.Configs)
		if !ok {
			return fmt.Errorf("expected *admissionregistrationv1.MutatingWebhookConfiguration, got %T", value)
		}

		if err := extensionshootwebhook.ReconcileWebhookConfig(
			ctx,
			a.client,
			network.Namespace,
			DefaultAddOptions.WebhookServerNamespace,
			cilium.Name,
			ShootWebhooksResourceName,
			*webhookConfig,
			cluster,
		); err != nil {
			return fmt.Errorf("could not reconcile shoot webhooks: %w", err)
		}
	}

	// Create shoot chart renderer
	chartRenderer, err := a.chartRendererFactory.NewChartRendererForShoot(cluster.Shoot.Spec.Kubernetes.Version)
	if err != nil {
		return fmt.Errorf("could not create chart renderer for shoot '%s': %w", network.Namespace, err)
	}

	configMap, err := getCiliumConfigMap(ctx, a.client, cluster)
	if err != nil {
		return fmt.Errorf("error getting cilium configMap: %w", err)
	}

	ciliumChart, err := chartspkg.RenderCiliumChart(chartRenderer, networkConfig, network, cluster, getIPAMMode(configMap), getConfigMapHash(configMap))
	if err != nil {
		return err
	}

	data := map[string][]byte{chartspkg.CiliumConfigKey: ciliumChart}
	if err := managedresources.CreateForShoot(ctx, a.client, network.Namespace, CiliumConfigManagedResourceName, "extension-networking-cilium", false, data); err != nil {
		return err
	}

	if err := applyMonitoringConfig(ctx, a.client, a.chartApplier, network, false); err != nil {
		return err
	}

	return a.updateProviderStatus(ctx, network, networkConfig)
}

func getCiliumConfigMap(ctx context.Context, cl client.Client, cluster *extensionscontroller.Cluster) (*corev1.ConfigMap, error) {
	// Cannot retrieve config map of hibernated clusters => use empty config map instead
	if extensionscontroller.IsHibernated(cluster) {
		return &corev1.ConfigMap{}, nil
	}
	_, shootClient, err := util.NewClientForShoot(ctx, cl, cluster.ObjectMeta.Name, client.Options{}, extensionsconfig.RESTOptions{})
	if err != nil {
		return nil, fmt.Errorf("could not create shoot client: %w", err)
	}
	configmap := &corev1.ConfigMap{}
	_ = shootClient.Get(ctx, client.ObjectKey{Namespace: "kube-system", Name: "cilium-config"}, configmap)
	return configmap, nil
}

func getIPAMMode(configMap *corev1.ConfigMap) string {
	if configMap != nil {
		if ipamMode, ok := configMap.Data["ipam"]; ok {
			return ipamMode
		}
	}
	return "kubernetes"
}

func getConfigMapHash(configMap *corev1.ConfigMap) string {
	if configMap != nil {
		return utils.ComputeConfigMapChecksum(configMap.Data)
	}
	return ""
}
