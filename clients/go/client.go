package capitalist

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const DefaultBaseURL = "https://api2.capitalist.net"

type Client struct {
	apiKey     string
	apiSecret  string
	baseURL    string
	httpClient *http.Client
}

type Option func(*Client)

func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = strings.TrimRight(baseURL, "/")
	}
}

func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		if httpClient != nil {
			c.httpClient = httpClient
		}
	}
}

func NewClient(apiKey, apiSecret string, options ...Option) *Client {
	client := &Client{
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		baseURL:    DefaultBaseURL,
		httpClient: http.DefaultClient,
	}

	for _, option := range options {
		option(client)
	}

	return client
}

func (c *Client) ReadWhitelist(ctx context.Context) ([]string, error) {
	var result []string
	err := c.requestJSON(ctx, http.MethodGet, "/v1/whitelist", nil, nil, &result)
	return result, err
}

func (c *Client) AddWhitelist(ctx context.Context, ip []string) (any, error) {
	var result any
	err := c.requestJSON(ctx, http.MethodPost, "/v1/whitelist", map[string]any{"ip": ip}, nil, &result)
	return result, err
}

func (c *Client) RemoveWhitelist(ctx context.Context, ip []string) (any, error) {
	var result any
	err := c.requestJSON(ctx, http.MethodPost, "/v1/whitelist/remove", map[string]any{"ip": ip}, nil, &result)
	return result, err
}

func (c *Client) ListAccounts(ctx context.Context, currency Currency) ([]Account, error) {
	var result []Account
	err := c.requestJSON(ctx, http.MethodGet, "/v1/account/list", nil, url.Values{"currency": {string(currency)}}, &result)
	return result, err
}

func (c *Client) CreateExchange(ctx context.Context, request ExchangeRequest) error {
	return c.requestJSON(ctx, http.MethodPost, "/v1/exchange", request, nil, nil)
}

func (c *Client) GetRate(ctx context.Context, from, to Currency) (float64, error) {
	var result float64
	err := c.requestJSON(ctx, http.MethodGet, "/v1/rate", nil, url.Values{"from": {string(from)}, "to": {string(to)}}, &result)
	return result, err
}

func (c *Client) CreatePayment(ctx context.Context, payment PaymentRequest) (PaymentCreateResponse, error) {
	var result PaymentCreateResponse
	err := c.requestJSON(ctx, http.MethodPost, "/v1/payment", payment, nil, &result)
	return result, err
}

func (c *Client) GetPaymentByDocumentID(ctx context.Context, documentID any) (PaymentStatus, error) {
	var result PaymentStatus
	err := c.requestJSON(ctx, http.MethodGet, "/v1/payment/document/"+url.PathEscape(fmt.Sprint(documentID)), nil, nil, &result)
	return result, err
}

func (c *Client) GetPaymentByUserRequestID(ctx context.Context, userRequestID string) (PaymentStatus, error) {
	var result PaymentStatus
	err := c.requestJSON(ctx, http.MethodGet, "/v1/payment/"+url.PathEscape(userRequestID), nil, nil, &result)
	return result, err
}

func (c *Client) GetOrders(ctx context.Context, filters OrdersFilters) (any, error) {
	var result any
	err := c.requestJSON(ctx, http.MethodGet, "/v1/orders", nil, queryFromStruct(filters), &result)
	return result, err
}

func (c *Client) GetTransactions(ctx context.Context, filters TransactionsFilters) (any, error) {
	var result any
	err := c.requestJSON(ctx, http.MethodGet, "/v1/transactions", nil, queryFromStruct(filters), &result)
	return result, err
}

func (c *Client) StartKYC(ctx context.Context, request KYCStartRequest) (KYCStartResponse, error) {
	var result KYCStartResponse
	err := c.requestJSON(ctx, http.MethodPost, "/v1/kyc/start", request, nil, &result)
	return result, err
}

func (c *Client) GetKYCStatus(ctx context.Context, kycExternalUserID, sort string) (KYCStatus, error) {
	var result KYCStatus
	path := "/v1/kyc/status/" + url.PathEscape(kycExternalUserID) + "/" + url.PathEscape(sort)
	err := c.requestJSON(ctx, http.MethodGet, path, nil, nil, &result)
	return result, err
}

