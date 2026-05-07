package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	capitalist "github.com/lshelper/capitalist-api2/clients/go"
	"github.com/lshelper/capitalist-api2/clients/go/examples/internal/exampleclient"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatal("usage: go run ./examples/create-exchange FROM_ACCOUNT TO_ACCOUNT AMOUNT")
	}

	amount, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		log.Fatalf("invalid amount: %v", err)
	}

	client := exampleclient.New()
	err = client.CreateExchange(context.Background(), capitalist.ExchangeRequest{
		FromAccount: os.Args[1],
		ToAccount:   os.Args[2],
		Amount:      amount,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("exchange created")
}
