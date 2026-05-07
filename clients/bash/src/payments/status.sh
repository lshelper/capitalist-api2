#!/usr/bin/env bash

capitalist_api2_get_payment_by_document_id() {
  capitalist_api2_request GET "/v1/payment/document/$(capitalist_api2_urlencode "$1")"
}

capitalist_api2_get_payment_by_user_request_id() {
  capitalist_api2_request GET "/v1/payment/$(capitalist_api2_urlencode "$1")"
}
