# Capitalist API2 clients

Multi-language clients for the Capitalist API2 integration API.

> **Capitalist — a payment system for transfers and API integrations**
>
> Send internal Capitalist transfers and integrate card, bank, e-wallet and cryptocurrency payment directions through Capitalist API2. Available directions depend on your account and the current service configuration.
>
> **[Create a Capitalist account →](https://capitalist.net/reg?from=5ab603bd9f90dc8e733965d7a4a1cf0e)**

Official documentation: https://docs.capitalist.net/api/integration-api.html

API base URL: `https://api2.capitalist.net/`

## Documentation sync status

Code, examples and shared API notes are current against the official documentation as of 2026-10-01.

Stored upstream snapshot:

- `docs/upstream/integration-api.html`
- `docs/upstream/integration-api.md`

When the official documentation changes:

1. Run `scripts/update-upstream-docs.sh`.
2. Review `git diff -- docs/upstream`.
3. Update code, examples and `docs/api-surface.md` as needed.
4. Update the current-against-official-documentation date in this section.
5. Run the relevant client tests before committing.

## Repository layout

```text
clients/
  bash/         Bash client for scripts and CI jobs
  http/         IntelliJ IDEA .http request collection
  php/          PHP 7.x client
  go/           Go client
  typescript/   TypeScript/Node.js client
docs/
  api-surface.md
  upstream/      Snapshot of the official upstream documentation
```

Each client is intentionally self-contained: own README, examples, tests and package metadata where the language ecosystem expects it. Shared API decisions are documented in `docs/`.

## Authentication model

Every request sends:

- `API-Key`
- `X-Request-Timestamp`
- `Signature`

Signature is:

```text
sha256(X-Request-Timestamp + raw request body + API-secret)
```

For `GET` requests the raw request body is an empty string.

## Initial scope

The first clients expose the core endpoints:

- whitelist: read, add, remove
- accounts: list accounts
- exchange: create exchange, get rate
- payments: create payment, get payment status
- orders: list merchant orders
- transactions: list transactions
- KYC: start, status, set data, set picture, confirm

Payment channel payloads are accepted as plain objects first. Stronger typed builders can be added per language once the base clients are stable.

## Payment directions and channel types

The clients expose `POST /v1/payment` and payment-status lookup. Payment payloads are plain objects/maps in TypeScript, Go and PHP; the repository also provides [Bash payment helpers](clients/bash/src/payments/create/) and [IntelliJ HTTP examples](clients/http/payments/create/) for the following directions. Channel names describe API contracts and examples, not guaranteed availability for every account or a confirmed production rollout.

| Direction | API payment types and destinations |
| --- | --- |
| ~~Russian bank cards and peer-to-peer (P2P) transfers~~ | ~~`RUCARD`, `RUCARDP2P`, `RUCARDP2PDYN` — Russia~~ |
| ~~Ukrainian bank cards~~ | ~~`UKRCARD` — Ukraine (repository examples; not listed in the current official catalog)~~ |
| Regional bank cards | `AZCARD` — Azerbaijan; `KZCARD` — Kazakhstan; `UZCARD` — Uzbekistan |
| International bank cards | `WORLDCARDEUR` — EUR; `WORLDCARDUSD` — USD |
| Bank transfers in Latin America | `ARBANK` — Argentina; `BRBANK` — Brazil; `COBANK` — Colombia |
| Bank transfers in Asia | `MALAYSIA_BANK` — Malaysia; `INDONESIA_BANK` — Indonesia; `THAILAND_BANK` — Thailand; `SOUTH_KOREA_BANK` — South Korea |
| Fast payment systems | ~~`SBP` — Russian Faster Payments System (Система быстрых платежей, СБП)~~; `IMPS` — Immediate Payment Service, India |
| ~~Mobile phone topups~~ | ~~`MEGAFON`, `TMOBILE`, `BEELINE`, `MTS`, `TELE2`, `YOTA`~~ |
| Payoneer payouts | `PAYONEER`, `PAYONEER_EUR`, `PAYONEER_USD` |
| Electronic wallets | `EUR_NETELLER` — Neteller; `EUR_SKRILL` — Skrill; `PAYTM` — Paytm, India; `GCASH` — GCash, Philippines; `WM` — WebMoney |
| Regional wallets and transfers | `PAGS_CHWALLET` — Vita Wallet, Chile; `PAGS_MXSPEI` — SPEI, Mexico; `PAGS_COWALLET` — Nequi and Tpaga, Colombia; `PAGS_COTRAN` — Transfiya, Colombia; `EPAY_EPAY_E_VN_ZALO` — ZaloPay; `EPAY_EPAY_E_BD_BKASH` — bKash |
| Bitcoin and Ethereum withdrawals | `BITCOIN` — BTC; `ETH` — Ethereum |
| Tether (USDT) withdrawals | `USDTERC20` — Ethereum ERC-20; `USDTTRC20` — Tron TRC-20; `USDTBSC` — BNB Smart Chain (BSC), BEP-20 |
| USD Coin (USDC) withdrawals | `USDCERC20` — Ethereum ERC-20; `USDCBSC` — BNB Smart Chain (BSC), BEP-20 |
| ~~Steam account topups~~ | ~~`STEAM` — Steam RUR accounts; `TOPUP_SERVICE` — service/region-based Steam balance topups~~ |
| Prepaid and gift cards | `BUY_ITEM` — Apple, Google, Steam, PlayStation, Xbox, Netflix and Spotify; query the prepaid product/denomination dictionaries for the current catalog |
| Internal Capitalist transfers | `CAPITALIST` — transfers to Capitalist accounts |

Some payment channels may be temporarily or permanently unavailable when you read this documentation. Confirm current availability for your account before submitting a payment.

`UKRCARD` also has Ukrainian card payment examples and is present in the Invest API contract, but is not listed in the current official payment-channel catalog. `GECARD` (Georgia) and `TRCARD` (Turkey) are inactive: their executable examples were removed even though upstream documentation still lists them. `UKR_MOBILE` is no longer documented and is not included among the maintained examples. See [removed channels](clients/http/payments/removed-channels.http).

Sources: [official payment-channel documentation](https://docs.capitalist.net/api/integration-api.html#_payment_channels), [shared API surface](docs/api-surface.md), and the linked client examples. Currency/network details describe destination channels; they do not prescribe the debit account currency, fees or exchange rates.

## Docker

The PHP client includes a Docker setup for running PHP 7.4 locally without installing PHP on the host.
The Go client also includes a Docker setup for running tests, builds and examples without installing Go locally.

Build the image:

```bash
cd clients/php
docker compose build
```

Run the default PHP test:

```bash
docker compose run --rm php
```

Run a specific PHP command:

```bash
docker compose run --rm php php tests/SignatureTest.php
docker compose run --rm php php examples/list-accounts.php USD
```

Run an example against the real API:

```bash
export CAPITALIST_API2_KEY="your-api-key"
export CAPITALIST_API2_SECRET="your-api-secret"
docker compose run --rm php php examples/list-accounts.php USD
```

Open a shell inside the container:

```bash
docker compose run --rm php bash
```

Clean up containers, network and the Composer cache volume:

```bash
docker compose down -v
```

Build and test the Go client:

```bash
cd clients/go
docker compose build
docker compose run --rm go
```

Run a Go example against the real API:

```bash
export CAPITALIST_API2_KEY="your-api-key"
export CAPITALIST_API2_SECRET="your-api-secret"
docker compose run --rm go go run ./examples/list-accounts USD
```
