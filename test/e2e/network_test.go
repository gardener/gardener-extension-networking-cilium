// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package e2e_test

import (
	"context"
	"path/filepath"
	"time"

	"github.com/gardener/gardener-extension-networking-cilium/test/templates"
	"github.com/gardener/gardener/pkg/client/kubernetes"
	"github.com/gardener/gardener/test/utils/shoots/access"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("Network Extension Tests", Label("Network"), func() {
	f := defaultShootCreationFramework()
	f.Shoot = defaultShoot("e2e-default")

	It("Create Shoot, Test Networking, Delete Shoot", Label("good-case"), func() {
		By("Create Shoot")
		ctx, cancel := context.WithTimeout(parentCtx, 15*time.Minute)
		defer cancel()
		Expect(f.CreateShootAndWaitForCreation(ctx, false)).To(Succeed())
		f.Verify()

		By("Test Networking")
		ctx, cancel = context.WithTimeout(parentCtx, 15*time.Minute)
		defer cancel()
		values := struct {
			HelmDeployNamespace string
			KubeVersion         string
		}{
			templates.NetworkingTestNamespace,
			f.Shoot.Spec.Kubernetes.Version,
		}

		var err error
		f.GardenClient, err = kubernetes.NewClientFromFile("", f.ShootFramework.Config.GardenerConfig.GardenerKubeconfig,
			kubernetes.WithClientOptions(client.Options{Scheme: kubernetes.GardenScheme}),
			kubernetes.WithAllowedUserFields([]string{kubernetes.AuthTokenFile}),
			kubernetes.WithDisabledCachedClient(),
		)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		f.ShootFramework.ShootClient, err = access.CreateShootClientFromAdminKubeconfig(ctx, f.GardenClient, f.Shoot)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		resourceDir, err := filepath.Abs(filepath.Join(".."))
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		f.TemplatesDir = filepath.Join(resourceDir, "templates")
		f.Logger.Info("Templates Dir: ", f.TemplatesDir)

		err = f.RenderAndDeployTemplate(ctx, f.ShootFramework.ShootClient, templates.NetworkingTestDeploymentName, values)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		err = f.ShootFramework.WaitUntilDeploymentIsReady(ctx, "networking-test", templates.NetworkingTestNamespace, f.ShootFramework.ShootClient)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		By("Networking-test deployment was deployed successfully!")

		By("Delete Shoot")
		ctx, cancel = context.WithTimeout(parentCtx, 15*time.Minute)
		defer cancel()
		Expect(f.DeleteShootAndWaitForDeletion(ctx, f.Shoot)).To(Succeed())
	})
})
