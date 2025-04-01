// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package mutator

import (
	"context"
	"fmt"

	extensionswebhook "github.com/gardener/gardener/extensions/pkg/webhook"
	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NewShootMutator returns a new instance of a shoot mutator.
func NewShootMutator() extensionswebhook.Mutator {
	return &shoot{}
}

type shoot struct{}

// Mutate mutates the given shoot object.
func (s *shoot) Mutate(ctx context.Context, new, old client.Object) error {
	shoot, ok := new.(*gardencorev1beta1.Shoot)
	if !ok {
		return fmt.Errorf("wrong object type %T", new)
	}

	if old != nil {
		// Leave existing clusters as is
		return nil
	}

	if shoot.Spec.Kubernetes.KubeProxy == nil {
		shoot.Spec.Kubernetes.KubeProxy = &gardencorev1beta1.KubeProxyConfig{}
	}
	shoot.Spec.Kubernetes.KubeProxy.Enabled = ptr.To(false)

	return nil
}
