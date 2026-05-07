#!/usr/bin/env bash

capitalist_api2_start_kyc_json() {
  capitalist_api2_request POST "/v1/kyc/start" "$1"
}

capitalist_api2_get_kyc_status() {
  capitalist_api2_request GET "/v1/kyc/status/$(capitalist_api2_urlencode "$1")/$(capitalist_api2_urlencode "$2")"
}

capitalist_api2_get_kyc_status_by_uuid() {
  capitalist_api2_request GET "/v1/kyc/statusByUuid/$(capitalist_api2_urlencode "$1")"
}

capitalist_api2_set_kyc_data_json() {
  capitalist_api2_request POST "/v1/kyc/setData/$(capitalist_api2_urlencode "$1")" "$2"
}

capitalist_api2_set_kyc_picture_jpeg() {
  local uuid="$1"
  local type="$2"
  local file="$3"

  capitalist_api2_request_file POST "/v1/kyc/setPicture/$(capitalist_api2_urlencode "$uuid")/$(capitalist_api2_urlencode "$type")" "$file" "image/jpeg"
}

capitalist_api2_confirm_kyc() {
  capitalist_api2_request POST "/v1/kyc/confirm/$(capitalist_api2_urlencode "$1")" "{}"
}
