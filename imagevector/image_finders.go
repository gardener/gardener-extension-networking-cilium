// Copyright (c) 2023 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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
