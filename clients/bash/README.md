# Capitalist API2 Bash client

Small Bash client for scripts, CI jobs and manual operations.

## Usage

```bash
export CAPITALIST_API2_KEY="your-api-key"
export CAPITALIST_API2_SECRET="your-api-secret"

source ./src/capitalist-api2.sh

capitalist_api2_list_accounts "USD"
```

The client requires `curl` and either `openssl` or `shasum`.

## Test

```bash
./tests/signature_test.sh
```
