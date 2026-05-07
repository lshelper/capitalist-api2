#!/usr/bin/env bash

# WORLDCARDUSD payment
capitalist_api2_create_payment_worldcardusd() {
  local user_request_id="${1:-worldcardusd-$(capitalist_api2_now_ms)}"
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
  "comment": "WORLDCARDUSD payout example",
  "payload": {
    "type": "WORLDCARDUSD",
    "account": "1234567890123456",
    "name": "John",
    "surname": "Smith",
    "cardCurrency": "USD",
    "birthDate": "1990-01-01",
    "address": "123 Main Street",
    "addressCountryCode": "USA",
    "addressCity": "New York",
    "cardExpiryMonth": "12",
    "cardExpiryYear": "2025"
  }
}
JSON
)"
}
