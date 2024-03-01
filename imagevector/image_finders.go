// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package imagevector

import (
	"github.com/gardener/gardener/pkg/utils/imagevector"
	"k8s.io/apimachinery/pkg/util/runtime"

	"github.com/gardener/gardener-extension-networking-cilium/pkg/cilium"
)

func findImage(name string, opts ...imagevector.FindOptionFunc) string {
	image, err := imageVector.FindImage(name, opts...)
	runtime.Must(err)
	return image.String()
}

// CiliumAgentImage returns the Cilium Image.
func CiliumAgentImage() string {
	return findImage(cilium.CiliumAgentImageName)
}

// CiliumOperatorImage returns the Cilium Operator image.
func CiliumOperatorImage() string {
	return findImage(cilium.CiliumOperatorImageName)
}

// CiliumHubbleRelayImage returns the Cilium Hubble image.
func CiliumHubbleRelayImage() string {
	return findImage(cilium.HubbleRelayImageName)
}

// CiliumHubbleUIImage returns the Cilium Hubble UI image.
func CiliumHubbleUIImage() string {
	return findImage(cilium.HubbleUIImageName)
}

// CiliumHubbleUIBackendImage returns the Cilium Hubble UI Backend image.
func CiliumHubbleUIBackendImage() string {
	return findImage(cilium.HubbleUIBackendImageName)
}

// CiliumCertGenImage returns the Cilium Cert Gen image.
func CiliumCertGenImage() string {
	return findImage(cilium.CertGenImageName)
}

// CiliumKubeProxyImage returns the kube-proxy image.
func CiliumKubeProxyImage(kubernetesVersion string) string {
	return findImage(cilium.KubeProxyImageName, imagevector.RuntimeVersion(kubernetesVersion), imagevector.TargetVersion(kubernetesVersion))
}
