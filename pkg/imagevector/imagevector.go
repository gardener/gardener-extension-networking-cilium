// Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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
	"strings"

	"github.com/gardener/gardener-extension-networking-cilium/pkg/cilium"

	"github.com/gardener/gardener/pkg/utils/imagevector"
	"github.com/gobuffalo/packr/v2"
	"k8s.io/apimachinery/pkg/util/runtime"
)

var imageVector imagevector.ImageVector

func init() {
	box := packr.New("charts", "../../charts")

	imagesYaml, err := box.FindString("images.yaml")
	runtime.Must(err)

	imageVector, err = imagevector.Read(strings.NewReader(imagesYaml))
	runtime.Must(err)

	imageVector, err = imagevector.WithEnvOverride(imageVector)
	runtime.Must(err)
}

// ImageVector is the image vector that contains all the needed images.
func ImageVector() imagevector.ImageVector {
	return imageVector
}

// CiliumImage returns the Cilium Image.
func CiliumImage() string {
	image, err := imageVector.FindImage(cilium.CiliumImageName)
	runtime.Must(err)
	return image.String()
}

// CiliumOperatorImage returns the Cilium Operator image.
func CiliumOperatorImage() string {
	image, err := imageVector.FindImage(cilium.CiliumOperatorImageName)
	runtime.Must(err)
	return image.String()
}

// CiliumDockerPluginImage returns the Cilium Docker Plugin image.
func CiliumDockerPluginImage() string {
	image, err := imageVector.FindImage(cilium.CiliumDockerPluginImageName)
	runtime.Must(err)
	return image.String()
}

// CiliumEtcdOperator returns the Cilium-etcd-Operator image.
func CiliumEtcdOperatorImage() string {
	image, err := imageVector.FindImage(cilium.CiliumEtcdOperatorImageName)
	runtime.Must(err)
	return image.String()
}

// CiliumHubbleImage returns the Cilium Hubble image.
func CiliumHubbleImage() string {
	image, err := imageVector.FindImage(cilium.HubbleImageName)
	runtime.Must(err)
	return image.String()
}

// CiliumHubbleUIImage returns the Cilium Hubble UI image.
func CiliumHubbleUIImage() string {
	image, err := imageVector.FindImage(cilium.HubbleUIImageName)
	runtime.Must(err)
	return image.String()
}

// IstioPilotImage returns the Istio Pilot image.
func IstioPilotImage() string {
	image, err := imageVector.FindImage(cilium.IstioPilotImageName)
	runtime.Must(err)
	return image.String()
}

// IstioProxyImage returns the Istio Proxy image.
func IstioProxyImage() string {
	image, err := imageVector.FindImage(cilium.IstioProxyImageName)
	runtime.Must(err)
	return image.String()
}
