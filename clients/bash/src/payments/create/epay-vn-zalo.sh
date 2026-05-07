#!/usr/bin/env bash

# EPAY_EPAY_E_VN_ZALO payment
capitalist_api2_create_payment_epay_vn_zalo() {
  local user_request_id="${1:-epay-vn-zalo-$(capitalist_api2_now_ms)}"
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
  "comment": "EPAY_EPAY_E_VN_ZALO payout example",
  "payload": {
    "type": "EPAY_EPAY_E_VN_ZALO",
    "receiver": {
      "surName": "Tran",
      "givName": "Anh",
      "phone": "0901234567"
    },
    "sender": {
      "surName": "Smith",
      "country": "US",
      "address": "123 Main St",
      "purposeOfRemittance": "Family support"
    },
    "account": "123456789",
    "receiveAmount": "100000",
    "receiveCurrency": "USD"
  }
}
JSON
)"
}
