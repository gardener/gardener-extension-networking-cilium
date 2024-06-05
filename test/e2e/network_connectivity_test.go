// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package e2e_test

import (
	"context"
	"path/filepath"
	"time"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/gardener/gardener-extension-networking-cilium/test/templates"
	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	"github.com/gardener/gardener/pkg/client/kubernetes"
	gardenerutils "github.com/gardener/gardener/pkg/utils/gardener"
	"github.com/gardener/gardener/test/utils/access"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("Network Extension Tests", Label("Network"), func() {
	f := defaultShootCreationFramework()
	f.Shoot = defaultShoot("con-test")

	It("Create Shoot, Test Network (Cilium Connectivity), Delete Shoot", Label("good-case"), func() {
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
			templates.NetworkConnectivityTestNamespace,
			f.Shoot.Spec.Kubernetes.Version,
		}

		var err error
		f.GardenClient, err = kubernetes.NewClientFromFile("", f.ShootFramework.Config.GardenerConfig.GardenerKubeconfig,
			kubernetes.WithClientOptions(client.Options{Scheme: kubernetes.GardenScheme}),
			kubernetes.WithAllowedUserFields([]string{kubernetes.AuthTokenFile}),
			kubernetes.WithDisabledCachedClient(),
		)
		Expect(err).NotTo(HaveOccurred())

		shootKubeconfigSecret := &corev1.Secret{}
		gardenClient := f.GardenClient.Client()
		err = gardenClient.Get(ctx, client.ObjectKey{Namespace: f.Shoot.Namespace, Name: gardenerutils.ComputeShootProjectResourceName(f.Shoot.Name, gardenerutils.ShootProjectSecretSuffixKubeconfig)}, shootKubeconfigSecret)
		Expect(err).NotTo(HaveOccurred())

		f.ShootFramework.ShootClient, err = access.CreateShootClientFromAdminKubeconfig(ctx, f.GardenClient, f.Shoot)
		Expect(err).NotTo(HaveOccurred())

		newShootKubeconfigSecret := &corev1.Secret{ObjectMeta: v1.ObjectMeta{
			Name:      "kubeconfig",
			Namespace: values.HelmDeployNamespace},
			Data: map[string][]byte{"kubeconfig": shootKubeconfigSecret.Data["kubeconfig"]},
		}
		err = f.ShootFramework.ShootClient.Client().Create(ctx, newShootKubeconfigSecret)
		Expect(err).NotTo(HaveOccurred())

		resourceDir, err := filepath.Abs(filepath.Join(".."))
		Expect(err).NotTo(HaveOccurred())
		f.TemplatesDir = filepath.Join(resourceDir, "templates")

		err = f.RenderAndDeployTemplate(ctx, f.ShootFramework.ShootClient, templates.NetworkConnectivityTestJobName, values)
		Expect(err).NotTo(HaveOccurred())
		time.Sleep(30 * time.Second)
		err = f.ShootFramework.WaitUntilPodIsRunningWithLabels(
			ctx,
			labels.SelectorFromSet(labels.Set{v1beta1constants.LabelApp: "networking-test"}),
			templates.NetworkConnectivityTestNamespace,
			f.ShootFramework.ShootClient,
		)
		Expect(err).NotTo(HaveOccurred())
		By("Network-test Job was deployed successfully!")

		By("Check if network-test Job fails or succeeds!")

		job := &batchv1.Job{}
		var succeeded bool
		for {
			err = f.ShootFramework.ShootClient.Client().Get(ctx, client.ObjectKey{Namespace: values.HelmDeployNamespace, Name: "network-test"}, job)
			Expect(err).NotTo(HaveOccurred())

			if job.Status.Succeeded > 0 {
				succeeded = true
				break
			}

			if job.Status.Failed > 0 {
				succeeded = false
				break
			}

			time.Sleep(1 * time.Second)
		}

		By("Dump networking-test logs!")
		err = f.ShootFramework.DumpLogsForPodsWithLabelsInNamespace(ctx, f.ShootFramework.ShootClient, values.HelmDeployNamespace, client.MatchingLabels{v1beta1constants.LabelApp: "networking-test"})
		Expect(err).NotTo(HaveOccurred())

		By("Delete Shoot")
		ctx, cancel = context.WithTimeout(parentCtx, 15*time.Minute)
		defer cancel()
		Expect(f.DeleteShootAndWaitForDeletion(ctx, f.Shoot)).To(Succeed())

		By("Network Test (Cilium Connectivity) status")
		Expect(succeeded).To(BeTrue())
	})
})
