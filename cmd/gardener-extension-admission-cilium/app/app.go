// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"context"
	"fmt"
	"os"

	controllercmd "github.com/gardener/gardener/extensions/pkg/controller/cmd"
	"github.com/gardener/gardener/extensions/pkg/util"
	webhookcmd "github.com/gardener/gardener/extensions/pkg/webhook/cmd"
	"github.com/gardener/gardener/pkg/apis/core/install"
	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	gardenerhealthz "github.com/gardener/gardener/pkg/healthz"
	"github.com/gardener/gardener/pkg/logger"
	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	componentbaseconfig "k8s.io/component-base/config/v1alpha1"
	"k8s.io/component-base/version"
	"k8s.io/component-base/version/verflag"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/cluster"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	runtimelog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	admissioncmd "github.com/gardener/gardener-extension-networking-cilium/pkg/admission/cmd"
	ciliuminstall "github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium/install"
	"github.com/gardener/gardener-extension-networking-cilium/pkg/cilium"
)

// Name is a constant for the name of this component.
const Name = "gardener-extension-admission-cilium"

// NewAdmissionCommand creates a new command for running a Cilium admission webhook.
func NewAdmissionCommand(ctx context.Context) *cobra.Command {
	var (
		restOpts = &controllercmd.RESTOptions{}
		mgrOpts  = &controllercmd.ManagerOptions{
			LeaderElection:          true,
			LeaderElectionID:        controllercmd.LeaderElectionNameID(Name),
			LeaderElectionNamespace: os.Getenv("LEADER_ELECTION_NAMESPACE"),
			WebhookServerPort:       443,
			HealthBindAddress:       ":8081",
			WebhookCertDir:          "/tmp/admission-cilium-cert",
		}
		// options for the webhook server
		webhookServerOptions = &webhookcmd.ServerOptions{
			Namespace: os.Getenv("WEBHOOK_CONFIG_NAMESPACE"),
		}
		webhookSwitches = admissioncmd.GardenWebhookSwitchOptions()
		webhookOptions  = webhookcmd.NewAddToManagerOptions(
			Name,
			"",
			nil,
			webhookServerOptions,
			webhookSwitches,
		)

		aggOption = controllercmd.NewOptionAggregator(
			restOpts,
			mgrOpts,
			webhookOptions,
		)
	)

	cmd := &cobra.Command{
		Use: fmt.Sprintf("admission-%s", cilium.Name),

		RunE: func(cmd *cobra.Command, args []string) error {
			verflag.PrintAndExitIfRequested()

			log, err := logger.NewZapLogger(logger.InfoLevel, logger.FormatJSON)
			if err != nil {
				return fmt.Errorf("error instantiating zap logger: %w", err)
			}
			runtimelog.SetLogger(log)

			log.Info("Starting "+Name, "version", version.Get())

			if gardenKubeconfig := os.Getenv("GARDEN_KUBECONFIG"); gardenKubeconfig != "" {
				log.Info("Getting rest config for garden from GARDEN_KUBECONFIG", "path", gardenKubeconfig)
				restOpts.Kubeconfig = gardenKubeconfig
			}

			if err := aggOption.Complete(); err != nil {
				return fmt.Errorf("error completing options: %w", err)
			}
			util.ApplyClientConnectionConfigurationToRESTConfig(&componentbaseconfig.ClientConnectionConfiguration{
				QPS:   100.0,
				Burst: 130,
			}, restOpts.Completed().Config)

			managerOptions := mgrOpts.Completed().Options()

			// Operators can enable the source cluster option via SOURCE_CLUSTER environment variable.
			// In-cluster config will be used if no SOURCE_KUBECONFIG is specified.
			//
			// The source cluster is for instance used by Gardener's certificate controller, to maintain certificate
			// secrets in a different cluster ('runtime-garden') than the cluster where the webhook configurations
			// are maintained ('virtual-garden').
			var sourceClusterConfig *rest.Config
			if sourceClusterEnabled := os.Getenv("SOURCE_CLUSTER"); sourceClusterEnabled != "" {
				var err error
				sourceClusterConfig, err = clientcmd.BuildConfigFromFlags("", os.Getenv("SOURCE_KUBECONFIG"))
				if err != nil {
					return err
				}
				managerOptions.LeaderElectionConfig = sourceClusterConfig
			} else {
				// Restrict the cache for secrets to the configured namespace to avoid the need for cluster-wide list/watch permissions.
				managerOptions.Cache = cache.Options{
					ByObject: map[client.Object]cache.ByObject{
						&corev1.Secret{}: {Namespaces: map[string]cache.Config{webhookOptions.Server.Completed().Namespace: {}}},
					},
				}
			}

			mgr, err := manager.New(restOpts.Completed().Config, managerOptions)
			if err != nil {
				return fmt.Errorf("could not instantiate manager: %w", err)
			}

			install.Install(mgr.GetScheme())

			if err := ciliuminstall.AddToScheme(mgr.GetScheme()); err != nil {
				return fmt.Errorf("could not update manager scheme: %w", err)
			}

			var sourceCluster cluster.Cluster
			if sourceClusterConfig != nil {
				log.Info("Configuring source cluster option")
				sourceCluster, err = cluster.New(sourceClusterConfig, func(opts *cluster.Options) {
					opts.Logger = log
					opts.Cache.DefaultNamespaces = map[string]cache.Config{v1beta1constants.GardenNamespace: {}}
				})
				if err != nil {
					return err
				}

				if err := mgr.AddReadyzCheck("source-informer-sync", gardenerhealthz.NewCacheSyncHealthz(sourceCluster.GetCache())); err != nil {
					return err
				}

				if err = mgr.Add(sourceCluster); err != nil {
					return err
				}
			}

			log.Info("Setting up webhook server")
			if _, err := webhookOptions.Completed().AddToManager(ctx, mgr, sourceCluster); err != nil {
				return err
			}

			if err := mgr.AddReadyzCheck("informer-sync", gardenerhealthz.NewCacheSyncHealthz(mgr.GetCache())); err != nil {
				return fmt.Errorf("could not add readycheck for informers: %w", err)
			}

			if err := mgr.AddHealthzCheck("ping", healthz.Ping); err != nil {
				return fmt.Errorf("could not add healthcheck: %w", err)
			}

			if err := mgr.AddReadyzCheck("webhook-server", mgr.GetWebhookServer().StartedChecker()); err != nil {
				return fmt.Errorf("could not add readycheck of webhook to manager: %w", err)
			}

			return mgr.Start(ctx)
		},
	}

	aggOption.AddFlags(cmd.Flags())

	return cmd
}
