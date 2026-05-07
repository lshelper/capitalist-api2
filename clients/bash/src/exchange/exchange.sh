#!/usr/bin/env bash

capitalist_api2_create_exchange_json() {
  capitalist_api2_request POST "/v1/exchange" "$1"
}

capitalist_api2_get_rate() {
  local from to
  from="$(capitalist_api2_urlencode "$1")"
  to="$(capitalist_api2_urlencode "$2")"
  capitalist_api2_request GET "/v1/rate?from=${from}&to=${to}"
}
