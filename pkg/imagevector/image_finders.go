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

// CiliumEtcdOperator returns the Cilium-etcd-Operator image.
func CiliumEtcdOperatorImage() string {
	return findImage(cilium.CiliumETCDOperatorImageName)
}

// CiliumHubbleImage returns the Cilium Hubble image.
func CiliumHubbleImage() string {
	return findImage(cilium.HubbleImageName)
}

// CiliumHubbleUIImage returns the Cilium Hubble UI image.
func CiliumHubbleUIImage() string {
	return findImage(cilium.HubbleUIImageName)
}
