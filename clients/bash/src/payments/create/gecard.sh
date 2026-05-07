#!/usr/bin/env bash

# GECARD payment
capitalist_api2_create_payment_gecard() {
  local user_request_id="${1:-gecard-$(capitalist_api2_now_ms)}"
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
  "comment": "GECARD payout example",
  "payload": {
    "type": "GECARD",
    "account": "5234567890123456",
    "senderFirstname": "Giorgi",
    "senderLastname": "Piponia",
    "senderIdentityNumber": "01234567890",
    "receiverFirstname": "Nino",
    "receiverLastname": "Kupadze"
  }
}
JSON
)"
}
