import assert from 'node:assert/strict';
import test from 'node:test';
import { readFileSync } from 'node:fs';
import { spawnSync } from 'node:child_process';
import vm from 'node:vm';
const script = new URL('./imps-synthetic.mjs', import.meta.url);
const source = readFileSync(script, 'utf8');
const reason = vm.runInNewContext(`(${source.slice(source.indexOf('function safeApiReason'), source.indexOf('\nasync function main'))})`, {
  process: { env: { CAPITALIST_API2_KEY: 'test-private-key', CAPITALIST_API2_SECRET: 'test-private-secret', CAPITALIST_API2_FROM_ACCOUNT: 'test-source-account' } },
});
test('prints codec and nested validation errors', () => {
  assert.match(reason({ responseBody: 'Codec Error .payload.type.IMPS.inr_sum(missing)' }), /inr_sum\(missing\)/);
  assert.match(reason({ responseBody: '{"errors":{"destination.bankCode":["Required"]}}' }), /destination.bankCode: Required/);
  assert.match(reason({ responseBody: '{"error":{"message":"Invalid payload"}}' }), /Invalid payload/);
});
test('redacts credentials and account', () => {
  const result = reason({ responseBody: 'test-private-key test-private-secret test-source-account' });
  for (const value of ['test-private-key', 'test-private-secret', 'test-source-account']) assert(!result.includes(value));
});
test('accepts valid amounts offline and rejects invalid amounts', () => {
  for (const value of ['20', '20.25']) {
    const r = spawnSync(process.execPath, [script.pathname, '--amount', value], { encoding: 'utf8' });
    assert.equal(r.status, 0, r.stderr);
    assert(r.stdout.includes(`IMPS ${value} USD`));
    assert(r.stdout.includes('No network request sent'));
  }
  for (const value of ['0', '-20', '20.001', '1e2', 'NaN', 'Infinity', '9007199254740992']) {
    const r = spawnSync(process.execPath, [script.pathname, '--amount', value], { encoding: 'utf8' });
    assert.equal(r.status, 1);
    assert(!r.stdout.includes('Sending'));
  }
});
