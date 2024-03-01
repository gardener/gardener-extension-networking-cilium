// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	runtimelog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"

	"github.com/gardener/gardener-extension-networking-cilium/cmd/gardener-extension-networking-cilium/app"
)

func main() {
	cmd := app.NewControllerManagerCommand(signals.SetupSignalHandler())

	if err := cmd.Execute(); err != nil {
		runtimelog.Log.Error(err, "Error executing the main controller command")
	}
}
