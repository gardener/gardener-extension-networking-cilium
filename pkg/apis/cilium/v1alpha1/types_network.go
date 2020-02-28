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

package v1alpha1

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

// Encryption related configuration for cilium cluster
type Encryption struct {
	EncryptionEnabled bool   `json:"encryptionEnabled"`
	KeyFile           string `json:"keyFile"`
	MountPath         string `json:"mountPath"`
	SecretName        string `json:"secretName"`
	NodeEncryption    bool   `json:"nodeEncryption"`
	Interface         string `json:"interface"`
}

// EtcdConfig related configuration for cilium
type EtcdConfig struct {
	EtcdEnabled    bool     `json:"etcdEnabled"`
	EtcdManaged    bool     `json:"etcdManaged"`
	EtcdsslEnabled bool     `json:"etcdsslEnabled"`
	EtcdEndpoints  []string `json:"etcdEndPoints"`
	CAFile         string   `json:"caFile"`
	CertFile       string   `json:"certFile"`
	KeyFile        string   `json:"keyFile"`
}

// Masquerade configuration for cilium
type Masquerade struct {
	MasqueradeEnabled bool `json:"masqueradeEnabled"`
}

// Prometheus configuration for cilium
type Prometheus struct {
	PrometheusEnabled     bool `json:"prometheusEnabled"`
	Port                  int  `json:"port"`
	ServiceMonitorEnabled bool `json:"serviceMonitorEnabled"`
}

// OperatorPrometheus configuration for cilium
type OperatorPrometheus struct {
	OperatorPrometheusEnabled bool `json:"operatorprometheusEnabled"`
	Port                      int  `json:"port"`
}

// InstallIPTableRules configuration for cilium
type InstallIPTableRules struct {
	InstallIPTableRulesEnabled bool `json:"installIptableRulesEnabled"`
}

// IPvlan configuration for cilium
type IPvlan struct {
	IPvlanEnabled bool       `json:"ipvlanenabled"`
	MasterDevice  string     `json:"masterDevice"`
	IPVlanMode    IPVlanMode `json:"ipvlanMode"`
}

// Flannel configuration for cilium
type Flannel struct {
	FlannelEnabled  bool   `json:"flannelEnabled"`
	MasterDevice    string `json:"masterDevice"`
	UninstallOnExit bool   `json:"uninstallOnExit"`
}

// AutoDirectNodeRoutes configuration for cilium
type AutoDirectNodeRoutes struct {
	AutoDirectNodeRoutesEnabled bool `json:"autoDirectNodeRoutesEnabled"`
}

// ExternalIPs configuration for cilium
type ExternalIP struct {
	// ExternalIPenabled is used to define whether ExternalIP address is required or not.
	ExternalIPEnabled bool `json:"externalipEnabled"`
}

// Debug level option for cilium
type Debug struct {
	// Enabled is used to define whether Debug is required or not.
	Enabled bool `json:"debugEnabled"`
}

// Hubble enablement for cilium
type Hubble struct {
	// HubbleEnabled is used to define whether Hubble is required or not.
	HubbleEnabled   bool     `json:"hubbleEnabled"`
	ListenAddresses []string `json:"listenAddresses"`
	EventQueueSize  int32    `json:"eventQueueSize"`
	FlowBufferSize  int32    `json:"flowBufferSize"`
	MetricServer    string   `json:"metricServer"`
	Metrics         []string `json:"metrics"`
	UI              bool     `json:"uiEnabled"`
}

// Bpf configuration for cilium
type Bpf struct {
	// Bpf tunning for cilium.
	WaitForMount       bool   `json:"waitForMount"`
	PreAllocateMaps    bool   `json:"preallocateMaps"`
	CtTCPMax           int32  `json:"ctTcpMax"`
	CtAnyMax           int32  `json:"ctAnyMax"`
	MonitorAggregation string `json:"monitorAggregation"`
	MonitorFlags       string `json:"monitorFlags"`
}

// Nodeport enablement for cilium
type Nodeport struct {
	// NodeportEnabled is used to define whether Nodeport is required or not.
	NodeportEnabled bool   `json:"nodeportEnabled"`
	NodeportMode    string `json:"nodeportMode"`
	NodeportDevice  string `json:"nodeportDevice"`
	NodeportRange   string `json:"nodeportRange"`
}

