// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package shoot

import (
	extensionswebhook "github.com/gardener/gardener/extensions/pkg/webhook"
	"github.com/gardener/gardener/extensions/pkg/webhook/shoot"
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var (
	// DefaultAddOptions are the default AddOptions for AddToManager.
	DefaultAddOptions = AddOptions{}
)

// AddOptions are options to apply when adding the Cilium shoot webhook to the manager.
type AddOptions struct{}

var logger = log.Log.WithName("shoot-webhook")

// AddToManagerWithOptions creates a webhook with the given options and adds it to the manager.
func AddToManagerWithOptions(mgr manager.Manager, opts AddOptions) (*extensionswebhook.Webhook, error) {
	logger.Info("Adding webhook to manager")
	failurePolicyFail := admissionregistrationv1.Fail
	webhook, err := shoot.New(mgr, shoot.Args{
		Types: []extensionswebhook.Type{
			{Obj: &corev1.ConfigMap{}},
			{Obj: &appsv1.DaemonSet{}},
		},
		Mutator:       NewMutator(),
		FailurePolicy: &failurePolicyFail,
	})
	if err != nil {
		return nil, err
	}
	webhook.ObjectSelector = &metav1.LabelSelector{MatchLabels: map[string]string{"k8s-app": "node-local-dns"}}
	return webhook, nil
}

// AddToManager creates a webhook with the default options and adds it to the manager.
func AddToManager(mgr manager.Manager) (*extensionswebhook.Webhook, error) {
	return AddToManagerWithOptions(mgr, DefaultAddOptions)
}
