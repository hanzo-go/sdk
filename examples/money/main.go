// money — what do I have, and what have I spent?
//
// Operations: GET /v1/billing/balance (get_billing_balance) and
// GET /v1/billing/usage (get_billing_usage), through the budget capability.
//
// Neither takes an org: both derive the tenant from the validated principal, so
// a credential can only read its own money.
//
// Both declare the address and not the shape — the document publishes them with
// no `responses` — so the generated methods hand back the raw *http.Response
// with nothing to unmarshal into. budget models the two shapes in one place, so
// this reads Money rather than a map. When cloud's handlers declare their Out
// types, the modelling goes and the fields stay.
//
//	HANZO_CLIENT_ID=... HANZO_CLIENT_SECRET=... go run ./examples/money
package main

import (
	"context"
	"fmt"
	"log"

	hanzoai "github.com/hanzoai/go-sdk/v8"
	"github.com/hanzoai/go-sdk/v8/budget"
)

func main() {
	ctx := context.Background()
	client := hanzoai.New(hanzoai.Options{})

	balance, err := client.Budget.Balance(ctx)
	if err != nil {
		log.Fatalf("balance: %v", err)
	}
	fmt.Printf("wallet   %s available, %s reserved, in %s\n",
		usd(balance.Available), usd(balance.Reserved), balance.Account)

	// The free-call allowance is a COUNT and the wallet is a SUM. They never
	// stand in for each other: a funded org can still be out of free calls.
	left, err := client.Budget.Left(ctx)
	if err != nil {
		log.Fatalf("allowance: %v", err)
	}
	if left.Left == nil {
		fmt.Printf("calls    unbounded on %s\n", left.Plan)
	} else {
		fmt.Printf("calls    %d of %d left this %s on %s\n",
			*left.Left, left.Limit, left.Window, left.Plan)
	}

	spent, err := client.Budget.Spent(ctx, budget.Filter{})
	if err != nil {
		log.Fatalf("usage: %v", err)
	}
	fmt.Printf("charges  %d\n", spent.Total)
	for _, charge := range spent.Items {
		fmt.Printf("  %s  %-28s %s\n",
			charge.At.Format("2006-01-02 15:04"), charge.Model, usd(charge.Amount))
	}
}

// usd renders a USD amount. Money counts the currency's minor units and is
// never a float, so the decimal point is put back only to print it — and only
// for a currency whose minor unit is a hundredth, which is why this reads the
// currency rather than assuming one.
func usd(m hanzoai.Money) string {
	return fmt.Sprintf("%d.%02d %s", m.Minor/100, m.Minor%100, m.Currency)
}
