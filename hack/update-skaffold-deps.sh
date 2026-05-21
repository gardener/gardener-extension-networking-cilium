#!/usr/bin/env bash
#
# SPDX-FileCopyrightText: 2025 SAP SE or an SAP affiliate company and Gardener contributors
#
# SPDX-License-Identifier: Apache-2.0

set -e

repo_root="$(git rev-parse --show-toplevel)"
export GARDENER_HACK_DIR="$(go list -m -f "{{.Dir}}" github.com/gardener/gardener)/hack"
$repo_root/hack/check-skaffold-deps.sh update
