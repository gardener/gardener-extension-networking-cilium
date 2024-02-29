// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"fmt"

	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"

	"github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium/install"
	ciliumv1alpha1 "github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium/v1alpha1"
)

var (
	// Scheme is a scheme with the types relevant for Network actuators.
	Scheme *runtime.Scheme

	decoder runtime.Decoder
)

func init() {
	Scheme = runtime.NewScheme()
	utilruntime.Must(install.AddToScheme(Scheme))

	decoder = serializer.NewCodecFactory(Scheme, serializer.EnableStrict).UniversalDecoder()
}

// CiliumNetworkConfigFromNetworkResource extracts the NetworkConfig from the
// ProviderConfig section of the given Network resource.
func CiliumNetworkConfigFromNetworkResource(network *extensionsv1alpha1.Network) (*ciliumv1alpha1.NetworkConfig, error) {
	config := &ciliumv1alpha1.NetworkConfig{}
	if network.Spec.ProviderConfig != nil && network.Spec.ProviderConfig.Raw != nil {
		if _, _, err := decoder.Decode(network.Spec.ProviderConfig.Raw, nil, config); err != nil {
			return nil, err
		}
		return config, nil
	}
	return nil, fmt.Errorf("provider config is not set on the network resource")
}
