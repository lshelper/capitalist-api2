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

## API documentation sync: 2026-10-01

See `../../docs/api-surface.md` for the updated transaction history contract, nested IMPS payload, and unresolved `inr_sum`/recipient requirements.

## Synthetic IMPS example

Build the client, then run the offline mock (no network or credentials required):

```bash
npm run build
node examples/imps-synthetic.mjs
```

The example defaults to 100 USD; use `--amount 20` to select 20 USD (positive decimal, at most two fractional digits). Recipient details are synthetic and taken from the public documentation. To send one request manually, first configure these environment variables outside the repository:

- `CAPITALIST_API2_BASE_URL`: explicitly selected HTTPS API origin, without `/v1`.
- `CAPITALIST_API2_KEY` and `CAPITALIST_API2_SECRET`: credentials.
- `CAPITALIST_API2_FROM_ACCOUNT`: verified USD source account.
- `CAPITALIST_API2_USER_REQUEST_ID`: unique identifier for this operation.

Then run `node examples/imps-synthetic.mjs --amount 20 --send` to submit 20 USD. This submits a payment to the configured environment; synthetic details alone do not establish a provider sandbox. The script sends once, logs the request ID, submission status/document ID and a redacted error reason, and never retries. After a timeout or error, reconcile by the supplied request ID before another submission. API keys, server addresses and source account data are not saved by this script.

Run example regression checks with `node --test examples/imps-synthetic.test.mjs` after building the client.
