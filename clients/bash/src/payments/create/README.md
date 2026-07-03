# Payment channel helpers

Each `*.sh` file defines one `capitalist_api2_create_payment_<channel>` helper with an example payload matching the HTTP collection.

Most channel helpers accept:

```text
userRequestId accountFrom amount currency [payloadAccount]
```

All arguments are optional. By default `accountFrom` is read from `CAPITALIST_API2_FROM_ACCOUNT`, currency from `CAPITALIST_API2_CURRENCY` or `USD`, and the amount is `100`.

The prepaid helpers have extra payload arguments:

```text
capitalist_api2_create_payment_topup_service userRequestId accountFrom amount currency serviceId account quantity region
capitalist_api2_create_payment_buy_item userRequestId accountFrom amount currency productId denominationId
```

Calculate prepaid amounts from the dictionaries before sending real payments. For `TOPUP_SERVICE`, use `quantity * unfixedRate`; for `BUY_ITEM`, use the denomination `price`.
