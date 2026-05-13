<p>Packages:</p>
<ul>
<li>
<a href="#cilium.networking.extensions.gardener.cloud%2fv1alpha1">cilium.networking.extensions.gardener.cloud/v1alpha1</a>
</li>
</ul>

<h2 id="cilium.networking.extensions.gardener.cloud/v1alpha1">cilium.networking.extensions.gardener.cloud/v1alpha1</h2>
<p>

</p>

<h3 id="bgpcontrolplane">BGPControlPlane
</h3>


<p>
(<em>Appears on:</em><a href="#networkconfig">NetworkConfig</a>)
</p>

<p>
BGPControlPlane enables the BGP Control Plane
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
<code>enabled</code></br>
<em>
boolean
</em>
</td>
<td>
<p></p>
</td>
</tr>

</tbody>
</table>


<h3 id="bpfsocketlbhostnsonly">BPFSocketLBHostnsOnly
</h3>


<p>
(<em>Appears on:</em><a href="#networkconfig">NetworkConfig</a>)
</p>

<p>
BPFSocketLBHostnsOnly enablement for cilium
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
<code>enabled</code></br>
<em>
boolean
</em>
</td>
<td>
<p></p>
</td>
</tr>

</tbody>
</table>


<h3 id="cni">CNI
</h3>


<p>
(<em>Appears on:</em><a href="#networkconfig">NetworkConfig</a>)
</p>

<p>
CNI configuration for cilium
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
<code>exclusive</code></br>
<em>
boolean
</em>
</td>
<td>
<p>false indicates that cilium will not overwrite its CNI configuration.</p>
</td>
</tr>

</tbody>
</table>


<h3 id="egressgateway">EgressGateway
</h3>


<p>
(<em>Appears on:</em><a href="#networkconfig">NetworkConfig</a>)
</p>

<p>
EgressGateway enablement for cilium
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
<code>enabled</code></br>
<em>
boolean
</em>
</td>
<td>
<p></p>
</td>
</tr>

</tbody>
</table>


<h3 id="encryption">Encryption
</h3>


<p>
(<em>Appears on:</em><a href="#networkconfig">NetworkConfig</a>)
</p>

<p>

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
<code>mode</code></br>
<em>
<a href="#encryptionmode">EncryptionMode</a>
</em>
</td>
<td>
<p></p>
</td>
</tr>
<tr>
<td>
<code>enabled</code></br>
<em>
boolean
</em>
</td>
<td>
<p></p>
</td>
</tr>
<tr>
<td>
<code>nodeToNodeEnabled</code></br>
<em>
boolean
</em>
</td>
<td>
<p></p>
</td>
</tr>
<tr>
<td>
<code>strictMode</code></br>
<em>
boolean
</em>
</td>
<td>
<p>StrictMode enables StrictMode encryption.<br />Must be used with Mode "wireguard"<br />See https://docs.cilium.io/en/stable/security/network/encryption/#egress-traffic-to-not-yet-discovered-remote-endpoints-may-be-unencrypted for more information</p>
</td>
</tr>

</tbody>
</table>


<h3 id="encryptionmode">EncryptionMode
</h3>
<p><em>Underlying type: string</em></p>


<p>
(<em>Appears on:</em><a href="#encryption">Encryption</a>)
</p>

<p>

</p>


<h3 id="hubble">Hubble
</h3>


<p>
(<em>Appears on:</em><a href="#networkconfig">NetworkConfig</a>)
</p>

<p>
Hubble enablement for cilium
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
<code>enabled</code></br>
<em>
boolean
</em>
</td>
<td>
<p>Enabled defines whether hubble will be enabled for the cluster.</p>
</td>
</tr>

</tbody>
</table>


<h3 id="ipv4">IPv4
</h3>


<p>
(<em>Appears on:</em><a href="#networkconfig">NetworkConfig</a>)
</p>

