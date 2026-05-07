#!/usr/bin/env bash

# COBANK payment
capitalist_api2_create_payment_cobank() {
  local user_request_id="${1:-cobank-$(capitalist_api2_now_ms)}"
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
  "comment": "COBANK payout example",
  "payload": {
    "type": "COBANK",
    "taxId": "1234567890",
    "fullName": "Carlos Rodriguez",
    "email": "beneficiary@example.com",
    "phone": "573001234567"
  }
}
JSON
)"
}
