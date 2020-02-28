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
	"encoding/json"
	"fmt"

	ciliumv1alpha1 "github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium/v1alpha1"
	"github.com/gardener/gardener-extension-networking-cilium/pkg/cilium"
	"github.com/gardener/gardener-extension-networking-cilium/pkg/imagevector"

	//"k8s.io/apiserver/pkg/admission/configuration"

	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
)

type ciliumConfig struct {
	Agent     agent             `json:"agent"`
	Config    config            `json:"config"`
	Operator  operator          `json:"operator"`
	Preflight preflight         `json:"preflight"`
	Global    global            `json:"global"`
	NodeInit  nonGlobalNodeinit `json:"nodeinit"`
}

type nonGlobalNodeinit struct {
	RestartPods              bool `json:"restartPods"`
	ReconfigureKubelet       bool `json:"reconfigureKubelet"`
	RemoveCbrBridge          bool `json:"removeCbrBridge"`
	RevertReconfigureKubelet bool `json:"revertReconfigureKubelet"`
}

type global struct {
	Registry                   string                                      `json:"registry"`
	Tag                        string                                      `json:"tag"`
	PullPolicy                 string                                      `json:"always"`
	Etcd                       etcd                                        `json:"etcd"`
	IdentityAllocationMode     ciliumv1alpha1.IdentityAllocationModeOption `json:"identityAllocationMode"`
	Ipv4                       ipv4                                        `json:"ipv4"`
	Ipv6                       ipv6                                        `json:"ipv6"`
	Debug                      debug                                       `json:"debug"`
	Prometheus                 prometheus                                  `json:"prometheus"`
	OperatorPrometheus         operatorPrometheus                          `json:"operatorPrometheus"`
	EnableXTSocketFallback     bool                                        `json:"enableXTSocketFallback"`
	InstallIptablesRules       bool                                        `json:"installIptablesRules"`
	Masquerade                 bool                                        `json:"masquerade"`
	AutoDirectNodeRoutes       bool                                        `json:"autoDirectNodeRoutes"`
	EndpointRoutes             endpointRoutes                              `json:"endpointRoutes"`
	Cni                        cni                                         `json:"cni"`
	Cluster                    cluster                                     `json:"cluster"`
	Tunnel                     ciliumv1alpha1.TunnelMode                   `json:"tunnel"`
	ContainerRuntime           containerRuntime                            `json:"containerRuntime"`
	Bpf                        bpf                                         `json:"bpf"`
	Encryption                 encryption                                  `json:"encryption"`
	KubeProxyReplacement       ciliumv1alpha1.KubeProxyReplacementMode     `json:"kubeProxyReplacement"`
	K8sServiceHost             string                                      `json:"k8sServiceHost"`
	K8sServicePort             int32                                       `json:"k8sServicePort"`
	HostServices               hostServices                                `json:"hostServices"`
	NodePort                   nodePort                                    `json:"nodePort"`
	ExternalIPs                externalIPs                                 `json:"externalIPs"`
	Flannel                    flannel                                     `json:"flannel"`
	Ipvlan                     ipvlan                                      `json:"ipvlan"`
	DatapathMode               string                                      `json:"datapathMode"`
	L7Proxy                    l7Proxy                                     `json:"l7Proxy"`
	Pprof                      pprof                                       `json:"pprof"`
	LogSystemLoad              bool                                        `json:"logSystemLoad"`
	Sockops                    sockops                                     `json:"sockops"`
	K8s                        k8s                                         `json:"k8s"`
	Eni                        bool                                        `json:"eni"`
	EgressMasqueradeInterfaces string                                      `json:"egressMasqueradeInterfaces"`
	Azure                      azure                                       `json:"azure"`
	CleanState                 bool                                        `json:"cleanState"`
	CleanBpfState              bool                                        `json:"cleanBpfState"`
	Nodeinit                   nodeinit                                    `json:"nodeinit"`
	Daemon                     daemon                                      `json:"daemon"`
	WellKnownIdentities        wellKnownIdentities                         `json:"wellKnownIdentities"`
	Tls                        tls                                         `json:"tls"`
	RemoteNodeIdentity         bool                                        `json:"remoteNodeIdentity"`
	SynchronizeK8sNodes        bool                                        `json:"synchronizeK8sNodes"`
	Psp                        psp                                         `json:"psp"`
	PolicyAuditMode            bool                                        `json:"policyAuditMode"`
	Hubble                     hubble                                      `json:"hubble"`
}

