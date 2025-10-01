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
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
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
			SkipAccessingShoot: false,
			CommonConfig:       &framework.CommonConfig{},
		},
	})
}

func defaultShoot(generateName string) *gardencorev1beta1.Shoot {
	return &gardencorev1beta1.Shoot{
		ObjectMeta: metav1.ObjectMeta{
			Name: generateName,
			Annotations: map[string]string{
				v1beta1constants.AnnotationShootCloudConfigExecutionMaxDelaySeconds: "0",
				v1beta1constants.ShootDisableIstioTLSTermination:                    "true",
				v1beta1constants.ShootAlphaControlPlaneScaleDownDisabled:            "true",
				v1beta1constants.ShootAlphaControlPlaneVPNVPAUpdateDisabled:         "true",
			},
		},
		Spec: gardencorev1beta1.ShootSpec{
			Region:            "local",
			SecretBindingName: ptr.To("local"),
			CloudProfile: &gardencorev1beta1.CloudProfileReference{
				Name: "local",
				Kind: "CloudProfile",
			},
			Kubernetes: gardencorev1beta1.Kubernetes{
				Version: "1.30.0",
				Kubelet: &gardencorev1beta1.KubeletConfig{
					SerializeImagePulls: ptr.To(false),
					RegistryPullQPS:     ptr.To(int32(10)),
					RegistryBurst:       ptr.To(int32(10)),
				},
				KubeProxy: &gardencorev1beta1.KubeProxyConfig{
					Mode:    ptr.To(gardencorev1beta1.ProxyModeIPTables),
					Enabled: ptr.To(false),
				},
				VerticalPodAutoscaler: &gardencorev1beta1.VerticalPodAutoscaler{
					Enabled: false,
				},
				KubeAPIServer: &gardencorev1beta1.KubeAPIServerConfig{
					Autoscaling: &gardencorev1beta1.ControlPlaneAutoscaling{
						MinAllowed: map[corev1.ResourceName]resource.Quantity{
							corev1.ResourceCPU:    resource.MustParse("250m"),
							corev1.ResourceMemory: resource.MustParse("500Mi"),
						},
					},
				},
			},
			Networking: &gardencorev1beta1.Networking{
				Type:           ptr.To("cilium"),
				Nodes:          ptr.To("10.0.0.0/16"),
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
			SystemComponents: &gardencorev1beta1.SystemComponents{
				NodeLocalDNS: &gardencorev1beta1.NodeLocalDNS{
					Enabled: true,
				},
			},
		},
	}
}
