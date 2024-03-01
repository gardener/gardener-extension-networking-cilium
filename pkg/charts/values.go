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
func RenderCiliumChart(renderer chartrenderer.Interface, config *ciliumv1alpha1.NetworkConfig, network *extensionsv1alpha1.Network, cluster *extensionscontroller.Cluster, ipamMode, configMapHash string) ([]byte, error) {
	var release *chartrenderer.RenderedChart

	values, err := ComputeCiliumChartValues(config, network, cluster, ipamMode, configMapHash)
	if err != nil {
		return nil, err
	}

	release, err = renderer.RenderEmbeddedFS(charts.InternalChart, cilium.CiliumChartPath, cilium.ReleaseName, metav1.NamespaceSystem, values)
	if err != nil {
		return nil, err
	}

	newConfigMapHash, err := getConfigMapHash(release)
	if err != nil {
		return nil, err
	}

	if newConfigMapHash != configMapHash {
		// Render the charts with the new configMap hash.
		values, err := ComputeCiliumChartValues(config, network, cluster, ipamMode, newConfigMapHash)
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

func getConfigMapHash(release *chartrenderer.RenderedChart) (string, error) {
	configMapPath := "cilium/charts/config/templates/configmap.yaml"
	configMapData, ok := release.Files()[configMapPath]
	if !ok {
		return "", fmt.Errorf("configmap not found in the given path: %s", configMapPath)
	}

	return utils.ComputeConfigMapChecksum(configMapData), nil
}
