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
