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
	Tunnel                 ciliumv1alpha1.TunnelMode               `json:"tunnel"`
	IdentityAllocationMode ciliumv1alpha1.IdentityAllocationMode   `json:"identityAllocationMode"`
	KubeProxyReplacement   ciliumv1alpha1.KubeProxyReplacementMode `json:"kubeProxyReplacement"`
	Etcd                   etcd                                    `json:"etcd"`
	Ipv4                   ipv4                                    `json:"ipv4"`
	Ipv6                   ipv6                                    `json:"ipv6"`
	Debug                  debug                                   `json:"debug"`
	Prometheus             prometheus                              `json:"prometheus"`
	OperatorPrometheus     operatorPrometheus                      `json:"operatorPrometheus"`
	Psp                    psp                                     `json:"psp"`
	Images                 map[string]string                       `json:"images"`
	K8sServiceHost         string                                  `json:"k8sServiceHost"`
	K8sServicePort         int32                                   `json:"k8sServicePort"`
	NodePort               nodePort                                `json:"nodePort"`
	PodCIDR                string                                  `json:"podCIDR"`
	BPFSocketLBHostnsOnly  bpfSocketLBHostnsOnly                   `json:"bpfSocketLBHostnsOnly"`
	LocalRedirectPolicy    localRedirectPolicy                     `json:"localRedirectPolicy"`
	NodeLocalDNS           nodeLocalDNS                            `json:"nodeLocalDNS"`
	EgressGateway          egressGateway                           `json:"egressGateway"`
	IPv4NativeRoutingCIDR  string                                  `json:"ipv4NativeRoutingCIDR"`
	MTU                    int                                     `json:"mtu"`
	Devices                []string                                `json:"devices"`
	BPF                    bpf                                     `json:"bpf"`
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
	Enabled bool                        `json:"enabled"`
	Mode    ciliumv1alpha1.NodePortMode `json:"mode"`
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

// psp configuration for cilium
type psp struct {
	// psp is required or not
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
