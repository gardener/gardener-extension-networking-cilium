#!/bin/bash
#
# SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
#
# SPDX-License-Identifier: Apache-2.0

set -o errexit
set -o nounset
set -o pipefail

PROJECT_ROOT="$(dirname $0)"/..

# setup virtual GOPATH
source "$GARDENER_HACK_DIR"/vgopath-setup.sh

CODE_GEN_DIR=$(go list -m -f '{{.Dir}}' k8s.io/code-generator)

source "${CODE_GEN_DIR}/kube_codegen.sh"

rm -f $GOPATH/bin/*-gen

kube::codegen::gen_helpers \
  --boilerplate "${GARDENER_HACK_DIR}/LICENSE_BOILERPLATE.txt" \
  --extra-peer-dir k8s.io/apimachinery/pkg/apis/meta/v1 \
  --extra-peer-dir k8s.io/apimachinery/pkg/conversion \
  --extra-peer-dir k8s.io/component-base/config \
  --extra-peer-dir k8s.io/component-base/config/v1alpha1 \
  "${PROJECT_ROOT}/pkg/apis"