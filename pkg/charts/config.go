package charts

import ciliumv1alpha1 "github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium/v1alpha1"

type ciliumConfig struct {
	Requirements requirementsConfig `json:"requirements"`
	Global       globalConfig       `json:"global"`
}

type requirementsConfig struct {
	Agent     agent     `json:"agent"`
	Config    config    `json:"config"`
	Operator  operator  `json:"operator"`
	Preflight preflight `json:"preflight"`
	Hubble    hubble    `json:"hubble"`
}

type globalConfig struct {
	Tunnel                   ciliumv1alpha1.TunnelMode               `json:"tunnel"`
	IdentityAllocationMode   ciliumv1alpha1.IdentityAllocationMode   `json:"identityAllocationMode"`
	KubeProxyReplacement     ciliumv1alpha1.KubeProxyReplacementMode `json:"kubeProxyReplacement"`
	Etcd                     etcd                                    `json:"etcd"`
	Ipv4                     ipv4                                    `json:"ipv4"`
	Ipv6                     ipv6                                    `json:"ipv6"`
	Debug                    debug                                   `json:"debug"`
	Prometheus               prometheus                              `json:"prometheus"`
	OperatorHighAvailability operatorHighAvailability                `json:"operatorHighAvailability"`
	OperatorPrometheus       operatorPrometheus                      `json:"operatorPrometheus"`
	Images                   map[string]string                       `json:"images"`
	K8sServiceHost           string                                  `json:"k8sServiceHost"`
	K8sServicePort           int32                                   `json:"k8sServicePort"`
	NodePort                 nodePort                                `json:"nodePort"`
	PodCIDR                  string                                  `json:"podCIDR"`
	NodeCIDR                 string                                  `json:"nodeCIDR"`
	BPFSocketLBHostnsOnly    bpfSocketLBHostnsOnly                   `json:"bpfSocketLBHostnsOnly"`
	CNI                      cni                                     `json:"cni"`
	LocalRedirectPolicy      localRedirectPolicy                     `json:"localRedirectPolicy"`
	NodeLocalDNS             nodeLocalDNS                            `json:"nodeLocalDNS"`
	EgressGateway            egressGateway                           `json:"egressGateway"`
	IPv4NativeRoutingCIDR    string                                  `json:"ipv4NativeRoutingCIDR"`
	IPv6NativeRoutingCIDR    string                                  `json:"ipv6NativeRoutingCIDR"`
	MTU                      int                                     `json:"mtu"`
	Devices                  []string                                `json:"devices"`
	BPF                      bpf                                     `json:"bpf"`
	IPAM                     ipam                                    `json:"ipam"`
	SnatToUpstreamDNS        snatToUpstreamDNS                       `json:"snatToUpstreamDNS"`
	SnatOutOfCluster         snatOutOfCluster                        `json:"snatOutOfCluster"`
	AutoDirectNodeRoutes     bool                                    `json:"autoDirectNodeRoutes"`
	BGPControlPlane          bgpControlPlane                         `json:"bgpControlPlane"`
	ConfigMapHash            string                                  `json:"configMapHash"`
	ConfigMapLabelPrefixHash string                                  `json:"configMapLabelPrefixHash"`
	EnvoyConfig              envoyConfig                             `json:"envoyConfig"`
	GatewayAPI               gatewayAPI                              `json:"gatewayAPI"`
}

// etcd related configuration for cilium
type etcd struct {
	Enabled       bool   `json:"enabled"`
	Managed       bool   `json:"managed"`
	ClusterDomain string `json:"clusterDomain"`
	SSL           bool   `json:"ssl"`
	Endpoints     string `json:"endpoints"`
}

// prometheus configuration for cilium
type prometheus struct {
	Enabled        bool           `json:"enabled"`
	Port           int            `json:"port"`
	ServiceMonitor serviceMonitor `json:"serviceMonitor"`
}

type serviceMonitor struct {
	Enabled bool `json:"enabled"`
}

type operatorHighAvailability struct {
	Enabled bool `json:"enabled"`
}

