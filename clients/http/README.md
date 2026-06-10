# Capitalist API2 IntelliJ IDEA HTTP Client

Executable `.http` requests for the JetBrains HTTP Client in IntelliJ IDEA.

JetBrains documentation:

- https://www.jetbrains.com/help/idea/http-client-in-product-code-editor.html
- https://www.jetbrains.com/help/idea/http-client-variables.html

## Setup

Copy the example environment files:

```bash
cp http-client.env.json.example http-client.env.json
cp http-client.private.env.json.example http-client.private.env.json
```

Put non-secret values in `http-client.env.json` and credentials in `http-client.private.env.json`.

`http-client.private.env.json` is ignored by git.

## Run

Open `capitalist-api2.http` in IntelliJ IDEA and click the run gutter icon next to a request.

The pre-request handler imports `sign-capitalist.js` and sets per-request `timestamp` and `signature` variables.

The collection is split by operation area:

- `account/*.http` for account list and transaction history.
- `exchange/*.http` for exchange creation and rates.
- `kyc/*.http` for KYC start, status, data and confirmation.
- `merchant/*.http` for merchant orders.
- `payments/create/*.http` for payment creation by channel.
- `payments/status/*.http` for payment status lookup.
- `settings/*.http` for security settings exposed by the API, such as IP whitelist management.
- `wallets/*.http` for cryptocurrency deposit addresses.

Open the file for the operation you want to test and replace placeholder values before sending a real request.

Each request calculates:

```text
Signature = sha256(X-Request-Timestamp + raw request body + API-secret)
```

For `GET` requests, the body is an empty string.

The binary KYC picture upload endpoint is not included yet, because the signature must be calculated from the exact JPEG bytes. The JetBrains helper currently signs text request bodies.