<p>
IPv4 enablement for cilium
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
<code>enabled</code></br>
<em>
boolean
</em>
</td>
<td>
<p>Enabled indicates whether IPv4 is enabled or not.</p>
</td>
</tr>

</tbody>
</table>


<h3 id="ipv6">IPv6
</h3>


<p>
(<em>Appears on:</em><a href="#networkconfig">NetworkConfig</a>)
</p>

<p>
IPv6 enablement for cilium
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
<code>enabled</code></br>
<em>
boolean
</em>
</td>
<td>
<p>Enabled indicates whether IPv6 is enabled or not.</p>
</td>
</tr>

</tbody>
</table>


<h3 id="identityallocationmode">IdentityAllocationMode
</h3>
<p><em>Underlying type: string</em></p>


<p>
IdentityAllocationMode selects how identities are shared between cilium
nodes by setting how they are stored. The options are "crd" or "kvstore".
</p>


<h3 id="kubeproxy">KubeProxy
</h3>


<p>
(<em>Appears on:</em><a href="#networkconfig">NetworkConfig</a>)
</p>

<p>
KubeProxy configuration for cilium
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
<code>k8sServiceHost</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceHost specify the controlplane node IP Address.</p>
</td>
</tr>
<tr>
<td>
<code>k8sServicePort</code></br>
<em>
integer
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServicePort specify the kube-apiserver port number.</p>
</td>
</tr>

</tbody>
</table>


<h3 id="kubeproxyreplacementmode">KubeProxyReplacementMode
</h3>
<p><em>Underlying type: string</em></p>


<p>
KubeProxyReplacementMode defines which mode should kube-proxy run in.
More infromation here: https://docs.cilium.io/en/v1.7/gettingstarted/kubeproxy-free/
</p>


<h3 id="l2announcements">L2Announcements
</h3>


<p>
(<em>Appears on:</em><a href="#networkconfig">NetworkConfig</a>)
</p>

<p>
L2Announcements enables the L2 announcements feature.
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
<code>enabled</code></br>
<em>
boolean
</em>
</td>
<td>
<p>Enabled defines whether L2 announcements is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>leaseDuration</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.33/#duration-v1-meta">Duration</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>LeaseDuration is the maximum duration of the lease.</p>
</td>
</tr>
<tr>
<td>
<code>leaseRenewDeadline</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.33/#duration-v1-meta">Duration</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>LeaseRenewDeadline is the duration that the current leader will retry<br />refreshing the lease before giving up.</p>
</td>
</tr>
<tr>
<td>
<code>leaseRetryPeriod</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.33/#duration-v1-meta">Duration</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>LeaseRetryPeriod is the duration the clients should wait between retries.</p>
</td>
</tr>

</tbody>
</table>


<h3 id="loadbalancingmode">LoadBalancingMode
</h3>
<p><em>Underlying type: string</em></p>


<p>
(<em>Appears on:</em><a href="#networkconfig">NetworkConfig</a>)
</p>

<p>
LoadBalancingMode defines what load balancing mode to use for Cilium.
</p>


<h3 id="networkconfig">NetworkConfig
</h3>


