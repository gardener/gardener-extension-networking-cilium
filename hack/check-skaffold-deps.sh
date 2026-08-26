#!/usr/bin/env bash
#
# SPDX-FileCopyrightText: Copyright Contributors to the Gardener project
#
# SPDX-License-Identifier: Apache-2.0

set -e

operation="${1:-check}"

echo "> ${operation} Skaffold Dependencies"

success=true

function run() {
  if ! bash "$GARDENER_HACK_DIR"/check-skaffold-deps-for-binary.sh "$operation" --skaffold-file "$1" --binary "$2" --skaffold-config "$3"; then
    success=false
  fi
}

# skaffold.yaml
run "skaffold.yaml" "gardener-extension-networking-cilium" "networking-cilium"

if ! $success ; then
  exit 1
fi