// SockOps configuration for cilium
type SockOps struct {
	// SockOpsEnabled is used to define whether SockOps is required or not.
	SockOpsEnabled bool `json:"sockOpsEnabled"`
}

// HostServices configuration for cilium
type HostServices struct {
	// HostServices is used to define whether IPv4 address is required or not.
	HostServicesEnabled bool   `json:"hostServicesEnabled"`
	Protocols           string `json:"protocols"`
}

// CleanState configuration for cilium
type CleanState struct {
	// CleanStateEnabled is used to define whether CleanState is required or not on Startup.
	CleanStateEnabled bool `json:"cleanStateEnabled"`
}

// CleanBpfState configuration for cilium
type CleanBpfState struct {
	// CleanBpfStateEnabled is used to define whether CleanBpfState is required or not on Startup.
	CleanBpfStateEnabled bool `json:"cleanBpfStateEnabled"`
}

// Operator configuration for cilium
type Operator struct {
	// OperatorEnabled is used to define whether Operator is required or not.
	OperatorEnabled bool `json:"operatorEnabled"`
}

// Preflight configuration for cilium
type Preflight struct {
	// Preflight is used to define whether Preflight is required or not.
	PreflightEnabled bool   `json:"preflightEnabled"`
	PreCacheToFQDNS  string `json:"precacheToFQDNS"`
}

// IPv4 configuration for cilium
type IPv4 struct {
	// IPv4Enabled is used to define whether IPv4 address is required or not.
	IPv4Enabled bool `json:"ipv4Enabled"`
}

// IPv6 configuration for cilium
type IPv6 struct {
	// IPv6Enabled is used to define whether IPv6 address is required or not.
	IPv6Enabled bool `json:"ipv6Enabled"`
}

// ClusterName configuration for cilium
type ClusterName struct {
	// ClusterName is used to define the name of the cluster.
	ClusterName string `json:"clusterName"`
}

// KubeProxy configuration for cilium
type KubeProxy struct {
	// KubeProxyEnabled is used to set if KubeProxy is required or not.
	KubeProxyEnabled bool `json:"kubeProxyEnabled"`
}

// SynchronizeK8sNodes configuration for cilium
type SynchronizeK8sNodes struct {
	// SynchronizeK8sNodes is required or not.
	SynchronizeK8sNodesEnabled bool `json:"synchronizeK8sNodesEnabled"`
}

// RemoteNodeIdentity configuration for cilium
type RemoteNodeIdentity struct {
	// RemoteNodeIdentity is required or not.
	RemoteNodeIdentityEnabled bool `json:"remoteNodeIdentityEnabled"`
}

// TLSsecretsBackend configuration for cilium
type TLSsecretsBackend struct {
	// TLSsecretsBackend is required or not.
	TLSsecretsBackend string `json:"tlsSecretsBackend"`
}

// WellKnownIdentities configuration for cilium
type WellKnownIdentities struct {
	// WellKnownIdentities is required or not.
	WellKnownIdentitiesEnabled bool `json:"wellKnownIdentitiesEnabled"`
}

// PolicyAuditMode configuration for cilium
type PolicyAuditMode struct {
	// PolicyAuditMode is required or not
	PolicyAuditModeEnabled bool `json:"policyAuditModeEnabled"`
}

// Psp configuration for cilium
type Psp struct {
	// Psp is required or not
	PspEnabled bool `json:"pspEnabled"`
}

// K8SIPv4PodCIDR configuration for cilium
type K8SIPv4PodCIDR struct {
	// K8SIPv4PodCIDR is required or not
	K8SIPv4PodCIDREnabled bool `json:"k8sIPv4PodCIDREnabled"`
}

// NodeInit configuration for cilium
type NodeInit struct {
	// NodeInit is required or not
	NodeInitEnabled                   bool   `json:"nodeInitEnabled"`
	BootStrapFile                     string `json:"bootStrapFile"`
	NonGlobalRestartPods              bool   `json:"restartPods"`
	NonGlobalReconfigureKubelet       bool   `json:"reconfigureKubelet"`
	NonGlobalRemoveCbrBridge          bool   `json:"removeCbrBridge"`
	NonGlobalRevertReconfigureKubelet bool   `json:"revertReconfigureKubelet"`
}

