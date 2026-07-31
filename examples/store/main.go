// store — put, get, delete.
//
// A full key/value round trip against the Hanzo KV service.
//
// Operations: PUT /v1/kv/keys/{key} (kv_setKey), GET /v1/kv/keys/{key}
// (kv_getKey), DELETE /v1/kv/keys/{key} (kv_deleteKey).
//
// KV is org-scoped, so it needs X-Org-Id alongside the API key.
//
//	HANZO_API_KEY=sk-... HANZO_ORG_ID=my-org go run ./examples/store
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	hanzoai "github.com/hanzoai/go-sdk"
)

func main() {
	ctx := context.Background()

	cfg := hanzoai.NewConfig("")
	if org := os.Getenv("HANZO_ORG_ID"); org != "" {
		cfg.AddDefaultHeader("X-Org-Id", org)
	}
	client := hanzoai.NewAPIClient(cfg)

	const key = "hanzo-sdk-example"

	set, _, err := client.KeysAPI.KvSetKey(ctx, key).
		KvSetKeyRequest(hanzoai.KvSetKeyRequest{Value: "hello from the Go SDK"}).Execute()
	if err != nil {
		log.Fatalf("put %s: %v", key, err)
	}
	fmt.Printf("put     %s = %q\n", key, set.GetValue())

	got, _, err := client.KeysAPI.KvGetKey(ctx, key).Execute()
	if err != nil {
		log.Fatalf("get %s: %v", key, err)
	}
	fmt.Printf("get     %s = %q\n", key, got.GetValue())

	if _, _, err := client.KeysAPI.KvDeleteKey(ctx, key).Execute(); err != nil {
		log.Fatalf("delete %s: %v", key, err)
	}
	fmt.Printf("delete  %s\n", key)
}