// operatorPrometheus configuration for cilium
type operatorPrometheus struct {
	Enabled bool `json:"enabled"`
	Port    int  `json:"port"`
}

// Debug level option for cilium
type debug struct {
	// Debug Enabled is used to define whether Debug is required or not.
	Enabled bool `json:"enabled"`
}

// hubble enablement for cilium
type hubble struct {
	Enabled bool `json:"enabled"`
}

// nodePort enablement for cilium
type nodePort struct {
	// Nodeport Enabled is used to define whether Nodeport is required or not.
	Enabled             bool                        `json:"enabled"`
	Mode                ciliumv1alpha1.NodePortMode `json:"mode"`
	DirectRoutingDevice string                      `json:"directRoutingDevice"`
}

// operator configuration for cilium
type operator struct {
	// operator Enabled is used to define whether Operator is required or not.
	Enabled bool `json:"enabled"`
}

// preflight configuration for cilium
type preflight struct {
	// preflight is used to define whether Preflight is required or not.
	Enabled         bool   `json:"enabled"`
	ToFQDNSPreCache string `json:"toFQDNPreCache"`
}

// ipv4 configuration for cilium
type ipv4 struct {
	// ipv4 Enabled is used to define whether IPv4 address is required or not.
	Enabled bool `json:"enabled"`
}

// ipv6 configuration for cilium
type ipv6 struct {
	// ipv6 Enabled is used to define whether IPv6 address is required or not.
	Enabled bool `json:"enabled"`
}

// agent configuration for cilium
type agent struct {
	// agent is required or not
	Enabled        bool `json:"enabled"`
	SleepAfterInit bool `json:"sleepAfterInit"`
}

// config enable option for cilium
type config struct {
	// config is required or not
	Enabled bool `json:"enabled"`
}

type bpfSocketLBHostnsOnly struct {
	Enabled bool `json:"enabled"`
}

type cni struct {
	Exclusive bool `json:"exclusive"`
}
type localRedirectPolicy struct {
	Enabled bool `json:"enabled"`
}

type nodeLocalDNS struct {
	Enabled bool `json:"enabled"`
}

type egressGateway struct {
	Enabled bool `json:"enabled"`
}

type bpf struct {
	LoadBalancingMode ciliumv1alpha1.LoadBalancingMode `json:"lbMode"`
}

type ipam struct {
	Mode string `json:"mode"`
}

// snatToUpstreamDNS enables the masquerading of packets to the upstream dns server
type snatToUpstreamDNS struct {
	Enabled bool `json:"enabled"`
}

// snatOutOfCluster enables the masquerading of packets outside of the cluster
type snatOutOfCluster struct {
	Enabled bool `json:"enabled"`
}

// bgpControlPlane enables the BGP Control Plane
type bgpControlPlane struct {
	Enabled bool `json:"enabled"`
}

// envoyConfig enables CiliumEnvoyConfig CRD
type envoyConfig struct {
	Enabled          bool                        `json:"enabled"`
	SecretsNamespace EnvoyConfigSecretsNamespace `json:"secretsNamespace"`
	RetryInterval    string                      `json:"retryInterval"`
}

type EnvoyConfigSecretsNamespace struct {
	Create bool   `json:"create"`
	Name   string `json:"name"`
}

// gatewayAPI enables the Gateway API
type gatewayAPI struct {
	Enabled               bool                                     `json:"enabled"`
	EnableProxyProtocol   bool                                     `json:"enableProxyProtocol"`
	EnableAppProtocol     bool                                     `json:"enableAppProtocol"`
	EnableAlpn            bool                                     `json:"enableAlpn"`
	XffNumTrustedHops     int                                      `json:"xffNumTrustedHops"`
	ExternalTrafficPolicy string                                   `json:"externalTrafficPolicy"`
	GatewayClass          ciliumv1alpha1.GatewayAPIGatewayClass    `json:"gatewayClass"`
	SecretsNamespace      ciliumv1alpha1.GatewayAPISecretNamespace `json:"secretsNamespace"`
	HostNetwork           ciliumv1alpha1.GatewayAPIHostNetwork     `json:"hostNetwork"`
}
