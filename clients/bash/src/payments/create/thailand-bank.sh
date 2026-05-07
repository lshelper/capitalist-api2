#!/usr/bin/env bash

# THAILAND_BANK payment
capitalist_api2_create_payment_thailand_bank() {
  local user_request_id="${1:-thailand-bank-$(capitalist_api2_now_ms)}"
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
  "comment": "THAILAND_BANK payout example",
  "payload": {
    "type": "THAILAND_BANK",
    "destination": {
      "bankCode": "BKKBTHBK",
      "accountName": "Somchai Prasert",
      "accountNumber": "1234567890"
    },
    "recipient": {
      "firstName": "Somchai",
      "lastName": "Prasert",
      "email": "beneficiary@example.com",
      "phone": "66812345678",
      "zipCode": "10100",
      "city": "Bangkok",
      "address": "Sukhumvit Road",
      "country": "th"
    },
    "account": "1234567890"
  }
}
JSON
)"
}
