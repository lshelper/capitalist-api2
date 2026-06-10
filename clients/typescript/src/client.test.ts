import assert from 'node:assert/strict';
import test from 'node:test';
import { CapitalistApi2Client } from './client.js';

test('gets cryptocurrency deposit address', async () => {
  const seen: { url?: string; body?: BodyInit | null } = {};
  const client = new CapitalistApi2Client({
    apiKey: 'key',
    apiSecret: 'secret',
    baseUrl: 'https://example.test',
    fetchImpl: async (input, init) => {
      seen.url = String(input);
      seen.body = init?.body;
      return new Response('{"address":"wallet-address"}');
    },
  });

  const result = await client.getDepositAddress('USDTb');

  assert.deepEqual(result, { address: 'wallet-address' });
  assert.equal(seen.url, 'https://example.test/v1/depositAddress/USDTb');
  assert.equal(seen.body, undefined);
});

test('gets autoconverted USDT TRC-20 deposit address', async () => {
  const seen: { url?: string; body?: BodyInit | null } = {};
  const client = new CapitalistApi2Client({
    apiKey: 'key',
    apiSecret: 'secret',
    baseUrl: 'https://example.test',
    fetchImpl: async (input, init) => {
      seen.url = String(input);
      seen.body = init?.body;
      return new Response('{"address":"wallet-address"}');
    },
  });

  const result = await client.getAutoConvertedUSDTtDepositAddress('U0123504');

  assert.deepEqual(result, { address: 'wallet-address' });
  assert.equal(seen.url, 'https://example.test/v1/depositAddressAutoUSDTt/U0123504');
  assert.equal(seen.body, undefined);
});
