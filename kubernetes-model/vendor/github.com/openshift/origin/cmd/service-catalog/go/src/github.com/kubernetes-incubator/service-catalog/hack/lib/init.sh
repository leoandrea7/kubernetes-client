#!/bin/bash
#
# Copyright (C) 2015 Red Hat, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#         http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#


# This script is meant to be the entrypoint for OpenShift Bash scripts to import all of the support
# libraries at once in order to make Bash script preambles as minimal as possible. This script recur-
# sively `source`s *.sh files in this directory tree. As such, no files should be `source`ed outside
# of this script to ensure that we do not attempt to overwrite read-only variables.

set -o errexit
set -o nounset
set -o pipefail

# os::util::absolute_path returns the absolute path to the directory provided
function os::util::absolute_path() {
	local relative_path="$1"
	local absolute_path

	pushd "${relative_path}" >/dev/null
	relative_path="$( pwd )"
	if [[ -h "${relative_path}" ]]; then
		absolute_path="$( readlink "${relative_path}" )"
	else
		absolute_path="${relative_path}"
	fi
	popd >/dev/null

	echo "${absolute_path}"
}
readonly -f os::util::absolute_path

# find the absolute path to the root of the Origin source tree
init_source="$( dirname "${BASH_SOURCE}" )/../.."
OS_ROOT="$( os::util::absolute_path "${init_source}" )"
export OS_ROOT
cd "${OS_ROOT}"

library_files=( $( find "${OS_ROOT}/hack/lib" -type f -name '*.sh' -not -path '*/hack/lib/init.sh' ) )
# TODO(skuzmets): Move the contents of the following files into respective library files.
library_files+=( "${OS_ROOT}/hack/common.sh" )
library_files+=( "${OS_ROOT}/hack/util.sh" )

for library_file in "${library_files[@]}"; do
	source "${library_file}"
done

unset library_files library_file init_source

# all of our Bash scripts need to have the stacktrace
# handler installed to deal with errors
os::log::stacktrace::install

# All of our Bash scripts need to have access to the
# binaries that we build so we don't have to find
# them before every invocation.
os::util::environment::update_path_var

if [[ -z "${OS_TMP_ENV_SET-}" ]]; then
	os::util::environment::setup_tmpdir_vars "$( basename "$0" ".sh" )"
fi