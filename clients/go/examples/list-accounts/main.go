package main

import (
	"context"
	"fmt"
	"log"
	"os"

	capitalist "github.com/lshelper/capitalist-api2/clients/go"
	"github.com/lshelper/capitalist-api2/clients/go/examples/internal/exampleclient"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./examples/list-accounts USD")
	}

	client := exampleclient.New()
	accounts, err := client.ListAccounts(context.Background(), capitalist.Currency(os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}

	for _, account := range accounts {
		fmt.Printf("%s %s %s %.8f\n", account.Number, account.Currency, account.Name, account.Balance)
	}
}
