<p>Packages:</p>
<ul>
<li>
<a href="#cilium.networking.extensions.config.gardener.cloud%2fv1alpha1">cilium.networking.extensions.config.gardener.cloud/v1alpha1</a>
</li>
</ul>
<h2 id="cilium.networking.extensions.config.gardener.cloud/v1alpha1">cilium.networking.extensions.config.gardener.cloud/v1alpha1</h2>
<p>
<p>Package v1alpha1 contains the Cilium networking configuration API resources.</p>
</p>
Resource Types:
<ul></ul>
<h3 id="cilium.networking.extensions.config.gardener.cloud/v1alpha1.ControllerConfiguration">ControllerConfiguration
</h3>
<p>
<p>ControllerConfiguration defines the configuration for the Cilium networking extension.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>clientConnection</code></br>
<em>
<a href="https://godoc.org/k8s.io/component-base/config/v1alpha1#ClientConnectionConfiguration">
Kubernetes v1alpha1.ClientConnectionConfiguration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ClientConnection specifies the kubeconfig file and client connection
settings for the proxy server to use when communicating with the apiserver.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<p><em>
Generated with <a href="https://github.com/ahmetb/gen-crd-api-reference-docs">gen-crd-api-reference-docs</a>
</em></p>
