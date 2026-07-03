#!/usr/bin/env bash

# TOPUP_SERVICE payment
capitalist_api2_create_payment_topup_service() {
  local user_request_id="${1:-topup-service-$(capitalist_api2_now_ms)}"
  local account_from="${2:-${CAPITALIST_API2_FROM_ACCOUNT:-}}"
  local amount="${3:-100}"
  local currency="${4:-${CAPITALIST_API2_CURRENCY:-USD}}"
  local service_id="${5:-34}"
  local account="${6:-myNastyAccountId_notNickName}"
  local quantity="${7:-100}"
  local region="${8:-Macau}"

  capitalist_api2_require_value CAPITALIST_API2_FROM_ACCOUNT "$account_from" || return 1

  capitalist_api2_create_payment_json "$(cat <<JSON
{
  "userRequestId": "$(capitalist_api2_json_escape "$user_request_id")",
  "accountFrom": "$(capitalist_api2_json_escape "$account_from")",
  "amount": ${amount},
  "currency": "$(capitalist_api2_json_escape "$currency")",
  "comment": "TOPUP_SERVICE payment example",
  "payload": {
    "serviceId": ${service_id},
    "account": "$(capitalist_api2_json_escape "$account")",
    "quantity": ${quantity},
    "region": "$(capitalist_api2_json_escape "$region")",
    "type": "TOPUP_SERVICE"
  }
}
JSON
)"
}