<p>
NetworkConfig is a struct representing the configmap for the cilium
networking plugin
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
<code>debug</code></br>
<em>
boolean
</em>
</td>
<td>
<em>(Optional)</em>
<p>Debug configuration to be enabled or not</p>
</td>
</tr>
<tr>
<td>
<code>kubeproxy</code></br>
<em>
<a href="#kubeproxy">KubeProxy</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>KubeProxy configuration to be enabled or not</p>
</td>
</tr>
<tr>
<td>
<code>hubble</code></br>
<em>
<a href="#hubble">Hubble</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Hubble configuration to be enabled or not</p>
</td>
</tr>
<tr>
<td>
<code>tunnel</code></br>
<em>
<a href="#tunnelmode">TunnelMode</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>TunnelMode configuration, it should be 'vxlan', 'geneve' or 'disabled'</p>
</td>
</tr>
<tr>
<td>
<code>store</code></br>
<em>
<a href="#store">Store</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Store can be either Kubernetes or etcd.</p>
</td>
</tr>
<tr>
<td>
<code>ipv4</code></br>
<em>
<a href="#ipv4">IPv4</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Enable IPv4</p>
</td>
</tr>
<tr>
<td>
<code>ipv6</code></br>
<em>
<a href="#ipv6">IPv6</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Enable IPv6</p>
</td>
</tr>
<tr>
<td>
<code>bpfSocketLBHostnsOnly</code></br>
<em>
<a href="#bpfsocketlbhostnsonly">BPFSocketLBHostnsOnly</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>BPFSocketLBHostnsOnly flag to be enabled or not</p>
</td>
</tr>
<tr>
<td>
<code>cni</code></br>
<em>
<a href="#cni">CNI</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>CNI configuration for cilium</p>
</td>
</tr>
<tr>
<td>
<code>egressGateway</code></br>
<em>
<a href="#egressgateway">EgressGateway</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>EgressGateway enablement for cilium</p>
</td>
</tr>
<tr>
<td>
<code>mtu</code></br>
<em>
integer
</em>
</td>
<td>
<em>(Optional)</em>
<p>MTU overwrites the auto-detected MTU of the underlying network</p>
</td>
</tr>
<tr>
<td>
<code>devices</code></br>
<em>
string array
</em>
</td>
<td>
<em>(Optional)</em>
<p>Devices is the list of devices facing cluster/external network</p>
</td>
</tr>
<tr>
<td>
<code>directRoutingDevice</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>DirectRoutingDevice is the device used for direct routing between Cilium nodes</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancingMode</code></br>
<em>
<a href="#loadbalancingmode">LoadBalancingMode</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>LoadBalancingMode configuration, it should be 'snat', 'dsr' or 'hybrid'</p>
</td>
</tr>
<tr>
<td>
<code>l2Announcements</code></br>
<em>
<a href="#l2announcements">L2Announcements</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>L2Announcements enables the L2 announcements feature</p>
</td>
</tr>
<tr>
<td>
<code>ipv4NativeRoutingCIDREnabled</code></br>
<em>
boolean
</em>
</td>
<td>
<em>(Optional)</em>
<p>IPv4NativeRoutingCIDRMode will set the ipv4 native routing cidr from the network configs node's cidr if enabled.</p>
</td>
</tr>
<tr>
<td>
<code>ipv6NativeRoutingCIDREnabled</code></br>
<em>
boolean
</em>
</td>
<td>
<em>(Optional)</em>
<p>IPv6NativeRoutingCIDRMode will set the ipv6 native routing cidr from the network configs node's cidr if enabled.</p>
</td>
</tr>
<tr>
<td>
<code>overlay</code></br>
<em>
<a href="#overlay">Overlay</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Overlay enables the network overlay</p>
</td>
</tr>
<tr>
<td>
<code>snatToUpstreamDNS</code></br>
<em>
<a href="#snattoupstreamdns">SnatToUpstreamDNS</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>SnatToUpstreamDNS enables the masquerading of packets to the upstream dns server</p>
</td>
</tr>
<tr>
<td>
<code>snatOutOfCluster</code></br>
<em>
<a href="#snatoutofcluster">SnatOutOfCluster</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>SnatOutOfCluster enables the masquerading of packets outside of the cluster</p>
</td>
</tr>
<tr>
<td>
<code>enableIpv4Masquerade</code></br>
<em>
boolean
</em>
</td>
<td>
<em>(Optional)</em>
<p>EnableIPv4Masquerade masquerades packets from endpoints leaving the host with BPF instead of iptables if Snat is not enabled</p>
</td>
</tr>
<tr>
<td>
<code>enableIpv6Masquerade</code></br>
<em>
boolean
</em>
</td>
<td>
<em>(Optional)</em>
<p>EnableIPv6Masquerade masquerades packets from endpoints leaving the host with BPF instead of iptables if Snat is not enabled</p>
</td>
</tr>
<tr>
<td>
<code>enableBPFMasquerade</code></br>
<em>
boolean
</em>
</td>
<td>
<em>(Optional)</em>
<p>EnableBPFMasquerade masquerades packets from endpoints leaving the host with BPF instead of iptables if Snat is not enabled</p>
</td>
</tr>
<tr>
<td>
<code>bgpControlPlane</code></br>
<em>
<a href="#bgpcontrolplane">BGPControlPlane</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>BGPControlPlane enables the BGP Control Plane</p>
</td>
</tr>
<tr>
<td>
<code>policyAuditMode</code></br>
<em>
boolean
</em>
</td>
<td>
<em>(Optional)</em>
<p>PolicyAuditMode enables non-drop mode for installed policies</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code></br>
<em>
<a href="#encryption">Encryption</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Encryption handles traffic encryption configuration</p>
</td>
</tr>

