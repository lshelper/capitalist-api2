#!/usr/bin/env bash

capitalist_api2_get_prepaid_services() {
  capitalist_api2_request GET "/v1/prepaid2/services"
}

capitalist_api2_get_prepaid_regions() {
  capitalist_api2_request GET "/v1/prepaid2/regions"
}

capitalist_api2_get_prepaid_products() {
  capitalist_api2_request GET "/v1/prepaid2/products"
}

capitalist_api2_get_prepaid_denominations() {
  local product_id
  product_id="$(capitalist_api2_urlencode "$1")"
  capitalist_api2_request GET "/v1/prepaid2/denominations?productid=${product_id}"
}
