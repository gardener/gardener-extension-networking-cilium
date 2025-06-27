// SPDX-FileCopyrightText: 2025 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"fmt"

	"k8s.io/apimachinery/pkg/util/validation/field"

	apiscilium "github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium"
	ciliumv1alpha1 "github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium/v1alpha1"
	ciliumvalidation "github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium/validation"
)

// ValidateNetworkConfig validates the given network configuration.
func ValidateNetworkConfig(networkConfig *ciliumv1alpha1.NetworkConfig) error {
	internalNetworkConfig := &apiscilium.NetworkConfig{}
	if err := ciliumv1alpha1.Convert_v1alpha1_NetworkConfig_To_cilium_NetworkConfig(networkConfig, internalNetworkConfig, nil); err != nil {
		return fmt.Errorf("could not convert network config: %w", err)
	}

	if errList := ciliumvalidation.ValidateNetworkConfig(internalNetworkConfig, field.NewPath("spec", "providerConfig")); len(errList) != 0 {
		return fmt.Errorf("invalid network config: %w", errList.ToAggregate())
	}

	return nil
}
