#!/usr/bin/env bash
set -eu

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
. "${SCRIPT_DIR}/../src/capitalist-api2.sh"

capitalist_api2_list_accounts "${1:-USD}"
