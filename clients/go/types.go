package capitalist

type Currency string

const (
	CurrencyRUR   Currency = "RUR"
	CurrencyUSD   Currency = "USD"
	CurrencyEUR   Currency = "EUR"
	CurrencyUSDT  Currency = "USDT"
	CurrencyUSDTt Currency = "USDTt"
	CurrencyETH   Currency = "ETH"
	CurrencyUSDC  Currency = "USDC"
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
