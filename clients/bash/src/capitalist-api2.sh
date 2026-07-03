#!/usr/bin/env bash

CAPITALIST_API2_SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

. "${CAPITALIST_API2_SCRIPT_DIR}/core.sh"
. "${CAPITALIST_API2_SCRIPT_DIR}/settings/whitelist.sh"
. "${CAPITALIST_API2_SCRIPT_DIR}/account/accounts.sh"
. "${CAPITALIST_API2_SCRIPT_DIR}/exchange/exchange.sh"
. "${CAPITALIST_API2_SCRIPT_DIR}/merchant/orders.sh"
. "${CAPITALIST_API2_SCRIPT_DIR}/wallets/deposit.sh"
. "${CAPITALIST_API2_SCRIPT_DIR}/prepaid2/prepaid2.sh"
. "${CAPITALIST_API2_SCRIPT_DIR}/payments/create.sh"
. "${CAPITALIST_API2_SCRIPT_DIR}/payments/status.sh"
. "${CAPITALIST_API2_SCRIPT_DIR}/kyc/kyc.sh"
