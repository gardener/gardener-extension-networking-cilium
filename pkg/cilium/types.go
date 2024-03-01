// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package cilium

import (
	"path/filepath"

	"github.com/gardener/gardener-extension-networking-cilium/charts"
)

const (
	// Name defines the extension name.
	Name = "networking-cilium"

	// CiliumAgentImageName defines the agent name.
	CiliumAgentImageName = "cilium-agent"
	// CiliumOperatorImageName defines cilium operator's image name.
	CiliumOperatorImageName = "cilium-operator"

	// HubbleRelayImageName defines the Hubble image name.
	HubbleRelayImageName = "hubble-relay"
	// HubbleUIImageName defines the UI image name.
	HubbleUIImageName = "hubble-ui"
	// HubbleUIBackendImageName defines the UI Backend image name.
	HubbleUIBackendImageName = "hubble-ui-backend"

	// CertGenImageName defines certificate generation image name.
	CertGenImageName = "certgen"

	// KubeProxyImageName defines the kube-proxy image name.
	KubeProxyImageName = "kube-proxy"

	// MonitoringChartName
	MonitoringName = "cilium-monitoring-config"

	// ReleaseName is the name of the Cilium Release.
	ReleaseName = "cilium"
)

var (
	// CiliumChartPath is the path for internal Cilium Chart
	CiliumChartPath = filepath.Join(charts.InternalChartsPath, "cilium")

	// CiliumMonitoringChartPath  path for internal Cilium monitoring chart
	CiliumMonitoringChartPath = filepath.Join(charts.InternalChartsPath, "cilium-monitoring")
)
