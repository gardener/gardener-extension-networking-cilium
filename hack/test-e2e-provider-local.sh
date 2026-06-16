#!/bin/bash

# SPDX-FileCopyrightText: 2025 SAP SE or an SAP affiliate company and Gardener contributors
#
# SPDX-License-Identifier: Apache-2.0

# trigger test run - delete later -4
set -o nounset
set -o pipefail
set -o errexit

repo_root="$(readlink -f $(dirname ${0})/..)"

if [[ ! -d "$repo_root/gardener" ]]; then
  git clone https://github.com/axel7born/gardener.git
fi

# gardener_version=$(go list -m -f '{{.Version}}' github.com/gardener/gardener)
cd "$repo_root/gardener"
git checkout "enh/machine-pod-apparmor-unconfined"
source "$repo_root/gardener/hack/ci-common.sh"

echo ">>>>>>>>>>>>>>>>>>>> kind-up"
make kind-up
trap '{
  cd "$repo_root/gardener"
  export_artifacts "gardener-local"
  make kind-down
}' EXIT
export KUBECONFIG=$repo_root/gardener/dev-setup/kubeconfigs/seed/kubeconfig
echo "<<<<<<<<<<<<<<<<<<<< kind-up done"

echo ">>>>>>>>>>>>>>>>>>>> gardener-up"
make gardener-up
echo "<<<<<<<<<<<<<<<<<<<< gardener-up done"

cd $repo_root
echo ">>>>>>>>>>>>>>>>>>>> extension-up"
make extension-up
echo "<<<<<<<<<<<<<<<<<<<< extension-up done"

export KUBECONFIG=$repo_root/gardener/dev-setup/kubeconfigs/virtual-garden/kubeconfig
export REPO_ROOT=$repo_root

# reduce flakiness in contended pipelines
export GOMEGA_DEFAULT_EVENTUALLY_TIMEOUT=5s
export GOMEGA_DEFAULT_EVENTUALLY_POLLING_INTERVAL=200ms
# if we're running low on resources, it might take longer for tested code to do something "wrong"
# poll for 5s to make sure, we're not missing any wrong action
export GOMEGA_DEFAULT_CONSISTENTLY_DURATION=5s
export GOMEGA_DEFAULT_CONSISTENTLY_POLLING_INTERVAL=200ms

ginkgo --timeout=1h --v --show-node-events "$@" $repo_root/test/e2e/...

echo ">>>>>>>>>>>>>>>>>>>> kind-down"
cd "$repo_root/gardener"
make kind-down
echo "<<<<<<<<<<<<<<<<<<<< kind-down done"