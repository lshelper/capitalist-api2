# Payment channel requests

Open one of these `.http` files in IntelliJ IDEA and run the request you need.

- `create/*.http` contains payment creation examples split by payment channel.
- `status/*.http` contains payment lookup requests by document ID or user request ID.

Each request signs the exact JSON body with:

```text
sha256(X-Request-Timestamp + raw request body + API-secret)
```

Some channels may be unavailable for a specific account or at a specific time. Replace placeholder payload values before running real payouts.

The official documentation no longer lists `UKR_MOBILE`; the example request was removed from this collection. Use the `prepaid2/*.http` dictionary requests to calculate amounts before running `TOPUP_SERVICE` or `BUY_ITEM` payment examples.
