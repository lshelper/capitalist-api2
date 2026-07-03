# Capitalist API2 TypeScript client

## Usage

```ts
import { CapitalistApi2Client } from './src';

const client = new CapitalistApi2Client({
  apiKey: process.env.CAPITALIST_API2_KEY!,
  apiSecret: process.env.CAPITALIST_API2_SECRET!,
});

const accounts = await client.listAccounts('USD');
const depositAddress = await client.getDepositAddress('USDTb');
const prepaidServices = await client.getPrepaidServices();
```

## Callback verification

Use the raw callback body exactly as it was received:

```ts
import { verifySignature } from './src';

const valid = verifySignature(timestampHeader, rawBody, apiSecret, signatureHeader);
```

## Test

```bash
npm install
npm test
```
