// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package e2e_test

import (
	"context"
	"os"

	ciliumv1alpha1 "github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium/v1alpha1"
	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	"github.com/gardener/gardener/test/framework"
	. "github.com/onsi/ginkgo/v2"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	jsonserializer "k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/utils/ptr"
)

var (
	parentCtx context.Context
	encoder   runtime.Encoder = &jsonserializer.Serializer{}
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

func defaultOverlayCiliumConfig() *ciliumv1alpha1.NetworkConfig {
	return &ciliumv1alpha1.NetworkConfig{
		Hubble: &ciliumv1alpha1.Hubble{
			Enabled: true,
		},
		Overlay: &ciliumv1alpha1.Overlay{
			Enabled: true,
		},
	}
}

func wireguardCiliumConfig() *ciliumv1alpha1.NetworkConfig {
	config := defaultOverlayCiliumConfig()
	config.Encryption = &ciliumv1alpha1.Encryption{
		Enabled: true,
		Mode:    ciliumv1alpha1.EncryptionModeWireguard,
	}
	return config
}

func defaultShoot(generateName string, ciliumConfig *ciliumv1alpha1.NetworkConfig) *gardencorev1beta1.Shoot {
	rawConfig, err := runtime.Encode(encoder, ciliumConfig)
	if err != nil {
		panic(err)
	}
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
				ProviderConfig: &runtime.RawExtension{Raw: rawConfig},
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
