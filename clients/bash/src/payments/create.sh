#!/usr/bin/env bash

capitalist_api2_create_payment_json() {
  capitalist_api2_request POST "/v1/payment" "$1"
}
