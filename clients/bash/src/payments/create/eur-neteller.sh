#!/usr/bin/env bash

# EUR_NETELLER payment
capitalist_api2_create_payment_eur_neteller() {
  local user_request_id="${1:-eur-neteller-$(capitalist_api2_now_ms)}"
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
  "comment": "EUR_NETELLER payout example",
  "payload": {
    "type": "EUR_NETELLER",
    "wallet_email": "beneficiary@example.com",
    "name": "John",
    "surname": "Doe",
    "country": "USA",
    "city": "New York",
    "address": "123 Main St",
    "zip": "10001",
    "code": "123456",
    "phone": "12125551234",
    "client_ip": "192.168.1.1",
    "state": "NY"
  }
}
JSON
)"
}
