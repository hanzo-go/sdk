// money — what do I have, and what have I spent?
//
// Operations: GET /v1/billing/balance (billing_billingBalance) and
// GET /v1/billing/usage (billing_billingUsage).
//
//	HANZO_API_KEY=sk-... go run ./examples/money
package main

import (
	"context"
	"fmt"
	"log"

	hanzoai "github.com/hanzoai/go-sdk"
)

func main() {
	ctx := context.Background()
	client := hanzoai.NewClient("")

	balance, _, err := client.BillingAPI.BillingBillingBalance(ctx).Execute()
	if err != nil {
		log.Fatalf("balance: %v", err)
	}
	fmt.Printf("balance    %d\n", balance.GetBalance())
	fmt.Printf("holds      %d\n", balance.GetHolds())
	fmt.Printf("available  %d\n", balance.GetAvailable())

	usage, _, err := client.BillingAPI.BillingBillingUsage(ctx).Execute()
	if err != nil {
		log.Fatalf("usage: %v", err)
	}
	fmt.Printf("\nusage      %+v\n", usage)
}
