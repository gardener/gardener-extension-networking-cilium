// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package install

import (
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium"
	"github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium/v1alpha1"
)

var (
	schemeBuilder = runtime.NewSchemeBuilder(
		v1alpha1.AddToScheme,
		cilium.AddToScheme,
		setVersionPriority,
	)

	// AddToScheme adds all APIs to the scheme.
	AddToScheme = schemeBuilder.AddToScheme
)

func setVersionPriority(scheme *runtime.Scheme) error {
	return scheme.SetVersionPriority(v1alpha1.SchemeGroupVersion)
}
