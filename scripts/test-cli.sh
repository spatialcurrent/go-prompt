#!/bin/bash

# =================================================================
#
# Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
# Released as open source under the MIT License.  See LICENSE file.
#
# =================================================================

set -euo pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

testString() {
  local expected='world'
  local output=$(echo 'world' | goprompt --question 'hello')
  assertEquals "unexpected output" "$(echo -e "${expected}")" "${output}"
}

testJSON() {
  local expected='{"hello": "world"}'
  local output=$(echo '{"hello": "world"}' | goprompt --question 'hello')
  assertEquals "unexpected output" "$(echo -e "${expected}")" "${output}"
}

oneTimeSetUp() {
  echo "Setting up"
  echo "Using temporary directory at ${SHUNIT_TMPDIR}"
}

oneTimeTearDown() {
  echo "Tearing Down"
}

# Load shUnit2.
. "${DIR}/shunit2"
