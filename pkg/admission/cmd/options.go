// SPDX-FileCopyrightText: Copyright Contributors to the Gardener project
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	webhookcmd "github.com/gardener/gardener/extensions/pkg/webhook/cmd"

	"github.com/gardener/gardener-extension-networking-cilium/pkg/admission/mutator"
	"github.com/gardener/gardener-extension-networking-cilium/pkg/admission/validator"
)

// GardenWebhookSwitchOptions are the webhookcmd.SwitchOptions for the admission webhooks.
func GardenWebhookSwitchOptions() *webhookcmd.SwitchOptions {
	return webhookcmd.NewSwitchOptions(
		webhookcmd.Switch(validator.Name, validator.New),
		webhookcmd.Switch(mutator.Name, mutator.New),
	)
}
