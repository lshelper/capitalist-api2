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
	if len(os.Args) < 3 {
		log.Fatal("usage: go run ./examples/get-rate FROM_CURRENCY TO_CURRENCY")
	}

	client := exampleclient.New()
	rate, err := client.GetRate(context.Background(), capitalist.Currency(os.Args[1]), capitalist.Currency(os.Args[2]))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s/%s %.8f\n", os.Args[1], os.Args[2], rate)
}
