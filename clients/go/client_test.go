package capitalist

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDepositAddress(t *testing.T) {
	var gotPath string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		if r.ContentLength > 0 {
			t.Fatalf("ContentLength = %d, want 0", r.ContentLength)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"address":"wallet-address"}`))
	}))
	defer server.Close()

	client := NewClient("key", "secret", WithBaseURL(server.URL))
	result, err := client.GetDepositAddress(context.Background(), CurrencyUSDTb)
	if err != nil {
		t.Fatal(err)
	}

	if gotPath != "/v1/depositAddress/USDTb" {
		t.Fatalf("path = %s, want /v1/depositAddress/USDTb", gotPath)
	}
	if result.Address != "wallet-address" {
		t.Fatalf("address = %s, want wallet-address", result.Address)
	}
}

func TestGetAutoConvertedUSDTtDepositAddress(t *testing.T) {
	var gotPath string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		if r.ContentLength > 0 {
			t.Fatalf("ContentLength = %d, want 0", r.ContentLength)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"address":"wallet-address"}`))
	}))
	defer server.Close()

	client := NewClient("key", "secret", WithBaseURL(server.URL))
	result, err := client.GetAutoConvertedUSDTtDepositAddress(context.Background(), "U0123504")
	if err != nil {
		t.Fatal(err)
	}

	if gotPath != "/v1/depositAddressAutoUSDTt/U0123504" {
		t.Fatalf("path = %s, want /v1/depositAddressAutoUSDTt/U0123504", gotPath)
	}
	if result.Address != "wallet-address" {
		t.Fatalf("address = %s, want wallet-address", result.Address)
	}
}

func TestGetPrepaidServices(t *testing.T) {
	var gotPath string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		if r.ContentLength > 0 {
			t.Fatalf("ContentLength = %d, want 0", r.ContentLength)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"id":3,"name":"Steam","description":"Top up","hasRegion":false,"type":"unfixed","unfixedCurrency":"KZT","unfixedRate":0.00205482}]`))
	}))
	defer server.Close()

	client := NewClient("key", "secret", WithBaseURL(server.URL))
	result, err := client.GetPrepaidServices(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if gotPath != "/v1/prepaid2/services" {
		t.Fatalf("path = %s, want /v1/prepaid2/services", gotPath)
	}
	if result[0].ID != 3 {
		t.Fatalf("id = %d, want 3", result[0].ID)
	}
	if result[0].Type != "unfixed" {
		t.Fatalf("type = %s, want unfixed", result[0].Type)
	}
}

func TestGetPrepaidRegions(t *testing.T) {
	var gotPath string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		if r.ContentLength > 0 {
			t.Fatalf("ContentLength = %d, want 0", r.ContentLength)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`["Asia","Macau"]`))
	}))
	defer server.Close()

	client := NewClient("key", "secret", WithBaseURL(server.URL))
	result, err := client.GetPrepaidRegions(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if gotPath != "/v1/prepaid2/regions" {
		t.Fatalf("path = %s, want /v1/prepaid2/regions", gotPath)
	}
	if len(result) != 2 || result[1] != "Macau" {
		t.Fatalf("regions = %#v, want Asia/Macau", result)
	}
}

func TestGetPrepaidProducts(t *testing.T) {
	var gotPath string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		if r.ContentLength > 0 {
			t.Fatalf("ContentLength = %d, want 0", r.ContentLength)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"id":98,"name":"PlayStation Store Wallet United Kingdom","category":"Gift Card, PlayStation","description":"PlayStation card","instruction":"Redeem instructions"}]`))
	}))
	defer server.Close()

	client := NewClient("key", "secret", WithBaseURL(server.URL))
	result, err := client.GetPrepaidProducts(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if gotPath != "/v1/prepaid2/products" {
		t.Fatalf("path = %s, want /v1/prepaid2/products", gotPath)
	}
	if result[0].ID != 98 {
		t.Fatalf("id = %d, want 98", result[0].ID)
	}
	if result[0].Category != "Gift Card, PlayStation" {
		t.Fatalf("category = %s, want Gift Card, PlayStation", result[0].Category)
	}
}

func TestGetPrepaidDenominations(t *testing.T) {
	var gotPath string
	var gotQuery string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotQuery = r.URL.RawQuery
		if r.ContentLength > 0 {
			t.Fatalf("ContentLength = %d, want 0", r.ContentLength)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"id":527,"name":"10 GBP","price":13.45,"stock":5,"value":"10 GBP"}]`))
	}))
	defer server.Close()

	client := NewClient("key", "secret", WithBaseURL(server.URL))
	result, err := client.GetPrepaidDenominations(context.Background(), 98)
	if err != nil {
		t.Fatal(err)
	}

	if gotPath != "/v1/prepaid2/denominations" {
		t.Fatalf("path = %s, want /v1/prepaid2/denominations", gotPath)
	}
	if gotQuery != "productid=98" {
		t.Fatalf("query = %s, want productid=98", gotQuery)
	}
	if result[0].Price != 13.45 {
		t.Fatalf("price = %f, want 13.45", result[0].Price)
	}
}

func TestGetPaymentByDocumentIDParsesCryptoMetadata(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"state":"EXECUTED","fee":0,"documentId":123,"amount":10,"currency":"USDTb","type":"USDTBSC","accountFrom":"U0123504","payload":"{}","userRequestId":"request-1","txId":"tx-hash","dstAddress":"wallet-address"}`))
	}))
	defer server.Close()

	client := NewClient("key", "secret", WithBaseURL(server.URL))
	result, err := client.GetPaymentByDocumentID(context.Background(), 123)
	if err != nil {
		t.Fatal(err)
	}

	if result.TxID == nil || *result.TxID != "tx-hash" {
		t.Fatalf("txId = %#v, want tx-hash", result.TxID)
	}
	if result.DstAddress == nil || *result.DstAddress != "wallet-address" {
		t.Fatalf("dstAddress = %#v, want wallet-address", result.DstAddress)
	}
}

func TestGetTransactionsParsesCryptoMetadata(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"transactions":[{"transactionId":8367664,"createDate":"2014-02-18T21:47:29.452Z","executeDate":"2014-02-18T21:47:39.543Z","type":"OUT","state":"EXECUTED","amount":500,"currency":"USD","planDate":"2014-02-18T21:21:53.314Z","version":1,"txId":"tx-hash","dstAddress":"wallet-address"}],"count":1}`))
	}))
	defer server.Close()

	client := NewClient("key", "secret", WithBaseURL(server.URL))
	result, err := client.GetTransactions(context.Background(), TransactionsFilters{})
	if err != nil {
		t.Fatal(err)
	}

	if result.Transactions[0].TxID == nil || *result.Transactions[0].TxID != "tx-hash" {
		t.Fatalf("txId = %#v, want tx-hash", result.Transactions[0].TxID)
	}
	if result.Transactions[0].DstAddress == nil || *result.Transactions[0].DstAddress != "wallet-address" {
		t.Fatalf("dstAddress = %#v, want wallet-address", result.Transactions[0].DstAddress)
	}
}
