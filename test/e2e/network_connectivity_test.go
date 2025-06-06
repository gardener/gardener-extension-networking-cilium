// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package e2e_test

import (
	"context"
	"path/filepath"

	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	"github.com/gardener/gardener/test/utils/access"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/gardener/gardener-extension-networking-cilium/test/templates"
)

var _ = Describe("Network Extension Tests", Label("Network"), func() {
	f := defaultShootCreationFramework()
	f.Shoot = defaultShoot("con-test")

	It("Create Shoot, Test Network (Cilium Connectivity), Delete Shoot", Label("good-case"), func() {
		By("Create Shoot")
		ctx, cancel := context.WithTimeout(parentCtx, defaultTimeout)
		defer cancel()
		Expect(f.CreateShootAndWaitForCreation(ctx, false)).To(Succeed())
		f.Verify()

		By("Create Deny-All Network Policy")
		denyAllPolicy := &networkingv1.NetworkPolicy{
			ObjectMeta: v1.ObjectMeta{Name: "deny-all", Namespace: templates.NetworkConnectivityTestNamespace},
			Spec: networkingv1.NetworkPolicySpec{
				PodSelector: v1.LabelSelector{},
				PolicyTypes: []networkingv1.PolicyType{networkingv1.PolicyTypeIngress, networkingv1.PolicyTypeEgress},
				Ingress:     []networkingv1.NetworkPolicyIngressRule{},
				Egress:      []networkingv1.NetworkPolicyEgressRule{},
			},
		}
		err := f.ShootFramework.ShootClient.Client().Create(ctx, denyAllPolicy)
		Expect(err).NotTo(HaveOccurred())

		By("Test Networking")
		ctx, cancel = context.WithTimeout(parentCtx, defaultTimeout)
		defer cancel()
		values := struct {
			HelmDeployNamespace string
			KubeVersion         string
		}{
			templates.NetworkConnectivityTestNamespace,
			f.Shoot.Spec.Kubernetes.Version,
		}

		const expirationSeconds int64 = 1 * 3600 // 1h
		shootAdminKubeconfig, err := access.RequestAdminKubeconfigForShoot(ctx, f.GardenClient, f.Shoot, ptr.To(expirationSeconds))
		Expect(err).NotTo(HaveOccurred())

		newShootKubeconfigSecret := &corev1.Secret{ObjectMeta: v1.ObjectMeta{
			Name:      "kubeconfig",
			Namespace: values.HelmDeployNamespace},
			Data: map[string][]byte{"kubeconfig": shootAdminKubeconfig},
		}
		err = f.ShootFramework.ShootClient.Client().Create(ctx, newShootKubeconfigSecret)
		Expect(err).NotTo(HaveOccurred())

		resourceDir, err := filepath.Abs(filepath.Join(".."))
		Expect(err).NotTo(HaveOccurred())
		f.TemplatesDir = filepath.Join(resourceDir, "templates")

		By("Deploy network-test Job")
		err = f.RenderAndDeployTemplate(ctx, f.ShootFramework.ShootClient, templates.NetworkConnectivityTestJobName, values)
		Expect(err).NotTo(HaveOccurred())
		Eventually(func() error {
			err = f.ShootFramework.WaitUntilPodIsRunningWithLabels(
				ctx,
				labels.SelectorFromSet(labels.Set{v1beta1constants.LabelApp: "networking-test"}),
				templates.NetworkConnectivityTestNamespace,
				f.ShootFramework.ShootClient,
			)
			return err
		}).WithTimeout(defaultTimeout).Should(Succeed())

		By("Check if network-test Job fails or succeeds!")

		job := &batchv1.Job{}
		var succeeded bool
		Eventually(func() bool {
			keepGoing := true
			err = f.ShootFramework.ShootClient.Client().Get(ctx, client.ObjectKey{Namespace: values.HelmDeployNamespace, Name: "network-test"}, job)
			Expect(err).NotTo(HaveOccurred())

			if job.Status.Succeeded > 0 {
				succeeded = true
				keepGoing = false
			}

			if job.Status.Failed > 0 {
				succeeded = false
				keepGoing = false
			}
			return keepGoing
		}).WithTimeout(defaultTimeout).Should(BeFalse())

		By("Dump networking-test logs!")
		err = f.ShootFramework.DumpLogsForPodsWithLabelsInNamespace(ctx, f.ShootFramework.ShootClient, values.HelmDeployNamespace, client.MatchingLabels{v1beta1constants.LabelApp: "networking-test"})
		Expect(err).NotTo(HaveOccurred())

		By("Delete Shoot")
		ctx, cancel = context.WithTimeout(parentCtx, defaultTimeout)
		defer cancel()
		Expect(f.DeleteShootAndWaitForDeletion(ctx, f.Shoot)).To(Succeed())

		By("Network Test (Cilium Connectivity) status")
		Expect(succeeded).To(BeTrue())
	})
})
