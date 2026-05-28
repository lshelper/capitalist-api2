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

## Callback verification

Verify incoming payment callbacks against the raw request body:

```bash
capitalist_api2_verify_signature "$timestamp_header" "$raw_body" "$api_secret" "$signature_header"
```

## Payment creation helpers

Payment channel helpers are loaded from `src/payments/create/*.sh`.

All channel helpers accept:

```text
userRequestId accountFrom amount currency [payloadAccount]
```

Arguments are optional and default to:

- `userRequestId`: `<channel>-$(capitalist_api2_now_ms)`
- `accountFrom`: `CAPITALIST_API2_FROM_ACCOUNT`
- `amount`: `100`
- `currency`: `CAPITALIST_API2_CURRENCY` or `USD`
- `payloadAccount`: channel-specific account only when the sample needs it

Example:

```bash
export CAPITALIST_API2_FROM_ACCOUNT="U0123504"
export CAPITALIST_API2_TO_ACCOUNT="U123456789"
capitalist_api2_create_payment_capitalist
```

## Source layout

`src/capitalist-api2.sh` is the compatibility entrypoint and loads the focused modules:

- `src/account/` for account list and transaction history.
- `src/exchange/` for exchange creation and rates.
- `src/kyc/` for KYC start, status, data, picture and confirmation.
- `src/merchant/` for merchant orders.
- `src/payments/create.sh` and `src/payments/create/` for payment creation.
- `src/payments/status.sh` for payment status lookup.
- `src/settings/` for IP whitelist management.

## Test

```bash
./tests/signature_test.sh
```
