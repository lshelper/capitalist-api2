# Payment channel helpers

Each `*.sh` file defines one `capitalist_api2_create_payment_<channel>` helper with an example payload matching the HTTP collection.

Helpers accept:

```text
userRequestId accountFrom amount currency [payloadAccount]
```

All arguments are optional. By default `accountFrom` is read from `CAPITALIST_API2_FROM_ACCOUNT`, currency from `CAPITALIST_API2_CURRENCY` or `USD`, and the amount is `100`.
