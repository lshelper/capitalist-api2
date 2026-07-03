export type Currency =
  | 'RUR'
  | 'USD'
  | 'EUR'
  | 'USDT'
  | 'USDTt'
  | 'USDTb'
  | 'ETH'
  | 'USDC'
  | 'USDCb'
  | 'BTC'
  | string;

export interface ClientOptions {
  apiKey: string;
  apiSecret: string;
  baseUrl?: string;
  fetchImpl?: typeof fetch;
}

export interface Account {
  number: string;
  currency: Currency;
  name: string;
  balance: number;
}

export interface ExchangeRequest {
  fromAccount: string;
  toAccount: string;
  amount: number;
}

export interface DepositAddressResponse {
  address: string;
}

export interface PaymentRequest {
  userRequestId: string;
  accountFrom: string;
  amount: number;
  currency?: Currency;
  comment?: string;
  callbackUrl?: string;
  payload: Record<string, unknown>;
}

export interface PaymentCreateResponse {
  documentId: number;
}

export interface PaymentStatus {
  state: 'PENDING' | 'EXECUTED' | 'DECLINED' | string;
  fee: number;
  documentId: number;
  comment?: string | null;
  amount: number;
  currency: Currency;
  type: string;
  accountFrom: string;
  payload: string;
  userRequestId: string;
  callbackUrl?: string | null;
  txId?: string;
  dstAddress?: string;
}

export interface PaymentCallback {
  state: 'EXECUTED' | 'DECLINED' | string;
  fee: number;
  documentId: number;
  comment?: string | null;
  amount: number;
  currency: Currency;
  type: string;
  accountFrom: string;
  userRequestId: string;
  callbackUrl: string;
}

export interface ListFilters {
  limit?: number;
  offset?: number;
  periodStart?: string;
  periodEnd?: string;
}

export interface OrdersFilters extends ListFilters {
  orderNumber?: string;
  merchantName?: string;
  orderId?: number;
  merchantId?: number;
}

export interface TransactionsFilters extends ListFilters {
  transactionId?: number;
}

export interface Transaction {
  transactionId: number;
  createDate: string;
  executeDate?: string;
  type: string;
  state: string;
  amount: number;
  currency: Currency;
  planDate?: string;
  version: number;
  txId?: string;
  dstAddress?: string;
}

export interface TransactionsResponse {
  transactions: Transaction[];
  count: number;
}

export interface PrepaidService {
  description: string;
  hasRegion: boolean;
  id: number;
  name: string;
  type: string;
  unfixedCurrency?: Currency;
  unfixedRate?: number | string;
}

export interface PrepaidProduct {
  category: string;
  description: string;
  id: number;
  instruction: string;
  name: string;
}

export interface PrepaidDenomination {
  id: number;
  name: string;
  price: number;
  stock: number;
  value: string;
}

export interface KycStartRequest {
  kycExternalUserId: string;
  callbackUrl?: string;
  sort: string;
}

export interface KycStartResponse {
  urlForUser: string;
  uuid: string;
}

export interface KycStatus {
  kycExternalUserId: string;
  status: 'INITIATED' | 'OPENED' | 'COMPLETED' | 'APPROVED' | 'DECLINED' | 'EXPIRED' | string;
  reason?: string;
  uuid: string;
}
