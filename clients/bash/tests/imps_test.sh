#!/usr/bin/env bash
set -eu
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
source "$ROOT/src/core.sh"
source "$ROOT/src/payments/create/imps.sh"
# Capture the helper body locally: never invoke curl or the API.
capitalist_api2_create_payment_json() {
  printf '%s' "$1" | python3 -c 'import json,sys; p=json.load(sys.stdin); d=p["payload"]; assert d == {"type":"IMPS","destination":{"bankCode":"SBIN0001234","accountName":"Raj Kumar","accountNumber":"0123456789012345","bankName":"Punjab National Bank"},"recipient":{"phone":"12125551234"}}; assert p["amount"] == 100; assert "inr_sum" not in d'
}
capitalist_api2_create_payment_imps test-request example-account 100 USD
printf 'IMPS payload test passed\n'
