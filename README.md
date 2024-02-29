# [Gardener Extension for cilium Networking](https://gardener.cloud)
[![REUSE status](https://api.reuse.software/badge/github.com/gardener/gardener-extension-networking-cilium)](https://api.reuse.software/info/github.com/gardener/gardener-extension-networking-cilium)
[![CI Build status](https://concourse.ci.gardener.cloud/api/v1/teams/gardener/pipelines/gardener-extension-networking-cilium-master/jobs/master-head-update-job/badge)](https://concourse.ci.gardener.cloud/teams/gardener/pipelines/gardener-extension-networking-cilium-master/jobs/master-head-update-job)
[![Go Report Card](https://goreportcard.com/badge/github.com/gardener/gardener-extension-networking-cilium)](https://goreportcard.com/report/github.com/gardener/gardener-extension-networking-cilium)

This controller operates on the [`Network`](https://github.com/gardener/gardener/blob/master/docs/proposals/03-networking-extensibility.md#gardener-network-extension) resource in the `extensions.gardener.cloud/v1alpha1` API group. It manages those objects that are requesting [cilium Networking](https://cilium.io/) configuration (`.spec.type=cilium`):

```yaml
apiVersion: extensions.gardener.cloud/v1alpha1
kind: Network
metadata:
  name: cilium-network
  namespace: shoot--foo--bar
spec:
  type: cilium
  podCIDR: 10.244.0.0/16
  serviceCIDR:  10.96.0.0/24
  providerConfig:
    apiVersion: cilium.networking.extensions.gardener.cloud/v1alpha1
    kind: NetworkConfig
#    hubble:
#      enabled: true
#    store: kubernetes
```

Please find [a concrete example](example/20-network.yaml) in the `example` folder. All the `cilium` specific configuration
should be configured in the `providerConfig` section. If additional configuration is required, it should be added to
the `networking-cilium` chart in `controllers/networking-cilium/charts/internal/cilium/values.yaml` and corresponding code
parts should be adapted (for example in `controllers/networking-cilium/pkg/charts/utils.go`).

Once the network resource is applied, the `networking-cilium` controller would then create all the necessary `managed-resources` which should be picked
up by the [gardener-resource-manager](https://github.com/gardener/gardener-resource-manager) which will then apply all the
network extensions resources to the shoot cluster.

Finally after successful reconciliation an output similar to the one below should be expected.

```yaml
  status:
    lastOperation:
      description: Successfully reconciled network
      lastUpdateTime: "..."
      progress: 100
      state: Succeeded
      type: Reconcile
    observedGeneration: 1
```

----

## How to start using or developing this extension controller locally

You can run the controller locally on your machine by executing `make start`. Please make sure to have the `kubeconfig` pointed to the cluster you want to connect to.
Static code checks and tests can be executed by running `make verify`. We are using Go modules for Golang package dependency management and [Ginkgo](https://github.com/onsi/ginkgo)/[Gomega](https://github.com/onsi/gomega) for testing.

## Feedback and Support

Feedback and contributions are always welcome. Please report bugs or suggestions as [GitHub issues](https://github.com/gardener/gardener-extension-networking-cilium/issues) or join our [Slack channel #gardener](https://kubernetes.slack.com/messages/gardener) (please invite yourself to the Kubernetes workspace [here](http://slack.k8s.io)).

## Learn more!

Please find further resources about out project here:

* [Our landing page gardener.cloud](https://gardener.cloud/)
* ["Gardener, the Kubernetes Botanist" blog on kubernetes.io](https://kubernetes.io/blog/2018/05/17/gardener/)
* ["Gardener Project Update" blog on kubernetes.io](https://kubernetes.io/blog/2019/12/02/gardener-project-update/)
* [Gardener Extensions Golang library](https://godoc.org/github.com/gardener/gardener/extensions/pkg)
* [GEP-1 (Gardener Enhancement Proposal) on extensibility](https://github.com/gardener/gardener/blob/master/docs/proposals/01-extensibility.md)
* [Extensibility API documentation](https://github.com/gardener/gardener/tree/master/docs/extensions)
* [Docs for `cilium` user](https://docs.cilium.io/)
