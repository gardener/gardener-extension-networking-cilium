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
	"k8s.io/utils/ptr"

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

	if err := validateSupported(fldPath.Child("tunnel"), networkConfig.TunnelMode, sets.New[apiscilium.TunnelMode](apiscilium.VXLan, apiscilium.Geneve, apiscilium.Disabled)); err != nil {
		allErrs = append(allErrs, err)
	}
	if err := validateSupported(fldPath.Child("store"), networkConfig.Store, sets.New[apiscilium.Store](apiscilium.Kubernetes)); err != nil {
		allErrs = append(allErrs, err)
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

	if networkConfig.LoadBalancer != nil {
		allErrs = append(allErrs, ValidateLoadBalancer(networkConfig.LoadBalancer, networkConfig.TunnelMode, ptr.Deref(networkConfig.Overlay, apiscilium.Overlay{}).Enabled, fldPath.Child("loadBalancer"))...)
	}

	if err := validateLoadBalancingMode(networkConfig.LoadBalancingMode, fldPath.Child("loadBalancingMode")); err != nil {
		allErrs = append(allErrs, err)
	}

	if networkConfig.Encryption != nil {
		allErrs = append(allErrs, ValidateEncryption(networkConfig.Encryption, fldPath.Child("encryption"))...)
	}

	return allErrs
}

func ValidateLoadBalancer(lb *apiscilium.LoadBalancer, tunnelMode *apiscilium.TunnelMode, overlayEnabled bool, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	if err := validateLoadBalancingMode(lb.Mode, fldPath.Child("mode")); err != nil {
		allErrs = append(allErrs, err)
	}
	if err := validateSupported(fldPath.Child("acceleration"), lb.Acceleration, sets.New(apiscilium.AccelerationBestEffort, apiscilium.AccelerationNative, apiscilium.AccelerationDisabled)); err != nil {
		allErrs = append(allErrs, err)
	}
	if err := validateSupported(fldPath.Child("algorithm"), lb.Algorithm, sets.New(apiscilium.LoadBalancerAlgorithmMaglev, apiscilium.LoadBalancerAlgorithmRandom)); err != nil {
		allErrs = append(allErrs, err)
	}

	if lb.DSRDispatch != nil {
		fldPath := fldPath.Child("dsrDispatch")
		if err := validateSupported(fldPath, lb.DSRDispatch, sets.New(apiscilium.DSRDispatchGeneve, apiscilium.DSRDispatchIPIP, apiscilium.DSRDispatchIPOption)); err != nil {
			allErrs = append(allErrs, err)
		}

		if overlayEnabled && tunnelMode != nil && *lb.DSRDispatch == apiscilium.DSRDispatchGeneve && *tunnelMode != apiscilium.Geneve {
			allErrs = append(allErrs, field.Invalid(fldPath, *lb.DSRDispatch, fmt.Sprintf("dsrDispatch geneve can't be used with tunnelMode %s", *tunnelMode)))
		}
	}

	return allErrs
}

func validateLoadBalancingMode(mode *apiscilium.LoadBalancingMode, fldPath *field.Path) *field.Error {
	return validateSupported(fldPath, mode, sets.New(apiscilium.SNAT, apiscilium.DSR, apiscilium.Hybrid))
}

func validateSupported[T ~string](fldPath *field.Path, val *T, set sets.Set[T]) *field.Error {
	if val != nil && !set.Has(*val) {
		return field.Invalid(fldPath, *val, fmt.Sprintf("unsupported value %v, supported values are %q", *val, set.UnsortedList()))
	}
	return nil
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
