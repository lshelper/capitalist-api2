#!/usr/bin/env bash

# BUY_ITEM payment
capitalist_api2_create_payment_buy_item() {
  local user_request_id="${1:-buy-item-$(capitalist_api2_now_ms)}"
  local account_from="${2:-${CAPITALIST_API2_FROM_ACCOUNT:-}}"
  local amount="${3:-13.45}"
  local currency="${4:-${CAPITALIST_API2_CURRENCY:-USD}}"
  local product_id="${5:-14}"
  local denomination_id="${6:-59}"

  capitalist_api2_require_value CAPITALIST_API2_FROM_ACCOUNT "$account_from" || return 1

  capitalist_api2_create_payment_json "$(cat <<JSON
{
  "userRequestId": "$(capitalist_api2_json_escape "$user_request_id")",
  "accountFrom": "$(capitalist_api2_json_escape "$account_from")",
  "amount": ${amount},
  "currency": "$(capitalist_api2_json_escape "$currency")",
  "comment": "BUY_ITEM payment example",
  "payload": {
    "productId": ${product_id},
    "denominationId": ${denomination_id},
    "type": "BUY_ITEM"
  }
}
JSON
)"
}
