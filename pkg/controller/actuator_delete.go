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
	"time"

	resourcemanager "github.com/gardener/gardener-resource-manager/pkg/manager"
	extensionscontroller "github.com/gardener/gardener/extensions/pkg/controller"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	managedresources "github.com/gardener/gardener/pkg/utils/managedresources"
)

// Delete implements Network.Actuator.
func (a *actuator) Delete(ctx context.Context, network *extensionsv1alpha1.Network, cluster *extensionscontroller.Cluster) error {
	// First delete the monitoring configuration
	if err := applyMonitoringConfig(ctx, a.client, a.chartApplier, network, true); err != nil {
		return err
	}

	// Then delete the managed resource along with its secrets
	if err := resourcemanager.
		NewSecret(a.client).
		WithNamespacedName(network.Namespace, CiliumConfigSecretName).
		Delete(ctx); err != nil {
		return err
	}
	if err := resourcemanager.
		NewManagedResource(a.client).
		WithNamespacedName(network.Namespace, CiliumConfigSecretName).
		Delete(ctx); err != nil {
		return err
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()
	return managedresources.WaitUntilManagedResourceDeleted(timeoutCtx, a.client, network.Namespace, CiliumConfigSecretName)
}
