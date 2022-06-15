// Copyright (c) 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

	re := regexp.MustCompile(`bind.*`)
	configmap.Data["Corefile"] = re.ReplaceAllString(configmap.Data["Corefile"], "bind 0.0.0.0")
	re = regexp.MustCompile(`health.*`)
	configmap.Data["Corefile"] = re.ReplaceAllString(configmap.Data["Corefile"], "health")

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
