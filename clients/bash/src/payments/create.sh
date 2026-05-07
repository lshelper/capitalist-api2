#!/usr/bin/env bash

capitalist_api2_create_payment_json() {
  capitalist_api2_request POST "/v1/payment" "$1"
}

CAPITALIST_API2_PAYMENTS_CREATE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/create" && pwd)"

for capitalist_api2_payments_create_file in "${CAPITALIST_API2_PAYMENTS_CREATE_DIR}"/*.sh; do
  . "$capitalist_api2_payments_create_file"
done

unset capitalist_api2_payments_create_file
unset CAPITALIST_API2_PAYMENTS_CREATE_DIR
