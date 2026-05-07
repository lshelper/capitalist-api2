#!/usr/bin/env bash

# PAYONEER_EUR payment
capitalist_api2_create_payment_payoneer_eur() {
  local user_request_id="${1:-payoneer-eur-$(capitalist_api2_now_ms)}"
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
  "comment": "PAYONEER_EUR payout example",
  "payload": {
    "type": "PAYONEER_EUR",
    "account": "beneficiary@example.com"
  }
}
JSON
)"
}
