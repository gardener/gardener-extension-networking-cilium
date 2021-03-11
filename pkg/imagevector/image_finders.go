package imagevector

import (
	"github.com/gardener/gardener-extension-networking-cilium/pkg/cilium"
	"k8s.io/apimachinery/pkg/util/runtime"
)

func findImage(name string) string {
	image, err := imageVector.FindImage(name)
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

// CiliumNodeInitImage returns the Cilium Node Init image.
func CiliumNodeInitImage() string {
	return findImage(cilium.CiliumNodeInitImageName)
}

// CiliumPreflightImage returns the Cilium Preflight image.
func CiliumPreflightImage() string {
	return findImage(cilium.CiliumPreflightImageName)
}

// CiliumEtcdOperatorImage returns the Cilium-etcd-Operator image.
func CiliumEtcdOperatorImage() string {
	return findImage(cilium.CiliumETCDOperatorImageName)
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

// CiliumEnvoyImage returns the Envoy image.
func CiliumEnvoyImage() string {
	return findImage(cilium.EnvoyImageName)
}
