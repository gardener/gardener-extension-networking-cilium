# Using the Networking Cilium extension with Gardener as end-user

The [`core.gardener.cloud/v1beta1.Shoot` resource](https://github.com/gardener/gardener/blob/master/example/90-shoot.yaml) declares a `networking` field that is meant to contain network-specific configuration.

In this document we are describing how this configuration looks like for Cilium and provide an example `Shoot` manifest with minimal configuration that you can use to create a cluster.

## Cilium Hubble

Hubble is a fully distributed networking and security observability platform build on top of Cilium and BPF. It is optional and is deployed to the cluster when enabled in the `NetworkConfig`.
If the dashboard is not externally exposed
```
kubectl port-forward -n kube-system deployment/hubble-ui 8081
```
can be used to acess it locally.

## Example `NetworkingConfig` manifest

An example `NetworkingConfig` for the Cilium extension looks as follows:

```yaml
apiVersion: cilium.networking.extensions.gardener.cloud/v1alpha1
kind: NetworkConfig
hubble:
  enabled: true
#debug: false
#psp: true
#tunnel: vxlan
#store: kubernetes
```

## `NetworkingConfig` options

The `hubble.enabled` field describes whether hubble should be deployed into the cluster or not (default).

The `debug` field describes whether you want to run cilium in debug mode or not (default), change this value to `true` to use debug mode.

The `psp` field describes whether `cilium-operator` and `cilium-agent` shall be deployed with pod security policies or not (default).

The `tunnel` field describes the encapsulation mode for communication between nodes. Possible values are `vxlan` (default), `geneve` or `disabled`.

The `store` field describes which backend to use to store the identities. Can be either `etcd` (kvstore) or `kubernetes` (crd) (default).

## Example `Shoot` manifest

Please find below an example `Shoot` manifest with cilium networking configuration:

```yaml
apiVersion: core.gardener.cloud/v1beta1
kind: Shoot
metadata:
  name: aws-cilium
  namespace: garden-dev
spec:
  networking:
    type: cilium
    providerConfig:
      apiVersion: cilium.networking.extensions.gardener.cloud/v1alpha1
      kind: NetworkConfig
      hubble:
        enabled: true
    pods: 100.96.0.0/11
    nodes: 10.250.0.0/16
    services: 100.64.0.0/13
  ...
```

If you would like to see a provider specific shoot example, please check out the documentation of the well-known extensions. A list of them can be found [here](https://github.com/gardener/gardener/tree/master/extensions#infrastructure-provider).
