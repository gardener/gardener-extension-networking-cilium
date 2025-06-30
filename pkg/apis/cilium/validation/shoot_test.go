// SPDX-FileCopyrightText: 2025 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package validation_test

import (
	"github.com/gardener/gardener/pkg/apis/core"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	gomegatypes "github.com/onsi/gomega/types"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/utils/ptr"

	"github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium/validation"
	"github.com/gardener/gardener-extension-networking-cilium/pkg/cilium"
)

var _ = Describe("Shoot validation", func() {
	DescribeTable("#ValidateNetworking",
		func(networking *core.Networking, matcher gomegatypes.GomegaMatcher) {
			Expect(validation.ValidateNetworking(networking, field.NewPath("spec", "networking"))).To(matcher)
		},

		Entry("should return error because type is missing", &core.Networking{},
			ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("spec.networking.type")})))),
		Entry("should return error because type is invalid", &core.Networking{Type: ptr.To("foo")},
			ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("spec.networking.type")})))),
		Entry("should succeed because type is correct", &core.Networking{Type: ptr.To(cilium.Type)},
			BeEmpty()),
		Entry("should return error because pods attribute is invalid", &core.Networking{Type: ptr.To(cilium.Type), Pods: ptr.To("foo")},
			ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("spec.networking.pods")})))),
		Entry("should succeed because pods attribute is correct", &core.Networking{Type: ptr.To(cilium.Type), Pods: ptr.To("10.0.0.0/23")},
			BeEmpty()),
		Entry("should return error because nodes attribute is invalid", &core.Networking{Type: ptr.To(cilium.Type), Nodes: ptr.To("foo")},
			ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("spec.networking.nodes")})))),
		Entry("should succeed because nodes attribute is correct", &core.Networking{Type: ptr.To(cilium.Type), Nodes: ptr.To("10.0.0.0/23")},
			BeEmpty()),
		Entry("should return error because services attribute is invalid", &core.Networking{Type: ptr.To(cilium.Type), Services: ptr.To("foo")},
			ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("spec.networking.services")})))),
		Entry("should succeed because services attribute is correct", &core.Networking{Type: ptr.To(cilium.Type), Services: ptr.To("10.0.0.0/23")},
			BeEmpty()),
		Entry("should return error because ipFamilies attribute is invalid (single entry)", &core.Networking{Type: ptr.To(cilium.Type), IPFamilies: []core.IPFamily{"foo"}},
			ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("spec.networking.ipFamilies")})))),
		Entry("should return error because ipFamilies attribute is invalid (multiple entries)", &core.Networking{Type: ptr.To(cilium.Type), IPFamilies: []core.IPFamily{"foo", "bar", "baz"}},
			ConsistOf(
				PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("spec.networking.ipFamilies")})),
				PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("spec.networking.ipFamilies")})),
			)),
		Entry("should return error because ipFamilies attribute is invalid (duplicate entry)", &core.Networking{Type: ptr.To(cilium.Type), IPFamilies: []core.IPFamily{core.IPFamilyIPv6, core.IPFamilyIPv6}},
			ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{"Field": Equal("spec.networking.ipFamilies")})))),
		Entry("should succeed for single stack IPv4", &core.Networking{Type: ptr.To(cilium.Type), IPFamilies: []core.IPFamily{core.IPFamilyIPv4}},
			BeEmpty()),
		Entry("should succeed for single stack IPv4", &core.Networking{Type: ptr.To(cilium.Type), IPFamilies: []core.IPFamily{core.IPFamilyIPv6}},
			BeEmpty()),
		Entry("should succeed for dual-stack (v4/v6)", &core.Networking{Type: ptr.To(cilium.Type), IPFamilies: []core.IPFamily{core.IPFamilyIPv4, core.IPFamilyIPv6}},
			BeEmpty()),
		Entry("should succeed for dual-stack (v6/v4", &core.Networking{Type: ptr.To(cilium.Type), IPFamilies: []core.IPFamily{core.IPFamilyIPv6, core.IPFamilyIPv4}},
			BeEmpty()),
	)
})
