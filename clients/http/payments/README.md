# Payment channel requests

Open one of these `.http` files in IntelliJ IDEA and run the request you need.

- `create/*.http` contains payment creation examples split by payment channel.
- `status/*.http` contains payment lookup requests by document ID or user request ID.

Each request signs the exact JSON body with:

```text
sha256(X-Request-Timestamp + raw request body + API-secret)
```

Some channels may be unavailable for a specific account or at a specific time. Replace placeholder payload values before running real payouts.
