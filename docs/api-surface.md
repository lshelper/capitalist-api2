# API surface

This file keeps the language clients aligned.

## Base URL

`https://api2.capitalist.net/`

## Authentication

Headers:

- `API-Key`
- `X-Request-Timestamp`
- `Signature`

Signature:

```text
sha256(timestamp + rawBody + apiSecret)
```

Repository test vector, based on the documentation example:

```text
timestamp=1678901234567
secret=somePrivateApiSecret
body={"type":"PAYONEER","accountFrom":"U0123504","payload":{"account":"[email protected]"},"userRequestId":"1745844029437","callbackUrl":"https://some-domain.com/callbacksFromCap","amount":100}
signature=57bd25d836c076c24a81f1f088f405b5fd371918ba96f9a28df72d39860396b6
```

Note: the official HTML page obfuscates the email inside the sample JSON. Because the signature depends on the exact raw body, this repository uses the visible placeholder string as its cross-language test vector. The upstream page currently publishes a different example signature, but it cannot be reproduced from the visible JSON sample.

## Endpoints

| Operation | Method | Path |
| --- | --- | --- |
| Read whitelist | GET | `/v1/whitelist` |
| Add IPs to whitelist | POST | `/v1/whitelist` |
| Remove IPs from whitelist | POST | `/v1/whitelist/remove` |
| List accounts | GET | `/v1/account/list?currency={currency}` |
| Create exchange | POST | `/v1/exchange` |
| Get exchange rate | GET | `/v1/rate?from={from}&to={to}` |
| Create payment | POST | `/v1/payment` |
| Get payment by document ID | GET | `/v1/payment/document/{documentId}` |
| Get payment by user request ID | GET | `/v1/payment/{userRequestId}` |
| Get orders | GET | `/v1/orders` |
| Get transactions | GET | `/v1/transactions` |
| Start KYC | POST | `/v1/kyc/start` |
| Get KYC status | GET | `/v1/kyc/status/{kycExternalUserId}/{sort}` |
| Get KYC status by UUID | GET | `/v1/kyc/statusByUuid/{uuid}` |
| Set KYC data | POST | `/v1/kyc/setData/{uuid}` |
| Set KYC picture | POST | `/v1/kyc/setPicture/{uuid}/{type}` |
| Confirm KYC | POST | `/v1/kyc/confirm/{uuid}` |

## Payment callbacks

When a payment request includes `callbackUrl`, Capitalist sends a `POST` request to that URL after the payment reaches a final status.

Callback headers:

- `X-Request-Timestamp`
- `Signature`

Callback signature:

```text
sha256(timestamp + rawCallbackBody + apiSecret)
```

Important: verify the signature against the exact raw callback body bytes received from the HTTP request. Do not parse and re-serialize JSON before verification.

Callback body fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| `state` | string | Yes | Final document state, such as `EXECUTED` or `DECLINED`. |
| `fee` | number | Yes | Transaction fee amount. |
| `documentId` | integer | Yes | Transaction document identifier. |
| `comment` | string | No | Additional transaction information. |
| `amount` | number | Yes | Transaction amount. |
| `currency` | string | Yes | Transaction currency code. |
| `type` | string | Yes | Payment channel type. |
| `accountFrom` | string | Yes | Source account identifier. |
| `userRequestId` | string | Yes | Unique request identifier supplied when creating the payment. |
| `callbackUrl` | string | Yes | Callback URL used for notifications. |

## Payment channels

The API accepts payment-specific data in the `payload` object. This repository keeps examples for the documented channel types in Bash and IntelliJ HTTP Client form.

Cryptocurrency channel types documented upstream:

- `BITCOIN`
- `ETH`
- `USDCERC20`
- `USDTERC20`
- `USDTTRC20`
- `USDCBSC`
- `USDTBSC`

## Rate limiting

The upstream documentation recommends using callbacks instead of polling payment status more than 20 times per minute. Excessively frequent API calls may be throttled.

## Client design rules

- Preserve raw JSON body for signature calculation.
- Preserve raw callback body for callback signature verification.
- Keep credentials out of logs and examples.
- Treat non-2xx HTTP responses as client errors.
- Parse JSON when possible, but allow empty `200` responses.
- Keep method names consistent across languages.

## Language clients

| Client | Location | Notes |
| --- | --- | --- |
| Bash | `clients/bash` | Sourceable shell functions around `curl`. |
| IntelliJ IDEA HTTP Client | `clients/http` | Executable `.http` collection with pre-request SHA-256 signing. |
| PHP 7.x | `clients/php` | PSR-4 package using `ext-curl`. |
| Go | `clients/go` | Go client using only the standard library. |
| TypeScript/Node.js | `clients/typescript` | Strict TypeScript client using `fetch`. |
