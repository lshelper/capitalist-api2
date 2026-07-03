package capitalist

type Currency string

const (
	CurrencyRUR   Currency = "RUR"
	CurrencyUSD   Currency = "USD"
	CurrencyEUR   Currency = "EUR"
	CurrencyUSDT  Currency = "USDT"
	CurrencyUSDTt Currency = "USDTt"
	CurrencyUSDTb Currency = "USDTb"
	CurrencyETH   Currency = "ETH"
	CurrencyUSDC  Currency = "USDC"
	CurrencyUSDCb Currency = "USDCb"
	CurrencyBTC   Currency = "BTC"
)

type Account struct {
	Number   string   `json:"number"`
	Currency Currency `json:"currency"`
	Name     string   `json:"name"`
	Balance  float64  `json:"balance"`
}

type ExchangeRequest struct {
	FromAccount string  `json:"fromAccount"`
	ToAccount   string  `json:"toAccount"`
	Amount      float64 `json:"amount"`
}

type DepositAddressResponse struct {
	Address string `json:"address"`
}

type PaymentRequest struct {
	UserRequestID string         `json:"userRequestId"`
	AccountFrom   string         `json:"accountFrom"`
	Amount        float64        `json:"amount"`
	Currency      Currency       `json:"currency,omitempty"`
	Comment       string         `json:"comment,omitempty"`
	CallbackURL   string         `json:"callbackUrl,omitempty"`
	Payload       map[string]any `json:"payload"`
}

type PaymentCreateResponse struct {
	DocumentID int64 `json:"documentId"`
}

type PaymentStatus struct {
	State         string   `json:"state"`
	Fee           float64  `json:"fee"`
	DocumentID    int64    `json:"documentId"`
	Comment       *string  `json:"comment,omitempty"`
	Amount        float64  `json:"amount"`
	Currency      Currency `json:"currency"`
	Type          string   `json:"type"`
	AccountFrom   string   `json:"accountFrom"`
	Payload       string   `json:"payload"`
	UserRequestID string   `json:"userRequestId"`
	CallbackURL   *string  `json:"callbackUrl,omitempty"`
	TxID          *string  `json:"txId,omitempty"`
	DstAddress    *string  `json:"dstAddress,omitempty"`
}

type PaymentCallback struct {
	State         string   `json:"state"`
	Fee           float64  `json:"fee"`
	DocumentID    int64    `json:"documentId"`
	Comment       *string  `json:"comment,omitempty"`
	Amount        float64  `json:"amount"`
	Currency      Currency `json:"currency"`
	Type          string   `json:"type"`
	AccountFrom   string   `json:"accountFrom"`
	UserRequestID string   `json:"userRequestId"`
	CallbackURL   string   `json:"callbackUrl"`
}

type ListFilters struct {
	Limit       int    `url:"limit"`
	Offset      int    `url:"offset"`
	PeriodStart string `url:"periodStart"`
	PeriodEnd   string `url:"periodEnd"`
}

type OrdersFilters struct {
	ListFilters
	OrderNumber  string `url:"orderNumber"`
	MerchantName string `url:"merchantName"`
	OrderID      int64  `url:"orderId"`
	MerchantID   int64  `url:"merchantId"`
}

type TransactionsFilters struct {
	ListFilters
	TransactionID int64 `url:"transactionId"`
}

type Transaction struct {
	TransactionID int64    `json:"transactionId"`
	CreateDate    string   `json:"createDate"`
	ExecuteDate   string   `json:"executeDate,omitempty"`
	Type          string   `json:"type"`
	State         string   `json:"state"`
	Amount        float64  `json:"amount"`
	Currency      Currency `json:"currency"`
	PlanDate      string   `json:"planDate,omitempty"`
	Version       int64    `json:"version"`
	TxID          *string  `json:"txId,omitempty"`
	DstAddress    *string  `json:"dstAddress,omitempty"`
}

type TransactionsResponse struct {
	Transactions []Transaction `json:"transactions"`
	Count        int64         `json:"count"`
}

type PrepaidService struct {
	Description     string   `json:"description"`
	HasRegion       bool     `json:"hasRegion"`
	ID              int64    `json:"id"`
	Name            string   `json:"name"`
	Type            string   `json:"type"`
	UnfixedCurrency Currency `json:"unfixedCurrency,omitempty"`
	UnfixedRate     any      `json:"unfixedRate,omitempty"`
}

type PrepaidProduct struct {
	Category    string `json:"category"`
	Description string `json:"description"`
	ID          int64  `json:"id"`
	Instruction string `json:"instruction"`
	Name        string `json:"name"`
}

type PrepaidDenomination struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int64   `json:"stock"`
	Value string  `json:"value"`
}

type KYCStartRequest struct {
	KYCExternalUserID string `json:"kycExternalUserId"`
	CallbackURL       string `json:"callbackUrl,omitempty"`
	Sort              string `json:"sort"`
}

type KYCStartResponse struct {
	URLForUser string `json:"urlForUser"`
	UUID       string `json:"uuid"`
}

type KYCStatus struct {
	KYCExternalUserID string `json:"kycExternalUserId"`
	Status            string `json:"status"`
	Reason            string `json:"reason,omitempty"`
	UUID              string `json:"uuid"`
}
