// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package charts

import (
	"fmt"
	"net/netip"
	"net/url"

	extensionscontroller "github.com/gardener/gardener/extensions/pkg/controller"
	"github.com/gardener/gardener/pkg/apis/core/v1beta1/helper"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"

	"github.com/gardener/gardener-extension-networking-cilium/imagevector"
	ciliumv1alpha1 "github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium/v1alpha1"
	"github.com/gardener/gardener-extension-networking-cilium/pkg/cilium"
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
	KubeProxyReplacement:   ciliumv1alpha1.KubeProxyReplacementFalse,
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
	EgressGateway: egressGateway{
		Enabled: false,
	},
	Prometheus: prometheus{
		Enabled: true,
		Port:    9090,
		ServiceMonitor: serviceMonitor{
			Enabled: false,
		},
	},
	OperatorHighAvailability: operatorHighAvailability{
		Enabled: true,
	},
	OperatorPrometheus: operatorPrometheus{
		Enabled: true,
		Port:    6942,
	},
	Images: map[string]string{
		cilium.CiliumAgentImageName:    imagevector.CiliumAgentImage(),
		cilium.CiliumEnvoyImageName:    imagevector.CiliumEnvoyImage(),
		cilium.CiliumOperatorImageName: imagevector.CiliumOperatorImage(),

		cilium.HubbleRelayImageName:     imagevector.CiliumHubbleRelayImage(),
		cilium.HubbleUIImageName:        imagevector.CiliumHubbleUIImage(),
		cilium.HubbleUIBackendImageName: imagevector.CiliumHubbleUIBackendImage(),
		cilium.CertGenImageName:         imagevector.CiliumCertGenImage(),
	},
	PodCIDR:  "",
	NodeCIDR: "",
	BPFSocketLBHostnsOnly: bpfSocketLBHostnsOnly{
		Enabled: false,
	},
	CNI: cni{
		Exclusive: true,
	},
	LocalRedirectPolicy: localRedirectPolicy{
		Enabled: false,
	},
	NodeLocalDNS: nodeLocalDNS{
		Enabled: false,
	},
	MTU:                   0, // --> means auto detection (default)
	Devices:               nil,
	IPv4NativeRoutingCIDR: "",
	BPF: bpf{
		LoadBalancingMode: ciliumv1alpha1.SNAT,
	},
	IPAM: ipam{
		Mode: "kubernetes",
	},
	SnatToUpstreamDNS: snatToUpstreamDNS{
		Enabled: false,
	},
	SnatOutOfCluster: snatOutOfCluster{
		Enabled: false,
	},
	EnableIPv4Masquerade: true,
	EnableIPv6Masquerade: false,
	EnableBPFMasquerade:  true,
	AutoDirectNodeRoutes: false,
	BGPControlPlane: bgpControlPlane{
		Enabled: false,
	},
	ConfigMapHash:            "",
	ConfigMapLabelPrefixHash: "",
}

func newGlobalConfig() globalConfig {
	return defaultGlobalConfig
}

func newRequirementsConfig() requirementsConfig {
	return defaultCiliumConfig
}

// ComputeCiliumChartValues computes the values for the cilium chart.
func ComputeCiliumChartValues(config *ciliumv1alpha1.NetworkConfig, network *extensionsv1alpha1.Network, cluster *extensionscontroller.Cluster, ipamMode, configMapHash, configMapLabelPrefixHash string) (*ciliumConfig, error) {
	requirementsConfig, globalConfig, err := generateChartValues(config, network, cluster, ipamMode, configMapHash, configMapLabelPrefixHash)
	if err != nil {
		return nil, fmt.Errorf("error when generating config values %w", err)
	}

	return &ciliumConfig{
		Requirements: requirementsConfig,
		Global:       globalConfig,
	}, nil
}

