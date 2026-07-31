// money — what do I have, and what have I spent?
//
// Operations: GET /v1/billing/balance (cloud_get_v1_billing_balance) and
// GET /v1/billing/usage (cloud_get_v1_billing_usage).
//
// Both currently declare only a `default` response with no content schema in
// hanzo.yaml, so the generated methods hand back the raw *http.Response and
// there is nothing to unmarshal into. This decodes the JSON body directly.
// When the spec regains response schemas for these two, the decode goes away
// and the typed value is returned instead.
//
//	HANZO_API_KEY=sk-... go run ./examples/money
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	hanzoai "github.com/hanzoai/go-sdk"
)

func main() {
	ctx := context.Background()
	client := hanzoai.NewClient("")

	balance, err := client.BillingAPI.CloudGetV1BillingBalance(ctx).Execute()
	if err != nil {
		log.Fatalf("balance: %v", err)
	}
	fmt.Printf("balance  %s\n", decode(balance))

	usage, err := client.BillingAPI.CloudGetV1BillingUsage(ctx).Execute()
	if err != nil {
		log.Fatalf("usage: %v", err)
	}
	fmt.Printf("usage    %s\n", decode(usage))
}

// decode renders an untyped JSON response body.
func decode(resp *http.Response) string {
	defer resp.Body.Close()
	var body any
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		log.Fatalf("decode %s: %v", resp.Request.URL.Path, err)
	}
	out, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("encode: %v", err)
	}
	return string(out)
}
