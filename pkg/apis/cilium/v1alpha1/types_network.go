// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IdentityAllocationMode selects how identities are shared between cilium
// nodes by setting how they are stored. The options are "crd" or "kvstore".
type IdentityAllocationMode string

const (
	// CRD defines the crd IdentityAllocationMode type.
	CRD IdentityAllocationMode = "crd"
	// KVStore defines the kvstore IdentityAllocationMode type.
	KVStore IdentityAllocationMode = "kvstore"
)

// TunnelMode defines what tunnel mode to use for Cilium.
type TunnelMode string

const (
	// VXLan defines the vxlan tunnel mode
	VXLan TunnelMode = "vxlan"
	// Geneve defines the geneve tunnel mode.
	Geneve TunnelMode = "geneve"
	// Disabled defines the disabled tunnel mode.
	Disabled TunnelMode = "disabled"
)

// LoadBalancingMode defines what load balancing mode to use for Cilium.
type LoadBalancingMode string

const (
	// SNAT defines the snat load balancing mode
	SNAT LoadBalancingMode = "snat"
	// DSR defines the dsr load balancing mode
	DSR LoadBalancingMode = "dsr"
	// Hybrid defines the hybrid load balancing mode
	Hybrid LoadBalancingMode = "hybrid"
)

// KubeProxyReplacementMode defines which mode should kube-proxy run in.
// More infromation here: https://docs.cilium.io/en/v1.7/gettingstarted/kubeproxy-free/
type KubeProxyReplacementMode string

const (
	// Strict defines the strict kube-proxy replacement mode
	Strict KubeProxyReplacementMode = "strict"
	// Probe defines the probe kube-proxy replacement mode
	Probe KubeProxyReplacementMode = "probe"
	// Partial defines the partial kube-proxy replacement mode
	Partial KubeProxyReplacementMode = "partial"
	// Disabled defines the disabled kube-proxy replacement mode
	KubeProxyReplacementDisabled KubeProxyReplacementMode = "disabled"
	// KubeProxyReplacementTrue defines the true kube-proxy replacement mode
	KubeProxyReplacementTrue KubeProxyReplacementMode = "true"
	// KubeProxyReplacementFalse defines the false kube-proxy replacement mode
	KubeProxyReplacementFalse KubeProxyReplacementMode = "false"
)

// NodePortMode defines how NodePort services are enabled.
type NodePortMode string

const (
	// Hybrid defines the hybrid nodeport mode.
	Hybird NodePortMode = "hybrid"
)

// Store defines the kubernetes storage backend
type Store string

const (
	// Kubernetes defines the kubernetes CRD store type
	Kubernetes Store = "kubernetes"
)

// Hubble enablement for cilium
type Hubble struct {
	// Enabled indicates whether hubble is enabled or not.
	Enabled bool `json:"enabled"`
}

// IPv4 enablement for cilium
type IPv4 struct {
	// Enabled indicates whether IPv4 is enabled or not.
	Enabled bool `json:"enabled"`
}

// IPv6 enablement for cilium
type IPv6 struct {
	// Enabled indicates whether IPv6 is enabled or not.
	Enabled bool `json:"enabled"`
}

// BPFSocketLBHostnsOnly enablement for cilium
type BPFSocketLBHostnsOnly struct {
	Enabled bool `json:"enabled"`
}

// CNI configuration for cilium
type CNI struct {
	// false indicates that cilium will not overwrite its CNI configuration.
	Exclusive bool `json:"exclusive"`
}

// EgressGateway enablement for cilium
type EgressGateway struct {
	Enabled bool `json:"enabled"`
}

// Nodeport enablement for cilium
type Nodeport struct {
	// Enabled is used to define whether Nodeport is required or not.
	Enabled bool `json:"nodePortEnabled"`
	// Mode is the mode of NodePort feature
	Mode NodePortMode `json:"nodePortMode"`
}

// KubeProxy configuration for cilium
type KubeProxy struct {
	// ServiceHost specify the controlplane node IP Address.
	// +optional
	ServiceHost *string `json:"k8sServiceHost,omitempty"`
	// ServicePort specify the kube-apiserver port number.
	// +optional
	ServicePort *int32 `json:"k8sServicePort,omitempty"`
}