func (c *Client) GetKYCStatusByUUID(ctx context.Context, uuid string) (KYCStatus, error) {
	var result KYCStatus
	err := c.requestJSON(ctx, http.MethodGet, "/v1/kyc/statusByUuid/"+url.PathEscape(uuid), nil, nil, &result)
	return result, err
}

func (c *Client) SetKYCData(ctx context.Context, uuid string, data map[string]any) error {
	return c.requestJSON(ctx, http.MethodPost, "/v1/kyc/setData/"+url.PathEscape(uuid), data, nil, nil)
}

func (c *Client) SetKYCPicture(ctx context.Context, uuid, pictureType string, jpeg []byte) error {
	path := "/v1/kyc/setPicture/" + url.PathEscape(uuid) + "/" + url.PathEscape(pictureType)
	return c.requestRaw(ctx, http.MethodPost, path, jpeg, "image/jpeg", nil, nil)
}

func (c *Client) ConfirmKYC(ctx context.Context, uuid string) error {
	return c.requestJSON(ctx, http.MethodPost, "/v1/kyc/confirm/"+url.PathEscape(uuid), map[string]any{}, nil, nil)
}

func (c *Client) requestJSON(ctx context.Context, method, path string, body any, query url.Values, out any) error {
	var rawBody []byte
	var err error

	if body != nil {
		rawBody, err = json.Marshal(body)
		if err != nil {
			return fmt.Errorf("encode request body: %w", err)
		}
	}

	return c.requestRaw(ctx, method, path, rawBody, "application/json", query, out)
}

func (c *Client) requestRaw(ctx context.Context, method, path string, body []byte, contentType string, query url.Values, out any) error {
	requestURL, err := c.buildURL(path, query)
	if err != nil {
		return err
	}

	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	signature := CalculateSignature(timestamp, body, c.apiSecret)
	var reader io.Reader
	if len(body) > 0 {
		reader = bytes.NewReader(body)
	}

	req, err := http.NewRequestWithContext(ctx, method, requestURL, reader)
	if err != nil {
		return err
	}

	req.Header.Set("API-Key", c.apiKey)
	req.Header.Set("X-Request-Timestamp", timestamp)
	req.Header.Set("Signature", signature)
	if len(body) > 0 {
		req.Header.Set("Content-Type", contentType)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &APIError{
			Message:    extractErrorMessage(responseBody),
			StatusCode: resp.StatusCode,
			Body:       string(responseBody),
		}
	}

	if out == nil || len(responseBody) == 0 || string(responseBody) == "null" {
		return nil
	}

	if err := json.Unmarshal(responseBody, out); err != nil {
		if outString, ok := out.(*string); ok {
			*outString = string(responseBody)
			return nil
		}
		if outAny, ok := out.(*any); ok {
			*outAny = string(responseBody)
			return nil
		}
		return err
	}

	return nil
}

func (c *Client) buildURL(path string, query url.Values) (string, error) {
	parsed, err := url.Parse(c.baseURL + path)
	if err != nil {
		return "", err
	}

	if query != nil {
		parsed.RawQuery = query.Encode()
	}

	return parsed.String(), nil
}

func extractErrorMessage(body []byte) string {
	var payload struct {
		Error string `json:"error"`
	}

	if err := json.Unmarshal(body, &payload); err == nil && payload.Error != "" {
		return payload.Error
	}

	if len(body) > 0 {
		return string(body)
	}

	return "Capitalist API request failed"
}

func queryFromStruct(value any) url.Values {
	result := url.Values{}
	addQueryValues(result, reflect.ValueOf(value))
	return result
}

func addQueryValues(result url.Values, value reflect.Value) {
	if !value.IsValid() {
		return
	}
	if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return
		}
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return
	}

	valueType := value.Type()
	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i)
		fieldType := valueType.Field(i)

		if fieldType.Anonymous {
			addQueryValues(result, fieldValue)
			continue
		}

		key := fieldType.Tag.Get("url")
		if key == "" || key == "-" || fieldValue.IsZero() {
			continue
		}

		result.Set(key, fmt.Sprint(fieldValue.Interface()))
	}
}
