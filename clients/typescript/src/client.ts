import { calculateSignature } from './signature.js';
import { CapitalistApi2Error } from './errors.js';
import type {
  Account,
  ClientOptions,
  Currency,
  DepositAddressResponse,
  ExchangeRequest,
  KycStartRequest,
  KycStartResponse,
  KycStatus,
  OrdersFilters,
  PaymentCreateResponse,
  PaymentRequest,
  PaymentStatus,
  PrepaidDenomination,
  PrepaidProduct,
  PrepaidService,
  TransactionsFilters,
  TransactionsResponse,
} from './types.js';

const DEFAULT_BASE_URL = 'https://api2.capitalist.net';

export class CapitalistApi2Client {
  private readonly apiKey: string;
  private readonly apiSecret: string;
  private readonly baseUrl: string;
  private readonly fetchImpl: typeof fetch;

  constructor(options: ClientOptions) {
    this.apiKey = options.apiKey;
    this.apiSecret = options.apiSecret;
    this.baseUrl = (options.baseUrl ?? DEFAULT_BASE_URL).replace(/\/+$/, '');
    this.fetchImpl = options.fetchImpl ?? fetch;
  }

  readWhitelist(): Promise<string[]> {
    return this.request('GET', '/v1/whitelist');
  }

  addWhitelist(ip: string[]): Promise<unknown> {
    return this.request('POST', '/v1/whitelist', { ip });
  }

  removeWhitelist(ip: string[]): Promise<unknown> {
    return this.request('POST', '/v1/whitelist/remove', { ip });
  }

  listAccounts(currency: Currency): Promise<Account[]> {
    return this.request('GET', '/v1/account/list', undefined, { currency });
  }

  createExchange(request: ExchangeRequest): Promise<null> {
    return this.request('POST', '/v1/exchange', request);
  }

  getRate(from: Currency, to: Currency): Promise<number> {
    return this.request('GET', '/v1/rate', undefined, { from, to });
  }

  createPayment(payment: PaymentRequest): Promise<PaymentCreateResponse> {
    return this.request('POST', '/v1/payment', payment);
  }

  getPaymentByDocumentId(documentId: number | string): Promise<PaymentStatus> {
    return this.request('GET', `/v1/payment/document/${encodeURIComponent(documentId)}`);
  }

  getPaymentByUserRequestId(userRequestId: string): Promise<PaymentStatus> {
    return this.request('GET', `/v1/payment/${encodeURIComponent(userRequestId)}`);
  }

  getOrders(filters: OrdersFilters = {}): Promise<unknown> {
    return this.request('GET', '/v1/orders', undefined, asQuery(filters));
  }

  getTransactions(filters: TransactionsFilters = {}): Promise<TransactionsResponse> {
    return this.request('GET', '/v1/transactions', undefined, asQuery(filters));
  }

  getDepositAddress(currency: Currency): Promise<DepositAddressResponse> {
    return this.request('GET', `/v1/depositAddress/${encodeURIComponent(currency)}`);
  }

  getAutoConvertedUSDTtDepositAddress(account: string): Promise<DepositAddressResponse> {
    return this.request('GET', `/v1/depositAddressAutoUSDTt/${encodeURIComponent(account)}`);
  }

  getPrepaidServices(): Promise<PrepaidService[]> {
    return this.request('GET', '/v1/prepaid2/services');
  }

  getPrepaidRegions(): Promise<string[]> {
    return this.request('GET', '/v1/prepaid2/regions');
  }

  getPrepaidProducts(): Promise<PrepaidProduct[]> {
    return this.request('GET', '/v1/prepaid2/products');
  }

  getPrepaidDenominations(productId: number | string): Promise<PrepaidDenomination[]> {
    return this.request('GET', '/v1/prepaid2/denominations', undefined, { productid: productId });
  }

  startKyc(request: KycStartRequest): Promise<KycStartResponse> {
    return this.request('POST', '/v1/kyc/start', request);
  }

  getKycStatus(kycExternalUserId: string, sort: string): Promise<KycStatus> {
    return this.request(
      'GET',
      `/v1/kyc/status/${encodeURIComponent(kycExternalUserId)}/${encodeURIComponent(sort)}`,
    );
  }

  getKycStatusByUuid(uuid: string): Promise<KycStatus> {
    return this.request('GET', `/v1/kyc/statusByUuid/${encodeURIComponent(uuid)}`);
  }

  setKycData(uuid: string, data: Record<string, unknown>): Promise<null> {
    return this.request('POST', `/v1/kyc/setData/${encodeURIComponent(uuid)}`, data);
  }

  setKycPicture(uuid: string, type: string, jpeg: BodyInit): Promise<null> {
    return this.requestRaw(
      'POST',
      `/v1/kyc/setPicture/${encodeURIComponent(uuid)}/${encodeURIComponent(type)}`,
      jpeg,
      'image/jpeg',
    );
  }

  confirmKyc(uuid: string): Promise<null> {
    return this.request('POST', `/v1/kyc/confirm/${encodeURIComponent(uuid)}`, {});
  }

  private request<T>(
    method: string,
    path: string,
    body?: unknown,
    query?: Record<string, string | number | undefined>,
  ): Promise<T> {
    const rawBody = body === undefined ? '' : JSON.stringify(body);
    return this.requestRaw<T>(method, path, rawBody, 'application/json', query);
  }

  private async requestRaw<T>(
    method: string,
    path: string,
    body: BodyInit | string = '',
    contentType = 'application/json',
    query?: Record<string, string | number | undefined>,
  ): Promise<T> {
    const rawBody = body;
    const timestamp = Date.now().toString();
    const signature = calculateSignature(timestamp, rawBodyToSignable(rawBody), this.apiSecret);
    const url = this.buildUrl(path, query);
    const headers: Record<string, string> = {
      'API-Key': this.apiKey,
      'X-Request-Timestamp': timestamp,
      Signature: signature,
    };

    if (body !== '') {
      headers['Content-Type'] = contentType;
    }

    const response = await this.fetchImpl(url, {
      method,
      headers,
      body: body === '' ? undefined : body,
    });
    const responseBody = await response.text();

    if (!response.ok) {
      throw new CapitalistApi2Error(extractErrorMessage(responseBody), response.status, responseBody);
    }

    if (responseBody === '' || responseBody === 'null') {
      return null as T;
    }

    try {
      return JSON.parse(responseBody) as T;
    } catch {
      return responseBody as T;
    }
  }

  private buildUrl(path: string, query?: Record<string, string | number | undefined>): string {
    const url = new URL(`${this.baseUrl}${path}`);

    for (const [key, value] of Object.entries(query ?? {})) {
      if (value !== undefined) {
        url.searchParams.set(key, String(value));
      }
    }

    return url.toString();
  }
}

function rawBodyToSignable(body: BodyInit | string): string | ArrayBuffer | ArrayBufferView {
  if (typeof body === 'string') {
    return body;
  }

  if (body instanceof ArrayBuffer) {
    return body;
  }

  if (ArrayBuffer.isView(body)) {
    return body;
  }

  return '';
}

function asQuery<T extends object>(value: T): Record<string, string | number | undefined> {
  return value as Record<string, string | number | undefined>;
}

function extractErrorMessage(responseBody: string): string {
  try {
    const parsed = JSON.parse(responseBody) as { error?: unknown };
    if (typeof parsed.error === 'string') {
      return parsed.error;
    }
  } catch {
    // The API normally returns JSON errors, but keep the original body for diagnostics.
  }

  return responseBody || 'Capitalist API request failed';
}
