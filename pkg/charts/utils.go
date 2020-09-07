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
	"fmt"

	ciliumv1alpha1 "github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium/v1alpha1"
	"github.com/gardener/gardener-extension-networking-cilium/pkg/cilium"
	"github.com/gardener/gardener-extension-networking-cilium/pkg/imagevector"
)

var defaultCiliumConfig = requirementsConfig{
	Agent: agent{
		Enabled:        true,
		SleepAfterInit: false,
	},
	Config: config{
		Enabled: true,
	},
	Operator: operator{
		Enabled: true,
	},
	Preflight: preflight{
		Enabled:         false,
		ToFQDNSPreCache: "",
	},
}

var defaultGlobalConfig = globalConfig{
	IdentityAllocationMode: ciliumv1alpha1.CRD,
	Tunnel:                 ciliumv1alpha1.VXLan,
	KubeProxyReplacement:   ciliumv1alpha1.Probe,
	Etcd: etcd{
		Enabled: false,
		Managed: false,
	},
	Ipv4: ipv4{
		Enabled: true,
	},
	Ipv6: ipv6{
		Enabled: false,
	},
	Debug: debug{
		Enabled: false,
	},
	Prometheus: prometheus{
		Enabled: true,
		Port:    9090,
		ServiceMonitor: serviceMonitor{
			Enabled: false,
		},
	},
	OperatorPrometheus: operatorPrometheus{
		Enabled: true,
		Port:    6942,
	},
	Psp: psp{
		Enabled: true,
	},
	Images: map[string]string{
		cilium.CiliumAgentImageName:        imagevector.CiliumAgentImage(),
		cilium.CiliumOperatorImageName:     imagevector.CiliumOperatorImage(),
		cilium.CiliumETCDOperatorImageName: imagevector.CiliumEtcdOperatorImage(),
		cilium.CiliumNodeInitImageName:     imagevector.CiliumNodeInitImage(),
		cilium.CiliumPreflightImageName:    imagevector.CiliumPreflightImage(),

		cilium.HubbleRelayImageName: imagevector.CiliumHubbleRelayImage(),
		cilium.HubbleUIImageName:    imagevector.CiliumHubbleUIImage(),
	},
}

func newGlobalConfig() globalConfig {
	return defaultGlobalConfig
}

func newRequirementsConfig() requirementsConfig {
	return defaultCiliumConfig
}

// ComputeCiliumChartValues computes the values for the cilium chart.
func ComputeCiliumChartValues(config *ciliumv1alpha1.NetworkConfig) (*ciliumConfig, error) {
	requirementsConfig, globalConfig, err := generateChartValues(config)
	if err != nil {
		return nil, fmt.Errorf("error when generating config values %v", err)
	}
	return &ciliumConfig{
		Requirements: requirementsConfig,
		Global:       globalConfig,
	}, nil
}

func generateChartValues(config *ciliumv1alpha1.NetworkConfig) (requirementsConfig, globalConfig, error) {
	var (
		requirementsConfig = newRequirementsConfig()
		globalConfig       = newGlobalConfig()
	)

	if config == nil {
		return requirementsConfig, globalConfig, nil
	}

	// Settings for Kube-Proxy disabled and using the HostService option
	// Also need to configure KubeProxy
	// TODO: enabling this feature assumes that Gardener won't deploy kube-proxy in the control-plane.
	if config.KubeProxy != nil && config.KubeProxy.Enabled != nil && *config.KubeProxy.Enabled {
		globalConfig.KubeProxyReplacement = ciliumv1alpha1.Strict

		if config.KubeProxy.ServiceHost != nil && config.KubeProxy.ServicePort != nil {
			globalConfig.K8sServiceHost = *config.KubeProxy.ServiceHost
			globalConfig.K8sServicePort = *config.KubeProxy.ServicePort
		}

		globalConfig.NodePort.Enabled = true
	}

	// check if PSPs are enabled
	if config.PSPEnabled != nil {
		globalConfig.Psp.Enabled = *config.PSPEnabled
	}

	// If Hubble enabled
	if config.Hubble != nil && config.Hubble.Enabled {
		requirementsConfig.Hubble.Enabled = config.Hubble.Enabled
	}

	// If ETCD enabled
	if config.Store != nil {
		if *config.Store == ciliumv1alpha1.ETCD {
			globalConfig.Etcd.Enabled = true
			globalConfig.Etcd.Managed = true
		}
	}

	return requirementsConfig, globalConfig, nil
}
