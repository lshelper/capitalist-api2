#!/usr/bin/env bash

capitalist_api2_read_whitelist() {
  capitalist_api2_request GET "/v1/whitelist"
}

capitalist_api2_add_whitelist_json() {
  capitalist_api2_request POST "/v1/whitelist" "$1"
}

capitalist_api2_remove_whitelist_json() {
  capitalist_api2_request POST "/v1/whitelist/remove" "$1"
}
