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
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	ciliumv1alpha1 "github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium/v1alpha1"
)

func (a *actuator) updateProviderStatus(
	ctx context.Context,
	network *extensionsv1alpha1.Network,
	config *ciliumv1alpha1.NetworkConfig,
) error {
	status, err := a.ComputeNetworkStatus(config)
	if err != nil {
		return err
	}

	patch := client.MergeFrom(network.DeepCopy())
	network.Status.ProviderStatus = &runtime.RawExtension{Object: status}
	network.Status.LastOperation = extensionscontroller.LastOperation(gardencorev1beta1.LastOperationTypeReconcile,
		gardencorev1beta1.LastOperationStateSucceeded,
		100,
		"Cilium was configured successfully")
	var ipFamilies []extensionsv1alpha1.IPFamily
	if config.IPv4 != nil {
		ipFamilies = append(ipFamilies, extensionsv1alpha1.IPFamilyIPv4)
	}
	if config.IPv6 != nil {
		ipFamilies = append(ipFamilies, extensionsv1alpha1.IPFamilyIPv6)
	}
	network.Status.IPFamilies = ipFamilies

	if err := a.client.Status().Patch(ctx, network, patch); err != nil {
		return &reconcilerutils.RequeueAfterError{
			Cause:        err,
			RequeueAfter: 30 * time.Second,
		}
	}
	return nil
}

func (a *actuator) ComputeNetworkStatus(networkConfig *ciliumv1alpha1.NetworkConfig) (*ciliumv1alpha1.NetworkStatus, error) {
	var (
		status = &ciliumv1alpha1.NetworkStatus{
			TypeMeta: StatusTypeMeta,
		}
	)

	return status, nil
}