// pprof configuration for cilium
type pprof struct {
	Enabled bool `json:"enabled"`
}

// Container Runtime option for cilium
type containerRuntime struct {
	Integration string `json:"integration"`
}

// encryption related configuration for cilium cluster
type encryption struct {
	Enabled        bool   `json:"enabled"`
	KeyFile        string `json:"keyFile"`
	MountPath      string `json:"mountPath"`
	SecretName     string `json:"secretName"`
	NodeEncryption bool   `json:"nodeEncryption"`
	Interface      string `json:"interface"`
}

// etcd related configuration for cilium
type etcd struct {
	Enabled       bool   `json:"enabled"`
	Managed       bool   `json:"managed"`
	ClusterDomain string `json:"clusterDomain"`
	Ssl           bool   `json:"ssl"`
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

// installIptableRules configuration for cilium
type installIptableRules struct {
	Enabled bool `json:"enabled"`
}

// ipvlan configuration for cilium
type ipvlan struct {
	Enabled      bool   `json:"enabled"`
	MasterDevice string `json:"masterDevice"`
}

// l7Proxy configuration for cilium
type l7Proxy struct {
	Enabled bool `json:"enabled"`
}

// flannel configuration for cilium
type flannel struct {
	Enabled         bool   `json:"enabled"`
	MasterDevice    string `json:"masterDevice"`
	UninstallOnExit bool   `json:"uninstallOnExit"`
}

// autoDirectNodeRoutes configuration for cilium
type autoDirectNodeRoutes struct {
	Enabled bool `json:"enabled"`
}

// externalIPs configuration for cilium
type externalIPs struct {
	// externalIP enabled is used to define whether ExternalIP address is required or not.
	Enabled bool `json:"enabled"`
}

// Debug level option for cilium
type debug struct {
	// Debug Enabled is used to define whether Debug is required or not.
	Enabled bool `json:"enabled"`
}

// hubble enablement for cilium
type hubble struct {
	// hubble is used to define whether Hubble is required or not.
	ListenAddresses []string `json:"listenAddresses"`
	EventQueueSize  int32    `json:"eventQueueSize"`
	FlowBufferSize  int32    `json:"flowBufferSize"`
	MetricServer    string   `json:"metricServer"`
	Metrics         []string `json:"metrics"`
	UIEnabled       bool     `json:"uiEnabled"`
}

// bpf configuration for cilium
type bpf struct {
	// bpf tunning for cilium.
	WaitForMount       bool   `json:"waitForMount"`
	PreallocateMaps    bool   `json:"preallocateMaps"`
	CtTcpMax           int32  `json:"ctTcpMax"`
	CtAnyMax           int32  `json:"ctAnyMax"`
	MonitorAggregation string `json:"monitorAggregation"`
	MonitorInterval    string `json:"monitorInterval"`
	MonitorFlags       string `json:"monitorFlags"`
}

// nodePort enablement for cilium
type nodePort struct {
	// Nodeport Enabled is used to define whether Nodeport is required or not.
	Enabled bool   `json:"enabled"`
	Mode    string `json:"mode"`
}

// sockops configuration for cilium
type sockops struct {
	// sockops Enabled is used to define whether Sockops is required or not.
	Enabled bool `json:"enabled"`
}

// hostServices configuration for cilium
type hostServices struct {
	// hostServices is used to define whether IPv4 address is required or not.
	Enabled   bool   `json:"enabled"`
	Protocols string `json:"protocols"`
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
	TofqdnsPreCache string `json:"tofqdnsPreCache"`
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

// cluster configuration for cilium
type cluster struct {
	// ClusterName is used to define the name of the cluster.
	Name string `json:"name"`
}

// wellKnownIdentities configuration for cilium
type wellKnownIdentities struct {
	// wellKnownIdentities enabled for cilium.
	Enabled bool `json:"enabled"`
}

// tls configuration for cilium
type tls struct {
	// tls is required or not.
	SecretsBackend string `json:"secretsBackend"`
}

// psp configuration for cilium
type psp struct {
	// psp is required or not
	Enabled bool `json:"enabled"`
}

// k8sIPv4PodCIDR configuration for cilium
type k8s struct {
	// k8s is required or not
	RequireIPv4PodCIDR bool `json:"requireIPv4PodCIDR"`
}

// nodeInit configuration for cilium
type nodeinit struct {
	// nodeInit is required or not
	Enabled       bool   `json:"enabled"`
	BootstrapFile string `json:"bootstrapFile"`
}

// azure configuration for cilium
type azure struct {
	Enabled        bool   `json:"enabled"`
	ResourceGroup  string `json:"resourceGroup"`
	SubscriptionID string `json:"subscriptionID"`
	TenantID       string `json:"tenantID"`
	ClientID       string `json:"clientID"`
	ClientSecret   string `json:"clientSecret"`
}

// cni configuration for cilium
type cni struct {
	Install              bool   `json:"install"`
	ChainingMode         string `json:"chainingMode"`
	CustomConf           bool   `json:"customConf"`
	ConfPath             string `json:"confPath"`
	BinPath              string `json:"binPath"`
	ConfigMapKey         string `json:"configMapKey"`
	ConfFileMountPath    string `json:"confFileMountPath"`
	HostConfDirMountPath string `json:"hostConfDirMountPath"`
	ConfigMap            string `json:"configMap"`
}

// endpointRoutes configuration for cilium
type endpointRoutes struct {
	// endpointRoutes is required or not
	Enabled bool `json:"enabled"`
}

// agent configuration for cilium
type agent struct {
	// agent is required or not
	Enabled              bool `json:"enabled"`
	SleepAfterInit       bool `json:"sleepAfterInit"`
	KeepDeprecatedLabels bool `json:"keepDeprecatedLabels"`
}

// config enable option for cilium
type config struct {
	// config is required or not
	Enabled bool `json:"enabled"`
}

// daemon option for cilium
type daemon struct {
	// daemon is required or not
	RunPath string `json:"runPath"`
}

var defaultCiliumConfig = ciliumConfig{
	Agent: agent{
		Enabled:              true,
		SleepAfterInit:       false,
		KeepDeprecatedLabels: false,
	},
	Config: config{
		Enabled: true,
	},
	Operator: operator{
		Enabled: true,
	},
	Preflight: preflight{
		Enabled:         false,
		TofqdnsPreCache: "",
	},
	Global: global{
		Registry:   "docker.io/cilium",
		Tag:        "latest",
		PullPolicy: "Always",
		Etcd: etcd{
			Enabled:       false,
			Managed:       false,
			ClusterDomain: "cluster.local",
			Endpoints:     "",
			Ssl:           false,
		},
		IdentityAllocationMode: "crd",
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
			Enabled: false,
			Port:    9090,
			ServiceMonitor: serviceMonitor{
				Enabled: false,
			},
		},
		OperatorPrometheus: operatorPrometheus{
			Enabled: false,
			Port:    6942,
		},
		EnableXTSocketFallback: true,
		InstallIptablesRules:   true,
		Masquerade:             true,
		AutoDirectNodeRoutes:   false,
		EndpointRoutes: endpointRoutes{
			Enabled: false,
		},
		Cni: cni{
			Install:              true,
			ChainingMode:         "none",
			CustomConf:           false,
			ConfPath:             "/etc/cni/net.d",
			BinPath:              "/opt/cni/bin",
			ConfigMapKey:         "cni-config",
			ConfFileMountPath:    "/tmp/cni-configuration",
			HostConfDirMountPath: "/host/etc/cni/net.d",
			ConfigMap:            "cni-config",
		},
		Cluster: cluster{
			Name: "default",
		},
		Tunnel: "vxlan",
		ContainerRuntime: containerRuntime{
			Integration: "none",
		},
		Bpf: bpf{
			WaitForMount:       false,
			PreallocateMaps:    false,
			CtTcpMax:           524288,
			CtAnyMax:           262144,
			MonitorAggregation: "medium",
			MonitorInterval:    "5s",
			MonitorFlags:       "all",
		},
		Encryption: encryption{
			Enabled:        false,
			KeyFile:        "keys",
			MountPath:      "/etc/ipsec",
			SecretName:     "cilium-ipsec-keys",
			NodeEncryption: false,
			Interface:      "eth0",
		},
		KubeProxyReplacement: "probe",
		K8sServiceHost:       "",
		K8sServicePort:       2030,
		HostServices: hostServices{
			Enabled:   false,
			Protocols: "tcp,udp",
		},
		NodePort: nodePort{
			Enabled: false,
			Mode:    "hybrid",
		},
		ExternalIPs: externalIPs{
			Enabled: false,
		},
		Flannel: flannel{
			Enabled:         false,
			MasterDevice:    "cni0",
			UninstallOnExit: false,
		},
		Ipvlan: ipvlan{
			Enabled:      false,
			MasterDevice: "eth0",
		},
		DatapathMode: "ipvlan",
		L7Proxy: l7Proxy{
			Enabled: false,
		},
		Pprof: pprof{
			Enabled: false,
		},
		LogSystemLoad: false,
		Sockops: sockops{
			Enabled: false,
		},
		K8s: k8s{
			RequireIPv4PodCIDR: false,
		},
		Eni:                        false,
		EgressMasqueradeInterfaces: "eth0",
		Azure: azure{
			Enabled: false,
		},
		CleanState:    false,
		CleanBpfState: false,
		Nodeinit: nodeinit{
			Enabled:       false,
			BootstrapFile: "/tmp/cilium-bootstrap-time",
		},
		Daemon: daemon{
			RunPath: "/var/run/cilium",
		},
		WellKnownIdentities: wellKnownIdentities{
			Enabled: false,
		},
		Tls: tls{
			SecretsBackend: "local",
		},
		RemoteNodeIdentity:  true,
		SynchronizeK8sNodes: true,
		Psp: psp{
			Enabled: false,
		},
		PolicyAuditMode: false,
		Hubble: hubble{
			ListenAddresses: nil,
			EventQueueSize:  200,
			FlowBufferSize:  300,
			MetricServer:    "",
			Metrics:         nil,
			UIEnabled:       true,
		},
	},
	NodeInit: nonGlobalNodeinit{
		RestartPods:              false,
		ReconfigureKubelet:       false,
		RemoveCbrBridge:          false,
		RevertReconfigureKubelet: false,
	},
}

func newCiliumConfig() ciliumConfig {
	return defaultCiliumConfig
}

func (c *ciliumConfig) toMap() (map[string]interface{}, error) {
	bytes, err := json.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("could not marshal cilium config: %v", err)
	}
	var configMap map[string]interface{}
	err = json.Unmarshal(bytes, &configMap)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal cilium config: %v", err)
	}
	return configMap, nil
}

