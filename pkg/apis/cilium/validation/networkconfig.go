// SPDX-FileCopyrightText: 2025 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package validation

import (
	"fmt"
	"net"
	"regexp"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"

	apiscilium "github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium"
)

const (
	deviceFormat    = "[^/\\s]{1,15}"
	deviceMaxLength = 15
)

var deviceRegexp = regexp.MustCompile("^" + deviceFormat + "$")

// ValidateNetworkConfig validates the network config.
func ValidateNetworkConfig(networkConfig *apiscilium.NetworkConfig, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	allErrs = append(allErrs, ValidateNetworkConfigKubeProxy(networkConfig.KubeProxy, fldPath.Child("kubeproxy"))...)

	allowedTunnelModes := sets.New[apiscilium.TunnelMode](apiscilium.VXLan, apiscilium.Geneve, apiscilium.Disabled)
	if networkConfig.TunnelMode != nil && !allowedTunnelModes.Has(*networkConfig.TunnelMode) {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("tunnel"), *networkConfig.TunnelMode, fmt.Sprintf("unsupported value %q for tunnel, supported values are [%q, %q, %q]", *networkConfig.TunnelMode, apiscilium.VXLan, apiscilium.Geneve, apiscilium.Disabled)))
	}

	allowedStores := sets.New[apiscilium.Store](apiscilium.Kubernetes)
	if networkConfig.Store != nil && !allowedStores.Has(*networkConfig.Store) {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("store"), *networkConfig.Store, fmt.Sprintf("unsupported value %q for store, supported values are [%q]", *networkConfig.Store, apiscilium.Kubernetes)))
	}

	// It is hard to put valid bounds on MTU, but negative values are definitively invalid.
	if networkConfig.MTU != nil && *networkConfig.MTU < 0 {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("mtu"), *networkConfig.MTU, "mtu must be a positive integer"))
	}

	for i, device := range networkConfig.Devices {
		allErrs = append(allErrs, ValidateDevice(device, fldPath.Child("devices").Index(i))...)
	}

	if networkConfig.DirectRoutingDevice != nil {
		allErrs = append(allErrs, ValidateDevice(*networkConfig.DirectRoutingDevice, fldPath.Child("directRoutingDevice"))...)
	}

	allowedLoadBalancingModes := sets.New[apiscilium.LoadBalancingMode](apiscilium.SNAT, apiscilium.DSR, apiscilium.Hybrid)
	if networkConfig.LoadBalancingMode != nil && !allowedLoadBalancingModes.Has(*networkConfig.LoadBalancingMode) {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("loadBalancingMode"), *networkConfig.LoadBalancingMode, fmt.Sprintf("unsupported value %q for loadBalancingMode, supported values are [%q, %q, %q]", *networkConfig.LoadBalancingMode, apiscilium.SNAT, apiscilium.DSR, apiscilium.Hybrid)))
	}

	if networkConfig.Encryption != nil {
		allErrs = append(allErrs, ValidateEncryption(networkConfig.Encryption, fldPath.Child("encryption"))...)
	}

	return allErrs
}

// ValidateNetworkConfigKubeProxy validates the kube-proxy configuration in the network config.
func ValidateNetworkConfigKubeProxy(kubeProxy *apiscilium.KubeProxy, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	if kubeProxy == nil {
		return allErrs
	}

	// ServiceHost can be an IP address or a DNS name.
	if kubeProxy.ServiceHost != nil {
		if net.ParseIP(*kubeProxy.ServiceHost) == nil {
			for _, err := range validation.IsDNS1123Subdomain(*kubeProxy.ServiceHost) {
				allErrs = append(allErrs, field.Invalid(fldPath.Child("k8sServiceHost"), *kubeProxy.ServiceHost, fmt.Sprintf("serviceHost is neither a valid IP address nor a valid domain name: %q", err)))
			}
		}
	}

	if kubeProxy.ServicePort != nil {
		if *kubeProxy.ServicePort < 1 || *kubeProxy.ServicePort > 65535 {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("k8sServicePort"), *kubeProxy.ServicePort, fmt.Sprintf("servicePort must be between 1 and 65535, got %d", *kubeProxy.ServicePort)))
		}
	}

	return allErrs
}

// ValidateDevice validates a linux device name.
func ValidateDevice(device string, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	if len(device) == 0 {
		allErrs = append(allErrs, field.Invalid(fldPath, device, "device name cannot be empty"))
	} else if len(device) > deviceMaxLength {
		allErrs = append(allErrs, field.Invalid(fldPath, device, fmt.Sprintf("device name cannot be longer than %d characters", deviceMaxLength)))
	} else if !deviceRegexp.MatchString(device) {
		allErrs = append(allErrs, field.Invalid(fldPath, device, fmt.Sprintf("device name must match the pattern %q", deviceFormat)))
	}

	return allErrs
}

func ValidateEncryption(enc *apiscilium.Encryption, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	if !enc.Enabled {
		return allErrs
	}
	if enc.StrictMode {
		if enc.Mode != apiscilium.EncryptionModeWireguard {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("mode"), enc.Mode, "strict mode can only be used with wireguard as encyption mode"))
		}
	}
	return allErrs
}
