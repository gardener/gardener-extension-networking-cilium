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

package controller

import (
	"context"
	"fmt"
	"time"

	extensionscontroller "github.com/gardener/gardener/extensions/pkg/controller"
	extensionswebhookshoot "github.com/gardener/gardener/extensions/pkg/webhook/shoot"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"github.com/gardener/gardener/pkg/utils/managedresources"
	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/gardener/gardener-extension-networking-cilium/pkg/cilium"
)

// Delete implements Network.Actuator.
func (a *actuator) Delete(ctx context.Context, _ logr.Logger, network *extensionsv1alpha1.Network, cluster *extensionscontroller.Cluster) error {
	// First delete the monitoring configuration
	if err := applyMonitoringConfig(ctx, a.client, a.chartApplier, network, true); err != nil {
		return err
	}

	// Then delete the managed resource along with its secrets
	if err := managedresources.Delete(ctx, a.client, network.Namespace, CiliumConfigManagedResourceName, true); err != nil {
		return err
	}

	if a.atomicShootWebhookConfig != nil {
		networkPolicy := extensionswebhookshoot.GetNetworkPolicyMeta(network.Namespace, cilium.Name)
		if err := a.client.Delete(ctx, networkPolicy); client.IgnoreNotFound(err) != nil {
			return fmt.Errorf("could not delete network policy for shoot webhooks in namespace '%s': %w", network.Namespace, err)
		}

		if err := managedresources.Delete(ctx, a.client, network.Namespace, ShootWebhooksResourceName, false); err != nil {
			return fmt.Errorf("could not delete managed resource containing shoot webhook '%s': %w", ShootWebhooksResourceName, err)
		}

		timeoutCtx4, cancel := context.WithTimeout(ctx, 2*time.Minute)
		defer cancel()
		if err := managedresources.WaitUntilDeleted(timeoutCtx4, a.client, network.Namespace, ShootWebhooksResourceName); err != nil {
			return fmt.Errorf("error while waiting for managed resource containing shoot webhook '%s' to be deleted: %w", ShootWebhooksResourceName, err)
		}
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()
	return managedresources.WaitUntilDeleted(timeoutCtx, a.client, network.Namespace, CiliumConfigManagedResourceName)
}
