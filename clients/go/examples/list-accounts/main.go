package main

import (
	"context"
	"fmt"
	"log"
	"os"

	capitalist "github.com/lshelper/capitalist-api2/clients/go"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./examples/list-accounts USD")
	}

	apiKey := os.Getenv("CAPITALIST_API2_KEY")
	apiSecret := os.Getenv("CAPITALIST_API2_SECRET")
	if apiKey == "" || apiSecret == "" {
		log.Fatal("set CAPITALIST_API2_KEY and CAPITALIST_API2_SECRET")
	}

	client := capitalist.NewClient(apiKey, apiSecret)
	if baseURL := os.Getenv("CAPITALIST_API2_BASE_URL"); baseURL != "" {
		client = capitalist.NewClient(apiKey, apiSecret, capitalist.WithBaseURL(baseURL))
	}

	accounts, err := client.ListAccounts(context.Background(), capitalist.Currency(os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}

	for _, account := range accounts {
		fmt.Printf("%s %s %s %.8f\n", account.Number, account.Currency, account.Name, account.Balance)
	}
}
