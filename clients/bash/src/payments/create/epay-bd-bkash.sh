#!/usr/bin/env bash

# EPAY_EPAY_E_BD_BKASH payment
capitalist_api2_create_payment_epay_bd_bkash() {
  local user_request_id="${1:-epay-bd-bkash-$(capitalist_api2_now_ms)}"
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
  "comment": "EPAY_EPAY_E_BD_BKASH payout example",
  "payload": {
    "type": "EPAY_EPAY_E_BD_BKASH",
    "receiver": {
      "surName": "Rahman",
      "givName": "Ali",
      "accountNo": "017XXXXXXXX"
    },
    "sender": {
      "givName": "John",
      "surName": "Doe",
      "area": "Dhaka",
      "phone": "8801XXXXXXXX",
      "nationality": "US",
      "idType": "Passport",
      "idNumber": "AB1234567",
      "birthday": "1985-01-01",
      "city": "NYC",
      "zipCode": "10001",
      "country": "USA",
      "address": "123 Main St",
      "purposeOfRemittance": "Family Support"
    },
    "account": "017XXXXXXXX",
    "receiveAmount": "5000",
    "receiveCurrency": "USD"
  }
}
JSON
)"
}
