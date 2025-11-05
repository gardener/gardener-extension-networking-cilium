// SPDX-FileCopyrightText: 2025 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package validation_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	gomegatypes "github.com/onsi/gomega/types"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/utils/ptr"

	apiscilium "github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium"
	"github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium/validation"
)

var _ = Describe("Network validation", func() {
	DescribeTable("#ValidateNetworkConfig",
		func(networkConfig *apiscilium.NetworkConfig, fldPath *field.Path, matcher gomegatypes.GomegaMatcher) {
			Expect(validation.ValidateNetworkConfig(networkConfig, fldPath)).To(matcher)
		},

		Entry("should succeed with empty config", &apiscilium.NetworkConfig{}, field.NewPath("config"),
			BeEmpty()),
		Entry("should return error with incorrect kubeproxy config", &apiscilium.NetworkConfig{KubeProxy: &apiscilium.KubeProxy{ServiceHost: ptr.To("-foo"), ServicePort: ptr.To(int32(-1))}}, field.NewPath("config"),
			ConsistOf(
				PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("config.kubeproxy.k8sServiceHost")})),
				PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("config.kubeproxy.k8sServicePort")})),
			)),
		Entry("should succeed with valid kubeproxy config", &apiscilium.NetworkConfig{KubeProxy: &apiscilium.KubeProxy{ServiceHost: ptr.To("foo"), ServicePort: ptr.To(int32(1))}}, field.NewPath("config"),
			BeEmpty()),
		Entry("should return error with incorrect tunnel config", &apiscilium.NetworkConfig{TunnelMode: ptr.To(apiscilium.TunnelMode("ipip"))}, field.NewPath("config"),
			ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("config.tunnel")})))),
		Entry("should succeed with valid tunnel config", &apiscilium.NetworkConfig{TunnelMode: ptr.To(apiscilium.Disabled)}, field.NewPath("config"),
			BeEmpty()),
		Entry("should return error with incorrect store config", &apiscilium.NetworkConfig{Store: ptr.To(apiscilium.Store("etcd"))}, field.NewPath("config"),
			ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("config.store")})))),
		Entry("should succeed with valid store config", &apiscilium.NetworkConfig{Store: ptr.To(apiscilium.Kubernetes)}, field.NewPath("config"),
			BeEmpty()),
		Entry("should return error with incorrect MTU config", &apiscilium.NetworkConfig{MTU: ptr.To(-1)}, field.NewPath("config"),
			ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("config.mtu")})))),
		Entry("should succeed with valid MTU config", &apiscilium.NetworkConfig{MTU: ptr.To(1400)}, field.NewPath("config"),
			BeEmpty()),
		Entry("should return error with incorrect device config", &apiscilium.NetworkConfig{Devices: []string{"", "abc/def"}}, field.NewPath("config"),
			ConsistOf(
				PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("config.devices[0]")})),
				PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("config.devices[1]")})),
			)),
		Entry("should succeed with valid device config", &apiscilium.NetworkConfig{Devices: []string{"dev-123", "dev-456"}}, field.NewPath("config"),
			BeEmpty()),
		Entry("should return error with incorrect routing device config", &apiscilium.NetworkConfig{DirectRoutingDevice: ptr.To("dev 123")}, field.NewPath("config"),
			ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("config.directRoutingDevice")})))),
		Entry("should succeed with valid device config", &apiscilium.NetworkConfig{DirectRoutingDevice: ptr.To("dev-123")}, field.NewPath("config"),
			BeEmpty()),
		Entry("should return error with incorrect load balancing config", &apiscilium.NetworkConfig{LoadBalancingMode: ptr.To(apiscilium.LoadBalancingMode("my-mode"))}, field.NewPath("config"),
			ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("config.loadBalancingMode")})))),
		Entry("should succeed with valid load balancing config", &apiscilium.NetworkConfig{LoadBalancingMode: ptr.To(apiscilium.DSR)}, field.NewPath("config"),
			BeEmpty()),
	)

	DescribeTable("#ValidateNetworkConfigKubeProxy",
		func(kubeProxy *apiscilium.KubeProxy, fldPath *field.Path, matcher gomegatypes.GomegaMatcher) {
			Expect(validation.ValidateNetworkConfigKubeProxy(kubeProxy, fldPath)).To(matcher)
		},

		Entry("should succeed with nil", nil, field.NewPath("kubeproxy"),
			BeEmpty()),
		Entry("should succeed with empty struct", &apiscilium.KubeProxy{}, field.NewPath("kubeproxy"),
			BeEmpty()),
		Entry("should return error because domain name starts with invalid character", &apiscilium.KubeProxy{ServiceHost: ptr.To("-invalid.domain")}, field.NewPath("kubeproxy"),
			ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("kubeproxy.k8sServiceHost")})))),
		Entry("should return error because domain name has invalid characters", &apiscilium.KubeProxy{ServiceHost: ptr.To("inv_alid.do$main")}, field.NewPath("kubeproxy"),
			ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("kubeproxy.k8sServiceHost")})))),
		Entry("should return error because domain name has invalid characters", &apiscilium.KubeProxy{ServiceHost: ptr.To("inv_alid.do$main")}, field.NewPath("kubeproxy"),
			ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("kubeproxy.k8sServiceHost")})))),
		Entry("should return error because domain name is too long", &apiscilium.KubeProxy{ServiceHost: ptr.To("12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234")}, field.NewPath("kubeproxy"),
			ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("kubeproxy.k8sServiceHost")})))),
		Entry("should succeed with IPv4 address", &apiscilium.KubeProxy{ServiceHost: ptr.To("10.11.12.13")}, field.NewPath("kubeproxy"),
			BeEmpty()),
		Entry("should succeed with IPv6 address", &apiscilium.KubeProxy{ServiceHost: ptr.To("2001:db8::1234")}, field.NewPath("kubeproxy"),
			BeEmpty()),
		Entry("should succeed with valid domain name", &apiscilium.KubeProxy{ServiceHost: ptr.To("service.endpoint.local")}, field.NewPath("kubeproxy"),
			BeEmpty()),
		Entry("should return error with zero port", &apiscilium.KubeProxy{ServicePort: ptr.To(int32(0))}, field.NewPath("kubeproxy"),
			ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("kubeproxy.k8sServicePort")})))),
		Entry("should return error with negative port", &apiscilium.KubeProxy{ServicePort: ptr.To(int32(-12345))}, field.NewPath("kubeproxy"),
			ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("kubeproxy.k8sServicePort")})))),
		Entry("should return error with high port", &apiscilium.KubeProxy{ServicePort: ptr.To(int32(123456))}, field.NewPath("kubeproxy"),
			ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("kubeproxy.k8sServicePort")})))),
		Entry("should succeed with valid port", &apiscilium.KubeProxy{ServicePort: ptr.To(int32(12345))}, field.NewPath("kubeproxy"),
			BeEmpty()),
	)

	DescribeTable("#ValidateDevice",
		func(device string, fldPath *field.Path, matcher gomegatypes.GomegaMatcher) {
			Expect(validation.ValidateDevice(device, fldPath)).To(matcher)
		},

		Entry("should return error because empty device names are not allowed", "", field.NewPath("device"),
			ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("device")})))),
		Entry("should return error because device name is too long", "1234567890123456", field.NewPath("device"),
			ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("device")})))),
		Entry("should return error because whitespace is not allowed", "dev 1234", field.NewPath("device"),
			ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("device")})))),
		Entry("should return error because '/' is not allowed", "dev/1234", field.NewPath("device"),
			ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("device")})))),
		Entry("should succeed", "dev-1234", field.NewPath("device"),
			BeEmpty()),
	)
	DescribeTable("#ValidateEncryption",
		func(enc *apiscilium.Encryption, matcher gomegatypes.GomegaMatcher) {
			Expect(validation.ValidateEncryption(enc, field.NewPath("encryption"))).To(matcher)
		},

		Entry("should succeed if encryption is not enabled", &apiscilium.Encryption{Enabled: false}, BeEmpty()),
		Entry("should error if encryption mode is not wireguard but strict mode is used", &apiscilium.Encryption{Enabled: true, Mode: apiscilium.EncryptionMode("foo"), StrictMode: true},
			ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("encryption.mode")}))),
		),
	)
})
