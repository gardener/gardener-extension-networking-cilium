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

package charts

import (
	extensionscontroller "github.com/gardener/gardener/extensions/pkg/controller"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"github.com/gardener/gardener/pkg/chartrenderer"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/gardener/gardener-extension-networking-cilium/charts"
	ciliumv1alpha1 "github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium/v1alpha1"
	"github.com/gardener/gardener-extension-networking-cilium/pkg/cilium"
)

// CiliumConfigKey defines the cilium configmap key.
const CiliumConfigKey = "config.yaml"

// RenderCiliumChart renders the cilium chart with the given values.
func RenderCiliumChart(renderer chartrenderer.Interface, config *ciliumv1alpha1.NetworkConfig, network *extensionsv1alpha1.Network, cluster *extensionscontroller.Cluster, ipamMode string) ([]byte, error) {
	values, err := ComputeCiliumChartValues(config, network, cluster, ipamMode)
	if err != nil {
		return nil, err
	}

	release, err := renderer.RenderEmbeddedFS(charts.InternalChart, cilium.CiliumChartPath, cilium.ReleaseName, metav1.NamespaceSystem, values)
	if err != nil {
		return nil, err
	}

	return release.Manifest(), nil
}
