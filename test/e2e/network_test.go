// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package e2e_test

import (
	"context"
	"path/filepath"
	"time"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/gardener/gardener-extension-networking-cilium/test/templates"
	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	"github.com/gardener/gardener/pkg/client/kubernetes"
	gardenerutils "github.com/gardener/gardener/pkg/utils/gardener"
	"github.com/gardener/gardener/test/framework"
	"github.com/gardener/gardener/test/utils/access"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("Network Extension Tests", Label("Network"), func() {
	f := defaultShootCreationFramework()
	f.Shoot = defaultShoot("ping-test")

	It("Create Shoot, Test Network (Ping), Delete Shoot", Label("good-case"), func() {
		By("Create Shoot")
		ctx, cancel := context.WithTimeout(parentCtx, 15*time.Minute)
		defer cancel()
		Expect(f.CreateShootAndWaitForCreation(ctx, false)).To(Succeed())
		f.Verify()

		By("Test Network")
		ctx, cancel = context.WithTimeout(parentCtx, 15*time.Minute)
		defer cancel()
		succeeded := testNetwork(ctx, f)

		By("Delete Shoot")
		ctx, cancel = context.WithTimeout(parentCtx, 15*time.Minute)
		defer cancel()
		Expect(f.DeleteShootAndWaitForDeletion(ctx, f.Shoot)).To(Succeed())

		By("Network Test (Ping) status")
		Expect(succeeded).To(BeTrue())
	})
})

func testNetwork(ctx context.Context, f *framework.ShootCreationFramework) bool {
	values := struct {
		HelmDeployNamespace string
		KubeVersion         string
	}{
		templates.NetworkTestNamespace,
		f.Shoot.Spec.Kubernetes.Version,
	}

	var err error
	f.GardenClient, err = kubernetes.NewClientFromFile("", f.ShootFramework.Config.GardenerConfig.GardenerKubeconfig,
		kubernetes.WithClientOptions(client.Options{Scheme: kubernetes.GardenScheme}),
		kubernetes.WithAllowedUserFields([]string{kubernetes.AuthTokenFile}),
		kubernetes.WithDisabledCachedClient(),
	)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())

	shootKubeconfigSecret := &corev1.Secret{}
	gardenClient := f.GardenClient.Client()
	gardenClient.Get(ctx, client.ObjectKey{Namespace: f.Shoot.Namespace, Name: gardenerutils.ComputeShootProjectResourceName(f.Shoot.Name, gardenerutils.ShootProjectSecretSuffixKubeconfig)}, shootKubeconfigSecret)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())

	f.ShootFramework.ShootClient, err = access.CreateShootClientFromAdminKubeconfig(ctx, f.GardenClient, f.Shoot)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())

	newShootKubeconfigSecret := &corev1.Secret{ObjectMeta: v1.ObjectMeta{
		Name:      "kubeconfig",
		Namespace: values.HelmDeployNamespace},
		Data: map[string][]byte{"kubeconfig": shootKubeconfigSecret.Data["kubeconfig"]},
	}
	err = f.ShootFramework.ShootClient.Client().Create(ctx, newShootKubeconfigSecret)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())

	resourceDir, err := filepath.Abs(filepath.Join(".."))
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	f.TemplatesDir = filepath.Join(resourceDir, "templates")

	err = f.RenderAndDeployTemplate(ctx, f.ShootFramework.ShootClient, templates.NetworkTestName, values)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	time.Sleep(30 * time.Second)
	err = f.ShootFramework.WaitUntilDaemonSetIsRunning(
		ctx,
		f.ShootFramework.ShootClient.Client(),
		"host-network",
		values.HelmDeployNamespace,
	)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())

	err = f.ShootFramework.WaitUntilDaemonSetIsRunning(
		ctx,
		f.ShootFramework.ShootClient.Client(),
		"pod-network",
		values.HelmDeployNamespace,
	)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())

	By("Network-test daemonsets were deployed successfully!")

	By("Check if network-test fails or succeeds!")

	// wait one minute until results are collected
	time.Sleep(60 * time.Second)

	podListHostNetwork, err := framework.GetPodsByLabels(ctx, labels.SelectorFromSet(map[string]string{
		v1beta1constants.LabelApp: "host-network",
	}), f.ShootFramework.ShootClient, values.HelmDeployNamespace)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())

	podListPodNetwork, err := framework.GetPodsByLabels(ctx, labels.SelectorFromSet(map[string]string{
		v1beta1constants.LabelApp: "pod-network",
	}), f.ShootFramework.ShootClient, values.HelmDeployNamespace)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())

	containerStatuses := []corev1.ContainerStatus{}
	for _, pod := range podListHostNetwork.Items {
		containerStatuses = append(containerStatuses, pod.Status.ContainerStatuses...)
	}

	for _, pod := range podListPodNetwork.Items {
		containerStatuses = append(containerStatuses, pod.Status.ContainerStatuses...)
	}
	succeeded := true
	for _, containerStatus := range containerStatuses {
		if containerStatus.RestartCount > 0 {
			succeeded = false
			break
		}
	}

	return succeeded
}
