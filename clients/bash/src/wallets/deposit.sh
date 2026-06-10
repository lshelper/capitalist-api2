#!/usr/bin/env bash

capitalist_api2_get_deposit_address() {
  capitalist_api2_request GET "/v1/depositAddress/$(capitalist_api2_urlencode "$1")"
}

capitalist_api2_get_auto_usdtt_deposit_address() {
  capitalist_api2_request GET "/v1/depositAddressAutoUSDTt/$(capitalist_api2_urlencode "$1")"
}
