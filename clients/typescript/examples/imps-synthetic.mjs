#!/usr/bin/env node
import assert from 'node:assert/strict';
import { CapitalistApi2Client } from '../dist/index.js';

// Synthetic data from the public IMPS documentation; never store credentials here.
const payload = {
  type: 'IMPS',
  destination: {
    bankCode: 'SBIN0001234',
    accountName: 'Raj Kumar',
    accountNumber: '0123456789012345',
    bankName: 'Punjab National Bank',
  },
  recipient: { phone: '12125551234' },
};

function required(name) {
  const value = process.env[name]?.trim();
  if (!value) throw new Error(`Set ${name}`);
  return value;
}

async function main() {
  const args = process.argv.slice(2);
  if (args.length > 1 || (args.length === 1 && args[0] !== '--send')) {
    throw new Error('Usage: node examples/imps-synthetic.mjs [--send]');
  }
  const send = args[0] === '--send';
  const payment = {
    userRequestId: send ? required('CAPITALIST_API2_USER_REQUEST_ID') : 'synthetic-imps-mock',
    accountFrom: send ? required('CAPITALIST_API2_FROM_ACCOUNT') : 'SYNTHETIC_USD_SOURCE',
    amount: 100,
    currency: 'USD',
    comment: 'Synthetic IMPS example',
    payload,
  };

  if (!send) {
    const client = new CapitalistApi2Client({
      baseUrl: 'https://example.test',
      apiKey: 'mock-key',
      apiSecret: 'mock-secret',
      fetchImpl: async (input, init) => {
        assert.equal(String(input), 'https://example.test/v1/payment');
        assert.equal(init?.method, 'POST');
        assert.deepEqual(JSON.parse(String(init?.body)), payment);
        return new Response('{"documentId":1}');
      },
    });
    assert.equal((await client.createPayment(payment)).documentId, 1);
    console.log('Mock passed: IMPS 100 USD with synthetic details. No network request sent.');
    return;
  }

  const baseUrl = new URL(required('CAPITALIST_API2_BASE_URL'));
  if (baseUrl.protocol !== 'https:' || baseUrl.username || baseUrl.password ||
      baseUrl.pathname !== '/' || baseUrl.search || baseUrl.hash) {
    throw new Error('CAPITALIST_API2_BASE_URL must be an HTTPS origin without /v1, credentials, query or fragment');
  }
  const client = new CapitalistApi2Client({
    baseUrl: baseUrl.origin,
    apiKey: required('CAPITALIST_API2_KEY'),
    apiSecret: required('CAPITALIST_API2_SECRET'),
  });
  console.log('Sending one IMPS request for 100 USD. Automatic retries are disabled.');
  try {
    const result = await client.createPayment(payment);
    console.log(JSON.stringify({ submitted: true, documentId: result.documentId }));
  } catch (error) {
    // Never print raw API errors: they may contain account or provider data.
    console.error(JSON.stringify({ submitted: 'unconfirmed', httpStatus: error.status ?? null }));
    console.error('Reconcile using CAPITALIST_API2_USER_REQUEST_ID before submitting again.');
    process.exitCode = 1;
  }
}

main().catch(() => {
  console.error('Example failed. Check arguments, required environment variables and the client build.');
  process.exitCode = 1;
});
