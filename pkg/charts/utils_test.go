package charts

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	ciliumv1alpha1 "github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium/v1alpha1"
)

var _ = Describe("#applyEncryptionConfig", func() {
	Describe("wireguard", func() {
		var config *ciliumv1alpha1.NetworkConfig
		BeforeEach(func() {
			config = &ciliumv1alpha1.NetworkConfig{
				Encryption: &ciliumv1alpha1.Encryption{
					Mode:    ciliumv1alpha1.EncryptionModeWireguard,
					Enabled: true,
				},
			}
		})
		Describe("vxlan tunnel config with strict mode encryption", func() {
			It("should set AllowRemoteNodeIdentities", func() {
				cfg := &globalConfig{
					Tunnel: ciliumv1alpha1.VXLan,
				}
				config.Encryption.StrictMode = true
				Expect(applyEncryptionConfig(cfg, config)).ShouldNot(HaveOccurred())
				Expect(cfg.Encryption.Wireguard.StrictMode.AllowRemoteNodeIdentities).To(BeTrue())
			})
		})
		Describe("overlapping node & pod CIDR with direct routing and strict mode encryption", func() {
			It("should set AllowRemoteNodeIdentities", func() {
				config.Encryption.StrictMode = true
				config.Overlay = &ciliumv1alpha1.Overlay{
					Enabled: false,
				}
				cfg := &globalConfig{
					Tunnel:   ciliumv1alpha1.Disabled,
					PodCIDR:  "10.0.0.0/16",
					NodeCIDR: "10.0.0.128/17",
				}
				Expect(applyEncryptionConfig(cfg, config)).ShouldNot(HaveOccurred())
				Expect(cfg.Encryption.Wireguard.StrictMode.AllowRemoteNodeIdentities).To(BeTrue())
			})
		})
	})
})
