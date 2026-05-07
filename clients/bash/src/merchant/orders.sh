#!/usr/bin/env bash

capitalist_api2_get_orders() {
  capitalist_api2_request GET "/v1/orders${1:-}"
}

capitalist_api2_get_merchant_orders() {
  capitalist_api2_get_orders "${1:-}"
}

capitalist_api2_get_merchant_orders_by_merchant() {
  local merchant_id merchant_name
  merchant_id="$(capitalist_api2_urlencode "$1")"
  merchant_name="$(capitalist_api2_urlencode "$2")"
  capitalist_api2_get_orders "?merchantId=${merchant_id}&merchantName=${merchant_name}"
}

capitalist_api2_get_merchant_orders_by_order() {
  local order_id order_number
  order_id="$(capitalist_api2_urlencode "$1")"
  order_number="$(capitalist_api2_urlencode "$2")"
  capitalist_api2_get_orders "?orderId=${order_id}&orderNumber=${order_number}"
}
