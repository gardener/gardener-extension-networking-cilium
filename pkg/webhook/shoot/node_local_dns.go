// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package shoot

import (
	"context"
	"regexp"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func (m *mutator) mutateNodeLocalDNSConfigMap(ctx context.Context, configmap *corev1.ConfigMap) error {
	if configmap.Data == nil {
		configmap.Data = make(map[string]string, 1)
	}

	re1 := regexp.MustCompile(`bind.*`)
	configmap.Data["Corefile"] = re1.ReplaceAllString(configmap.Data["Corefile"], "bind 0.0.0.0")
	re2 := regexp.MustCompile(`health.*(:[0-9]+)`)
	configmap.Data["Corefile"] = re2.ReplaceAllString(configmap.Data["Corefile"], "health $1")

	return nil
}

func (m *mutator) mutateNodeLocalDNSDaemonSet(ctx context.Context, daemonset *appsv1.DaemonSet) error {
	if daemonset.Spec.Template.Spec.HostNetwork {
		daemonset.Spec.Template.Spec.HostNetwork = false
	}

	ciliumArgs := []string{"-skipteardown=true", "-setupinterface=false", "-setupiptables=false"}
	for k, v := range daemonset.Spec.Template.Spec.Containers {
		if v.Name == "node-cache" {
			daemonset.Spec.Template.Spec.Containers[k].Args = append(daemonset.Spec.Template.Spec.Containers[k].Args, ciliumArgs...)
			daemonset.Spec.Template.Spec.Containers[k].LivenessProbe.HTTPGet.Host = ""
			break
		}
	}
	return nil

}