</tbody>
</table>


<h3 id="networkstatus">NetworkStatus
</h3>


<p>
NetworkStatus contains information about created Network resources.
</p>


<h3 id="nodeportmode">NodePortMode
</h3>
<p><em>Underlying type: string</em></p>


<p>
(<em>Appears on:</em><a href="#nodeport">Nodeport</a>)
</p>

<p>
NodePortMode defines how NodePort services are enabled.
</p>


<h3 id="nodeport">Nodeport
</h3>


<p>
Nodeport enablement for cilium
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
<code>nodePortEnabled</code></br>
<em>
boolean
</em>
</td>
<td>
<p>Enabled is used to define whether Nodeport is required or not.</p>
</td>
</tr>
<tr>
<td>
<code>nodePortMode</code></br>
<em>
<a href="#nodeportmode">NodePortMode</a>
</em>
</td>
<td>
<p>Mode is the mode of NodePort feature</p>
</td>
</tr>

</tbody>
</table>


<h3 id="overlay">Overlay
</h3>


<p>
(<em>Appears on:</em><a href="#networkconfig">NetworkConfig</a>)
</p>

<p>
Overlay configuration for cilium
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
<code>enabled</code></br>
<em>
boolean
</em>
</td>
<td>
<p>Enabled enables the network overlay.</p>
</td>
</tr>
<tr>
<td>
<code>createPodRoutes</code></br>
<em>
boolean
</em>
</td>
<td>
<em>(Optional)</em>
<p>CreatePodRoutes installs routes to pods on all cluster nodes.<br />This will only work if the cluster nodes share a single L2 network.</p>
</td>
</tr>

</tbody>
</table>


<h3 id="snatoutofcluster">SnatOutOfCluster
</h3>


<p>
(<em>Appears on:</em><a href="#networkconfig">NetworkConfig</a>)
</p>

<p>
SnatOutOfCluster enables the masquerading of packets outside of the cluster
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
<code>enabled</code></br>
<em>
boolean
</em>
</td>
<td>
<p></p>
</td>
</tr>

</tbody>
</table>


<h3 id="snattoupstreamdns">SnatToUpstreamDNS
</h3>


<p>
(<em>Appears on:</em><a href="#networkconfig">NetworkConfig</a>)
</p>

<p>
SnatToUpstreamDNS enables the masquerading of packets to the upstream dns server
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
<code>enabled</code></br>
<em>
boolean
</em>
</td>
<td>
<p></p>
</td>
</tr>

</tbody>
</table>


<h3 id="store">Store
</h3>
<p><em>Underlying type: string</em></p>


<p>
(<em>Appears on:</em><a href="#networkconfig">NetworkConfig</a>)
</p>

<p>
Store defines the kubernetes storage backend
</p>


<h3 id="tunnelmode">TunnelMode
</h3>
<p><em>Underlying type: string</em></p>


<p>
(<em>Appears on:</em><a href="#networkconfig">NetworkConfig</a>)
</p>

<p>
TunnelMode defines what tunnel mode to use for Cilium.
</p>


