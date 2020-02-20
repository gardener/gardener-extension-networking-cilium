// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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
	VXLan TunnelMode = "vxlan"
	Geneve TunnelMode = "geneve"
	None TunnelMode = "disabled"
)

type IPAMOptionMode string

const (
	Crd IPAMOptionMode = "crd"
	Eni IPAMOptionMode = "eni"
)

// IPv4 configuration for cilium
type IPv4 struct {
	// Enabled is used to define whether IPv4 address is required or not.
	V4_enabled bool `json:"v4_enabled"`
}

// IPv6 configuration for cilium
type IPv6 struct {
        // Enabled is used to define whether IPv4 address is required or not.
        V6_enabled bool `json:"v6_enabled"`
}

// NodeSpec is the configuration specific to a node
type NodeSpec struct {
        // IPAM is the address management specification. This section can be
        // populated by a user or it can be automatically populated by an IPAM
        // operator
        //
        // +optional
	IPv4 IPv4 `json:"ipv4,omitempty"`
        IPv6 IPv6 `json:"ipv6,omitempty"`
        IPAMOptionMode IPAMOptionMode `json:"default,omitempty"`
        TunnelMode TunnelMode `json:"vxlan,omitempty"`

        IPAM IPAMSpec `json:"ipam,omitempty"`
}

// IPAMSpec is the IPAM specification of the node
type IPAMSpec struct {
        // Pool is the list of IPs available to the node for allocation. When
        // an IP is used, the IP will remain on this list but will be added to
        // Status.IPAM.InUse
        //
        // +optional
        Pool map[string]AllocationIP `json:"pool,omitempty"`
}

// AllocationIP is an IP available for allocation or already allocated
type AllocationIP struct {
        // Owner is the owner of the IP, this field is set if the IP has been
        // allocated. It will be set to the pod name or another identifier
        // representing the usage of the IP
        //
        // The owner field is left blank for an entry in Spec.IPAM.Pool
        // and filled out as the IP is used and also added to
        // Status.IPAM.InUse.
        //
        // +optional
        Owner string `json:"owner,omitempty"`

        // Resource is set for both available and allocated IPs, it represents
        // what resource the IP is associated with, e.g. in combination with
        // AWS ENI, this will refer to the ID of the ENI
        //
        // +optional
        Resource string `json:"resource,omitempty"`
}


// NodeStatus is the status of a node
type NodeStatus struct {
        // [...]

        // IPAM is the IPAM status of the node
        //
        // +optional
        IPAM IPAMStatus `json:"ipam,omitempty"`
}

// IPAMStatus is the IPAM status of a node
type IPAMStatus struct {
        // InUse lists all IPs out of Spec.IPAM.Pool which have been
        // allocated and are in use.
        //
        // +optional
        InUse map[string]AllocationIP `json:"used,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkConfig is a struct representing the configmap for the cilium
// networking plugin
type CiliumNetwork struct {
	metav1.TypeMeta   `json:",inline"`
        metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec is the specification of the node
        Spec NodeSpec `json:"spec"`

        // Status it the status of the node
        Status NodeStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkStatus contains information about created Network resources.
type CiliumNetworkStatus struct {
	metav1.TypeMeta `json:",inline"`
}
