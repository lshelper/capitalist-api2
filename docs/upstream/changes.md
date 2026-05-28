# Upstream changes observed on 2026-05-28

Compared with the repository API surface before storing an upstream snapshot, the current official documentation includes or emphasizes:

- Payment callbacks, including callback headers and signature verification.
- Callback body fields for final payment status notifications.
- `USDCBSC` and `USDTBSC` cryptocurrency channel types.
- A rate limiting recommendation to use callbacks instead of polling statuses more than 20 times per minute.
- A published signature example that differs from the repository test vector. The visible JSON sample contains an obfuscated email value, so the published hash is not reproducible from the visible body.
