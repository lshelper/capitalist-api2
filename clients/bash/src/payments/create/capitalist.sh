#!/usr/bin/env bash

# CAPITALIST internal payment
capitalist_api2_create_payment_capitalist() {
  local user_request_id="${1:-capitalist-$(capitalist_api2_now_ms)}"
  local account_from="${2:-${CAPITALIST_API2_FROM_ACCOUNT:-}}"
  local amount="${3:-100}"
  local currency="${4:-${CAPITALIST_API2_CURRENCY:-USD}}"
  local payload_account="${5:-${CAPITALIST_API2_TO_ACCOUNT:-}}"

  capitalist_api2_require_value CAPITALIST_API2_FROM_ACCOUNT "$account_from" || return 1
  capitalist_api2_require_value CAPITALIST_API2_TO_ACCOUNT "$payload_account" || return 1

  capitalist_api2_create_payment_json "$(cat <<JSON
{
  "userRequestId": "$(capitalist_api2_json_escape "$user_request_id")",
  "accountFrom": "$(capitalist_api2_json_escape "$account_from")",
  "amount": ${amount},
  "currency": "$(capitalist_api2_json_escape "$currency")",
  "comment": "CAPITALIST payout example",
  "payload": {
    "type": "CAPITALIST",
    "account": "$(capitalist_api2_json_escape "$payload_account")"
  }
}
JSON
)"
}
