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
