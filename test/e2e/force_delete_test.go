// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package e2e_test

import (
	"context"
	"fmt"

	ciliumv1alpha1 "github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium/v1alpha1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTableSubtree("Network Extension Tests", Label("Network"), func(shootSuffix string, config *ciliumv1alpha1.NetworkConfig) {
	f := defaultShootCreationFramework()
	f.Shoot = defaultShoot(fmt.Sprintf("forcedel-%s", shootSuffix), config)
	It("Create Shoot, Test Network, Force Delete Shoot", Label("force-delete"), func() {
		By("Create Shoot")
		ctx, cancel := context.WithTimeout(parentCtx, defaultTimeout)
		defer cancel()
		Expect(f.CreateShootAndWaitForCreation(ctx, false)).To(Succeed())
		f.Verify()

		ctx, cancel = context.WithTimeout(parentCtx, defaultTimeout)
		defer cancel()
		succeeded := testNetwork(ctx, f)

		By("Wait for Shoot to be force-deleted")
		ctx, cancel = context.WithTimeout(parentCtx, defaultTimeout)
		defer cancel()
		Expect(f.ForceDeleteShootAndWaitForDeletion(ctx, f.Shoot)).To(Succeed())

		By("Network Test status")
		Expect(succeeded).To(BeTrue())
	})
},
	Entry("overlay config", "dft", defaultOverlayCiliumConfig()),
	Entry("wireguard config", "wg", wireguardCiliumConfig()),
)