// ComputeCiliumChartValues computes the values for the cilium chart.
func ComputeCiliumChartValues(network *extensionsv1alpha1.Network, config *ciliumv1alpha1.NetworkConfig) (map[string]interface{}, error) {
	typedConfig, err := generateChartValues(config)
	if err != nil {
		return nil, fmt.Errorf("error when generating cilium config: %v", err)
	}
	ciliumConfig, err := typedConfig.toMap()
	if err != nil {
		return nil, fmt.Errorf("could not convert cilium config: %v", err)
	}
	ciliumChartValues := map[string]interface{}{
		"images": map[string]interface{}{
			cilium.CiliumImageName:             imagevector.CiliumImage(),
			cilium.CiliumOperatorImageName:     imagevector.CiliumOperatorImage(),
			cilium.CiliumDockerPluginImageName: imagevector.CiliumDockerPluginImage(),
			cilium.CiliumEtcdOperatorImageName: imagevector.CiliumEtcdOperatorImage(),
			cilium.HubbleImageName:             imagevector.CiliumHubbleImage(),
			cilium.HubbleUIImageName:           imagevector.CiliumHubbleUIImage(),
			cilium.IstioPilotImageName:         imagevector.IstioPilotImage(),
			cilium.IstioProxyImageName:         imagevector.IstioProxyImage(),
		},
		"global": map[string]string{
			"podCIDR": network.Spec.PodCIDR,
		},
		"config": ciliumConfig,
	}
	return ciliumChartValues, nil
}

