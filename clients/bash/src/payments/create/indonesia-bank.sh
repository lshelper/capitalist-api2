#!/usr/bin/env bash

# INDONESIA_BANK payment
capitalist_api2_create_payment_indonesia_bank() {
  local user_request_id="${1:-indonesia-bank-$(capitalist_api2_now_ms)}"
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
  "comment": "INDONESIA_BANK payout example",
  "payload": {
    "type": "INDONESIA_BANK",
    "destination": {
      "bankCode": "CENAIDJA",
      "accountName": "Budi Santoso",
      "accountNumber": "1234567890"
    },
    "recipient": {
      "firstName": "Budi",
      "lastName": "Santoso",
      "email": "beneficiary@example.com",
      "phone": "628123456789",
      "zipCode": "10110",
      "city": "Jakarta",
      "address": "Jalan Sudirman",
      "country": "id"
    },
    "account": "1234567890"
  }
}
JSON
)"
}
