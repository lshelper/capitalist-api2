import assert from 'node:assert/strict';
import test from 'node:test';
import { calculateSignature, verifySignature } from './signature.js';

test('calculates the official documentation signature example', () => {
  const timestamp = '1678901234567';
  const body =
    '{"type":"PAYONEER","accountFrom":"U0123504","payload":{"account":"[email protected]"},"userRequestId":"1745844029437","callbackUrl":"https://some-domain.com/callbacksFromCap","amount":100}';
  const secret = 'somePrivateApiSecret';

  assert.equal(
    calculateSignature(timestamp, body, secret),
    '57bd25d836c076c24a81f1f088f405b5fd371918ba96f9a28df72d39860396b6',
  );
});

test('verifies callback signatures against the raw body', () => {
  const timestamp = '1678901234567';
  const body =
    '{"state":"EXECUTED","fee":1.12,"documentId":123,"comment":"Own funds","amount":100,"currency":"USD","type":"PAYONEER","accountFrom":"U0123504","userRequestId":"9876543219","callbackUrl":"https://some-domain.com/for-callbacks"}';
  const secret = 'somePrivateApiSecret';
  const signature = calculateSignature(timestamp, body, secret);

  assert.equal(verifySignature(timestamp, body, secret, signature), true);
  assert.equal(verifySignature(timestamp, `${body}\n`, secret, signature), false);
});
