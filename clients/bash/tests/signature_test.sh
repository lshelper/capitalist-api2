#!/usr/bin/env bash
set -eu

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
CLIENT_DIR="$(cd "${SCRIPT_DIR}/.." && pwd)"

. "${CLIENT_DIR}/src/capitalist-api2.sh"

timestamp="1678901234567"
body='{"type":"PAYONEER","accountFrom":"U0123504","payload":{"account":"[email protected]"},"userRequestId":"1745844029437","callbackUrl":"https://some-domain.com/callbacksFromCap","amount":100}'
secret="somePrivateApiSecret"
expected="57bd25d836c076c24a81f1f088f405b5fd371918ba96f9a28df72d39860396b6"
actual="$(capitalist_api2_signature "$timestamp" "$body" "$secret")"

if [ "$actual" != "$expected" ]; then
  printf 'Expected %s, got %s\n' "$expected" "$actual" >&2
  exit 1
fi

printf 'signature ok\n'
