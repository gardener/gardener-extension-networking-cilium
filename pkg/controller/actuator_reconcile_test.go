// SPDX-FileCopyrightText: Contributors to the Gardener project
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
)

var _ = DescribeTable("#getIPAMMode",
	func(configMap *corev1.ConfigMap, expected string) {
		Expect(getIPAMMode(configMap)).To(Equal(expected))
	},
	Entry("nil ConfigMap", nil, "kubernetes"),
	Entry("empty ConfigMap", &corev1.ConfigMap{}, "kubernetes"),
	Entry("ConfigMap without ipam key", &corev1.ConfigMap{Data: map[string]string{"other": "value"}}, "kubernetes"),
	Entry("ipam=kubernetes", &corev1.ConfigMap{Data: map[string]string{"ipam": "kubernetes"}}, "kubernetes"),
	Entry("ipam=cluster-pool", &corev1.ConfigMap{Data: map[string]string{"ipam": "cluster-pool"}}, "cluster-pool"),
	Entry("unknown ipam value falls back to kubernetes", &corev1.ConfigMap{Data: map[string]string{"ipam": "unknown"}}, "kubernetes"),
)
