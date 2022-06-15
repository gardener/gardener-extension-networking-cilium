// Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package app

import (
	"context"
	"fmt"
	"os"

	ciliuminstall "github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium/install"
	"github.com/gardener/gardener-extension-networking-cilium/pkg/cilium"
	ciliumcmd "github.com/gardener/gardener-extension-networking-cilium/pkg/cmd"
	ciliumcontroller "github.com/gardener/gardener-extension-networking-cilium/pkg/controller"
	"github.com/gardener/gardener-extension-networking-cilium/pkg/healthcheck"
	"github.com/pkg/errors"

	"github.com/gardener/gardener/extensions/pkg/controller"
	controllercmd "github.com/gardener/gardener/extensions/pkg/controller/cmd"
	"github.com/gardener/gardener/extensions/pkg/util"
	webhookcmd "github.com/gardener/gardener/extensions/pkg/webhook/cmd"
	"github.com/spf13/cobra"
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// NewControllerManagerCommand creates a new command for running a Cilium controller.
func NewControllerManagerCommand(ctx context.Context) *cobra.Command {
	var (
		generalOpts = &controllercmd.GeneralOptions{}
		restOpts    = &controllercmd.RESTOptions{}
		mgrOpts     = &controllercmd.ManagerOptions{
			LeaderElection:             true,
			LeaderElectionResourceLock: resourcelock.LeasesResourceLock,
			LeaderElectionID:           controllercmd.LeaderElectionNameID(cilium.Name),
			LeaderElectionNamespace:    os.Getenv("LEADER_ELECTION_NAMESPACE"),
		}
		// options for the networking-cilium controller
		ciliumCtrlOpts = &controllercmd.ControllerOptions{
			MaxConcurrentReconciles: 5,
		}
		reconcileOpts = &controllercmd.ReconcilerOptions{
			IgnoreOperationAnnotation: true,
		}

		// options for the health care controller
		healthCheckCtrlOpts = &controllercmd.ControllerOptions{
			MaxConcurrentReconciles: 5,
		}

		configFileOpts = &ciliumcmd.ConfigOptions{}

		// options for the webhook server
		webhookServerOptions = &webhookcmd.ServerOptions{
			Namespace: os.Getenv("WEBHOOK_CONFIG_NAMESPACE"),
		}

		webhookSwitches = ciliumcmd.WebhookSwitchOptions()
		webhookOptions  = webhookcmd.NewAddToManagerOptions(cilium.Name, webhookServerOptions, webhookSwitches)

		aggOption = controllercmd.NewOptionAggregator(
			generalOpts,
			restOpts,
			mgrOpts,
			ciliumCtrlOpts,
			controllercmd.PrefixOption("healthcheck-", healthCheckCtrlOpts),
			reconcileOpts,
			configFileOpts,
			webhookOptions,
		)
	)

	cmd := &cobra.Command{
		Use: fmt.Sprintf("%s-controller-manager", cilium.Name),

		RunE: func(cmd *cobra.Command, args []string) error {
			if err := aggOption.Complete(); err != nil {
				return fmt.Errorf("error completing options: %w", err)
			}
			util.ApplyClientConnectionConfigurationToRESTConfig(configFileOpts.Completed().Config.ClientConnection, restOpts.Completed().Config)

			completedMgrOpts := mgrOpts.Completed().Options()
			completedMgrOpts.ClientDisableCacheFor = []client.Object{
				&corev1.Secret{},    // applied for ManagedResources
				&corev1.ConfigMap{}, // applied for monitoring config
			}

			mgr, err := manager.New(restOpts.Completed().Config, completedMgrOpts)
			if err != nil {
				return fmt.Errorf("could not instantiate manager: %w", err)
			}

			if err := controller.AddToScheme(mgr.GetScheme()); err != nil {
				return fmt.Errorf("could not update manager scheme: %w", err)
			}

			if err := ciliuminstall.AddToScheme(mgr.GetScheme()); err != nil {
				return fmt.Errorf("could not update manager scheme: %w", err)
			}

			reconcileOpts.Completed().Apply(&ciliumcontroller.DefaultAddOptions.IgnoreOperationAnnotation)
			ciliumCtrlOpts.Completed().Apply(&ciliumcontroller.DefaultAddOptions.Controller)
			configFileOpts.Completed().ApplyHealthCheckConfig(&healthcheck.AddOptions.HealthCheckConfig)
			healthCheckCtrlOpts.Completed().Apply(&healthcheck.AddOptions.Controller)

			_, shootWebhooks, err := webhookOptions.Completed().AddToManager(ctx, mgr)
			if err != nil {
				return errors.Wrap(err, "Could not add webhooks to manager")
			}

			ciliumcontroller.DefaultAddOptions.ShootWebhooks = shootWebhooks

			// Update shoot webhook configuration in case the webhook server port has changed.
			if err := mgr.Add(&shootWebhookReconciler{
				client:            mgr.GetClient(),
				webhookServerPort: mgr.GetWebhookServer().Port,
				shootWebhooks:     shootWebhooks,
			}); err != nil {
				return fmt.Errorf("error adding runnable for reconciling shoot webhooks in all namespaces: %w", err)
			}

			if err := ciliumcontroller.AddToManager(mgr); err != nil {
				return fmt.Errorf("could not add controllers to manager: %w", err)
			}

			if err := healthcheck.AddToManager(mgr); err != nil {
				return fmt.Errorf("could not add health check controller to manager: %w", err)
			}

			if err := mgr.Start(ctx); err != nil {
				return fmt.Errorf("error running manager: %w", err)
			}

			return nil
		},
	}

	aggOption.AddFlags(cmd.Flags())

	return cmd
}

type shootWebhookReconciler struct {
	client            client.Client
	webhookServerPort int
	shootWebhooks     []admissionregistrationv1.MutatingWebhook
}

func (s *shootWebhookReconciler) NeedLeaderElection() bool {
	return true
}

func (s *shootWebhookReconciler) Start(ctx context.Context) error {
	return ciliumcontroller.ReconcileShootWebhooksForAllNamespaces(ctx, s.client, cilium.Name, cilium.Type, s.webhookServerPort, s.shootWebhooks)
}
