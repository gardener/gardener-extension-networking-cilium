// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"context"
	"fmt"
	"slices"

	extensionsconfig "github.com/gardener/gardener/extensions/pkg/apis/config/v1alpha1"
	extensionscontroller "github.com/gardener/gardener/extensions/pkg/controller"
	"github.com/gardener/gardener/extensions/pkg/util"
	"github.com/gardener/gardener/extensions/pkg/webhook"
	extensionshootwebhook "github.com/gardener/gardener/extensions/pkg/webhook/shoot"
	gardenv1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	v1beta1helper "github.com/gardener/gardener/pkg/apis/core/v1beta1/helper"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	gardenerkubernetes "github.com/gardener/gardener/pkg/client/kubernetes"
	"github.com/gardener/gardener/pkg/utils"
	"github.com/gardener/gardener/pkg/utils/chart"
	"github.com/gardener/gardener/pkg/utils/managedresources"
	"github.com/go-logr/logr"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	monitoringv1alpha1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/utils/ptr"
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
			{
				Type: &corev1.ConfigMap{},
				Name: "cilium-dashboards",
			},
			{
				Type: &monitoringv1alpha1.ScrapeConfig{},
				Name: "shoot-cilium-agent",
			},
			{
				Type: &monitoringv1alpha1.ScrapeConfig{},
				Name: "shoot-cilium-hubble",
			},
			{
				Type: &monitoringv1alpha1.ScrapeConfig{},
				Name: "shoot-cilium-operator",
			},
			{
				Type: &monitoringv1.PrometheusRule{},
				Name: "shoot-cilium-agent",
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
		networkConfig = &ciliumv1alpha1.NetworkConfig{}
		err           error
	)

	if network.Spec.ProviderConfig != nil {
		networkConfig, err = CiliumNetworkConfigFromNetworkResource(network)
		if err != nil {
			return err
		}
	}

	ipFamilies := slices.Clone(network.Spec.IPFamilies)

	condition := v1beta1helper.GetCondition(cluster.Shoot.Status.Constraints, gardenv1beta1.ShootDualStackNodesMigrationReady)

	if condition != nil && condition.Status != gardenv1beta1.ConditionTrue {
		if len(ipFamilies) > 1 {
			ipFamilies = ipFamilies[:1]
		}
	}

	if networkConfig.Overlay != nil && networkConfig.Overlay.Enabled {
		if networkConfig.TunnelMode == nil || networkConfig.TunnelMode != nil && *networkConfig.TunnelMode == ciliumv1alpha1.Disabled {
			// use vxlan as default overlay network
			networkConfig.TunnelMode = ptr.To(ciliumv1alpha1.VXLan)
		}
		if slices.Contains(ipFamilies, extensionsv1alpha1.IPFamilyIPv4) {
			networkConfig.IPv4NativeRoutingCIDREnabled = ptr.To(false)
		}
		if slices.Contains(ipFamilies, extensionsv1alpha1.IPFamilyIPv6) {
			networkConfig.IPv6NativeRoutingCIDREnabled = ptr.To(false)
		}
	}
	if networkConfig.Overlay != nil && !networkConfig.Overlay.Enabled {
		networkConfig.TunnelMode = ptr.To(ciliumv1alpha1.Disabled)
		if slices.Contains(ipFamilies, extensionsv1alpha1.IPFamilyIPv4) {
			networkConfig.IPv4NativeRoutingCIDREnabled = ptr.To(true)
		}
		if slices.Contains(ipFamilies, extensionsv1alpha1.IPFamilyIPv6) {
			networkConfig.IPv6NativeRoutingCIDREnabled = ptr.To(true)
		}
	}

	networkConfig.IPv4 = &ciliumv1alpha1.IPv4{
		Enabled: slices.Contains(ipFamilies, extensionsv1alpha1.IPFamilyIPv4),
	}
	networkConfig.IPv6 = &ciliumv1alpha1.IPv6{
		Enabled: slices.Contains(ipFamilies, extensionsv1alpha1.IPFamilyIPv6),
	}

	if cluster.Shoot.Spec.Kubernetes.KubeProxy != nil &&
		ptr.Deref(cluster.Shoot.Spec.Kubernetes.KubeProxy.Enabled, false) &&
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
			ShootWebhooksResourceName,
			*webhookConfig,
			cluster,
			true,
		); err != nil {
			return fmt.Errorf("could not reconcile shoot webhooks: %w", err)
		}
	}

	// Create shoot chart renderer
	chartRenderer, err := a.chartRendererFactory.NewChartRendererForShoot(cluster.Shoot.Spec.Kubernetes.Version)
	if err != nil {
		return fmt.Errorf("could not create chart renderer for shoot '%s': %w", network.Namespace, err)
	}

	configMap, err := getConfigMap(ctx, a.client, cluster, "cilium-config")
	if err != nil {
		return fmt.Errorf("error getting cilium configMap: %w", err)
	}

	configMapLabelPrefix, err := getConfigMap(ctx, a.client, cluster, "label-prefix-conf")
	if err != nil {
		return fmt.Errorf("error getting cilium configMap: %w", err)
	}

	ciliumChart, err := chartspkg.RenderCiliumChart(chartRenderer, networkConfig, network, cluster, getIPAMMode(configMap), getConfigMapHash(configMap), getConfigMapHash(configMapLabelPrefix))
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

func getConfigMap(ctx context.Context, cl client.Client, cluster *extensionscontroller.Cluster, name string) (*corev1.ConfigMap, error) {
	// Cannot retrieve config map of hibernated clusters => use empty config map instead
	if extensionscontroller.IsHibernated(cluster) {
		return &corev1.ConfigMap{}, nil
	}
	_, shootClient, err := util.NewClientForShoot(ctx, cl, cluster.ObjectMeta.Name, client.Options{}, extensionsconfig.RESTOptions{})
	if err != nil {
		// No need to report the error as this is anyway only best effort. Some scenarios, e.g. autonomous shoot clusters,
		// might not have the gardener secret and hence cannot construct the shoot client here.
		// RenderCiliumChart(...) has the real fallback logic in case the config map was not found/changed.
		return &corev1.ConfigMap{}, nil
	}
	configmap := &corev1.ConfigMap{}
	_ = shootClient.Get(ctx, client.ObjectKey{Namespace: "kube-system", Name: name}, configmap)
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
