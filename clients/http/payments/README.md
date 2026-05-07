# Payment channel requests

Open one of these `.http` files in IntelliJ IDEA and run the request for the channel you need.

Each request signs the exact JSON body with:

```text
sha256(X-Request-Timestamp + raw request body + API-secret)
```

Some channels may be unavailable for a specific account or at a specific time. Replace placeholder payload values before running real payouts.