// Azure configuration for cilium
type Azure struct {
	AzureEnabled        bool   `json:"azureEnabled"`
	AzureResourceGroup  string `json:"azureResourceGroup"`
	AzureSubscriptionID string `json:"azureSubscriptionID"`
	AzuretenantID       string `json:"azureTenantID"`
	AzureClientID       string `json:"azureClientID"`
	AzureClientSecret   string `json:"azureClientSecret"`
}

// CNI configuration for cilium
type CNI struct {
	CniInstallEnabled       bool   `json:"cniInstallEnabled"`
	CniChainginMode         string `json:"cniChainingMode"`
	CniCustomConfigEnabled  bool   `json:"cniCustomConfigEnabled"`
	CniConfigPath           string `json:"cniConfigPath"`
	CniBinPath              string `json:"cnibinPath"`
	CniConfigMapKey         string `json:"cniConfigMapKey"`
	CniConfigMap            string `json:"cniConfigMap"`
	CniConfFileMountPath    string `json:"cniConfFileMountPath"`
	CniHostConfDirMountPath string `json:"cniHostConfDirMountPath"`
}

// Eni configuration for cilium
type Eni struct {
	// Eni is required or not
	EniEnabled                 bool   `json:"eniEnabled"`
	EgressMasqueradeInterfaces string `json:"egressMasqueradeInterfaces"`
}

// EndpointRoutes configuration for cilium
type EndpointRoutes struct {
	// EndpointRoutes is required or not
	EndpointRoutesEnabled bool `json:"endpointRoutesEnabled"`
}

// XTSocketFallback configuration for cilium
type XTSocketFallback struct {
	// XTSocketFallback is required or not
	XTSocketFallbackEnabled bool `json:"xtSocketFallbackEnabled"`
}

// Agent configuration for cilium
type Agent struct {
	// Agent is required or not
	AgentEnabled                     bool `json:"agentEnabled"`
	AgentSleepAfterInitEnabled       bool `json:"agentSleepAfterInitEnabled"`
	AgentKeepDeprecatedLabelsEnabled bool `json:"agentKeepDeprecatedLabelsEnabled"`
}

// Config enable option for cilium
type ConfigEnable struct {
	// ConfigEnable is required or not
	Enabled bool `json:"configEnabled"`
}

// DaemonRunPath option for cilium
type DaemonRunPath struct {
	// DaemonRunPath is required or not
	DaemonRunPath string `json:"daemonRunPath"`
}

