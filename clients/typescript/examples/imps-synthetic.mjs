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

function safeApiReason(error) {
  function messages(value, path = '', depth = 0) {
    if (depth > 5 || value == null) return [];
    if (typeof value === 'string') return [path ? `${path}: ${value}` : value];
    if (Array.isArray(value)) return value.flatMap(item => messages(item, path, depth + 1));
    if (typeof value === 'object') {
      return Object.entries(value).flatMap(([key, item]) =>
        messages(item, path ? `${path}.${key}` : key, depth + 1));
    }
    return [];
  }
  let reason;
  try {
    const body = JSON.parse(error.responseBody);
    if (typeof body === 'string') reason = body;
    else if (body && typeof body === 'object') {
      reason = ['error', 'message', 'reason', 'description', 'errors', 'details']
        .flatMap(key => messages(body[key], key)).concat(messages(body.model?.errors, 'model.errors'))
        .join('; ');
    }
  } catch {
    // Plain-text errors are useful too; strip markup from gateway HTML errors.
    reason = error.responseBody?.replace(/<[^>]*>/g, ' ').replace(/\s+/g, ' ').trim();
  }
  reason ||= error.responseBody ? 'API rejected the request without an error description' : error.message;
  reason ||= 'Request failed without an error description';
  for (const name of ['CAPITALIST_API2_KEY', 'CAPITALIST_API2_SECRET', 'CAPITALIST_API2_FROM_ACCOUNT']) {
    const value = process.env[name];
    if (value) reason = reason.split(value).join('[redacted]');
  }
  return reason
    .replace(/(?:api[-_ ]?key|api[-_ ]?secret|signature|authorization|password|token)\s*[:=]\s*[^\s,;]+/gi, '[redacted credential]')
    .replace(/[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}/gi, '[email]')
    .replace(/[A-Za-z0-9+/_=-]{24,}/g, '[redacted]')
    .replace(/\b\d{6,}\b/g, '[redacted]')
    .slice(0, 1000);
}

async function main() {
  const args = process.argv.slice(2);
  let send = false;
  let amountText = '100';
  let amountProvided = false;
  for (let i = 0; i < args.length; i += 1) {
    if (args[i] === '--send' && !send) send = true;
    else if (args[i] === '--amount' && !amountProvided) {
      amountText = args[++i] ?? '';
      amountProvided = true;
    } else throw new Error('Usage: node examples/imps-synthetic.mjs [--amount USD_AMOUNT] [--send]');
  }
  if (!/^(0|[1-9][0-9]*)(\.[0-9]{1,2})?$/.test(amountText)) {
    throw new Error('Amount must be a positive decimal with at most two fractional digits');
  }
  const cents = BigInt(amountText.replace('.', '') + '0'.repeat(2 - (amountText.split('.')[1]?.length ?? 0)));
  if (cents <= 0n || cents > BigInt(Number.MAX_SAFE_INTEGER)) {
    throw new Error('Amount is out of range');
  }
  const amount = Number(amountText);
  const payment = {
    userRequestId: send ? required('CAPITALIST_API2_USER_REQUEST_ID') : 'synthetic-imps-mock',
    accountFrom: send ? required('CAPITALIST_API2_FROM_ACCOUNT') : 'SYNTHETIC_USD_SOURCE',
    amount,
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
    console.log(`Mock passed: IMPS ${amountText} USD with synthetic details. No network request sent.`);
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
  console.log(JSON.stringify({ userRequestId: payment.userRequestId }));
  console.log(`Sending one IMPS request for ${amountText} USD. Automatic retries are disabled.`);
  try {
    const result = await client.createPayment(payment);
    console.log(JSON.stringify({ submitted: true, documentId: result.documentId }));
  } catch (error) {
    // Never print raw API errors: they may contain account or provider data.
    console.error(JSON.stringify({ submitted: 'unconfirmed', httpStatus: error.status ?? null, reason: safeApiReason(error), userRequestId: payment.userRequestId }));
    console.error('Reconcile using CAPITALIST_API2_USER_REQUEST_ID before submitting again.');
    process.exitCode = 1;
  }
}

main().catch((error) => {
  console.error(`Example failed: ${safeApiReason(error)}`);
  process.exitCode = 1;
});
