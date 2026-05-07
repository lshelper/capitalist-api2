# Capitalist API2 TypeScript client

## Usage

```ts
import { CapitalistApi2Client } from './src';

const client = new CapitalistApi2Client({
  apiKey: process.env.CAPITALIST_API2_KEY!,
  apiSecret: process.env.CAPITALIST_API2_SECRET!,
});

const accounts = await client.listAccounts('USD');
```

## Test

```bash
npm install
npm test
```