func generateChartValues(config *ciliumv1alpha1.NetworkConfig, network *extensionsv1alpha1.Network, cluster *extensionscontroller.Cluster, ipamMode, configMapHash, configMapLabelPrefixHash string) (requirementsConfig, globalConfig, error) {
	var (
		requirementsConfig = newRequirementsConfig()
		globalConfig       = newGlobalConfig()
	)

	globalConfig.ConfigMapHash = configMapHash
	globalConfig.ConfigMapLabelPrefixHash = configMapLabelPrefixHash

	if network.Spec.PodCIDR != "" {
		globalConfig.PodCIDR = network.Spec.PodCIDR
	}
	if cluster != nil && cluster.Shoot != nil && cluster.Shoot.Spec.Networking != nil && cluster.Shoot.Spec.Networking.Nodes != nil {
		globalConfig.NodeCIDR = *cluster.Shoot.Spec.Networking.Nodes
	}

	// The Cilium operator runs once per node. To safely enable HA, we must check guaranteed capacity (Minimum)
	if cluster != nil && cluster.Shoot != nil {
		countOfApplicableWorkerNodes := 0

		// Iterate over all worker groups and accumulate the guaranteed (minimum) count.
		// HA requires two or more guaranteed nodes. Anything less won’t reliably support a secondary operator instance.
		for _, k := range cluster.Shoot.Spec.Provider.Workers {
			countOfApplicableWorkerNodes += int(k.Minimum)
		}

		// If we can’t guarantee at least two nodes, don’t enable HA.
		if countOfApplicableWorkerNodes < 2 {
			globalConfig.OperatorHighAvailability.Enabled = false
		}
	}

	// Settings for Kube-Proxy disabled and using the HostService option
	// Also need to configure KubeProxy
	if cluster.Shoot.Spec.Kubernetes.KubeProxy != nil && cluster.Shoot.Spec.Kubernetes.KubeProxy.Enabled != nil && !*cluster.Shoot.Spec.Kubernetes.KubeProxy.Enabled {
		globalConfig.KubeProxyReplacement = ciliumv1alpha1.KubeProxyReplacementTrue
		globalConfig.KubeProxyReplacementHealthzBindAddr = "0.0.0.0:10256"
		globalConfig.Images[cilium.KubeProxyImageName] = imagevector.CiliumKubeProxyImage(cluster.Shoot.Spec.Kubernetes.Version)

		if config != nil && config.KubeProxy != nil && config.KubeProxy.ServiceHost != nil && config.KubeProxy.ServicePort != nil {
			globalConfig.K8sServiceHost = *config.KubeProxy.ServiceHost
			globalConfig.K8sServicePort = *config.KubeProxy.ServicePort
		} else {
			k8sServiceHost, err := getK8sServiceHost(cluster)
			if err != nil {
				return requirementsConfig, globalConfig, err
			}
			globalConfig.K8sServiceHost = k8sServiceHost
		}
		if globalConfig.K8sServiceHost == "" {
			return requirementsConfig, globalConfig, fmt.Errorf("required kubernetes service host missing while running without kube-proxy")
		}

		globalConfig.NodePort.Enabled = true
	}

	// If node local dns feature is enabled, enable local redirect policy
	if helper.IsNodeLocalDNSEnabled(cluster.Shoot.Spec.SystemComponents) {
		globalConfig.NodeLocalDNS.Enabled = true
		globalConfig.LocalRedirectPolicy.Enabled = true
	}

	if config == nil {
		return requirementsConfig, globalConfig, nil
	}

	// If Hubble enabled
	if config.Hubble != nil && config.Hubble.Enabled {
		requirementsConfig.Hubble.Enabled = config.Hubble.Enabled
	}

	// If ETCD enabled
	if config.Store != nil {
		if *config.Store != ciliumv1alpha1.Kubernetes {
			return requirementsConfig, globalConfig, fmt.Errorf("%s is not a supported value for field store", *config.Store)
		}
	}

	// check if IPv4 is enabled
	if config.IPv4 != nil {
		globalConfig.Ipv4.Enabled = config.IPv4.Enabled
	}
	// check if IPv6 is enabled
	if config.IPv6 != nil {
		globalConfig.Ipv6.Enabled = config.IPv6.Enabled
	}

	// check if BPFSocketLBHostnsOnly is enabled
	if config.BPFSocketLBHostnsOnly != nil {
		globalConfig.BPFSocketLBHostnsOnly.Enabled = config.BPFSocketLBHostnsOnly.Enabled
	}

	if config.CNI != nil {
		globalConfig.CNI.Exclusive = config.CNI.Exclusive
	}

	// check if tunnel mode is set
	if config.TunnelMode != nil {
		globalConfig.Tunnel = *config.TunnelMode
	}

	// check if debug is set
	if config.Debug != nil {
		globalConfig.Debug.Enabled = *config.Debug
	}

	// check if egress gateway is enabled
	if config.EgressGateway != nil {
		globalConfig.EgressGateway.Enabled = config.EgressGateway.Enabled
	}

	// check if mtu is set
	if config.MTU != nil {
		globalConfig.MTU = *config.MTU
	}

	// check if devices are set
	if len(config.Devices) > 0 {
		globalConfig.Devices = config.Devices
	}

	if config.DirectRoutingDevice != nil {
		globalConfig.NodePort.DirectRoutingDevice = *config.DirectRoutingDevice
	}

	// check if load balancing mode is set
	if config.LoadBalancingMode != nil {
		globalConfig.BPF = bpf{
			LoadBalancingMode: *config.LoadBalancingMode,
		}
	}

	// check if ipv4 native routing cidr is set
	if config.IPv4NativeRoutingCIDREnabled != nil && *config.IPv4NativeRoutingCIDREnabled {
		if cluster.Shoot.Spec.Networking.Pods == nil {
			return requirementsConfig, globalConfig, fmt.Errorf("pods cidr required for setting ipv4 native routing cidr was not yet set")
		}
		globalConfig.IPv4NativeRoutingCIDR = *cluster.Shoot.Spec.Networking.Pods
	}

	// check if ipv6 native routing cidr is set
	if config.IPv6NativeRoutingCIDREnabled != nil && *config.IPv6NativeRoutingCIDREnabled {
		if cluster.Shoot.Status.Networking.Pods == nil {
			return requirementsConfig, globalConfig, fmt.Errorf("pods cidr required for setting ipv6 native routing cidr was not yet set")
		}
		ipv6Pods, err := firstIPv6Range(cluster.Shoot.Status.Networking.Pods)
		if err != nil {
			return requirementsConfig, globalConfig, fmt.Errorf("failed to parse pods cidrs for setting ipv6 native routing cidr: %w", err)
		}
		globalConfig.IPv6NativeRoutingCIDR = ipv6Pods
	}

	if config.SnatToUpstreamDNS != nil && config.SnatToUpstreamDNS.Enabled {
		globalConfig.SnatToUpstreamDNS.Enabled = config.SnatToUpstreamDNS.Enabled
	}

	if config.SnatOutOfCluster != nil && config.SnatOutOfCluster.Enabled {
		globalConfig.SnatOutOfCluster.Enabled = config.SnatOutOfCluster.Enabled
	}

	if config.EnableIPv4Masquerade != nil {
		globalConfig.EnableIPv4Masquerade = *config.EnableIPv4Masquerade
	}

	if config.EnableIPv6Masquerade != nil {
		globalConfig.EnableIPv6Masquerade = *config.EnableIPv6Masquerade
	}

	if config.EnableBPFMasquerade != nil {
		globalConfig.EnableBPFMasquerade = *config.EnableBPFMasquerade
	}

	if config.Overlay != nil && !config.Overlay.Enabled && config.Overlay.CreatePodRoutes != nil {
		globalConfig.AutoDirectNodeRoutes = *config.Overlay.CreatePodRoutes
	}

	if config.BGPControlPlane != nil && config.BGPControlPlane.Enabled {
		globalConfig.BGPControlPlane.Enabled = config.BGPControlPlane.Enabled
	}

	globalConfig.IPAM.Mode = ipamMode

	return requirementsConfig, globalConfig, nil
}

