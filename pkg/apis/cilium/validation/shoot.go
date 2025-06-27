// SPDX-FileCopyrightText: 2025 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package validation

import (
	"fmt"
	"net"

	"github.com/gardener/gardener/pkg/apis/core"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/gardener/gardener-extension-networking-cilium/pkg/cilium"
)

// ValidateNetworking validates the network settings of a Shoot.
func ValidateNetworking(networking *core.Networking, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	if networking.Type == nil {
		allErrs = append(allErrs, field.Required(fldPath.Child("type"), "networking type is required"))
	} else if *networking.Type != cilium.Type {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("type"), networking.Type, fmt.Sprintf("networking type %q is not supported", *networking.Type)))
	}

	for _, network := range []struct {
		name string
		cidr *string
	}{
		{name: "pods", cidr: networking.Pods},
		{name: "nodes", cidr: networking.Nodes},
		{name: "services", cidr: networking.Services},
	} {
		if network.cidr != nil {
			if _, _, err := net.ParseCIDR(*network.cidr); err != nil {
				allErrs = append(allErrs, field.Invalid(fldPath.Child(network.name), *network.cidr, fmt.Sprintf("%s is not a valid CIDR: %q", network.name, *network.cidr)))
			}
		}
	}

	allowedIPFamilies := sets.New[core.IPFamily](core.IPFamilyIPv4, core.IPFamilyIPv6)
	if !allowedIPFamilies.HasAll(networking.IPFamilies...) {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("ipFamilies"), networking.IPFamilies, fmt.Sprintf("ipFamilies (%q) must be a subset of [%q, %q]", networking.IPFamilies, core.IPFamilyIPv4, core.IPFamilyIPv6)))
	}
	if len(networking.IPFamilies) > 2 {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("ipFamilies"), networking.IPFamilies, fmt.Sprintf("ipFamilies (%q) cannot have more than two entries", networking.IPFamilies)))
	} else if len(networking.IPFamilies) == 2 && networking.IPFamilies[0] == networking.IPFamilies[1] {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("ipFamilies"), networking.IPFamilies, fmt.Sprintf("ipFamilies (%q) cannot have duplicate entries", networking.IPFamilies)))
	}

	return allErrs
}