// GoogleGKE  option for cilium
type GoogleGKE struct {
	// GoogleGKE is required or not
	GoogleGkeEnabled bool `json:"googleGkeEnabled"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkConfig is a struct representing the configmap for the cilium
// networking plugin
type NetworkConfig struct {
	metav1.TypeMeta `json:",inline"`
	// Debug configuration to be enabled or not
	// +optional
	Debug *Debug `json:"debug,omitempty"`
	// GoogleGKE enabled or not
	GoogleGKE GoogleGKE `json:"googleGKE"`
	// NodeInit config for cilium
	NodeInit NodeInit `json:"nodeInit"`
	// Operator enabled or not
	// +optional
	Operator *Operator `json:"operator,omitempty"`
	// ClusterName for cluster
	// +optional
	ClusterName *ClusterName `json:"clustername,omitempty"`
	// Preflight configuration
	// +optional
	Preflight *Preflight `json:"preflight,omitempty"`
	// Cleanstate configuration
	// +optional
	CleanState *CleanState `json:"cleanstate,omitempty"`
	// CleanBpfstate configuraton
	// +optional
	CleanBpfState *CleanBpfState `json:"cleanbpfstate,omitempty"`
	// Bpf tunning configuraton
	// +optional
	Bpf *Bpf `json:"bpf,omitempty"`
	// HostServices configuration
	// +optional
	HostServices *HostServices `json:"hostservices,omitempty"`
	// SockOps configuration
	// +optional
	SockOps *SockOps `json:"sockops,omitempty"`
	// InstallIPTableRules configuration
	// +optional
	InstallIPTableRules *InstallIPTableRules `json:"installiptablerules,omitempty"`
	// Prometheus configuration
	// +optional
	Prometheus *Prometheus `json:"prometheus,omitempty"`
	// SynchronizeK8sNodes configuration
	SynchronizeK8sNodes SynchronizeK8sNodes `json:"synchronizeK8sNodes"`
	// RemoteNodeIdentity configuration
	RemoteNodeIdentity RemoteNodeIdentity `json:"remoteNodeIdentity"`
	// TLSsecretsBackend configuration
	TLSsecretsBackend TLSsecretsBackend `json:"tlsSecretsBackend"`
	// WellKnownIdentities configuration
	WellKnownIdentities WellKnownIdentities `json:"wellKnownIdentities"`
	// CNI configuration
	// +optional
	CNI *CNI `json:"cni"`
	// PSP configuration
	Psp Psp `json:"psp"`
	// Agent configuration
	Agent Agent `json:"agent"`
	// Config enable configuration
	ConfigEnable ConfigEnable `json:"config"`
	// EnableXTSocketFallback configuration
	XTSocketFallback XTSocketFallback `json:"enableXTSocketFallback"`
	// K8SIPv4PodCIDR is required or not
	K8SIPv4PodCIDR K8SIPv4PodCIDR `json:"k8sIPv4PodCIDR"`
	// PolicyAuditMode configuration
	PolicyAuditMode PolicyAuditMode `json:"policyAuditMode"`
	// DaemonRunPath configuration
	// +optional
	DaemonRunPath *DaemonRunPath `json:"daemonRunPath"`
	// Eni configuration
	Eni Eni `json:"eni"`
	// EndpointRoutes configuration
	EndpointRoutes EndpointRoutes `json:"endpointRoutes"`
	// Azure configuration
	// +optional
	Azure *Azure `json:"azure"`
	// OperatorPrometheus configuration
	// +optional
	OperatorPrometheus *OperatorPrometheus `json:"operatorprometheus,omitempty"`
	// KubeProxy configuration to be enabled or not
	// +optional
	KubeProxy *KubeProxy `json:"kubeproxy,omitempty"`
	// KubeProxyReplacementMode configuration. It should be
	// 'probe', 'strict', 'partial'
	// If KubeProxy is disabled it is required to configure the
	// KubeProxyReplacementMode
	// +optional
	KubeProxyReplacementMode *KubeProxyReplacementMode `json:"kubeproxyreplacementmode,omitempty"`
	// Etcd configuration. Configure managed or not and
	// then configure the etcd-secrets.
	EtcdConfig EtcdConfig `json:"etcdconfig"`
	// Masquerade configuraton to be enabled or not
	Masquerade Masquerade `json:"masquerade"`
	// IPvlan configuraton to be enabled or not
	// +optional
	IPvlan *IPvlan `json:"ipvlan,omitempty"`
	// Flannel configuraton to be enabled or not
	// +optional
	Flannel *Flannel `json:"flannel,omitempty"`
	// AutoDirectNodeRoutes configuraton to be enabled or not
	// +optional
	AutoDirectNodeRoutes *AutoDirectNodeRoutes `json:"autoDirectNodeRoutes"`
	// Encryption configuraton
	// +optional
	Encryption *Encryption `json:"encryption"`
	// Hubble configuration to be enabled or not
	// +optional
	Hubble *Hubble `json:"hubble,omitempty"`
	// Nodeport enablement for cilium
	// +optional
	Nodeport *Nodeport `json:"nodeport,omitempty"`
	// IPv4 configuration to be enabled or not
	IPv4 IPv4 `json:"ipv4"`
	// IPv6 configuration to be enabled or not
	// +optional
	IPv6 *IPv6 `json:"ipv6,omitempty"`
	// ExternalIP configuration to be enabled or not
	// +optional
	ExternalIP *ExternalIP `json:"externalip,omitempty"`
	// IPAMOptionMode configuration, it should be either
	// 'crd' or 'eni' or 'default' to be enabled or not
	// +optional
	IPAMOptionMode *IPAMOptionMode `json:"ipam-mode,omitempty"`
	// IdentityAllocationModeOption configuration, it should be either
	// 'crd' or 'kvstore'
	IdentityAllocationModeOption IdentityAllocationModeOption `json:"identityAllocationMode,omitempty"`
	// TunnelMode configuration, it should be 'vxlan', 'geneve' or 'disabled'
	TunnelMode TunnelMode `json:"tunnel"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkStatus contains information about created Network resources.
type NetworkStatus struct {
	metav1.TypeMeta `json:",inline"`
}
