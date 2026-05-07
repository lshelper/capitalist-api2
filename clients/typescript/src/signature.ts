import { createHash, timingSafeEqual } from 'node:crypto';

export type SignableBody = string | ArrayBuffer | ArrayBufferView;

export function calculateSignature(timestamp: string | number, body: SignableBody, apiSecret: string): string {
  const hash = createHash('sha256');
  hash.update(String(timestamp));

  if (typeof body === 'string') {
    hash.update(body);
  } else if (body instanceof ArrayBuffer) {
    hash.update(Buffer.from(body));
  } else {
    hash.update(Buffer.from(body.buffer as ArrayBuffer, body.byteOffset, body.byteLength));
  }

  hash.update(apiSecret);
  return hash.digest('hex');
}

export function verifySignature(
  timestamp: string | number,
  body: SignableBody,
  apiSecret: string,
  signature: string,
): boolean {
  const expected = Buffer.from(calculateSignature(timestamp, body, apiSecret), 'hex');
  const actual = Buffer.from(signature, 'hex');

  return expected.length === actual.length && timingSafeEqual(expected, actual);
}