func generateChartValues(config *ciliumv1alpha1.NetworkConfig) (*ciliumConfig, error) {
	c := newCiliumConfig()
	// IPVlan based setting for different modes of operation
	if (config.IPvlan.IPVlanMode == "PureL3") || (config.IPvlan.IPVlanMode == "L3S") || (config.IPvlan.IPVlanMode == "L3E") {
		c.Global.Ipvlan.Enabled = config.IPvlan.IPvlanEnabled
		// Required v4.12 or more recent Linux Kernel
		if config.IPvlan.IPVlanMode == "PureL3" {
			c.Global.DatapathMode = "ipvlan"
			c.Global.Ipvlan.MasterDevice = config.IPvlan.MasterDevice
			c.Global.Tunnel = "disabled"
		}
		// Required v4.12 or more recent Linux Kernel and
		// https://git.kernel.org/pub/scm/linux/kernel/git/netdev/net.git/commit/?id=d5256083f62e2720f75bb3c5a928a0afe47d6bc3
		// patch available in stable kernels v4.9.155, v4.14.98, v4.19.20 and v4.20.6 and higher
		if config.IPvlan.IPVlanMode == "L3S" {
			c.Global.DatapathMode = "ipvlan"
			c.Global.Ipvlan.MasterDevice = config.IPvlan.MasterDevice
			c.Global.Tunnel = "disabled"
			c.Global.InstallIptablesRules = false
			c.Global.L7Proxy.Enabled = false
			c.Global.AutoDirectNodeRoutes = true
		}
		// L3 mode with more efficient BPF-based masquerade instead of iptables-based
		if config.IPvlan.IPVlanMode == "L3E" {
			c.Global.DatapathMode = "ipvlan"
			c.Global.Ipvlan.MasterDevice = config.IPvlan.MasterDevice
			c.Global.Tunnel = "disabled"
			c.Global.Masquerade = true
			c.Global.InstallIptablesRules = false
			c.Global.AutoDirectNodeRoutes = true
		}
	}
	// Settings for Kube-Proxy disabled and using the HostService option
	// Also need to configrue KubeProxy
	if !(config.KubeProxy.KubeProxyEnabled) {
		// Strict mode
		if config.KubeProxyReplacementMode != nil {
			if *config.KubeProxyReplacementMode == "strict" {
				// Setting for kube-proxy free Kubernetes setup where cilium replaces all kube-proxy functionality
				c.Global.KubeProxyReplacement = *config.KubeProxyReplacementMode
				c.Global.K8sServiceHost = "HOST IP"
				c.Global.K8sServicePort = 2030
				c.Global.NodePort.Enabled = true
				c.Global.ExternalIPs.Enabled = true
				c.Global.HostServices.Enabled = config.HostServices.HostServicesEnabled
				if config.HostServices.Protocols == "tcp" {
					c.Global.HostServices.Protocols = config.HostServices.Protocols
				}
				if config.HostServices.Protocols == "udp" {
					c.Global.HostServices.Protocols = config.HostServices.Protocols
				}
				c.Global.HostServices.Protocols = "tcp,udp"
			}
			// Hybrid mode where it automatically determines which feature should be used or not used
			if *config.KubeProxyReplacementMode == "probe" {
				// Setting for kube-proxy in hybrid mode
				c.Global.KubeProxyReplacement = *config.KubeProxyReplacementMode
			}
			// Hybrid mode but manual setting is required on which features should be used or not
			// With just Nodeport and ExternalIPs
			if *config.KubeProxyReplacementMode == "partial" {
				// Setting for kube-proxy in hybrid mode but Manual configuration.
				c.Global.KubeProxyReplacement = *config.KubeProxyReplacementMode
				c.Global.NodePort.Enabled = true
				c.Global.ExternalIPs.Enabled = true
			}
		}

	} else {
		c.Global.KubeProxyReplacement = "disabled"
	}
	// Transparent Encryption configuration for cilium
	// Node to Node encryption is disabled or enabled can be configured by user.
	// Make sure we pass the Interface if configured to do so
	if config.Encryption.EncryptionEnabled {
		c.Global.Encryption.NodeEncryption = config.Encryption.NodeEncryption
		c.Global.Encryption.Enabled = config.Encryption.EncryptionEnabled
		c.Global.Encryption.Interface = "eth0"
	}
	// NodePort with DSR Mode
	// TODO: Revisit to provide nodePort.range and nodePort.device
	if config.TunnelMode == "disabled" && config.Nodeport.NodeportMode == "dsr" {
		// DSR Mode for Nodeport will work in this scenario when tunnel is disabled.
		c.Global.Tunnel = config.TunnelMode
		c.Global.AutoDirectNodeRoutes = true
		c.Global.KubeProxyReplacement = "strict"
		c.Global.NodePort.Mode = config.Nodeport.NodeportMode
		c.Global.K8sServiceHost = "HOST IP"
		c.Global.K8sServicePort = 2330
	}
	// Setup cilium in AWS ENI Mode
	if config.Eni.EniEnabled {
		c.Global.Eni = config.Eni.EniEnabled
		c.Global.EgressMasqueradeInterfaces = config.Eni.EgressMasqueradeInterfaces
		c.Global.Tunnel = "disabled"
		c.Global.Nodeinit.Enabled = true
	}
	// If Hubble enabled
	// TODO: Make sure you add the 'ui.enable' in here as well.
	if config.Hubble.HubbleEnabled {
		c.Global.Hubble.Metrics = config.Hubble.Metrics
		c.Global.Hubble.ListenAddresses = config.Hubble.ListenAddresses
		c.Global.Hubble.MetricServer = config.Hubble.MetricServer
		c.Global.Hubble.EventQueueSize = config.Hubble.EventQueueSize
		c.Global.Hubble.FlowBufferSize = config.Hubble.FlowBufferSize
		c.Global.Hubble.UIEnabled = config.Hubble.UI
	}
	// Flannel support is ther from Flannel 0.10.0 and Kubernetes version >= 1.9
	if config.Flannel.FlannelEnabled {
		c.Global.Flannel.Enabled = config.Flannel.FlannelEnabled
		c.Global.Flannel.MasterDevice = config.Flannel.MasterDevice
		c.Global.Flannel.UninstallOnExit = config.Flannel.UninstallOnExit
	}
	// Azure AKS settings for starting the cilium
	if config.Azure.AzureEnabled {
		c.Global.Azure.Enabled = config.Azure.AzureEnabled
		c.Global.Cni.ChainingMode = "generic-veth"
		c.Global.Cni.CustomConf = config.CNI.CniCustomConfigEnabled
		c.Global.Nodeinit.Enabled = config.NodeInit.NodeInitEnabled
		c.Global.Cni.ConfigMap = config.CNI.CniConfigMap
		c.Global.Tunnel = "disabled"
		c.Global.Masquerade = false
	}
	// AWS EKS settings
	if config.Eni.EniEnabled {
		c.Global.Eni = config.Eni.EniEnabled
		c.Global.EgressMasqueradeInterfaces = config.Eni.EgressMasqueradeInterfaces
		c.Global.Tunnel = "disabled"
		c.Global.Nodeinit.Enabled = true
	}
	// Google GKE
	if config.GoogleGKE.GoogleGkeEnabled {
		c.Global.Nodeinit.Enabled = config.GoogleGKE.GoogleGkeEnabled
		c.NodeInit.ReconfigureKubelet = config.NodeInit.NonGlobalReconfigureKubelet
		c.NodeInit.RemoveCbrBridge = config.NodeInit.NonGlobalRemoveCbrBridge
		c.Global.Cni.BinPath = config.CNI.CniBinPath
	}

	return &c, nil
}
