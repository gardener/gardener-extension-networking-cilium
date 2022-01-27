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
	"context"
	"fmt"

	ciliumv1alpha1 "github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium/v1alpha1"
	"github.com/gardener/gardener-extension-networking-cilium/pkg/charts"
	"github.com/gardener/gardener-extension-networking-cilium/pkg/cilium"

	extensionscontroller "github.com/gardener/gardener/extensions/pkg/controller"
	"github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	gardenerkubernetes "github.com/gardener/gardener/pkg/client/kubernetes"
	"github.com/gardener/gardener/pkg/utils/chart"
	"github.com/gardener/gardener/pkg/utils/managedresources/builder"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// CiliumConfigSecretName is the name of the secret used for the managed resource of networking cilium
	CiliumConfigSecretName = "extension-networking-cilium-config"
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

func applyMonitoringConfig(ctx context.Context, seedClient client.Client, chartApplier gardenerkubernetes.ChartApplier, network *extensionsv1alpha1.Network, deleteChart bool, useProjectedTokenMount bool) error {
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

	return ciliumControlPlaneMonitoringChart.Apply(ctx, chartApplier, network.Namespace, nil, "", "", map[string]interface{}{"useProjectedTokenMount": useProjectedTokenMount})
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

	// Create shoot chart renderer
	chartRenderer, err := a.chartRendererFactory.NewChartRendererForShoot(cluster.Shoot.Spec.Kubernetes.Version)
	if err != nil {
		return fmt.Errorf("could not create chart renderer for shoot '%s': %w", network.Namespace, err)
	}

	ciliumChart, err := charts.RenderCiliumChart(chartRenderer, networkConfig, network, cluster, a.useProjectedTokenMount)
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

	if err := applyMonitoringConfig(ctx, a.client, a.chartApplier, network, false, a.useProjectedTokenMount); err != nil {
		return err
	}

	return a.updateProviderStatus(ctx, network, networkConfig)
}
