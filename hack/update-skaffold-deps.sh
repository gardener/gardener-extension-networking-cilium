#!/usr/bin/env bash
#
# SPDX-FileCopyrightText: Contributors to the Gardener project
#
# SPDX-License-Identifier: Apache-2.0

set -e

repo_root="$(git rev-parse --show-toplevel)"
export GARDENER_HACK_DIR="$(go list -m -f "{{.Dir}}" github.com/gardener/gardener)/hack"
$repo_root/hack/check-skaffold-deps.sh update
