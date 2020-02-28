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

package cilium

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TunnelMode string

const (
	VXLan    TunnelMode = "vxlan"
	Geneve   TunnelMode = "geneve"
	NoTunnel TunnelMode = "disabled"
)

type IPVlanMode string

const (
	PureL3           IPVlanMode = "L3"
	L3S              IPVlanMode = "L3S"
	L3withMasquerade IPVlanMode = "L3E"
)

type KubeProxyReplacementMode string

const (
	Strict  KubeProxyReplacementMode = "strict"
	Probe   KubeProxyReplacementMode = "probe"
	Partial KubeProxyReplacementMode = "partial"
)

type IPAMOptionMode string

const (
	CRD        IPAMOptionMode = "crd"
	ENI        IPAMOptionMode = "eni"
	K8NODECIDR IPAMOptionMode = "default"
)

type IdentityAllocationModeOption string

const (
	Option_crd     IdentityAllocationModeOption = "crd"
	Option_kvstore IdentityAllocationModeOption = "kvstore"
)

// Encryption configuration for cilium
type Encryption struct {
	EncryptionEnabled bool
	KeyFile           string
	MountPath         string
	SecretName        string
	NodeEncryption    bool
	Interface         string
}

// EtcdConfig related configuration for cilium
type EtcdConfig struct {
	EtcdEnabled    bool
	EtcdManaged    bool
	EtcdsslEnabled bool
	EtcdEndpoints  []string
	CAFile         string
	CertFile       string
	KeyFile        string
}

// Masquerade configuration for cilium
type Masquerade struct {
	MasqueradeEnabled bool
}

// Prometheus configuration for cilium
type Prometheus struct {
	PrometheusEnabled     bool
	Port                  int
	ServiceMonitorEnabled bool
}

// OperatorPrometheus configuration for cilium
type OperatorPrometheus struct {
	OperatorPrometheusEnabled bool
	Port                      int
}

// InstallIPTableRules configuration for cilium
type InstallIPTableRules struct {
	InstallIPTableRulesEnabled bool
}

// IPvlan configuration for cilium
type IPvlan struct {
	IPvlanEnabled bool       `json:"ipvlanenabled"`
	MasterDevice  string     `json:"masterDevice"`
	IPVlanMode    IPVlanMode `json:"ipvlanMode"`
}

// Flannel configuration for cilium
type Flannel struct {
	FlannelEnabled  bool
	MasterDevice    string
	UninstallOnExit bool
}

// AutoDirectNodeRoutes configuration for cilium
type AutoDirectNodeRoutes struct {
	AutoDirectNodeRoutesEnabled bool
}

// ExternalIPs configuration for cilium
type ExternalIP struct {
	// ExternalIPenabled is used to define whether ExternalIP address is required or not.
	ExternalIPEnabled bool
}

// Debug level option for cilium
type Debug struct {
	// Enabled is used to define whether Debug is required or not.
	Enabled bool
}

// Hubble enablement for cilium
type Hubble struct {
	// HubbleEnabled is used to define whether Hubble is required or not.
	HubbleEnabled   bool
	ListenAddresses []string
	EventQueueSize  int32
	FlowBufferSize  int32
	MetricServer    string
	Metrics         []string
	UI              bool
}

// Bpf configuration for cilium
type Bpf struct {
	// Bpf tunning for cilium.
	WaitForMount       bool
	PreAllocateMaps    bool
	CtTCPMax           int32
	CtAnyMax           int32
	MonitorAggregation string
	MonitorFlags       string
}

// Nodeport enablement for cilium
type Nodeport struct {
	// NodeportEnabled is used to define whether Nodeport is required or not.
	NodeportEnabled bool
	NodeportMode    string
	NodeportDevice  string
	NodeportRange   string
}

// SockOps configuration for cilium
type SockOps struct {
	// SockOpsEnabled is used to define whether SockOps is required or not.
	SockOpsEnabled bool
}

// HostServices configuration for cilium
type HostServices struct {
	// HostServices is used to define whether IPv4 address is required or not.
	HostServicesEnabled bool
	Protocols           string
}

// CleanState configuration for cilium
type CleanState struct {
	// CleanStateEnabled is used to define whether CleanState is required or not on Startup.
	CleanStateEnabled bool
}

// CleanBpfState configuration for cilium
type CleanBpfState struct {
	// CleanBpfStateEnabled is used to define whether CleanBpfState is required or not on Startup.
	CleanBpfStateEnabled bool
}

