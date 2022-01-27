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

	"github.com/gardener/gardener/extensions/pkg/controller"
	controllercmd "github.com/gardener/gardener/extensions/pkg/controller/cmd"
	"github.com/gardener/gardener/extensions/pkg/util"
	"github.com/spf13/cobra"
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

		aggOption = controllercmd.NewOptionAggregator(
			generalOpts,
			restOpts,
			mgrOpts,
			ciliumCtrlOpts,
			controllercmd.PrefixOption("healthcheck-", healthCheckCtrlOpts),
			reconcileOpts,
			configFileOpts,
		)
	)

	cmd := &cobra.Command{
		Use: fmt.Sprintf("%s-controller-manager", cilium.Name),

		Run: func(cmd *cobra.Command, args []string) {
			if err := aggOption.Complete(); err != nil {
				controllercmd.LogErrAndExit(err, "Error completing options")
			}
			util.ApplyClientConnectionConfigurationToRESTConfig(configFileOpts.Completed().Config.ClientConnection, restOpts.Completed().Config)

			completedMgrOpts := mgrOpts.Completed().Options()
			completedMgrOpts.ClientDisableCacheFor = []client.Object{
				&corev1.Secret{},    // applied for ManagedResources
				&corev1.ConfigMap{}, // applied for monitoring config
			}

			mgr, err := manager.New(restOpts.Completed().Config, completedMgrOpts)
			if err != nil {
				controllercmd.LogErrAndExit(err, "Could not instantiate manager")
			}

			if err := controller.AddToScheme(mgr.GetScheme()); err != nil {
				controllercmd.LogErrAndExit(err, "Could not update manager scheme")
			}

			if err := ciliuminstall.AddToScheme(mgr.GetScheme()); err != nil {
				controllercmd.LogErrAndExit(err, "Could not update manager scheme")
			}

			useProjectedTokenMount, err := controller.UseServiceAccountTokenVolumeProjection(generalOpts.Completed().GardenerVersion)
			if err != nil {
				controllercmd.LogErrAndExit(err, "could not determine whether service account token volume projection should be used")
			}
			ciliumcontroller.DefaultAddOptions.UseProjectedTokenMount = useProjectedTokenMount

			reconcileOpts.Completed().Apply(&ciliumcontroller.DefaultAddOptions.IgnoreOperationAnnotation)
			ciliumCtrlOpts.Completed().Apply(&ciliumcontroller.DefaultAddOptions.Controller)
			configFileOpts.Completed().ApplyHealthCheckConfig(&healthcheck.AddOptions.HealthCheckConfig)
			healthCheckCtrlOpts.Completed().Apply(&healthcheck.AddOptions.Controller)

			if err := ciliumcontroller.AddToManager(mgr); err != nil {
				controllercmd.LogErrAndExit(err, "Could not add controllers to manager")
			}

			if err := healthcheck.AddToManager(mgr); err != nil {
				controllercmd.LogErrAndExit(err, "Could not add health check controller to manager")
			}

			if err := mgr.Start(ctx); err != nil {
				controllercmd.LogErrAndExit(err, "Error running manager")
			}
		},
	}

	aggOption.AddFlags(cmd.Flags())

	return cmd
}
