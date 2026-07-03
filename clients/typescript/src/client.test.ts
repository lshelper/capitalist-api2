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

test('gets prepaid services dictionary', async () => {
  const seen: { url?: string; body?: BodyInit | null } = {};
  const client = new CapitalistApi2Client({
    apiKey: 'key',
    apiSecret: 'secret',
    baseUrl: 'https://example.test',
    fetchImpl: async (input, init) => {
      seen.url = String(input);
      seen.body = init?.body;
      return new Response('[{"id":3,"name":"Steam","description":"Top up","hasRegion":false,"type":"unfixed","unfixedCurrency":"KZT","unfixedRate":0.00205482}]');
    },
  });

  const result = await client.getPrepaidServices();

  assert.equal(result[0]?.id, 3);
  assert.equal(result[0]?.type, 'unfixed');
  assert.equal(seen.url, 'https://example.test/v1/prepaid2/services');
  assert.equal(seen.body, undefined);
});

test('gets prepaid regions dictionary', async () => {
  const seen: { url?: string; body?: BodyInit | null } = {};
  const client = new CapitalistApi2Client({
    apiKey: 'key',
    apiSecret: 'secret',
    baseUrl: 'https://example.test',
    fetchImpl: async (input, init) => {
      seen.url = String(input);
      seen.body = init?.body;
      return new Response('["Asia","Macau"]');
    },
  });

  const result = await client.getPrepaidRegions();

  assert.deepEqual(result, ['Asia', 'Macau']);
  assert.equal(seen.url, 'https://example.test/v1/prepaid2/regions');
  assert.equal(seen.body, undefined);
});

test('gets prepaid products dictionary', async () => {
  const seen: { url?: string; body?: BodyInit | null } = {};
  const client = new CapitalistApi2Client({
    apiKey: 'key',
    apiSecret: 'secret',
    baseUrl: 'https://example.test',
    fetchImpl: async (input, init) => {
      seen.url = String(input);
      seen.body = init?.body;
      return new Response('[{"id":98,"name":"PlayStation Store Wallet United Kingdom","category":"Gift Card, PlayStation","description":"PlayStation card","instruction":"Redeem instructions"}]');
    },
  });

  const result = await client.getPrepaidProducts();

  assert.equal(result[0]?.id, 98);
  assert.equal(result[0]?.category, 'Gift Card, PlayStation');
  assert.equal(seen.url, 'https://example.test/v1/prepaid2/products');
  assert.equal(seen.body, undefined);
});

test('gets prepaid denominations dictionary by product ID', async () => {
  const seen: { url?: string; body?: BodyInit | null } = {};
  const client = new CapitalistApi2Client({
    apiKey: 'key',
    apiSecret: 'secret',
    baseUrl: 'https://example.test',
    fetchImpl: async (input, init) => {
      seen.url = String(input);
      seen.body = init?.body;
      return new Response('[{"id":527,"name":"10 GBP","price":13.45,"stock":5,"value":"10 GBP"}]');
    },
  });

  const result = await client.getPrepaidDenominations(98);

  assert.equal(result[0]?.id, 527);
  assert.equal(result[0]?.price, 13.45);
  assert.equal(seen.url, 'https://example.test/v1/prepaid2/denominations?productid=98');
  assert.equal(seen.body, undefined);
});

test('parses crypto metadata fields on payment status', async () => {
  const client = new CapitalistApi2Client({
    apiKey: 'key',
    apiSecret: 'secret',
    baseUrl: 'https://example.test',
    fetchImpl: async () =>
      new Response(
        JSON.stringify({
          state: 'EXECUTED',
          fee: 0,
          documentId: 123,
          amount: 10,
          currency: 'USDTb',
          type: 'USDTBSC',
          accountFrom: 'U0123504',
          payload: '{}',
          userRequestId: 'request-1',
          txId: 'tx-hash',
          dstAddress: 'wallet-address',
        }),
      ),
  });

  const result = await client.getPaymentByDocumentId(123);

  assert.equal(result.txId, 'tx-hash');
  assert.equal(result.dstAddress, 'wallet-address');
});

test('parses crypto metadata fields on transactions', async () => {
  const client = new CapitalistApi2Client({
    apiKey: 'key',
    apiSecret: 'secret',
    baseUrl: 'https://example.test',
    fetchImpl: async () =>
      new Response(
        JSON.stringify({
          transactions: [
            {
              transactionId: 8367664,
              createDate: '2014-02-18T21:47:29.452Z',
              executeDate: '2014-02-18T21:47:39.543Z',
              type: 'OUT',
              state: 'EXECUTED',
              amount: 500,
              currency: 'USD',
              planDate: '2014-02-18T21:21:53.314Z',
              version: 1,
              txId: 'tx-hash',
              dstAddress: 'wallet-address',
            },
          ],
          count: 1,
        }),
      ),
  });

  const result = await client.getTransactions();

  assert.equal(result.transactions[0]?.txId, 'tx-hash');
  assert.equal(result.transactions[0]?.dstAddress, 'wallet-address');
});
