// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	webhookcmd "github.com/gardener/gardener/extensions/pkg/webhook/cmd"
	extensionshootwebhook "github.com/gardener/gardener/extensions/pkg/webhook/shoot"

	shootwebhook "github.com/gardener/gardener-extension-networking-cilium/pkg/webhook/shoot"
)

// WebhookSwitchOptions are the webhookcmd.SwitchOptions for the cilium network extension webhooks.
func WebhookSwitchOptions() *webhookcmd.SwitchOptions {
	return webhookcmd.NewSwitchOptions(
		webhookcmd.Switch(extensionshootwebhook.WebhookName, shootwebhook.AddToManager),
	)
}
