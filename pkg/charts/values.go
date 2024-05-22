// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package charts

import (
	"fmt"

	extensionscontroller "github.com/gardener/gardener/extensions/pkg/controller"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"github.com/gardener/gardener/pkg/chartrenderer"
	"github.com/gardener/gardener/pkg/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/gardener/gardener-extension-networking-cilium/charts"
	ciliumv1alpha1 "github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium/v1alpha1"
	"github.com/gardener/gardener-extension-networking-cilium/pkg/cilium"
)

// CiliumConfigKey defines the cilium configmap key.
const CiliumConfigKey = "config.yaml"

// RenderCiliumChart renders the cilium chart with the given values.
func RenderCiliumChart(renderer chartrenderer.Interface, config *ciliumv1alpha1.NetworkConfig, network *extensionsv1alpha1.Network, cluster *extensionscontroller.Cluster, ipamMode, configMapHash, configMapLabelPrefixHash string) ([]byte, error) {
	var release *chartrenderer.RenderedChart

	values, err := ComputeCiliumChartValues(config, network, cluster, ipamMode, configMapHash, configMapLabelPrefixHash)
	if err != nil {
		return nil, err
	}

	release, err = renderer.RenderEmbeddedFS(charts.InternalChart, cilium.CiliumChartPath, cilium.ReleaseName, metav1.NamespaceSystem, values)
	if err != nil {
		return nil, err
	}

	configMapPath := "cilium/charts/config/templates/configmap.yaml"
	newConfigMapHash, err := getConfigMapHash(release, configMapPath)
	if err != nil {
		return nil, err
	}

	configMapPath = "cilium/charts/config/templates/label-prefix-configmap.yaml"
	newConfigMapLabelPrefixHash, err := getConfigMapHash(release, configMapPath)
	if err != nil {
		return nil, err
	}

	if newConfigMapHash != configMapHash || newConfigMapLabelPrefixHash != configMapLabelPrefixHash {
		// Render the charts with the new configMap hash.
		values, err := ComputeCiliumChartValues(config, network, cluster, ipamMode, newConfigMapHash, configMapLabelPrefixHash)
		if err != nil {
			return nil, err
		}

		release, err = renderer.RenderEmbeddedFS(charts.InternalChart, cilium.CiliumChartPath, cilium.ReleaseName, metav1.NamespaceSystem, values)
		if err != nil {
			return nil, err
		}
	}

	return release.Manifest(), nil
}

func getConfigMapHash(release *chartrenderer.RenderedChart, configMapPath string) (string, error) {
	configMapData, ok := release.Files()[configMapPath]
	if !ok {
		return "", fmt.Errorf("configmap not found in the given path: %s", configMapPath)
	}

	return utils.ComputeConfigMapChecksum(configMapData), nil
}
