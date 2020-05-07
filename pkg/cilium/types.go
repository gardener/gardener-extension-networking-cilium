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

package cilium

import "path/filepath"

const (
	Name = "networking-cilium"

	// ImageNames
	CiliumAgentImageName        = "cilium-agent"
	CiliumOperatorImageName     = "cilium-operator"
	CiliumETCDOperatorImageName = "cilium-etcd-operator"
	CiliumNodeInitImageName     = "cilium-nodeinit"
	CiliumPreflightImageName    = "cilium-preflight"

	HubbleImageName   = "hubble"
	HubbleUIImageName = "hubble-ui"

	// ReleaseName is the name of the Cilium Release
	ReleaseName = "cilium"
)

var (
	// ChartsPath is the path to the charts
	ChartsPath = filepath.Join("charts")
	// InternalChartsPath is the path to the internal charts
	InternalChartsPath = filepath.Join(ChartsPath, "internal")

	// ChartPath path for internal Cilium Chart
	ChartPath = filepath.Join(InternalChartsPath, "cilium")
)
