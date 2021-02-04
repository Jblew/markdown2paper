#!/usr/bin/env bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "${DIR}"
set -e

SRC_DIR=".."
BUILD_CMD="${SRC_DIR}/build.sh"
RUN_CMD="${SRC_DIR}/markdown2paper"

"${BUILD_CMD}"

"${RUN_CMD}" build
