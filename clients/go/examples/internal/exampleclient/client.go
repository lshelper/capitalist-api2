package exampleclient

import (
	"log"
	"os"

	capitalist "github.com/lshelper/capitalist-api2/clients/go"
)

func New() *capitalist.Client {
	apiKey := os.Getenv("CAPITALIST_API2_KEY")
	apiSecret := os.Getenv("CAPITALIST_API2_SECRET")
	if apiKey == "" || apiSecret == "" {
		log.Fatal("set CAPITALIST_API2_KEY and CAPITALIST_API2_SECRET")
	}

	options := []capitalist.Option{}
	if baseURL := os.Getenv("CAPITALIST_API2_BASE_URL"); baseURL != "" {
		options = append(options, capitalist.WithBaseURL(baseURL))
	}

	return capitalist.NewClient(apiKey, apiSecret, options...)
}
