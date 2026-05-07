<?php

require __DIR__ . '/../src/Signature.php';
require __DIR__ . '/../src/ApiException.php';
require __DIR__ . '/../src/Client.php';

use Capitalist\Api2\Client;

$client = new Client(getenv('CAPITALIST_API2_KEY'), getenv('CAPITALIST_API2_SECRET'));
$result = $client->listAccounts($argv[1] ?? 'USD');

echo json_encode($result, JSON_PRETTY_PRINT | JSON_UNESCAPED_SLASHES) . PHP_EOL;
