go-bankcode
---

![CI](https://github.com/knsh14/go-bankcode/workflows/CI/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/knsh14/go-bankcode.svg)](https://pkg.go.dev/github.com/knsh14/go-bankcode)

# About
unofficial [bankcode API]( https://bankcode-jp.com/ ) client for Go

# Example

```
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/knsh14/go-bankcode"
)

func main() {
	client, err := bankcode.NewClient(bankcode.WithAPIKey("BANKCODE_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	banks, err := client.ListBanks(context.Background(), &ListParameter{
		Limit: 10,
	})
	if err != nil {
		log.Fatal(err)
	}
	if len(banks.Data) > 0 {
		for _, bank := range banks.Data {
			fmt.Println(bank.Code, bank.Name)
		}
	}
}
```

# License
MIT License
