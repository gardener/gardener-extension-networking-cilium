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
	Enabled bool  `json:"prometheusEnabled"`
	Port    int32 `json:"port"`
}

// OperatorPrometheus configuration for cilium
type OperatorPrometheus struct {
	Enabled bool  `json:"operatorprometheusEnabled"`
	Port    int32 `json:"port"`
}

// ExternalIPs configuration for cilium
type ExternalIP struct {
	// ExternalIPenabled is used to define whether ExternalIP address is required or not.
	Enabled bool `json:"externalipEnabled"`
}

// Hubble enablement for cilium
type Hubble struct {
	Enabled bool     `json:"enabled"`
	UI      bool     `json:"ui"`
	Metrics []string `json:"metrics"`
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
	// Enabled specifies whether kubeproxy is disabled.
	// +optional
	Enabled *bool `json:"disabled"`
	// ServiceHost specify the controlplane node IP Address.
	// +optional
	ServiceHost *string `json:"k8sServiceHost,omitempty"`
	// ServicePort specify the kube-apiserver port number.
	// +optional
	ServicePort *int32 `json:"k8sServicePort,omitempty"`
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
	// Prometheus configuration
	// +optional
	Prometheus *Prometheus `json:"prometheus,omitempty"`
	// OperatorPrometheus configuration
	// +optional
	OperatorPrometheus *OperatorPrometheus `json:"operatorprometheus,omitempty"`
	// PSPEnabled configuration
	// +optional
	PSPEnabled *bool `json:"psp,omitempty"`
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
}