func getK8sServiceHost(cluster *extensionscontroller.Cluster) (string, error) {
	if cluster == nil {
		return "", fmt.Errorf("cluster missing when retrieving kubernetes service host")
	}
	if cluster.Shoot == nil {
		return "", fmt.Errorf("shoot missing when retrieving kubernetes service host")
	}
	if len(cluster.Shoot.Status.AdvertisedAddresses) == 0 {
		return "", fmt.Errorf("advertised addresses missing in shoot status when retrieving kubernetes service host")
	}
	for _, address := range cluster.Shoot.Status.AdvertisedAddresses {
		if address.Name == "external" {
			url, err := url.Parse(address.URL)
			if err != nil {
				return "", fmt.Errorf("error while parsing external kubernetes service host: %s", err.Error())
			}
			return url.Hostname(), nil
		}
	}
	return "", fmt.Errorf("external address not found among advertised adresses")
}

func firstIPv6Range(cidrs []string) (string, error) {
	for i, cidr := range cidrs {
		prefix, err := netip.ParsePrefix(cidr)
		if err != nil {
			return "", fmt.Errorf("error while parsing pod cidr (%s, index %d): %w", cidr, i, err)
		}
		if prefix.Addr().Is6() {
			return cidr, nil
		}
	}
	return "", fmt.Errorf("no valid pod IPv6 range found in %v", cidrs)
}
