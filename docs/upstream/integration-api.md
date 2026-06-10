<div id="header">

# Integration API Documentation

<div id="toc" class="toc2">

<div id="toctitle">

Table of Contents

</div>

- [1. Introduction](#_introduction)
- [2. Authentication](#_authentication)
  - [2.1. Authentication Headers](#_authentication_headers)
  - [2.2. Signature Calculation Example](#_signature_calculation_example)
    - [2.2.1. Input Data](#_input_data)
    - [2.2.2. Signature Calculation Steps](#_signature_calculation_steps)
- [3. Security: IP Whitelist](#_security_ip_whitelist)
  - [3.1. Read Whitelist](#_read_whitelist)
  - [3.2. Add IPs to Whitelist](#_add_ips_to_whitelist)
  - [3.3. Remove IPs from Whitelist](#_remove_ips_from_whitelist)
- [4. API Endpoints](#_api_endpoints)
  - [4.1. Account Management](#_account_management)
    - [4.1.1. List Accounts](#_list_accounts)
  - [4.2. Exchange Operations](#_exchange_operations)
    - [4.2.1. Create Currency Exchange](#_create_currency_exchange)
  - [4.3. Exchange Rate](#_exchange_rate)
    - [4.3.1. Get Exchange Rate](#_get_exchange_rate)
  - [4.4. Payment Operations](#_payment_operations)
    - [4.4.1. Create Payment](#_create_payment)
    - [4.4.2. Get Payment Status by Document ID](#_get_payment_status_by_document_id)
    - [4.4.3. Get Payment Status by User Request ID](#_get_payment_status_by_user_request_id)
  - [4.5. Merchant Orders](#_merchant_orders)
    - [4.5.1. Get Orders](#_get_orders)
  - [4.6. List of transactions](#_list_of_transactions)
    - [4.6.1. Get Transactions](#_get_transactions)
  - [4.7. Cryptocurrency wallets](#_cryptocurrency_wallets)
    - [4.7.1. Get address for cryptocurrency deposit](#_get_address_for_cryptocurrency_deposit)
    - [4.7.2. Get address for cryptocurrency deposit (in USDT TRC-20) with autoconversion](#_get_address_for_cryptocurrency_deposit_in_usdt_trc_20_with_autoconversion)
  - [4.8. KYC](#_kyc)
    - [4.8.1. Initiate new KYC](#_initiate_new_kyc)
    - [4.8.2. Query KYC status](#_query_kyc_status)
    - [4.8.3. Query KYC status by UUID](#_query_kyc_status_by_uuid)
    - [4.8.4. Fill KYC data](#_fill_kyc_data)
    - [4.8.5. Add KYC picture](#_add_kyc_picture)
    - [4.8.6. Send KYC for verification](#_send_kyc_for_verification)
- [5. Payment Callbacks](#_payment_callbacks)
  - [5.1. Callback Headers](#_callback_headers)
  - [5.2. Callback Body Fields](#_callback_body_fields)
  - [5.3. Example Callback Body](#_example_callback_body)
- [6. Payment Channels](#_payment_channels)
  - [6.1. Card Payments](#_card_payments)
    - [6.1.1. RUCARD - Russian Cards](#_rucard_russian_cards)
    - [6.1.2. RUCARDP2P / RUCARDP2PDYN - Russian P2P Card Transfers](#_rucardp2p_rucardp2pdyn_russian_p2p_card_transfers)
    - [6.1.3. TRCARD - Turkish Cards](#_trcard_turkish_cards)
    - [6.1.4. GECARD - Georgian Cards](#_gecard_georgian_cards)
    - [6.1.5. AZCARD - Azerbaijani Cards](#_azcard_azerbaijani_cards)
    - [6.1.6. KZCARD - Kazakhstani Cards](#_kzcard_kazakhstani_cards)
    - [6.1.7. UZCARD - Uzbekistani Cards](#_uzcard_uzbekistani_cards)
    - [6.1.8. WORLDCARDEUR / WORLDCARDUSD - International Cards](#_worldcardeur_worldcardusd_international_cards)
  - [6.2. Bank Transfers](#_bank_transfers)
    - [6.2.1. ARBANK - Argentinian Banks](#_arbank_argentinian_banks)
    - [6.2.2. BRBANK - Brazilian Banks](#_brbank_brazilian_banks)
    - [6.2.3. COBANK - Colombian Banks](#_cobank_colombian_banks)
    - [6.2.4. MALAYSIA_BANK - Malaysian Banks](#_malaysia_bank_malaysian_banks)
    - [6.2.5. INDONESIA_BANK - Indonesian Banks](#_indonesia_bank_indonesian_banks)
    - [6.2.6. THAILAND_BANK - Thai Banks](#_thailand_bank_thai_banks)
    - [6.2.7. SOUTH_KOREA_BANK - South Korean Banks](#_south_korea_bank_south_korean_banks)
  - [6.3. Mobile Payments](#_mobile_payments)
    - [6.3.1. MEGAFON / TMOBILE / BEELINE / MTS / TELE2 / YOTA](#_megafon_tmobile_beeline_mts_tele2_yota)
    - [6.3.2. UKR_MOBILE - Ukrainian Mobile Payments](#_ukr_mobile_ukrainian_mobile_payments)
  - [6.4. Fast Payment Systems](#_fast_payment_systems)
    - [6.4.1. SBP - Russian Fast Payment System / "Система Быстрых Платежей"](#_sbp_russian_fast_payment_system_система_быстрых_платежей)
    - [6.4.2. IMPS - Immediate Payment Service (India)](#_imps_immediate_payment_service_india)
  - [6.5. E-Wallets](#_e_wallets)
    - [6.5.1. PAYONEER / PAYONEER_EUR / PAYONEER_USD](#_payoneer_payoneer_eur_payoneer_usd)
    - [6.5.2. EUR_NETELLER / EUR_SKRILL](#_eur_neteller_eur_skrill)
    - [6.5.3. PAYTM](#_paytm)
    - [6.5.4. GCASH](#_gcash)
    - [6.5.5. WM](#_wm)
    - [6.5.6. PAGS_CHWALLET](#_pags_chwallet)
    - [6.5.7. PAGS_MXSPEI](#_pags_mxspei)
    - [6.5.8. PAGS_COWALLET](#_pags_cowallet)
    - [6.5.9. PAGS_COTRAN](#_pags_cotran)
    - [6.5.10. EPAY_EPAY_E_VN_ZALO](#_epay_epay_e_vn_zalo)
    - [6.5.11. EPAY_EPAY_E_BD_BKASH](#_epay_epay_e_bd_bkash)
  - [6.6. Cryptocurrency](#_cryptocurrency)
    - [6.6.1. BITCOIN / ETH / USDCERC20 / USDTERC20 / USDTTRC20 / USDCBSC / USDTBSC](#_bitcoin_eth_usdcerc20_usdterc20_usdttrc20_usdcbsc_usdtbsc)
  - [6.7. STEAM](#_steam)
    - [6.7.1. STEAM RUR ACCOUNTS](#_steam_rur_accounts)
  - [6.8. CAPITALIST](#_capitalist)
    - [6.8.1. CAPITALIST INTERNAL PAYMENTS](#_capitalist_internal_payments)
- [7. Error Handling](#_error_handling)
  - [7.1. Error Response](#_error_response)
  - [7.2. HTTP Status Codes](#_http_status_codes)
- [8. Reference](#_reference)
  - [8.1. Currency](#_currency)
- [9. Best Practices](#_best_practices)
  - [9.1. Security](#_security)
  - [9.2. Integration Tips](#_integration_tips)
  - [9.3. Rate Limiting](#_rate_limiting)
- [10. Coding Examples](#_coding_examples)
- [11. Support](#_support)

</div>

</div>

<div id="content">

<div class="sect1">

## 1. Introduction

<div class="sectionbody">

<div class="paragraph">

To work with API, you need to get API-Key and API-secret in your personal account.

</div>

<div class="paragraph">

**API URL:** <a href="https://api2.capitalist.net/" class="bare"><code>https://api2.capitalist.net/</code></a>

</div>

</div>

</div>

<div class="sect1">

## 2. Authentication

<div class="sectionbody">

<div class="paragraph">

With each API call, you need to pass the following headers:

</div>

<div class="sect2">

### 2.1. Authentication Headers

| Header | Description |
|----|----|
| `API-Key` | Your API key can be generated in your personal account under the ‘Security’ section (<a href="https://capitalist.net/security" class="bare">https://capitalist.net/security</a>). Please note that this requires enabling API access and activating Google 2FA on your account. |
| `X-Request-Timestamp` | Current timestamp in epoch milliseconds |
| `Signature` | SHA-256 hex hash of: `X-Request-Timestamp + request body + API-secret` |

</div>

<div class="sect2">

### 2.2. Signature Calculation Example

<div class="sect3">

#### 2.2.1. Input Data

<div class="ulist">

- **Request timestamp:** `1678901234567`

- **Request body:**

  <div class="listingblock">

  <div class="content">

  ``` highlightjs
  {"type":"PAYONEER","accountFrom":"U0123504","payload":{"account":"[email protected]"},"userRequestId":"1745844029437","callbackUrl":"https://some-domain.com/callbacksFromCap","amount":100}
  ```

  </div>

  </div>

- **API secret:** `somePrivateApiSecret`

</div>

</div>

<div class="sect3">

#### 2.2.2. Signature Calculation Steps

<div class="olist arabic">

1.  **Concatenate** the following components as a single string:

    <div class="listingblock">

    <div class="content">

        rawString = X-Request-Timestamp + requestBody + API-secret

    </div>

    </div>

2.  In this case:

    <div class="listingblock">

    <div class="content">

        1678901234567{"type":"PAYONEER","accountFrom":"U0123504","payload":{"account":"[email protected]"},"userRequestId":"1745844029437","callbackUrl":"https://some-domain.com/callbacksFromCap","amount":100}somePrivateApiSecret

    </div>

    </div>

3.  **Calculate SHA-256 hex** of the string above:

    <div class="listingblock">

    <div class="content">

    ``` highlightjs
    val signature = sha256Hex(input)
    ```

    </div>

    </div>

4.  **Resulting signature** (SHA-256 digest as a hex string):

    <div class="listingblock">

    <div class="content">

        2eec8f9fcedae5de9a4ebba8e4ef393ece034a1a14a28a51ee34812311e5b516

    </div>

    </div>

5.  **Final Header to send:**

    <div class="listingblock">

    <div class="content">

        Signature: 2eec8f9fcedae5de9a4ebba8e4ef393ece034a1a14a28a51ee34812311e5b516

    </div>

    </div>

</div>

</div>

</div>

</div>

</div>

<div class="sect1">

## 3. Security: IP Whitelist

<div class="sectionbody">

<div class="paragraph">

API methods may be restricted to a whitelist of IP addresses. Requests coming from IPs that are not in the whitelist will be rejected.

</div>

<div class="paragraph">

You can manage the whitelist via dedicated endpoints below. The request body accepts an array of strings with exact IPs or wildcard patterns for the last octet (e.g., `127.0.0.*`).

</div>

<div class="sect2">

### 3.1. Read Whitelist

<div class="paragraph">

**Endpoint:** `GET /v1/whitelist`

</div>

<div class="paragraph">

Returns an array of IPs.

</div>

##### Response Body

<div class="listingblock">

<div class="content">

``` highlightjs
["12.12.11.14"]
```

</div>

</div>

</div>

<div class="sect2">

### 3.2. Add IPs to Whitelist

<div class="paragraph">

**Endpoint:** `POST /v1/whitelist`

</div>

##### Request Body

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "ip": ["12.12.11.14"]
}
```

</div>

</div>

##### Request Fields

| Field | Type | Required | Description |
|----|----|----|----|
| `ip` | array\[string\] | Yes | List of IPs or patterns to allow (e.g., `"12.12.11.14"`, `"127.0.0.*"`). |

</div>

<div class="sect2">

### 3.3. Remove IPs from Whitelist

<div class="paragraph">

**Endpoint:** `POST /v1/whitelist/remove`

</div>

##### Request Body

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "ip": [
      "127.0.0.*",
      "127.0.0.1"
  ]
}
```

</div>

</div>

##### Request Fields

| Field | Type | Required | Description |
|----|----|----|----|
| `ip` | array\[string\] | Yes | List of IPs or patterns to remove from the whitelist. |

</div>

</div>

</div>

<div class="sect1">

## 4. API Endpoints

<div class="sectionbody">

<div class="sect2">

### 4.1. Account Management

<div class="sect3">

#### 4.1.1. List Accounts

<div class="paragraph">

**Endpoint:** `GET /v1/account/list`

</div>

<div class="paragraph">

Retrieves a list of accounts filtered by currency.

</div>

<div class="sect4">

##### Request Parameters

<table class="tableblock frame-all grid-all stretch" style="width:100%;">
<colgroup>
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th class="tableblock halign-left valign-top">Parameter</th>
<th class="tableblock halign-left valign-top">Type</th>
<th class="tableblock halign-left valign-top">Required</th>
<th class="tableblock halign-left valign-top">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><p><code>currency</code></p></td>
<td class="tableblock halign-left valign-top"><p>string</p></td>
<td class="tableblock halign-left valign-top"><p>Yes</p></td>
<td class="tableblock halign-left valign-top"><p>Currency code to filter accounts (e.g., USD, EUR, RUR, BTC, USDT, USDTt, USDC).</p>
<p>Note: USDT - for ERC-20, USDTt - for TRC-20R)</p></td>
</tr>
</tbody>
</table>

</div>

<div class="sect4">

##### Response

<div class="paragraph">

Returns an array of account objects.

</div>

<div class="listingblock">

<div class="content">

``` highlightjs
[
  {
    "number": "U0123504",
    "currency": "USD",
    "name": "Main Account",
    "balance": 1250.50
  }
]
```

</div>

</div>

</div>

<div class="sect4">

##### Response Fields

| Field      | Type   | Description               |
|------------|--------|---------------------------|
| `number`   | string | Account number identifier |
| `currency` | string | Account currency code     |
| `name`     | string | Account display name      |
| `balance`  | number | Current account balance   |

</div>

</div>

</div>

<div class="sect2">

### 4.2. Exchange Operations

<div class="sect3">

#### 4.2.1. Create Currency Exchange

<div class="paragraph">

**Endpoint:** `POST /v1/exchange`

</div>

<div class="paragraph">

Creates a currency conversion/exchange operation between accounts.

</div>

<div class="sect4">

##### Request Body

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "fromAccount": "U0123504",
  "toAccount": "E0123505",
  "amount": 100.00
}
```

</div>

</div>

</div>

<div class="sect4">

##### Request Fields

| Field         | Type   | Required | Description                            |
|---------------|--------|----------|----------------------------------------|
| `fromAccount` | string | Yes      | Source account number                  |
| `toAccount`   | string | Yes      | Destination account number             |
| `amount`      | number | Yes      | Amount to exchange from source account |

</div>

<div class="sect4">

##### Response

<div class="paragraph">

Returns `null` on success.

</div>

</div>

</div>

</div>

<div class="sect2">

### 4.3. Exchange Rate

<div class="sect3">

#### 4.3.1. Get Exchange Rate

<div class="paragraph">

**Endpoint:** `GET /v1/rate`

</div>

<div class="paragraph">

Retrieves the current exchange rate between two currencies.

</div>

<div class="sect4">

##### Request Parameters

| Parameter | Type   | Required | Description          |
|-----------|--------|----------|----------------------|
| `from`    | string | Yes      | Source currency code |
| `to`      | string | Yes      | Target currency code |

</div>

<div class="sect4">

##### Response

<div class="paragraph">

Returns the exchange rate as a number.

</div>

<div class="listingblock">

<div class="content">

``` highlightjs
1.0842
```

</div>

</div>

</div>

</div>

</div>

<div class="sect2">

### 4.4. Payment Operations

<div class="sect3">

#### 4.4.1. Create Payment

<div class="paragraph">

**Endpoint:** `POST /v1/payment`

</div>

<div class="paragraph">

Creates a new payment transaction.

</div>

<div class="paragraph">

The payment type is determined by the **payload** object — its structure defines the specific payment flow and validation rules.

</div>

<div class="paragraph">

Optionally, you may include a **callbackUrl** field. If provided, the system will send a callback request to this URL with the final payment status once the transaction is completed.

</div>

<div class="admonitionblock warning">

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><div class="title">
Warning
</div></td>
<td class="content">Not all payment channels listed in the documentation may be currently available.</td>
</tr>
</tbody>
</table>

</div>

<div class="sect4">

##### Request Body

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "userRequestId": "string",
  "accountFrom": "string",
  "amount": 0.1,
  "currency": "string",
  "comment": "string",
  "callbackUrl": "string",
  "payload": {
    "type": "SOME_BANK_TYPE",
    "account": "string",
    "taxId": "string",
    "cbu": "string"
  }
}
```

</div>

</div>

</div>

<div class="sect4">

##### Request Fields

| Field | Type | Required | Description |
|----|----|----|----|
| `userRequestId` | string | Yes | Your unique transaction identifier |
| `accountFrom` | string | Yes | Your account number from which funds will be withdrawn |
| `amount` | number | Yes | Transaction amount in account currency or in the specified currency |
| `currency` | string | No | Currency code (e.g., USD, EUR, RUR, BTC, USDT, USDTt, USDC). Note: USDT - for ERC-20, USDTt - for TRC-20 |
| `comment` | string | No | Optional transaction comment |
| `callbackUrl` | string | No | URL where the system will send callback notifications about transaction status |
| `payload` | object | Yes | Payment-specific data. Required fields depend on the payment channel type |

</div>

<div class="sect4">

##### Response

<div class="paragraph">

If the request is successful, the API returns a JSON object containing the unique document ID of the created payment.

</div>

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "documentId": 9007199254740991
}
```

</div>

</div>

</div>

<div class="sect4">

##### Response Fields

| Field        | Type    | Description                                       |
|--------------|---------|---------------------------------------------------|
| `documentId` | integer | Unique identifier of the created payment document |

</div>

</div>

<div class="sect3">

#### 4.4.2. Get Payment Status by Document ID

<div class="paragraph">

**Endpoint:** `GET /v1/payment/document/{documentId}`

</div>

<div class="paragraph">

Retrieves the current status and details of a payment by its document ID. The response also includes all payment fields as originally submitted, as well as the calculated fee associated with the transaction.

</div>

<div class="sect4">

##### Path Parameters

| Parameter | Type | Required | Description |
|----|----|----|----|
| `documentId` | integer | Yes | Unique document identifier returned when payment was created |

</div>

<div class="sect4">

##### Response

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "state": "EXECUTED",
  "fee": 1.12,
  "documentId": 123,
  "comment": "Own funds",
  "amount": 100.00,
  "currency": "USD",
  "type": "PAYONEER",
  "accountFrom": "U0123504",
  "payload": "{\"account\":\"[email protected]\",\"type\":\"PAYONEER\"}",
  "userRequestId": "9876543219",
  "callbackUrl": "https://some-domain.com/for-callbacks"
}
```

</div>

</div>

</div>

</div>

<div class="sect3">

#### 4.4.3. Get Payment Status by User Request ID

<div class="paragraph">

**Endpoint:** `GET /v1/payment/{userRequestId}`

</div>

<div class="paragraph">

Retrieves the current status and details of a payment by your unique request identifier. The response also includes all payment fields as originally submitted, as well as the calculated fee associated with the transaction.

</div>

<div class="sect4">

##### Path Parameters

| Parameter | Type | Required | Description |
|----|----|----|----|
| `userRequestId` | string | Yes | Your unique transaction identifier specified when creating the payment |

</div>

<div class="sect4">

##### Response

<div class="paragraph">

Same response format as "Get Payment Status by Document ID".

</div>

</div>

<div class="sect4">

##### Response Fields

| Field           | Type    | Description                                   |
|-----------------|---------|-----------------------------------------------|
| `state`         | string  | Payment state (PENDING, EXECUTED, DECLINED)   |
| `fee`           | number  | Transaction fee amount                        |
| `documentId`    | integer | Unique document identifier                    |
| `comment`       | string  | Transaction comment (may be null)             |
| `amount`        | number  | Transaction amount                            |
| `currency`      | string  | Transaction currency code                     |
| `type`          | string  | Payment channel type (e.g., RUCARD, PAYONEER) |
| `accountFrom`   | string  | Source account identifier                     |
| `payload`       | string  | JSON string containing payment-specific data  |
| `userRequestId` | string  | Your unique request identifier                |
| `callbackUrl`   | string  | Callback URL for notifications                |

</div>

</div>

</div>

<div class="sect2">

### 4.5. Merchant Orders

<div class="sect3">

#### 4.5.1. Get Orders

<div class="paragraph">

**Endpoint:** `GET /v1/orders`

</div>

<div class="paragraph">

Retrieves a list of orders using filters.

</div>

<div class="sect4">

##### Request Parameters

| Parameter | Type | Required | Description |
|----|----|----|----|
| `limit` | number | No | Maximum number of returned orders. Default value is 100. |
| `offset` | number | No | Offset of returned orders. Default value is 0. |
| `periodStart` | string | No | Date in ISO 8601 format |
| `periodEnd` | string | No | Date in ISO 8601 format |
| `orderNumber` | string | No | Order number identifier |
| `merchantName` | string | No | Merchant name |
| `orderId` | number | No | Order id identifier |
| `merchantId` | number | No | Merchant id identifier |

</div>

<div class="sect4">

##### Response

<div class="paragraph">

Returns an array of merchant orders.

</div>

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "orders": [
    {
      "merchantId": 163,
      "number": "utorg0012",
      "creationDate": "2020-08-31T18:49:23.670Z",
      "description": "My demo payment.",
      "currency": "BTC",
      "state": "PAID",
      "amount": 0.00161283,
      "paidAmount": 0.00161283,
      "cryptoCurrency": "BTC",
      "orderId": 800000012
    }
  ],
  "count": 123
}
```

</div>

</div>

</div>

<div class="sect4">

##### Response Fields

| Field    | Type            | Description      |
|----------|-----------------|------------------|
| `orders` | Array\[Object\] | List of orders   |
| `count`  | number          | Number of orders |

</div>

<div class="sect4">

##### Order Object Fields

| Field | Type | Description |
|----|----|----|
| `merchantId` | number | Merchant id identifier |
| `number` | string | Order number identifier |
| `creationDate` | string | Date of creation in ISO 8601 format. |
| `currency` | string | Currency code |
| `cryptoCurrency` | string | Crypto currency code |
| `state` | number | Order state. Possible values: NEW, PARTIALLYPAID, PAID, FAIL, REFUND, CHARGEBACK, CANCELLED. |
| `amount` | number | Order amount |
| `paidAmount` | number | Order paid amount |
| `orderId` | number | Order id identifier |
| `description` | string | Order description |

</div>

</div>

</div>

<div class="sect2">

### 4.6. List of transactions

<div class="sect3">

#### 4.6.1. Get Transactions

<div class="paragraph">

**Endpoint:** `GET /v1/transactions`

</div>

<div class="paragraph">

Retrieves a list of transactions using filters.

</div>

<div class="sect4">

##### Request Parameters

| Parameter | Type | Required | Description |
|----|----|----|----|
| `limit` | number | No | Maximum number of transactions to return. Default value is 100. |
| `offset` | number | No | Number of transactions to skip. Default value is 0. |
| `periodStart` | string | No | Start date of the period in ISO 8601 format. |
| `periodEnd` | string | No | End date of the period in ISO 8601 format. |
| `transactionId` | number | No | Unique identifier of the transaction. |

</div>

<div class="sect4">

##### Response

<div class="paragraph">

Returns an array of transactions.

</div>

<div class="listingblock">

<div class="content">

``` highlightjs
{
    "transactions": [
        {
            "transactionId": 8367664,
            "createDate": "2014-02-18T21:47:29.452Z",
            "executeDate": "2014-02-18T21:47:39.543Z",
            "type": "EXTERNALOUT",
            "state": "EXECUTED",
            "amount": 500.00000000,
            "currency": "USD",
            "planDate": "2014-02-18T21:21:53.314Z",
            "version": 1
        }
    ],
    "count": 123
}
```

</div>

</div>

</div>

<div class="sect4">

##### Response Fields

| Field          | Type            | Description            |
|----------------|-----------------|------------------------|
| `transactions` | Array\[Object\] | List of transactions   |
| `count`        | number          | Number of transactions |

</div>

<div class="sect4">

##### Transaction Object Fields

| Field           | Type   | Description                           |
|-----------------|--------|---------------------------------------|
| `transactionId` | number | Unique identifier of the transaction. |
| `createDate`    | string | Date of creation in ISO 8601 format.  |
| `executeDate`   | string | Date of execution in ISO 8601 format. |
| `type`          | string | Transaction type                      |
| `state`         | string | Transaction state                     |
| `amount`        | number | Transaction amount                    |
| `currency`      | string | Currency code                         |
| `planDate`      | string | Date of plan in ISO 8601 format.      |
| `version`       | number | Transaction version                   |

</div>

</div>

</div>

<div class="sect2">

### 4.7. Cryptocurrency wallets

<div class="sect3">

#### 4.7.1. Get address for cryptocurrency deposit

<div class="paragraph">

**Endpoint:** `GET /v1/depositAddress/{currency}`

</div>

<div class="paragraph">

Returns address for cryptocurrency deposits. IMPORTANT: Do not attempt to save or cache addresses. They are not permanent and will change. Maximum guarantee: the returned address will be valid for 8 hours.

</div>

<div class="sect4">

##### Required Fields

| Field | Type | Description |
|----|----|----|
| `currency` | string | Must be "BTC", "ETH", "USDT", "USDC", "USDTt", "USDCb" or "USDTb" |

<div class="paragraph">

NOTE:

</div>

<div class="ulist">

- BTC - for Bitcoin

- ETH - for Ethereum

- USDT - USDT on ERC-20 network (Ethereum)

- USDTt - USDT on TRC-20 network (Tron)

- USDC - USDC on ERC-20 network (Ethereum)

- USDTb - USDT on BEP-20 network (BSC)

- USDCb - USDC on BEP-20 network (BSC)

</div>

</div>

<div class="sect4">

##### Response

<div class="paragraph">

If the request is successful, the API returns the wallet address.

</div>

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "address": "some your address",
}
```

</div>

</div>

</div>

<div class="sect4">

##### Response Fields

| Field     | Type   | Description                                   |
|-----------|--------|-----------------------------------------------|
| `address` | string | The address you can use to topup your account |

</div>

</div>

<div class="sect3">

#### 4.7.2. Get address for cryptocurrency deposit (in USDT TRC-20) with autoconversion

<div class="paragraph">

**Endpoint:** `GET /v1/depositAddressAutoUSDTt/{account}`

</div>

<div class="paragraph">

Returns address for cryptocurrency deposits in USDT TRC-20 which will be automatically converted to given account. IMPORTANT: Do not attempt to save or cache addresses. They are not permanent and will change. Maximum guarantee: the returned address will be valid for 8 hours.

</div>

<div class="sect4">

##### Required Fields

| Field     | Type   | Description                                          |
|-----------|--------|------------------------------------------------------|
| `account` | string | Account code to which your deposits will be credited |

</div>

<div class="sect4">

##### Response

<div class="paragraph">

If the request is successful, the API returns the wallet address.

</div>

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "address": "some your address",
}
```

</div>

</div>

</div>

<div class="sect4">

##### Response Fields

| Field     | Type   | Description                                   |
|-----------|--------|-----------------------------------------------|
| `address` | string | The address you can use to topup your account |

</div>

</div>

</div>

<div class="sect2">

### 4.8. KYC

<div class="sect3">

#### 4.8.1. Initiate new KYC

<div class="paragraph">

**Endpoint:** `POST /v1/kyc/start`

</div>

<div class="paragraph">

Initiates new KYC.

</div>

<div class="admonitionblock warning">

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><div class="title">
Warning
</div></td>
<td class="content">Starting a new KYC for the same kycExternalUserId makes current KYC obsolete.</td>
</tr>
</tbody>
</table>

</div>

<div class="sect4">

##### Request Body

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "kycExternalUserId": "string",
  "callbackUrl": "string",
  "sort": "string"
}
```

</div>

</div>

</div>

<div class="sect4">

##### Request Fields

| Field | Type | Required | Description |
|----|----|----|----|
| `kycExternalUserId` | string | Yes | Your unique final user id |
| `callbackUrl` | string | No | URL to make POST when KYC status changes (see next method for callback fields) |
| `sort` | string | Yes | Sort of KYC you are making KYC for |

</div>

<div class="sect4">

##### Response

<div class="paragraph">

If the request is successful, the API returns a JSON object containing the URL to proceed with KYC.

</div>

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "urlForUser": "https://somewhere.to/someuuid",
  "uuid": "a-b-c"
}
```

</div>

</div>

</div>

<div class="sect4">

##### Response Fields

| Field        | Type   | Description                                       |
|--------------|--------|---------------------------------------------------|
| `urlForUser` | string | URL to proceed with KYC                           |
| `uuid`       | string | UUID of KYC for to proceed with KYC automatically |

</div>

</div>

<div class="sect3">

#### 4.8.2. Query KYC status

<div class="paragraph">

**Endpoint:** `GET /v1/kyc/status/{kycExternalUserId}/{sort}`

</div>

<div class="paragraph">

Gives current status of KYC.

</div>

<div class="paragraph">

INFO: Body of callback for KYC status change is the same as the result of this method call.

</div>

<div class="sect4">

##### Path Parameters

| Field               | Type   | Required | Description                     |
|---------------------|--------|----------|---------------------------------|
| `kycExternalUserId` | string | Yes      | Your unique final user id       |
| `sort`              | string | Yes      | KYC sort you are making KYC for |

</div>

<div class="sect4">

##### Response

<div class="paragraph">

If the request is successful, the API returns a JSON object containing the URL to proceed with KYC.

</div>

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "kycExternalUserId": "some your user unique id",
  "status": "FAILED",
  "reason": "photo is dirty",
  "uuid": "a-b-c"
}
```

</div>

</div>

</div>

<div class="sect4">

##### Response Fields

| Field | Type | Description |
|----|----|----|
| `kycExternalUserId` | string | Your unique final user id |
| `status` | string | Status of KYC. Possible values: INITIATED (just created), OPENED (in process of filling), COMPLETED (user has filled, not checked), APPROVED (internally approved), DECLINED, EXPIRED, ??? |
| `reason` | string (opt) | Explanation of Status of KYC. |
| `uuid` | string | UUID of KYC for to proceed with KYC automatically |

</div>

</div>

<div class="sect3">

#### 4.8.3. Query KYC status by UUID

<div class="paragraph">

**Endpoint:** `GET /v1/kyc/statusByUuid/{uuid}`

</div>

<div class="paragraph">

Gives current status of KYC.

</div>

<div class="paragraph">

INFO: Body of callback for KYC status change is the same as the result of this method call.

</div>

<div class="sect4">

##### Path Parameters

| Field  | Type   | Required | Description |
|--------|--------|----------|-------------|
| `uuid` | string | Yes      | UUID of KYC |

</div>

<div class="sect4">

##### Response

<div class="paragraph">

If the request is successful, the API returns a JSON object containing the URL to proceed with KYC.

</div>

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "kycExternalUserId": "some your user unique id",
  "status": "FAILED",
  "reason": "photo is dirty",
  "uuid": "a-b-c"
}
```

</div>

</div>

</div>

<div class="sect4">

##### Response Fields

| Field | Type | Description |
|----|----|----|
| `kycExternalUserId` | string | Your unique final user id |
| `status` | string | Status of KYC. Possible values: INITIATED (just created), OPENED (in process of filling), COMPLETED (user has filled, not checked), APPROVED (internally approved), DECLINED, EXPIRED, ??? |
| `reason` | string (opt) | Explanation of Status of KYC. |
| `uuid` | string | UUID of KYC for to proceed with KYC automatically |

</div>

</div>

<div class="sect3">

#### 4.8.4. Fill KYC data

<div class="paragraph">

**Endpoint:** `POST /v1/kyc/setData/{uuid}`

</div>

<div class="paragraph">

Fills basic data for KYC.

</div>

<div class="paragraph">

INFO: Body is JSON which format depends on sort of KYC.

</div>

<div class="sect4">

##### Path Parameters

| Field  | Type   | Required | Description |
|--------|--------|----------|-------------|
| `uuid` | string | Yes      | UUID of KYC |

</div>

<div class="sect4">

##### Response

<div class="paragraph">

If the request is successful, the API returns http code 200.

</div>

</div>

</div>

<div class="sect3">

#### 4.8.5. Add KYC picture

<div class="paragraph">

**Endpoint:** `POST /v1/kyc/setPicture/{uuid}/{type}`

</div>

<div class="paragraph">

Attaches a picture to KYC.

</div>

<div class="paragraph">

Body is picture (jpeg) binary.

</div>

<div class="sect4">

##### Path Parameters

| Field | Type | Required | Description |
|----|----|----|----|
| `uuid` | string | Yes | UUID of KYC |
| `type` | string | Yes | Type of picture (selfie / passport / etc). Needed types depends on product Id you are making KYC for. |

</div>

<div class="sect4">

##### Response

<div class="paragraph">

If the request is successful, the API returns http code 200.

</div>

</div>

</div>

<div class="sect3">

#### 4.8.6. Send KYC for verification

<div class="paragraph">

**Endpoint:** `POST /v1/kyc/confirm/{uuid}`

</div>

<div class="paragraph">

Sends filled KYC for verification.

</div>

<div class="sect4">

##### Path Parameters

| Field  | Type   | Required | Description |
|--------|--------|----------|-------------|
| `uuid` | string | Yes      | UUID of KYC |

</div>

<div class="sect4">

##### Response

<div class="paragraph">

If the request is successful, the API returns http code 200.

</div>

</div>

</div>

</div>

</div>

</div>

<div class="sect1">

## 5. Payment Callbacks

<div class="sectionbody">

<div class="paragraph">

The service will inform you about changes in the payment status using the `callbackUrl` link. A POST request will be sent to this URL when a final status is received.

</div>

<div class="sect2">

### 5.1. Callback Headers

| Header | Description |
|----|----|
| `X-Request-Timestamp` | Current timestamp in epoch milliseconds |
| `Signature` | SHA-256 hex hash of: `X-Request-Timestamp + request body + API-secret` |

<div class="admonitionblock note">

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><div class="title">
Note
</div></td>
<td class="content">The callback signature is calculated the same way as the request signature. Use this to verify the authenticity of the callback. When validating the callback, use the timestamp from the <strong>X-Request-Timestamp</strong> header together with your <strong>API-secret</strong> to recompute the signature and compare it with the value in the <strong>Signature</strong> header.</td>
</tr>
</tbody>
</table>

</div>

</div>

<div class="sect2">

### 5.2. Callback Body Fields

| Field | Type | Required | Description |
|----|----|----|----|
| `state` | string | Yes | Document final state (EXECUTED, DECLINED) |
| `fee` | number | Yes | Transaction fee amount |
| `documentId` | integer | Yes | Unique identifier of the transaction document |
| `comment` | string | No | Additional information or notes about the transaction |
| `amount` | number | Yes | Transaction amount |
| `currency` | string | Yes | Transaction currency code (e.g., USD, EUR) |
| `type` | string | Yes | Transaction type (e.g., RUCARD, PAYONEER) |
| `accountFrom` | string | Yes | Source account identifier |
| `userRequestId` | string | Yes | Unique request identifier from the user |
| `callbackUrl` | string | Yes | URL for sending callback notifications |

</div>

<div class="sect2">

### 5.3. Example Callback Body

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "state": "EXECUTED",
  "fee": 1.12,
  "documentId": 123,
  "comment": "Own funds",
  "amount": 100.00,
  "currency": "USD",
  "type": "PAYONEER",
  "accountFrom": "U0123504",
  "userRequestId": "9876543219",
  "callbackUrl": "https://some-domain.com/for-callbacks"
}
```

</div>

</div>

</div>

</div>

</div>

<div class="sect1">

## 6. Payment Channels

<div class="sectionbody">

<div class="paragraph">

The following payment channels are supported. Each channel has specific validation requirements for the `payload` field.

</div>

<div class="paragraph">

Each supported record type in the system has its own validation rules. These rules define which fields are required in the request payload when creating a new payment.

</div>

<div class="admonitionblock warning">

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><div class="title">
Warning
</div></td>
<td class="content">Not all payment channels listed in the documentation may be currently available.</td>
</tr>
</tbody>
</table>

</div>

<div class="sect2">

### 6.1. Card Payments

<div class="sect3">

#### 6.1.1. RUCARD - Russian Cards

<div class="paragraph">

Russian payment cards.

</div>

<div class="sect4">

##### Required Fields

| Field     | Type   | Description                                     |
|-----------|--------|-------------------------------------------------|
| `type`    | string | Must be "RUCARD"                                |
| `account` | string | Russian card number (16-19 digits, digits only) |
| `name`    | string | Cardholder first name                           |
| `surname` | string | Cardholder last name                            |
| `midname` | string | Cardholder middle name (patronymic)             |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "RUCARD",
  "name": "Ivan",
  "account": "1234567890123456",
  "surname": "Petrov",
  "midname": "Sergeevich"
}
```

</div>

</div>

</div>

</div>

<div class="sect3">

#### 6.1.2. RUCARDP2P / RUCARDP2PDYN - Russian P2P Card Transfers

<div class="paragraph">

Peer-to-peer transfers between Russian cards.

</div>

<div class="sect4">

##### Required Fields

| Field     | Type   | Description                                     |
|-----------|--------|-------------------------------------------------|
| `type`    | string | Must be "RUCARDP2P" or "RUCARDP2PDYN"           |
| `account` | string | Russian card number (16-19 digits, digits only) |
| `name`    | string | Cardholder first name                           |
| `surname` | string | Cardholder last name                            |
| `midname` | string | Cardholder middle name (patronymic)             |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "RUCARDP2P",
  "account": "1234567890123456",
  "name": "Ivan",
  "surname": "Petrov",
  "midname": "Sergeevich"
}
```

</div>

</div>

</div>

</div>

<div class="sect3">

#### 6.1.3. TRCARD - Turkish Cards

<div class="paragraph">

Payment cards issued by Turkish banks.

</div>

<div class="sect4">

##### Required Fields

| Field | Type | Description |
|----|----|----|
| `type` | string | Must be "TRCARD" |
| `account` | string | Card number (16 digits starting with 4, 5, or 9) |
| `senderFullname` | string | Sender’s full name as on card |
| `senderIdentityNumber` | string | Sender’s tax code number (11 digits) |
| `beneficiaryFullname` | string | Recipient’s full name |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "TRCARD",
  "account": "423411110123111",
  "senderFullname": "Ahmet Ymaz",
  "senderIdentityNumber": "12345678901",
  "beneficiaryFullname": "Mehmet Damir"
}
```

</div>

</div>

</div>

</div>

<div class="sect3">

#### 6.1.4. GECARD - Georgian Cards

<div class="paragraph">

Cards issued by Georgian banks.

</div>

<div class="sect4">

##### Required Fields

| Field                  | Type   | Description            |
|------------------------|--------|------------------------|
| `type`                 | string | Must be "GECARD"       |
| `account`              | string | Card number            |
| `senderFirstname`      | string | Sender’s first name    |
| `senderLastname`       | string | Sender’s last name     |
| `senderIdentityNumber` | string | Sender’s ID number     |
| `receiverFirstname`    | string | Recipient’s first name |
| `receiverLastname`     | string | Recipient’s last name  |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "GECARD",
  "account": "5234567890123456",
  "senderFirstname": "Giorgi",
  "senderLastname": "Piponia",
  "senderIdentityNumber": "01234567890",
  "receiverFirstname": "Nino",
  "receiverLastname": "Kupdze"
}
```

</div>

</div>

</div>

</div>

<div class="sect3">

#### 6.1.5. AZCARD - Azerbaijani Cards

<div class="paragraph">

Cards issued by Azerbaijani banks.

</div>

<div class="sect4">

##### Required Fields

| Field             | Type   | Description                                  |
|-------------------|--------|----------------------------------------------|
| `type`            | string | Must be "AZCARD"                             |
| `account`         | string | Card number (exact length depends on issuer) |
| `cardExpiryYear`  | string | 4-digit future year (e.g., 2033)             |
| `cardExpiryMonth` | string | 2-digit month (01-12)                        |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "AZCARD",
  "account": "1234567890123456",
  "cardExpiryYear": "2033",
  "cardExpiryMonth": "11"
}
```

</div>

</div>

</div>

</div>

<div class="sect3">

#### 6.1.6. KZCARD - Kazakhstani Cards

<div class="paragraph">

Payment cards issued by Kazakhstani banks.

</div>

<div class="sect4">

##### Required Fields

| Field     | Type   | Description               |
|-----------|--------|---------------------------|
| `type`    | string | Must be "KZCARD"          |
| `account` | string | Card number (digits only) |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "KZCARD",
  "account": "1234567890123456"
}
```

</div>

</div>

</div>

</div>

<div class="sect3">

#### 6.1.7. UZCARD - Uzbekistani Cards

<div class="paragraph">

Payment cards issued by Uzbekistani banks.

</div>

<div class="sect4">

##### Required Fields

| Field     | Type   | Description               |
|-----------|--------|---------------------------|
| `type`    | string | Must be "UZCARD"          |
| `account` | string | Card number (digits only) |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "UZCARD",
  "account": "1234567890123456"
}
```

</div>

</div>

</div>

</div>

<div class="sect3">

#### 6.1.8. WORLDCARDEUR / WORLDCARDUSD - International Cards

<div class="paragraph">

International payment cards in EUR or USD.

</div>

<div class="sect4">

##### Required Fields

| Field                | Type   | Description                              |
|----------------------|--------|------------------------------------------|
| `type`               | string | Must be "WORLDCARDEUR" or "WORLDCARDUSD" |
| `account`            | string | Card number (16-18 digits, no spaces)    |
| `name`               | string | Cardholder first name                    |
| `surname`            | string | Cardholder last name                     |
| `cardCurrency`       | string | Must match service (EUR/USD)             |
| `birthDate`          | string | Date in YYYY-MM-DD format                |
| `address`            | string | Full street address                      |
| `addressCountryCode` | string | Country code (ISO)                       |
| `addressCity`        | string | City name                                |
| `cardExpiryMonth`    | string | 2 digits (01-12)                         |
| `cardExpiryYear`     | string | 4 digits (e.g., 2025)                    |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "WORLDCARDEUR",
  "account": "1234567890123456",
  "name": "John",
  "surname": "Smith",
  "cardCurrency": "EUR",
  "birthDate": "1990-01-01",
  "address": "123 Main Street",
  "addressCountryCode": "USA",
  "addressCity": "New York",
  "cardExpiryMonth": "12",
  "cardExpiryYear": "2025"
}
```

</div>

</div>

</div>

</div>

</div>

<div class="sect2">

### 6.2. Bank Transfers

<div class="sect3">

#### 6.2.1. ARBANK - Argentinian Banks

<div class="paragraph">

Bank accounts in Argentina.

</div>

<div class="sect4">

##### Required Fields

| Field | Type | Description |
|----|----|----|
| `type` | string | Must be "ARBANK" |
| `account` | string | Beneficiary account ID (alphanumeric, exact length depends on bank) |
| `taxId` | string | Exactly 11 digits of beneficiary’s tax ID |
| `cbu` | string | 22-digit numeric code for routing transfers |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "ARBANK",
  "account": "12345",
  "taxId": "12345678901",
  "cbu": "1234567890123456789012"
}
```

</div>

</div>

</div>

</div>

<div class="sect3">

#### 6.2.2. BRBANK - Brazilian Banks

<div class="paragraph">

Bank accounts in Brazil.

</div>

<div class="sect4">

##### Required Fields

| Field | Type | Description |
|----|----|----|
| `type` | string | Must be "BRBANK" |
| `taxId` | string | 11 digits of taxpayer identification number (CPF format) |
| `name` | string | Full legal name |
| `email` | string | Valid email format |
| `additionalRemark` | string | Minimum 6 characters |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "BRBANK",
  "taxId": "12345678901",
  "name": "João Silva",
  "email": "[email protected]",
  "additionalRemark": "Own funds"
}
```

</div>

</div>

</div>

</div>

<div class="sect3">

#### 6.2.3. COBANK - Colombian Banks

<div class="paragraph">

Bank accounts in Colombia.

</div>

<div class="sect4">

##### Required Fields

| Field      | Type   | Description                                      |
|------------|--------|--------------------------------------------------|
| `type`     | string | Must be "COBANK"                                 |
| `taxId`    | string | 9 to 11 digits of taxpayer identification number |
| `fullName` | string | Full legal name                                  |
| `email`    | string | Valid email format                               |
| `phone`    | string | 11 to 15 digits                                  |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "COBANK",
  "taxId": "1234567890",
  "fullName": "Carlos Rodriguez",
  "email": "[email protected]",
  "phone": "573001234567"
}
```

</div>

</div>

</div>

</div>

<div class="sect3">

#### 6.2.4. MALAYSIA_BANK - Malaysian Banks

<div class="sect4">

##### Required Fields

| Field | Type | Description |
|----|----|----|
| `type` | string | Must be "MALAYSIA_BANK" |
| `account` | string | Bank account number (digits only) |
| `destination` | object | Bank details |
| `destination.bankCode` | string | Bank identifier code |
| `destination.accountName` | string | Name on the account |
| `destination.accountNumber` | string | Recipient account number |
| `recipient` | object | Recipient information |
| `recipient.firstName` | string | Recipient’s first name |
| `recipient.lastName` | string | Recipient’s last name |
| `recipient.email` | string | Recipient’s email |
| `recipient.phone` | string | Recipient’s phone number |
| `recipient.zipCode` | string | Postal code |
| `recipient.city` | string | Recipient’s city |
| `recipient.address` | string | Recipient’s street address |
| `recipient.country` | string | Lowercase ISO 2-letter country code (e.g., 'my') |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "MALAYSIA_BANK",
  "destination": {
    "bankCode": "MBBEMYKL",
    "accountName": "Ahmad Abdullah",
    "accountNumber": "1234567890"
  },
  "recipient": {
    "firstName": "Ahmad",
    "lastName": "Abdullah",
    "email": "[email protected]",
    "phone": "60123456789",
    "zipCode": "50000",
    "city": "Kuala Lumpur",
    "address": "Jalan Sultan Ismail",
    "country": "my"
  },
  "account": "1234567890"
}
```

</div>

</div>

</div>

</div>

<div class="sect3">

#### 6.2.5. INDONESIA_BANK - Indonesian Banks

<div class="sect4">

##### Required Fields

<div class="paragraph">

Same structure as MALAYSIA_BANK with `type` = "INDONESIA_BANK" and `recipient.country` = "id".

</div>

</div>

</div>

<div class="sect3">

#### 6.2.6. THAILAND_BANK - Thai Banks

<div class="sect4">

##### Required Fields

<div class="paragraph">

Same structure as MALAYSIA_BANK with `type` = "THAILAND_BANK" and `recipient.country` = "th".

</div>

</div>

</div>

<div class="sect3">

#### 6.2.7. SOUTH_KOREA_BANK - South Korean Banks

<div class="sect4">

##### Required Fields

<div class="paragraph">

Same structure as MALAYSIA_BANK with `type` = "SOUTH_KOREA_BANK" and `recipient.country` = "kr".

</div>

</div>

</div>

</div>

<div class="sect2">

### 6.3. Mobile Payments

<div class="sect3">

#### 6.3.1. MEGAFON / TMOBILE / BEELINE / MTS / TELE2 / YOTA

<div class="paragraph">

Digital wallets and Russian mobile carrier payments.

</div>

<div class="sect4">

##### Required Fields

| Field     | Type   | Description                                           |
|-----------|--------|-------------------------------------------------------|
| `type`    | string | Payment type (e.g., "MEGAFON", "TELE2", …​)            |
| `account` | string | Russian phone number (11 digits starting with 7 or 8) |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "MEGAFON",
  "account": "79001234567"
}
```

</div>

</div>

</div>

</div>

<div class="sect3">

#### 6.3.2. UKR_MOBILE - Ukrainian Mobile Payments

<div class="paragraph">

Payments via Ukrainian mobile operators.

</div>

<div class="sect4">

##### Required Fields

| Field     | Type   | Description                          |
|-----------|--------|--------------------------------------|
| `type`    | string | Must be "UKR_MOBILE"                 |
| `account` | string | Ukrainian phone number (digits only) |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "UKR_MOBILE",
  "account": "380501234567"
}
```

</div>

</div>

</div>

</div>

</div>

<div class="sect2">

### 6.4. Fast Payment Systems

<div class="sect3">

#### 6.4.1. SBP - Russian Fast Payment System / "Система Быстрых Платежей"

<div class="paragraph">

System for instant payments in Russia.

</div>

<div class="sect4">

##### Required Fields

| Field       | Type   | Description                        |
|-------------|--------|------------------------------------|
| `type`      | string | Must be "SBP"                      |
| `account`   | string | Russian phone number (digits only) |
| `email`     | string | Beneficiary’s email                |
| `nspk`      | string | NSPK member bank ID                |
| `firstName` | string | Beneficiary’s first name           |
| `lastName`  | string | Beneficiary’s last name            |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "SBP",
  "account": "79001234567",
  "email": "[email protected]",
  "nspk": "100000000004",
  "firstName": "Ivan",
  "lastName": "Petrov"
}
```

</div>

</div>

</div>

</div>

<div class="sect3">

#### 6.4.2. IMPS - Immediate Payment Service (India)

<div class="paragraph">

India’s real-time interbank electronic funds transfer system.

</div>

<div class="sect4">

##### Required Fields

| Field          | Type   | Description                                      |
|----------------|--------|--------------------------------------------------|
| `type`         | string | Must be "IMPS"                                   |
| `account`      | string | Account number (digits only)                     |
| `account_name` | string | Beneficiary account name                         |
| `bank_code`    | string | Indian Financial System Code (IFSC, digits only) |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "IMPS",
  "account": "1234567890",
  "account_name": "Raj Kumar",
  "bank_code": "SBIN0001234"
}
```

</div>

</div>

</div>

</div>

</div>

<div class="sect2">

### 6.5. E-Wallets

<div class="sect3">

#### 6.5.1. PAYONEER / PAYONEER_EUR / PAYONEER_USD

<div class="paragraph">

International digital payment platform for cross-border transfers.

</div>

<div class="sect4">

##### Required Fields

| Field     | Type   | Description                                           |
|-----------|--------|-------------------------------------------------------|
| `type`    | string | Must be "PAYONEER", "PAYONEER_EUR", or "PAYONEER_USD" |
| `account` | string | Beneficiary email address                             |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "PAYONEER_EUR",
  "account": "[email protected]"
}
```

</div>

</div>

</div>

</div>

<div class="sect3">

#### 6.5.2. EUR_NETELLER / EUR_SKRILL

<div class="paragraph">

Electronic wallets in euros via Neteller or Skrill.

</div>

<div class="sect4">

##### Required Fields

| Field | Type | Description |
|----|----|----|
| `type` | string | Must be "EUR_NETELLER" or "EUR_SKRILL" |
| `wallet_email` | string | Beneficiary email address |
| `name` | string | Legal name |
| `surname` | string | Legal surname |
| `country` | string | ISO country code |
| `city` | string | City of residence |
| `address` | string | Physical address |
| `zip` | string | Postal code |
| `code` | string | Verification code |
| `phone` | string | Phone number (digits only) |
| `client_ip` | string | Client’s IP address |
| `state` | string | Required if country is USA or CANADA (ISO ALPHA-2 state code) |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "EUR_NETELLER",
  "wallet_email": "[email protected]",
  "name": "John",
  "surname": "Doe",
  "country": "USA",
  "city": "New York",
  "address": "123 Main St",
  "zip": "10001",
  "code": "123456",
  "phone": "12125551234",
  "client_ip": "192.168.1.1",
  "state": "NY"
}
```

</div>

</div>

</div>

</div>

<div class="sect3">

#### 6.5.3. PAYTM

<div class="paragraph">

Indian digital payments and mobile wallet platform.

</div>

<div class="sect4">

##### Required Fields

| Field     | Type   | Description                  |
|-----------|--------|------------------------------|
| `type`    | string | Must be "PAYTM"              |
| `account` | string | Account number (digits only) |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "PAYTM",
  "account": "9876543210"
}
```

</div>

</div>

</div>

</div>

<div class="sect3">

#### 6.5.4. GCASH

<div class="paragraph">

Digital wallet and mobile payment service (Philippines).

</div>

<div class="sect4">

##### Required Fields

| Field | Type | Description |
|----|----|----|
| `type` | string | Must be "GCASH" |
| `receiver` | object | Recipient details |
| `receiver.surName` | string | Recipient’s surname |
| `receiver.givName` | string | Recipient’s given name |
| `receiver.phone` | string | Recipient’s mobile number (digits only) |
| `sender` | object | Sender details |
| `sender.purposeOfRemittance` | string | Purpose of transaction |
| `account` | string | GCash account number |
| `receiveAmount` | number | Amount to receive |
| `receiveCurrency` | string | Currency code |
| `channelCode` | string | Payment channel code |
| `channelName` | string | Payment channel name |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "GCASH",
  "receiver": {
    "surName": "Dela Cruz",
    "givName": "Juan",
    "phone": "639171234567"
  },
  "sender": {
    "purposeOfRemittance": "FAMILY SUPPORT"
  },
  "account": "639171234567",
  "receiveAmount": 5000.00,
  "receiveCurrency": "PHP",
  "channelCode": "73",
  "channelName": "GCASH"
}
```

</div>

</div>

</div>

</div>

<div class="sect3">

#### 6.5.5. WM

<div class="paragraph">

WebMoney payments.

</div>

<div class="sect4">

##### Required Fields

| Field     | Type   | Description                              |
|-----------|--------|------------------------------------------|
| `type`    | string | Must be "WM"                             |
| `account` | string | WebMoney wallet identifier, alphanumeric |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "WM",
  "account": "9876543210"
}
```

</div>

</div>

</div>

</div>

<div class="sect3">

#### 6.5.6. PAGS_CHWALLET

<div class="paragraph">

Vita wallet Chile payments.

</div>

<div class="sect4">

##### Required Fields

| Field     | Type   | Description             |
|-----------|--------|-------------------------|
| `type`    | string | Must be "PAGS_CHWALLET" |
| `account` | string | Beneficiary’s email     |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "PAGS_CHWALLET",
  "account": "[email protected]"
}
```

</div>

</div>

</div>

</div>

<div class="sect3">

#### 6.5.7. PAGS_MXSPEI

<div class="paragraph">

Mexican SPEI

</div>

<div class="sect4">

##### Required Fields

| Field          | Type   | Description                 |
|----------------|--------|-----------------------------|
| `type`         | string | Must be "PAGS_MXSPEI"       |
| `name`         | string | Beneficiary’s full name     |
| `bankCode`     | string | Bank code                   |
| `account`      | string | Account number              |
| `accountType`  | string | One of: clabe, phone, debit |
| `documentId`   | string | Document ID                 |
| `documentType` | string | One of: RFC, CURP           |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "PAGS_MXSPEI",
  "name": "John Doe",
  "bankCode": "123",
  "account": "9876543210",
  "accountType": "clabe",
  "documentId": "ABC123456",
  "documentType": "RFC"
}
```

</div>

</div>

</div>

</div>

<div class="sect3">

#### 6.5.8. PAGS_COWALLET

<div class="paragraph">

Colombian Nequi, Tpaga

</div>

<div class="sect4">

##### Required Fields

| Field     | Type   | Description             |
|-----------|--------|-------------------------|
| `type`    | string | Must be "PAGS_COWALLET" |
| `account` | string | Account number          |
| `channel` | string | One of: Nequi, Tpaga    |
| `email`   | string | Email of beneficiary    |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "PAGS_COWALLET",
  "account": "3001234567",
  "channel": "Nequi",
  "email": "[email protected]"
}
```

</div>

</div>

</div>

</div>

<div class="sect3">

#### 6.5.9. PAGS_COTRAN

<div class="paragraph">

Colombian Transfiya

</div>

<div class="sect4">

##### Required Fields

| Field | Type | Description |
|----|----|----|
| `type` | string | Must be "PAGS_COTRAN" |
| `account` | string | Account number. Must start with 57 and contain 12–14 digits |
| `name` | string | Beneficiary full name |
| `documentNumber` | string | CC: 6–10 digits. CE: max 12 chars. PEP: max 15 digits |
| `documentType` | string | One of: CC, CE, PEP |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "PAGS_COTRAN",
  "account": "5712345678901",
  "name": "Juan Perez",
  "documentNumber": "123456789",
  "documentType": "CC"
}
```

</div>

</div>

</div>

</div>

<div class="sect3">

#### 6.5.10. EPAY_EPAY_E_VN_ZALO

<div class="paragraph">

ZaloPay payments.

</div>

<div class="sect4">

##### Required Fields

| Field | Type | Description |
|----|----|----|
| `type` | string | Must be "EPAY_EPAY_E_VN_ZALO" |
| `receiver` | object | Receiver info (surName, givName, phone) |
| `sender` | object | Sender info (surName, country, address, purposeOfRemittance) |
| `account` | string | Receiver account |
| `receiveAmount` | string | Amount to be received |
| `receiveCurrency` | string | Currency code (e.g., USD) |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "EPAY_EPAY_E_VN_ZALO",
  "receiver": {
    "surName": "Tran",
    "givName": "Anh",
    "phone": "0901234567"
  },
  "sender": {
    "surName": "Smith",
    "country": "US",
    "address": "123 Main St",
    "purposeOfRemittance": "Family support"
  },
  "account": "123456789",
  "receiveAmount": "100000",
  "receiveCurrency": "USD"
}
```

</div>

</div>

</div>

</div>

<div class="sect3">

#### 6.5.11. EPAY_EPAY_E_BD_BKASH

<div class="paragraph">

bKash payments.

</div>

<div class="sect4">

##### Required Fields

| Field | Type | Description |
|----|----|----|
| `type` | string | Must be "EPAY_EPAY_E_BD_BKASH" |
| `receiver` | object | Receiver info (surName, givName, accountNo) |
| `sender` | object | Sender info (givName, surName, area, phone, nationality, idType, idNumber, birthday, city, zipCode, country, address, purposeOfRemittance) |
| `account` | string | Beneficiary account |
| `receiveAmount` | string | Amount to be received |
| `receiveCurrency` | string | Currency code (e.g., USD) |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "EPAY_EPAY_E_BD_BKASH",
  "receiver": {
    "surName": "Rahman",
    "givName": "Ali",
    "accountNo": "017XXXXXXXX"
  },
  "sender": {
    "givName": "John",
    "surName": "Doe",
    "area": "Dhaka",
    "phone": "8801XXXXXXXX",
    "nationality": "US",
    "idType": "Passport",
    "idNumber": "AB1234567",
    "birthday": "1985-01-01",
    "city": "NYC",
    "zipCode": "10001",
    "country": "USA",
    "address": "123 Main St",
    "purposeOfRemittance": "Family Support"
  },
  "account": "017XXXXXXXX",
  "receiveAmount": "5000",
  "receiveCurrency": "USD"
}
```

</div>

</div>

</div>

</div>

</div>

<div class="sect2">

### 6.6. Cryptocurrency

<div class="sect3">

#### 6.6.1. BITCOIN / ETH / USDCERC20 / USDTERC20 / USDTTRC20 / USDCBSC / USDTBSC

<div class="paragraph">

Cryptocurrency payments for Bitcoin, Ethereum, and various stablecoins.

</div>

<div class="sect4">

##### Required Fields

| Field | Type | Description |
|----|----|----|
| `type` | string | Must be "BITCOIN", "ETH", "USDCERC20", "USDTERC20", "USDTTRC20", "USDCBSC" or "USDTBSC" |
| `account` | string | Wallet address (minimum 26 characters) |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "BITCOIN",
  "account": "bc1qxy2kg342rsqtzq2n0yrf2493p83kkfjhx0wlh"
}
```

</div>

</div>

<div class="paragraph">

NOTE:

</div>

<div class="ulist">

- USDTERC20 - USDT on ERC-20 network (Ethereum)

- USDTTRC20 - USDT on TRC-20 network (Tron)

- USDCERC20 - USDC on ERC-20 network (Ethereum)

- USDTBSC - USDT on BEP-20 network (BSC)

- USDCBSC - USDC on BEP-20 network (BSC)

</div>

</div>

</div>

</div>

<div class="sect2">

### 6.7. STEAM

<div class="sect3">

#### 6.7.1. STEAM RUR ACCOUNTS

<div class="paragraph">

Top up Steam account balance in rubles

</div>

<div class="sect4">

##### Required Fields

| Field     | Type   | Description                            |
|-----------|--------|----------------------------------------|
| `account` | string | Steam account name, not Steam nickname |
| `ip`      | string | Ip address                             |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "steamAccountName",
  "ip": "179.195.60.108"
}
```

</div>

</div>

</div>

</div>

</div>

<div class="sect2">

### 6.8. CAPITALIST

<div class="sect3">

#### 6.8.1. CAPITALIST INTERNAL PAYMENTS

<div class="paragraph">

Send payment to another system user.

</div>

<div class="sect4">

##### Required Fields

| Field     | Type   | Description            |
|-----------|--------|------------------------|
| `type`    | string | Must be "CAPITALIST"   |
| `account` | string | Beneficiary account ID |

</div>

<div class="sect4">

##### Example Payload

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "type": "CAPITALIST",
  "account": "U123456789"
}
```

</div>

</div>

</div>

</div>

</div>

</div>

</div>

<div class="sect1">

## 7. Error Handling

<div class="sectionbody">

<div class="paragraph">

When an error occurs, the API returns a JSON object with an error description.

</div>

<div class="sect2">

### 7.1. Error Response

<div class="listingblock">

<div class="content">

``` highlightjs
{
  "error": "Error description message"
}
```

</div>

</div>

</div>

<div class="sect2">

### 7.2. HTTP Status Codes

| Status Code | Description                                          |
|-------------|------------------------------------------------------|
| 200         | Success - Request processed successfully             |
| 400         | Bad Request - Invalid parameters or validation error |

</div>

</div>

</div>

<div class="sect1">

## 8. Reference

<div class="sectionbody">

<div class="sect2">

### 8.1. Currency

| Currency Code | Description     |
|---------------|-----------------|
| RUR           | Russian Roubles |
| USD           | U.S. Dollars    |
| EUR           | European EURO   |
| USDT          | USDT (Erc 20)   |
| USDTt         | USDT (Trc 20)   |
| USDTb         | USDT (Bep 20)   |
| ETH           | Ethereum        |
| USDC          | USDC (Erc 20)   |
| USDCb         | USDC (Bep 20)   |
| BTC           | Bitcoin         |

</div>

</div>

</div>

<div class="sect1">

## 9. Best Practices

<div class="sectionbody">

<div class="sect2">

### 9.1. Security

<div class="olist arabic">

1.  **Keep your API secret secure** - Never commit it to version control or share it publicly

2.  **Validate callback signatures** - Always verify the signature of incoming callbacks to ensure they’re from the legitimate service

3.  **Use HTTPS** - All API calls must use HTTPS protocol

4.  **Implement idempotency** - Use unique `userRequestId` for each transaction to prevent duplicates

</div>

</div>

<div class="sect2">

### 9.2. Integration Tips

<div class="olist arabic">

1.  **Test with small amounts first** - Start with minimal transaction amounts during integration

2.  **Implement proper error handling** - Handle all possible error responses gracefully

3.  **Monitor callback endpoints** - Ensure your callback URL is always accessible and responds within timeout limits

4.  **Log all transactions** - Keep detailed logs of all API requests and responses for debugging

5.  **Check payment channel availability** - Not all payment channels may be available at all times

</div>

</div>

<div class="sect2">

### 9.3. Rate Limiting

<div class="paragraph">

If you want to poll transaction statuses more than 20 times per minute, consider integrating callbacks instead. You can create new transactions without any restrictions.

</div>

<div class="paragraph">

With excessively frequent API calls, the system will automatically throttle your requests.

</div>

</div>

</div>

</div>

<div class="sect1">

## 10. Coding Examples

<div class="sectionbody">

<div class="paragraph">

You can find working integration examples in our public GitHub repository:

</div>

<div class="ulist">

- Java example: <a href="https://github.com/capitalist-net/API-V2" class="bare">https://github.com/capitalist-net/API-V2</a>

</div>

</div>

</div>

<div class="sect1">

## 11. Support

<div class="sectionbody">

<div class="paragraph">

For technical support and questions about the API integration, please contact support through your personal account.

</div>

</div>

</div>

</div>

<div id="footer">

<div id="footer-text">

Last updated 2026-06-02 12:39:13 UTC

</div>

</div>
