// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

/**
	Overview
		- Tests the health checks for the networking-cilium extension.

	Prerequisites
		- A Shoot exists.

	Test-case:
		- Network CRD
			- HealthCondition Type: Shoot SystemComponentsHealthy
				1)  update the ManagedResource 'extension-networking-cilium-config' and verify health check conditions in the Network CRD status.

 **/

package healthcheck

import (
	"context"
	"fmt"
	"time"

	networkcontroller "github.com/gardener/gardener-extension-networking-cilium/pkg/controller"

	healthcheckoperation "github.com/gardener/gardener/extensions/test/integration/healthcheck"
	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	"github.com/gardener/gardener/test/framework"
	"github.com/onsi/ginkgo"
)

const (
	timeout = 5 * time.Minute
)

var _ = ginkgo.Describe("Networking-cilium integration test: health checks", func() {
	f := framework.NewShootFramework(nil)

	ginkgo.Context("Network", func() {
		ginkgo.Context("Condition type: ShootSystemComponentsHealthy", func() {
			f.Serial().Release().CIt(fmt.Sprintf("Network CRD should contain unhealthy condition due to ManagedResource '%s' is unhealthy", networkcontroller.CiliumConfigSecretName), func(ctx context.Context) {
				err := healthcheckoperation.NetworkHealthCheckWithManagedResource(ctx, timeout, f, networkcontroller.CiliumConfigSecretName, gardencorev1beta1.ShootSystemComponentsHealthy)
				framework.ExpectNoError(err)
			}, timeout)
		})
	})
})
