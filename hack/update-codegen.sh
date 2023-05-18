#!/usr/bin/env bash

# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

MODULE="github.com/meixiezichuan/clustertopology"
APIS_PKG=api
#代码生出输出
OUTPUT_PKG=generated

GROUP=edge
VERSION=v1
GROUP_VERSION=$GROUP:$VERSION

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

# kubebuilder生成的api目录结构code-generator无法直接使用
# rm -rf "${APIS_PKG}/${GROUP}" && mkdir -p "${APIS_PKG}/${GROUP}" && cp -r "${APIS_PKG}/${VERSION}" "${APIS_PKG}/${GROUP}"

bash "${CODEGEN_PKG}"/generate-groups.sh "client,lister,informer" \
${MODULE}/${OUTPUT_PKG} ${MODULE}/${APIS_PKG} \
${GROUP_VERSION} \
--go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt




