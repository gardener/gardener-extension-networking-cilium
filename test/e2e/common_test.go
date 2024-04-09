// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package e2e_test

import (
	"context"
	"os"

	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	"github.com/gardener/gardener/test/framework"
	. "github.com/onsi/ginkgo/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/utils/pointer"
	"k8s.io/utils/ptr"
)

var (
	parentCtx context.Context
)

var _ = BeforeEach(func() {
	parentCtx = context.Background()
})

const projectNamespace = "garden-local"

func defaultShootCreationFramework() *framework.ShootCreationFramework {
	kubeconfigPath := os.Getenv("KUBECONFIG")
	return framework.NewShootCreationFramework(&framework.ShootCreationConfig{
		GardenerConfig: &framework.GardenerConfig{
			ProjectNamespace:   projectNamespace,
			GardenerKubeconfig: kubeconfigPath,
			SkipAccessingShoot: true,
			CommonConfig:       &framework.CommonConfig{},
		},
	})
}

func defaultShoot(generateName string) *gardencorev1beta1.Shoot {
	return &gardencorev1beta1.Shoot{
		ObjectMeta: metav1.ObjectMeta{
			Name: generateName,
			Annotations: map[string]string{
				v1beta1constants.AnnotationShootInfrastructureCleanupWaitPeriodSeconds: "0",
				v1beta1constants.AnnotationShootCloudConfigExecutionMaxDelaySeconds:    "0",
			},
		},
		Spec: gardencorev1beta1.ShootSpec{
			Region:            "local",
			SecretBindingName: pointer.String("local"),
			CloudProfileName:  "local",
			Kubernetes: gardencorev1beta1.Kubernetes{
				Version:                     "1.26.0",
				EnableStaticTokenKubeconfig: pointer.Bool(true),
				Kubelet: &gardencorev1beta1.KubeletConfig{
					SerializeImagePulls: pointer.Bool(false),
					RegistryPullQPS:     pointer.Int32(10),
					RegistryBurst:       pointer.Int32(20),
				},
				KubeAPIServer: &gardencorev1beta1.KubeAPIServerConfig{},
				KubeProxy: &gardencorev1beta1.KubeProxyConfig{
					Mode:    ptr.To(gardencorev1beta1.ProxyModeIPTables),
					Enabled: pointer.Bool(false),
				},
			},
			Networking: &gardencorev1beta1.Networking{
				Type:           pointer.String("cilium"),
				ProviderConfig: &runtime.RawExtension{Raw: []byte(`{"apiVersion":"cilium.networking.extensions.gardener.cloud/v1alpha1","kind":"NetworkConfig","hubble":{"enabled":true},"overlay":{"enabled":true}}`)},
			},
			Provider: gardencorev1beta1.Provider{
				Type: "local",
				Workers: []gardencorev1beta1.Worker{{
					Name: "local",
					Machine: gardencorev1beta1.Machine{
						Type: "local",
					},
					CRI: &gardencorev1beta1.CRI{
						Name: gardencorev1beta1.CRINameContainerD,
					},
					Minimum: 2,
					Maximum: 2,
				}},
			},
		},
	}
}
