// hello — prove the key works.
//
// Operation: GET /v1/keys (get_keys), the caller's own API keys.
//
// This flow's whole job is to FAIL on a bad key, so the route is chosen by
// probing api.hanzo.ai rather than by reading the document. Three
// identity-shaped routes are disqualified for answering 200 to a caller
// carrying no credential at all — GET /v1/ai/account (type="anonymous-user"),
// GET /v1/iam/whoami and GET /v1/iam/account (both 200 {"status":"error"}) —
// and a hello built on any of them prints a cheerful identity for a key that
// would be refused everywhere else.
//
// /v1/keys answers 403 {"code":"forbidden","error":"sign in to manage API
// keys"} with no key and with a bogus one, while the nonsense sibling
// /v1/keys-zzq9 answers 404, so the 403 is this route refusing rather than a
// wildcard route swallowing the address.
//
// This replaces bot_authMe (GET /v1/bot/auth/me), which no longer resolves:
// cloud relays all of /v1/bot through one app.All("/v1/bot/*"), so the document
// carries /v1/bot/{wildcard1} and no operation at that address.
//
//	HANZO_API_KEY=sk-... go run ./examples/hello
package main

import (
	"context"
	"fmt"
	"log"

	hanzoai "github.com/hanzoai/go-sdk/v8"
)

func main() {
	client := hanzoai.NewClient("")

	listing, _, err := client.AccountAPI.GetAccountKeys(context.Background()).Execute()
	if err != nil {
		log.Fatalf("keys: %v", err)
	}

	keys := listing.Keys
	fmt.Printf("the key is accepted; this org holds %d key(s)\n", len(keys))
	for _, key := range keys {
		// Prefix is the recognizable head of a key, never enough to use one.
		fmt.Printf("  %-12s %s\n", key.GetType(), key.GetPrefix())
	}
}
