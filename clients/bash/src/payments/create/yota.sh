#!/usr/bin/env bash

# YOTA mobile payment
capitalist_api2_create_payment_yota() {
  local user_request_id="${1:-yota-$(capitalist_api2_now_ms)}"
  local account_from="${2:-${CAPITALIST_API2_FROM_ACCOUNT:-}}"
  local amount="${3:-100}"
  local currency="${4:-${CAPITALIST_API2_CURRENCY:-USD}}"

  capitalist_api2_require_value CAPITALIST_API2_FROM_ACCOUNT "$account_from" || return 1

  capitalist_api2_create_payment_json "$(cat <<JSON
{
  "userRequestId": "$(capitalist_api2_json_escape "$user_request_id")",
  "accountFrom": "$(capitalist_api2_json_escape "$account_from")",
  "amount": ${amount},
  "currency": "$(capitalist_api2_json_escape "$currency")",
  "comment": "YOTA payout example",
  "payload": {
    "type": "YOTA",
    "account": "79001234567"
  }
}
JSON
)"
}
