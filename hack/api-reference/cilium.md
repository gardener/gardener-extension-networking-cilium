<p>Packages:</p>
<ul>
<li>
<a href="#cilium.networking.extensions.gardener.cloud%2fv1alpha1">cilium.networking.extensions.gardener.cloud/v1alpha1</a>
</li>
</ul>
<h2 id="cilium.networking.extensions.gardener.cloud/v1alpha1">cilium.networking.extensions.gardener.cloud/v1alpha1</h2>
<p>
<p>Package v1alpha1 contains the configuration of the Cilium Network Extension.</p>
</p>
Resource Types:
<ul><li>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>
</li></ul>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig
</h3>
<p>
<p>NetworkConfig is a struct representing the configmap for the cilium
networking plugin</p>
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
<code>apiVersion</code></br>
string</td>
<td>
<code>
cilium.networking.extensions.gardener.cloud/v1alpha1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
string
</td>
<td><code>NetworkConfig</code></td>
</tr>
<tr>
<td>
<code>debug</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.Debug">
Debug
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Debug configuration to be enabled or not</p>
</td>
</tr>
<tr>
<td>
<code>googleGKE</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.GoogleGKE">
GoogleGKE
</a>
</em>
</td>
<td>
<p>GoogleGKE enabled or not</p>
</td>
</tr>
<tr>
<td>
<code>nodeInit</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NodeInit">
NodeInit
</a>
</em>
</td>
<td>
<p>NodeInit config for cilium</p>
</td>
</tr>
<tr>
<td>
<code>operator</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.Operator">
Operator
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Operator enabled or not</p>
</td>
</tr>
<tr>
<td>
<code>clustername</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.ClusterName">
ClusterName
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ClusterName for cluster</p>
</td>
</tr>
<tr>
<td>
<code>preflight</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.Preflight">
Preflight
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Preflight configuration</p>
</td>
</tr>
<tr>
<td>
<code>cleanstate</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.CleanState">
CleanState
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Cleanstate configuration</p>
</td>
</tr>
<tr>
<td>
<code>cleanbpfstate</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.CleanBpfState">
CleanBpfState
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>CleanBpfstate configuraton</p>
</td>
</tr>
<tr>
<td>
<code>bpf</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.Bpf">
Bpf
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Bpf tunning configuraton</p>
</td>
</tr>
<tr>
<td>
<code>hostservices</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.HostServices">
HostServices
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>HostServices configuration</p>
</td>
</tr>
<tr>
<td>
<code>sockops</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.SockOps">
SockOps
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>SockOps configuration</p>
</td>
</tr>
<tr>
<td>
<code>installiptablerules</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.InstallIPTableRules">
InstallIPTableRules
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>InstallIPTableRules configuration</p>
</td>
</tr>
<tr>
<td>
<code>prometheus</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.Prometheus">
Prometheus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Prometheus configuration</p>
</td>
</tr>
<tr>
<td>
<code>synchronizeK8sNodes</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.SynchronizeK8sNodes">
SynchronizeK8sNodes
</a>
</em>
</td>
<td>
<p>SynchronizeK8sNodes configuration</p>
</td>
</tr>
<tr>
<td>
<code>remoteNodeIdentity</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.RemoteNodeIdentity">
RemoteNodeIdentity
</a>
</em>
</td>
<td>
<p>RemoteNodeIdentity configuration</p>
</td>
</tr>
<tr>
<td>
<code>tlsSecretsBackend</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.TLSsecretsBackend">
TLSsecretsBackend
</a>
</em>
</td>
<td>
<p>TLSsecretsBackend configuration</p>
</td>
</tr>
<tr>
<td>
<code>wellKnownIdentities</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.WellKnownIdentities">
WellKnownIdentities
</a>
</em>
</td>
<td>
<p>WellKnownIdentities configuration</p>
</td>
</tr>
<tr>
<td>
<code>cni</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.CNI">
CNI
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>CNI configuration</p>
</td>
</tr>
<tr>
<td>
<code>psp</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.Psp">
Psp
</a>
</em>
</td>
<td>
<p>PSP configuration</p>
</td>
</tr>
<tr>
<td>
<code>agent</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.Agent">
Agent
</a>
</em>
</td>
<td>
<p>Agent configuration</p>
</td>
</tr>
<tr>
<td>
<code>config</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.ConfigEnable">
ConfigEnable
</a>
</em>
</td>
<td>
<p>Config enable configuration</p>
</td>
</tr>
<tr>
<td>
<code>enableXTSocketFallback</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.XTSocketFallback">
XTSocketFallback
</a>
</em>
</td>
<td>
<p>EnableXTSocketFallback configuration</p>
</td>
</tr>
<tr>
<td>
<code>k8sIPv4PodCIDR</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.K8SIPv4PodCIDR">
K8SIPv4PodCIDR
</a>
</em>
</td>
<td>
<p>K8SIPv4PodCIDR is required or not</p>
</td>
</tr>
<tr>
<td>
<code>policyAuditMode</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.PolicyAuditMode">
PolicyAuditMode
</a>
</em>
</td>
<td>
<p>PolicyAuditMode configuration</p>
</td>
</tr>
<tr>
<td>
<code>daemonRunPath</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.DaemonRunPath">
DaemonRunPath
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>DaemonRunPath configuration</p>
</td>
</tr>
<tr>
<td>
<code>eni</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.Eni">
Eni
</a>
</em>
</td>
<td>
<p>Eni configuration</p>
</td>
</tr>
<tr>
<td>
<code>endpointRoutes</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.EndpointRoutes">
EndpointRoutes
</a>
</em>
</td>
<td>
<p>EndpointRoutes configuration</p>
</td>
</tr>
<tr>
<td>
<code>azure</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.Azure">
Azure
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Azure configuration</p>
</td>
</tr>
<tr>
<td>
<code>operatorprometheus</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.OperatorPrometheus">
OperatorPrometheus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>OperatorPrometheus configuration</p>
</td>
</tr>
<tr>
<td>
<code>kubeproxy</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.KubeProxy">
KubeProxy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>KubeProxy configuration to be enabled or not</p>
</td>
</tr>
<tr>
<td>
<code>kubeproxyreplacementmode</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.KubeProxyReplacementMode">
KubeProxyReplacementMode
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>KubeProxyReplacementMode configuration. It should be
&lsquo;probe&rsquo;, &lsquo;strict&rsquo;, &lsquo;partial&rsquo;
If KubeProxy is disabled it is required to configure the
KubeProxyReplacementMode</p>
</td>
</tr>
<tr>
<td>
<code>etcdconfig</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.EtcdConfig">
EtcdConfig
</a>
</em>
</td>
<td>
<p>Etcd configuration. Configure managed or not and
then configure the etcd-secrets.</p>
</td>
</tr>
<tr>
<td>
<code>masquerade</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.Masquerade">
Masquerade
</a>
</em>
</td>
<td>
<p>Masquerade configuraton to be enabled or not</p>
</td>
</tr>
<tr>
<td>
<code>ipvlan</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.IPvlan">
IPvlan
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>IPvlan configuraton to be enabled or not</p>
</td>
</tr>
<tr>
<td>
<code>flannel</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.Flannel">
Flannel
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Flannel configuraton to be enabled or not</p>
</td>
</tr>
<tr>
<td>
<code>autoDirectNodeRoutes</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.AutoDirectNodeRoutes">
AutoDirectNodeRoutes
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>AutoDirectNodeRoutes configuraton to be enabled or not</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.Encryption">
Encryption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Encryption configuraton</p>
</td>
</tr>
<tr>
<td>
<code>hubble</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.Hubble">
Hubble
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Hubble configuration to be enabled or not</p>
</td>
</tr>
<tr>
<td>
<code>nodeport</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.Nodeport">
Nodeport
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Nodeport enablement for cilium</p>
</td>
</tr>
<tr>
<td>
<code>ipv4</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.IPv4">
IPv4
</a>
</em>
</td>
<td>
<p>IPv4 configuration to be enabled or not</p>
</td>
</tr>
<tr>
<td>
<code>ipv6</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.IPv6">
IPv6
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>IPv6 configuration to be enabled or not</p>
</td>
</tr>
<tr>
<td>
<code>externalip</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.ExternalIP">
ExternalIP
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ExternalIP configuration to be enabled or not</p>
</td>
</tr>
<tr>
<td>
<code>ipam-mode</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.IPAMOptionMode">
IPAMOptionMode
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>IPAMOptionMode configuration, it should be either
&lsquo;crd&rsquo; or &lsquo;eni&rsquo; or &lsquo;default&rsquo; to be enabled or not</p>
</td>
</tr>
<tr>
<td>
<code>identityAllocationMode</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.IdentityAllocationModeOption">
IdentityAllocationModeOption
</a>
</em>
</td>
<td>
<p>IdentityAllocationModeOption configuration, it should be either
&lsquo;crd&rsquo; or &lsquo;kvstore&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>tunnel</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.TunnelMode">
TunnelMode
</a>
</em>
</td>
<td>
<p>TunnelMode configuration, it should be &lsquo;vxlan&rsquo;, &lsquo;geneve&rsquo; or &lsquo;disabled&rsquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.Agent">Agent
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>Agent configuration for cilium</p>
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
<code>agentEnabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>Agent is required or not</p>
</td>
</tr>
<tr>
<td>
<code>agentSleepAfterInitEnabled</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>agentKeepDeprecatedLabelsEnabled</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.AutoDirectNodeRoutes">AutoDirectNodeRoutes
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>AutoDirectNodeRoutes configuration for cilium</p>
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
<code>autoDirectNodeRoutesEnabled</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.Azure">Azure
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>Azure configuration for cilium</p>
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
<code>azureEnabled</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>azureResourceGroup</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>azureSubscriptionID</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>azureTenantID</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>azureClientID</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>azureClientSecret</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.Bpf">Bpf
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>Bpf configuration for cilium</p>
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
<code>waitForMount</code></br>
<em>
bool
</em>
</td>
<td>
<p>Bpf tunning for cilium.</p>
</td>
</tr>
<tr>
<td>
<code>preallocateMaps</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>ctTcpMax</code></br>
<em>
int32
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>ctAnyMax</code></br>
<em>
int32
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>monitorAggregation</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>monitorFlags</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.CNI">CNI
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>CNI configuration for cilium</p>
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
<code>cniInstallEnabled</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>cniChainingMode</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>cniCustomConfigEnabled</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>cniConfigPath</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>cnibinPath</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>cniConfigMapKey</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>cniConfigMap</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>cniConfFileMountPath</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>cniHostConfDirMountPath</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.CleanBpfState">CleanBpfState
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>CleanBpfState configuration for cilium</p>
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
<code>cleanBpfStateEnabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>CleanBpfStateEnabled is used to define whether CleanBpfState is required or not on Startup.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.CleanState">CleanState
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>CleanState configuration for cilium</p>
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
<code>cleanStateEnabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>CleanStateEnabled is used to define whether CleanState is required or not on Startup.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.ClusterName">ClusterName
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>ClusterName configuration for cilium</p>
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
<code>clusterName</code></br>
<em>
string
</em>
</td>
<td>
<p>ClusterName is used to define the name of the cluster.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.ConfigEnable">ConfigEnable
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>Config enable option for cilium</p>
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
<code>configEnabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>ConfigEnable is required or not</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.DaemonRunPath">DaemonRunPath
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>DaemonRunPath option for cilium</p>
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
<code>daemonRunPath</code></br>
<em>
string
</em>
</td>
<td>
<p>DaemonRunPath is required or not</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.Debug">Debug
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>Debug level option for cilium</p>
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
<code>debugEnabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>Enabled is used to define whether Debug is required or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.Encryption">Encryption
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>Encryption related configuration for cilium cluster</p>
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
<code>encryptionEnabled</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>keyFile</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>mountPath</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>secretName</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>nodeEncryption</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>interface</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.EndpointRoutes">EndpointRoutes
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>EndpointRoutes configuration for cilium</p>
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
<code>endpointRoutesEnabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>EndpointRoutes is required or not</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.Eni">Eni
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>Eni configuration for cilium</p>
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
<code>eniEnabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>Eni is required or not</p>
</td>
</tr>
<tr>
<td>
<code>egressMasqueradeInterfaces</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.EtcdConfig">EtcdConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>EtcdConfig related configuration for cilium</p>
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
<code>etcdEnabled</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>etcdManaged</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>etcdsslEnabled</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>etcdEndPoints</code></br>
<em>
[]string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>caFile</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>certFile</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>keyFile</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.ExternalIP">ExternalIP
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>ExternalIPs configuration for cilium</p>
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
<code>externalipEnabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>ExternalIPenabled is used to define whether ExternalIP address is required or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.Flannel">Flannel
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>Flannel configuration for cilium</p>
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
<code>flannelEnabled</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>masterDevice</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>uninstallOnExit</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.GoogleGKE">GoogleGKE
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>GoogleGKE  option for cilium</p>
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
<code>googleGkeEnabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>GoogleGKE is required or not</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.HostServices">HostServices
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>HostServices configuration for cilium</p>
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
<code>hostServicesEnabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>HostServices is used to define whether IPv4 address is required or not.</p>
</td>
</tr>
<tr>
<td>
<code>protocols</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.Hubble">Hubble
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>Hubble enablement for cilium</p>
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
<code>hubbleEnabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>HubbleEnabled is used to define whether Hubble is required or not.</p>
</td>
</tr>
<tr>
<td>
<code>listenAddresses</code></br>
<em>
[]string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>eventQueueSize</code></br>
<em>
int32
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>flowBufferSize</code></br>
<em>
int32
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>metricServer</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>metrics</code></br>
<em>
[]string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>uiEnabled</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.IPAMOptionMode">IPAMOptionMode
(<code>string</code> alias)</p></h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
</p>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.IPVlanMode">IPVlanMode
(<code>string</code> alias)</p></h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.IPvlan">IPvlan</a>)
</p>
<p>
</p>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.IPv4">IPv4
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>IPv4 configuration for cilium</p>
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
<code>ipv4Enabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>IPv4Enabled is used to define whether IPv4 address is required or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.IPv6">IPv6
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>IPv6 configuration for cilium</p>
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
<code>ipv6Enabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>IPv6Enabled is used to define whether IPv6 address is required or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.IPvlan">IPvlan
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>IPvlan configuration for cilium</p>
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
<code>ipvlanenabled</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>masterDevice</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>ipvlanMode</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.IPVlanMode">
IPVlanMode
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.IdentityAllocationModeOption">IdentityAllocationModeOption
(<code>string</code> alias)</p></h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
</p>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.InstallIPTableRules">InstallIPTableRules
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>InstallIPTableRules configuration for cilium</p>
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
<code>installIptableRulesEnabled</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.K8SIPv4PodCIDR">K8SIPv4PodCIDR
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>K8SIPv4PodCIDR configuration for cilium</p>
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
<code>k8sIPv4PodCIDREnabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>K8SIPv4PodCIDR is required or not</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.KubeProxy">KubeProxy
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>KubeProxy configuration for cilium</p>
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
<code>kubeProxyEnabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>KubeProxyEnabled is used to set if KubeProxy is required or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.KubeProxyReplacementMode">KubeProxyReplacementMode
(<code>string</code> alias)</p></h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
</p>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.Masquerade">Masquerade
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>Masquerade configuration for cilium</p>
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
<code>masqueradeEnabled</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkStatus">NetworkStatus
</h3>
<p>
<p>NetworkStatus contains information about created Network resources.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.NodeInit">NodeInit
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>NodeInit configuration for cilium</p>
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
<code>nodeInitEnabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>NodeInit is required or not</p>
</td>
</tr>
<tr>
<td>
<code>bootStrapFile</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>restartPods</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>reconfigureKubelet</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>removeCbrBridge</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>revertReconfigureKubelet</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.Nodeport">Nodeport
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>Nodeport enablement for cilium</p>
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
<code>nodeportEnabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>NodeportEnabled is used to define whether Nodeport is required or not.</p>
</td>
</tr>
<tr>
<td>
<code>nodeportMode</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>nodeportDevice</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>nodeportRange</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.Operator">Operator
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>Operator configuration for cilium</p>
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
<code>operatorEnabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>OperatorEnabled is used to define whether Operator is required or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.OperatorPrometheus">OperatorPrometheus
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>OperatorPrometheus configuration for cilium</p>
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
<code>operatorprometheusEnabled</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>port</code></br>
<em>
int
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.PolicyAuditMode">PolicyAuditMode
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>PolicyAuditMode configuration for cilium</p>
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
<code>policyAuditModeEnabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>PolicyAuditMode is required or not</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.Preflight">Preflight
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>Preflight configuration for cilium</p>
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
<code>preflightEnabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>Preflight is used to define whether Preflight is required or not.</p>
</td>
</tr>
<tr>
<td>
<code>precacheToFQDNS</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.Prometheus">Prometheus
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>Prometheus configuration for cilium</p>
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
<code>prometheusEnabled</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>port</code></br>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>serviceMonitorEnabled</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.Psp">Psp
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>Psp configuration for cilium</p>
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
<code>pspEnabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>Psp is required or not</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.RemoteNodeIdentity">RemoteNodeIdentity
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>RemoteNodeIdentity configuration for cilium</p>
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
<code>remoteNodeIdentityEnabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>RemoteNodeIdentity is required or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.SockOps">SockOps
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>SockOps configuration for cilium</p>
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
<code>sockOpsEnabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>SockOpsEnabled is used to define whether SockOps is required or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.SynchronizeK8sNodes">SynchronizeK8sNodes
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>SynchronizeK8sNodes configuration for cilium</p>
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
<code>synchronizeK8sNodesEnabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>SynchronizeK8sNodes is required or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.TLSsecretsBackend">TLSsecretsBackend
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>TLSsecretsBackend configuration for cilium</p>
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
<code>tlsSecretsBackend</code></br>
<em>
string
</em>
</td>
<td>
<p>TLSsecretsBackend is required or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.TunnelMode">TunnelMode
(<code>string</code> alias)</p></h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
</p>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.WellKnownIdentities">WellKnownIdentities
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>WellKnownIdentities configuration for cilium</p>
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
<code>wellKnownIdentitiesEnabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>WellKnownIdentities is required or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.XTSocketFallback">XTSocketFallback
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>XTSocketFallback configuration for cilium</p>
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
<code>xtSocketFallbackEnabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>XTSocketFallback is required or not</p>
</td>
</tr>
</tbody>
</table>
<hr/>
