#!/usr/bin/env bash

# USDTERC20 payment
capitalist_api2_create_payment_usdterc20() {
  local user_request_id="${1:-usdterc20-$(capitalist_api2_now_ms)}"
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
  "comment": "USDTERC20 payout example",
  "payload": {
    "type": "USDTERC20",
    "account": "0x0000000000000000000000000000000000000000"
  }
}
JSON
)"
}
