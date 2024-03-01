// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"context"
	"time"

	extensionscontroller "github.com/gardener/gardener/extensions/pkg/controller"
	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	reconcilerutils "github.com/gardener/gardener/pkg/controllerutils/reconciler"

	ciliumv1alpha1 "github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium/v1alpha1"
)

func (a *actuator) updateProviderStatus(
	ctx context.Context,
	network *extensionsv1alpha1.Network,
	config *ciliumv1alpha1.NetworkConfig,
) error {
	network.Status.LastOperation = extensionscontroller.LastOperation(gardencorev1beta1.LastOperationTypeReconcile,
		gardencorev1beta1.LastOperationStateSucceeded,
		100,
		"Cilium was configured successfully")

	if err := a.client.Status().Update(ctx, network); err != nil {
		return &reconcilerutils.RequeueAfterError{
			Cause:        err,
			RequeueAfter: 30 * time.Second,
		}
	}
	return nil
}
