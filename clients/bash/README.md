# Capitalist API2 Bash client

Small Bash client for scripts, CI jobs and manual operations.

## Usage

```bash
export CAPITALIST_API2_KEY="your-api-key"
export CAPITALIST_API2_SECRET="your-api-secret"

source ./src/capitalist-api2.sh

capitalist_api2_list_accounts "USD"
```

The client requires `curl` and either `openssl` or `shasum`.

## Source layout

`src/capitalist-api2.sh` is the compatibility entrypoint and loads the focused modules:

- `src/account/` for account list and transaction history.
- `src/exchange/` for exchange creation and rates.
- `src/kyc/` for KYC start, status, data, picture and confirmation.
- `src/merchant/` for merchant orders.
- `src/payments/create.sh` for payment creation.
- `src/payments/status.sh` for payment status lookup.
- `src/settings/` for IP whitelist management.

## Test

```bash
./tests/signature_test.sh
```
