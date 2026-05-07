package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	capitalist "github.com/lshelper/capitalist-api2/clients/go"
	"github.com/lshelper/capitalist-api2/clients/go/examples/internal/exampleclient"
)

func main() {
	if len(os.Args) < 6 {
		log.Fatal("usage: go run ./examples/create-rucard-payment FROM_ACCOUNT CARD_NUMBER NAME SURNAME MIDNAME [AMOUNT] [CURRENCY]")
	}

	amount := 100.0
	if len(os.Args) >= 7 {
		parsed, err := strconv.ParseFloat(os.Args[6], 64)
		if err != nil {
			log.Fatalf("invalid amount: %v", err)
		}
		amount = parsed
	}

	currency := capitalist.CurrencyUSD
	if len(os.Args) >= 8 {
		currency = capitalist.Currency(os.Args[7])
	}

	client := exampleclient.New()
	response, err := client.CreatePayment(context.Background(), capitalist.PaymentRequest{
		UserRequestID: fmt.Sprintf("rucard-%d", time.Now().UnixMilli()),
		AccountFrom:   os.Args[1],
		Amount:        amount,
		Currency:      currency,
		Comment:       "Go RUCARD payout example",
		Payload: map[string]any{
			"type":    "RUCARD",
			"account": os.Args[2],
			"name":    os.Args[3],
			"surname": os.Args[4],
			"midname": os.Args[5],
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("documentId=%d\n", response.DocumentID)
}
