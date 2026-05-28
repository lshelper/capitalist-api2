# Capitalist API2 clients

Multi-language clients for the Capitalist API2 integration API.

Official documentation: https://docs.capitalist.net/api/integration-api.html

API base URL: `https://api2.capitalist.net/`

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