// Operator configuration for cilium
type Operator struct {
	// OperatorEnabled is used to define whether Operator is required or not.
	OperatorEnabled bool
}

// Preflight configuration for cilium
type Preflight struct {
	// Preflight is used to define whether Preflight is required or not.
	PreflightEnabled bool
	PreCacheToFQDNS  string
}

// IPv4 configuration for cilium
type IPv4 struct {
	// IPv4Enabled is used to define whether IPv4 address is required or not.
	IPv4Enabled bool
}

// IPv6 configuration for cilium
type IPv6 struct {
	// IPv6Enabled is used to define whether IPv6 address is required or not.
	IPv6Enabled bool
}

// ClusterName configuration for cilium
type ClusterName struct {
	// ClusterName is used to define the name of the cluster.
	ClusterName string
}

// KubeProxy configuration for cilium
type KubeProxy struct {
	// KubeProxyEnabled is used to set if KubeProxy is required or not.
	KubeProxyEnabled bool
}

// SynchronizeK8sNodes configuration for cilium
type SynchronizeK8sNodes struct {
	// SynchronizeK8sNodes is required or not.
	SynchronizeK8sNodesEnabled bool
}

// RemoteNodeIdentity configuration for cilium
type RemoteNodeIdentity struct {
	// RemoteNodeIdentity is required or not.
	RemoteNodeIdentityEnabled bool
}

// TLSsecretsBackend configuration for cilium
type TLSsecretsBackend struct {
	// TLSsecretsBackend is required or not.
	TLSsecretsBackend string
}

// WellKnownIdentities configuration for cilium
type WellKnownIdentities struct {
	// WellKnownIdentities is required or not.
	WellKnownIdentitiesEnabled bool
}

// PolicyAuditMode configuration for cilium
type PolicyAuditMode struct {
	// PolicyAuditMode is required or not
	PolicyAuditModeEnabled bool
}

// Psp configuration for cilium
type Psp struct {
	// Psp is required or not
	PspEnabled bool
}

// K8SIPv4PodCIDR configuration for cilium
type K8SIPv4PodCIDR struct {
	// K8SIPv4PodCIDR is required or not
	K8SIPv4PodCIDREnabled bool
}

// NodeInit configuration for cilium
type NodeInit struct {
	// NodeInit is required or not
	NodeInitEnabled                   bool
	BootStrapFile                     string
	NonGlobalRestartPods              bool
	NonGlobalReconfigureKubelet       bool
	NonGlobalRemoveCbrBridge          bool
	NonGlobalRevertReconfigureKubelet bool
}

// Azure configuration for cilium
type Azure struct {
	AzureEnabled        bool
	AzureResourceGroup  string
	AzureSubscriptionID string
	AzuretenantID       string
	AzureClientID       string
	AzureClientSecret   string
}

// CNI configuration for cilium
type CNI struct {
	CniInstallEnabled       bool
	CniChainginMode         string
	CniCustomConfigEnabled  bool
	CniConfigPath           string
	CniBinPath              string
	CniConfigMapKey         string
	CniConfigMap            string
	CniConfFileMountPath    string
	CniHostConfDirMountPath string
}

// Eni configuration for cilium
type Eni struct {
	// Eni is required or not
	EniEnabled                 bool
	EgressMasqueradeInterfaces string
}

// EndpointRoutes configuration for cilium
type EndpointRoutes struct {
	// EndpointRoutes is required or not
	EndpointRoutesEnabled bool
}

// XTSocketFallback configuration for cilium
type XTSocketFallback struct {
	// XTSocketFallback is required or not
	XTSocketFallbackEnabled bool
}

// Agent configuration for cilium
type Agent struct {
	// Agent is required or not
	AgentEnabled                     bool
	AgentSleepAfterInitEnabled       bool
	AgentKeepDeprecatedLabelsEnabled bool
}

// Config enable option for cilium
type ConfigEnable struct {
	// ConfigEnable is required or not
	Enabled bool
}

// DaemonRunPath option for cilium
type DaemonRunPath struct {
	// DaemonRunPath is required or not
	DaemonRunPath string
}

