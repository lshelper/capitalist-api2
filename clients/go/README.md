# Capitalist API2 Go client

## Usage

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	capitalist "github.com/lshelper/capitalist-api2/clients/go"
)

func main() {
	client := capitalist.NewClient(
		os.Getenv("CAPITALIST_API2_KEY"),
		os.Getenv("CAPITALIST_API2_SECRET"),
	)

	accounts, err := client.ListAccounts(context.Background(), capitalist.CurrencyUSD)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(accounts)
}
```

## Test

```bash
go test ./...
```

## Docker

Run the Go client without installing Go locally:

```bash
docker compose build
docker compose run --rm go
```

Run a specific command:

```bash
docker compose run --rm go go test ./...
docker compose run --rm go go run ./examples/list-accounts USD
```

For real API calls, pass credentials from your shell:

```bash
export CAPITALIST_API2_KEY="your-api-key"
export CAPITALIST_API2_SECRET="your-api-secret"
export CAPITALIST_API2_BASE_URL="https://api2.capitalist.net"
docker compose run --rm go go run ./examples/list-accounts USD
```

Open a shell in the Go container:

```bash
docker compose run --rm go sh
```

Clean up Docker resources for this client:

```bash
docker compose down -v
```