// Overlay configuration for cilium
type Overlay struct {
	// Enabled enables the network overlay.
	Enabled bool `json:"enabled"`
	// CreatePodRoutes installs routes to pods on all cluster nodes.
	// This will only work if the cluster nodes share a single L2 network.
	// +optional
	CreatePodRoutes *bool `json:"createPodRoutes,omitempty"`
}

// SnatToUpstreamDNS enables the masquerading of packets to the upstream dns server
type SnatToUpstreamDNS struct {
	Enabled bool `json:"enabled"`
}

// SnatOutOfCluster enables the masquerading of packets outside of the cluster
type SnatOutOfCluster struct {
	Enabled bool `json:"enabled"`
}

// BGPControlPlane enables the BGP Control Plane
type BGPControlPlane struct {
	Enabled bool `json:"enabled"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkConfig is a struct representing the configmap for the cilium
// networking plugin
type NetworkConfig struct {
	metav1.TypeMeta `json:",inline"`
	// Debug configuration to be enabled or not
	// +optional
	Debug *bool `json:"debug,omitempty"`
	// KubeProxy configuration to be enabled or not
	// +optional
	KubeProxy *KubeProxy `json:"kubeproxy,omitempty"`
	// Hubble configuration to be enabled or not
	// +optional
	Hubble *Hubble `json:"hubble,omitempty"`
	// TunnelMode configuration, it should be 'vxlan', 'geneve' or 'disabled'
	// +optional
	TunnelMode *TunnelMode `json:"tunnel,omitempty"`
	// Store can be either Kubernetes or etcd.
	// +optional
	Store *Store `json:"store,omitempty"`
	// Enable IPv4
	// +optional
	IPv4 *IPv4 `json:"ipv4,omitempty"`
	// Enable IPv6
	// +optional
	IPv6 *IPv6 `json:"ipv6,omitempty"`
	// BPFSocketLBHostnsOnly flag to be enabled or not
	// +optional
	BPFSocketLBHostnsOnly *BPFSocketLBHostnsOnly `json:"bpfSocketLBHostnsOnly,omitempty"`
	// CNI configuration for cilium
	// +optional
	CNI *CNI `json:"cni,omitempty"`
	// EgressGateway enablement for cilium
	// +optional
	EgressGateway *EgressGateway `json:"egressGateway,omitempty"`
	// MTU overwrites the auto-detected MTU of the underlying network
	// +optional
	MTU *int `json:"mtu,omitempty"`
	// Devices is the list of devices facing cluster/external network
	// +optional
	Devices []string `json:"devices,omitempty"`
	// DirectRoutingDevice is the device used for direct routing between Cilium nodes
	// +optional
	DirectRoutingDevice *string `json:"directRoutingDevice,omitempty"`
	// LoadBalancingMode configuration, it should be 'snat', 'dsr' or 'hybrid'
	// +optional
	LoadBalancingMode *LoadBalancingMode `json:"loadBalancingMode,omitempty"`
	// IPv4NativeRoutingCIDRMode will set the ipv4 native routing cidr from the network configs node's cidr if enabled.
	// +optional
	IPv4NativeRoutingCIDREnabled *bool `json:"ipv4NativeRoutingCIDREnabled,omitempty"`
	// IPv6NativeRoutingCIDRMode will set the ipv6 native routing cidr from the network configs node's cidr if enabled.
	// +optional
	IPv6NativeRoutingCIDREnabled *bool `json:"ipv6NativeRoutingCIDREnabled,omitempty"`
	// Overlay enables the network overlay
	// +optional
	Overlay *Overlay `json:"overlay,omitempty"`
	// SnatToUpstreamDNS enables the masquerading of packets to the upstream dns server
	// +optional
	SnatToUpstreamDNS *SnatToUpstreamDNS `json:"snatToUpstreamDNS,omitempty"`
	// SnatOutOfCluster enables the masquerading of packets outside of the cluster
	// +optional
	SnatOutOfCluster *SnatOutOfCluster `json:"snatOutOfCluster,omitempty"`
	// BGPControlPlane enables the BGP Control Plane
	// +optional
	BGPControlPlane *BGPControlPlane `json:"bgpControlPlane,omitempty"`
}
