#!/usr/bin/env bash

capitalist_api2_list_accounts() {
  local currency
  currency="$(capitalist_api2_urlencode "$1")"
  capitalist_api2_request GET "/v1/account/list?currency=${currency}"
}

capitalist_api2_get_transactions() {
  capitalist_api2_request GET "/v1/transactions${1:-}"
}

capitalist_api2_get_account_history() {
  capitalist_api2_get_transactions "${1:-}"
}

capitalist_api2_get_account_history_by_transaction_id() {
  local transaction_id
  transaction_id="$(capitalist_api2_urlencode "$1")"
  capitalist_api2_get_transactions "?transactionId=${transaction_id}"
}
