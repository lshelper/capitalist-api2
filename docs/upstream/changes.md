# Upstream changes observed on 2026-05-28

Compared with the repository API surface before storing an upstream snapshot, the current official documentation includes or emphasizes:

- Payment callbacks, including callback headers and signature verification.
- Callback body fields for final payment status notifications.
- `USDCBSC` and `USDTBSC` cryptocurrency channel types.
- A rate limiting recommendation to use callbacks instead of polling statuses more than 20 times per minute.
- A published signature example that differs from the repository test vector. The visible JSON sample contains an obfuscated email value, so the published hash is not reproducible from the visible body.

# Upstream changes observed on 2026-06-10

Compared with the 2026-05-28 snapshot, the current official documentation adds:

- A `Cryptocurrency wallets` endpoint group.
- `GET /v1/depositAddress/{currency}` for cryptocurrency deposit addresses.
- `GET /v1/depositAddressAutoUSDTt/{account}` for USDT TRC-20 deposit addresses with autoconversion to a target account.
- `USDTb` and `USDCb` currency codes for BEP-20/BSC deposit addresses.
- A warning that deposit addresses are not permanent, must not be cached, and are guaranteed for at most 8 hours.

# Upstream changes observed on 2026-07-03

Compared with the 2026-06-10 snapshot, the current official documentation adds or changes:

- `txId` and `dstAddress` response fields for crypto payment status and transaction history responses.
- `GET /v1/prepaid2/services` and `GET /v1/prepaid2/regions` dictionaries for `TOPUP_SERVICE` payments.
- `GET /v1/prepaid2/products` and `GET /v1/prepaid2/denominations?productid={productId}` dictionaries for `BUY_ITEM` prepaid card payments.
- `TOPUP_SERVICE` payment payloads for Steam account topups and `BUY_ITEM` payment payloads for prepaid cards.
- Expanded `USDTb` and `USDCb` descriptions to explicitly name BEP-20 as Binance Smart Chain.
- Removed the `UKR_MOBILE` Ukrainian mobile payment channel from the documented payment channels.
- Updated the transactions example to use transaction type `OUT` and include a crypto `txId`.

# Upstream changes observed on 2026-09-01

Compared with the 2026-07-03 snapshot, the official page changes only its generated `Last updated` timestamp from 2026-07-02 13:08:04 UTC to 2026-07-16 14:27:48 UTC. No endpoint, request or response field, authentication or signature rule, callback, currency, payment channel, limit, warning, or example content changed. The maintained clients, examples, and shared API surface remain current without code changes.

# Upstream changes observed on 2026-10-01

Compared with the 2026-09-01 snapshot (upstream timestamp now 2026-09-30 10:12:32 UTC):

- Transaction history filter `transactionId` changes from number to string.
- Transaction response replaces the old identifier/date/version fields with `userRequestId`, `docType`, `fee`, `san`, `documentId`, and nullable `comment`; the example also includes `dstAddress`.
- IMPS replaces flat snake_case fields with a nested `destination`, adds `bankName` and example `recipient.phone`, and corrects IFSC format.
- No endpoint, signature, callback, currency, or rate-limit changes were found.
- `inr_sum` remains undocumented; recipient requiredness and `san` semantics remain unspecified.
