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

type IdentityAllocationMode string

const (
	CRD     IdentityAllocationMode = "crd"
	KVStore IdentityAllocationMode = "kvstore"
)

type TunnelMode string

const (
	VXLan    TunnelMode = "vxlan"
	Geneve   TunnelMode = "geneve"
	Disabled TunnelMode = "disabled"
)

type KubeProxyReplacementMode string

const (
	Strict  KubeProxyReplacementMode = "strict"
	Probe   KubeProxyReplacementMode = "probe"
	Partial KubeProxyReplacementMode = "partial"
)

type NodePortMode string

const (
	Hybird NodePortMode = "hybrid"
)

type Store string

const (
	Kubernetes Store = "kubernetes"
	ETCD       Store = "etcd"
)

// Prometheus configuration for cilium
type Prometheus struct {
	Enabled bool
	Port    int32
}

// OperatorPrometheus configuration for cilium
type OperatorPrometheus struct {
	Enabled bool
	Port    int32
}

// InstallIPTableRules configuration for cilium
type InstallIPTableRules struct {
	Enabled bool
}

// ExternalIPs configuration for cilium
type ExternalIP struct {
	// ExternalIPenabled is used to define whether ExternalIP address is required or not.
	Enabled bool
}

// Hubble enablement for cilium
type Hubble struct {
	// Enabled defines whether hubble will be enabled for the cluster.
	Enabled bool
	// UI defines whether Hubble UI is enabled or not.
	UI bool
	// Metrics defined what metrics will be reported by hubble
	Metrics []string
}

// Nodeport enablement for cilium
type Nodeport struct {
	// Enabled is used to define whether Nodeport is required or not.
	Enabled bool
	// Mode is the mode of NodePort feature
	Mode NodePortMode
}

// KubeProxy configuration for cilium
type KubeProxy struct {
	// Enabled specifies whether kubeproxy is disabled.
	Enabled *bool
	// ServiceHost specify the controlplane node IP Address.
	ServiceHost *string
	// ServicePort specify the kube-apiserver port number.
	ServicePort *int32
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkConfig is a struct representing the configmap for the cilium
// networking plugin
type NetworkConfig struct {
	metav1.TypeMeta
	// Debug configuration to be enabled or not
	Debug *bool
	// Prometheus configuration
	Prometheus *Prometheus
	// PSPEnabled configuration
	PSPEnabled *bool
	// OperatorPrometheus configuration
	OperatorPrometheus *OperatorPrometheus
	// KubeProxy configuration to be enabled or not
	KubeProxy *KubeProxy
	// Hubble configuration to be enabled or not
	Hubble *Hubble
	// TunnelMode configuration, it should be 'vxlan', 'geneve' or 'disabled'
	TunnelMode *TunnelMode
	// Store can be either Kubernetes or etcd.
	Store *Store
}
