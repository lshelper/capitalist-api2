# Capitalist API2 PHP 7.x client

## Usage

```php
<?php

require __DIR__ . '/vendor/autoload.php';

use Capitalist\Api2\Client;

$client = new Client(getenv('CAPITALIST_API2_KEY'), getenv('CAPITALIST_API2_SECRET'));
$accounts = $client->listAccounts('USD');
$depositAddress = $client->getDepositAddress('USDTb');
```

## Callback verification

Verify incoming payment callbacks against the raw request body:

```php
use Capitalist\Api2\Signature;

$valid = Signature::verify($timestampHeader, $rawBody, $apiSecret, $signatureHeader);
```

## Test

```bash
composer dump-autoload
php tests/SignatureTest.php
```

## Docker

Run the PHP 7.4 client without installing PHP locally:

```bash
docker compose build
docker compose run --rm php
```

Build from scratch:

```bash
docker compose build --no-cache
```

Run a specific command:

```bash
docker compose run --rm php php tests/SignatureTest.php
docker compose run --rm php php examples/list-accounts.php USD
```

Install or refresh Composer autoload inside the mounted workspace:

```bash
docker compose run --rm php composer install
docker compose run --rm php composer dump-autoload
```

For real API calls, pass credentials from your shell:

```bash
export CAPITALIST_API2_KEY="your-api-key"
export CAPITALIST_API2_SECRET="your-api-secret"
docker compose run --rm php php examples/list-accounts.php USD
```

Open a shell in the PHP container:

```bash
docker compose run --rm php bash
```

Clean up Docker resources for this client:

```bash
docker compose down -v
```