// GoogleGKE  option for cilium
type GoogleGKE struct {
	// GoogleGKE is required or not
	GoogleGkeEnabled bool
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkConfig is a struct representing the configmap for the cilium
// networking plugin
type NetworkConfig struct {
	metav1.TypeMeta
	// Debug configuration to be enabled or not
	// +optional
	Debug *Debug
	// GoogleGKE enabled or not
	GoogleGKE GoogleGKE
	// NodeInit config for cilium
	NodeInit NodeInit
	// Operator enabled or not
	// +optional
	Operator *Operator
	// ClusterName for cluster
	// +optional
	ClusterName *ClusterName
	// Preflight configuration
	// +optional
	Preflight *Preflight
	// Cleanstate configuration
	// +optional
	CleanState *CleanState
	// CleanBpfstate configuraton
	// +optional
	CleanBpfState *CleanBpfState
	// Bpf tunning configuraton
	// +optional
	Bpf *Bpf
	// HostServices configuration
	// +optional
	HostServices *HostServices
	// SockOps configuration
	// +optional
	SockOps *SockOps
	// InstallIPTableRules configuration
	// +optional
	InstallIPTableRules *InstallIPTableRules
	// Prometheus configuration
	// +optional
	Prometheus *Prometheus
	// SynchronizeK8sNodes configuration
	SynchronizeK8sNodes SynchronizeK8sNodes
	// RemoteNodeIdentity configuration
	RemoteNodeIdentity RemoteNodeIdentity
	// TLSsecretsBackend configuration
	TLSsecretsBackend TLSsecretsBackend
	// WellKnownIdentities configuration
	WellKnownIdentities WellKnownIdentities
	// CNI configuration
	// +optional
	CNI *CNI
	// PSP configuration
	Psp Psp
	// Agent configuration
	Agent Agent
	// Config enable configuration
	ConfigEnable ConfigEnable
	// EnableXTSocketFallback configuration
	XTSocketFallback XTSocketFallback
	// K8SIPv4PodCIDR is required or not
	K8SIPv4PodCIDR K8SIPv4PodCIDR
	// PolicyAuditMode configuration
	PolicyAuditMode PolicyAuditMode
	// DaemonRunPath configuration
	// +optional
	DaemonRunPath *DaemonRunPath
	// Eni configuration
	Eni Eni
	// EndpointRoutes configuration
	EndpointRoutes EndpointRoutes
	// Azure configuration
	// +optional
	Azure *Azure
	// OperatorPrometheus configuration
	// +optional
	OperatorPrometheus *OperatorPrometheus
	// KubeProxy configuration to be enabled or not
	// +optional
	KubeProxy *KubeProxy
	// KubeProxyReplacementMode configuration. It should be
	// 'probe', 'strict', 'partial'
	// If KubeProxy is disabled it is required to configure the
	// KubeProxyReplacementMode
	// +optional
	KubeProxyReplacementMode *KubeProxyReplacementMode
	// Etcd configuration. Configure managed or not and
	// then configure the etcd-secrets.
	EtcdConfig EtcdConfig
	// Masquerade configuraton to be enabled or not
	Masquerade Masquerade
	// IPvlan configuraton to be enabled or not
	// +optional
	IPvlan *IPvlan
	// Flannel configuraton to be enabled or not
	// +optional
	Flannel *Flannel
	// AutoDirectNodeRoutes configuraton to be enabled or not
	// +optional
	AutoDirectNodeRoutes *AutoDirectNodeRoutes
	// Encryption configuraton
	// +optional
	Encryption *Encryption
	// Hubble configuration to be enabled or not
	// +optional
	Hubble *Hubble
	// Nodeport enablement for cilium
	// +optional
	Nodeport *Nodeport
	// IPv4 configuration to be enabled or not
	IPv4 IPv4
	// IPv6 configuration to be enabled or not
	// +optional
	IPv6 *IPv6
	// ExternalIP configuration to be enabled or not
	// +optional
	ExternalIP *ExternalIP
	// IPAMOptionMode configuration, it should be either
	// 'crd' or 'eni' or 'default' to be enabled or not
	// +optional
	IPAMOptionMode *IPAMOptionMode
	// IdentityAllocationModeOption configuration, it should be either
	// 'crd' or 'kvstore'
	IdentityAllocationModeOption IdentityAllocationModeOption
	// TunnelMode configuration, it should be 'vxlan', 'geneve' or 'disabled'
	TunnelMode TunnelMode
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkStatus contains information about created Network resources.
type NetworkStatus struct {
	metav1.TypeMeta
}
