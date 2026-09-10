// store — provision a KV store, read it back, delete it.
//
// Operations: POST /v1/provisioning/kv (post_provisioning_kv), GET /v1/provisioning/kv/{name} (get_provisioning_kv_by_name),
// DELETE /v1/provisioning/kv/{name} (delete_provisioning_kv_by_name).
//
// This is the provisioning plane, and it is the one that answers. The value
// plane — /v1/kv/keys/{key}, kv_setKey/kv_getKey/kv_deleteKey — is authored in
// the spec but is mounted nowhere: it replies 404 to GET and 405 to PUT and
// DELETE at api.hanzo.ai, while the provisioning plane replies 403. An example
// may only call what routes.
//
// KV is org-scoped. The org is the validated principal's own — it is read
// off the credential, never named in a request — so there is nothing to set.
//
//	HANZO_CLIENT_ID=... HANZO_CLIENT_SECRET=... go run ./examples/store
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	hanzoai "github.com/hanzoai/go-sdk/v8"
)

func main() {
	ctx := context.Background()

	client := hanzoai.New(hanzoai.Options{})

	// Names are org-unique, so a hardcoded one collides with the last run.
	name := fmt.Sprintf("sdk-example-%d", os.Getpid())

	created, _, err := client.ProvisioningAPI.PostProvisioningKv(ctx).
		ProvisionRequest(hanzoai.ProvisionRequest{Name: &name}).Execute()
	if err != nil {
		log.Fatalf("provision %s: %v", name, err)
	}
	// Delete in a defer, so a failed read still cleans up instead of leaving
	// the store behind for the next run to collide with.
	defer func() {
		if _, err := client.ProvisioningAPI.DeleteProvisioningKvByName(ctx, name).Execute(); err != nil {
			log.Fatalf("delete %s: %v", name, err)
		}
		fmt.Printf("delete   %s\n", name)
	}()
	fmt.Printf("create   %s (%s)\n", created.GetName(), created.GetStatus())

	got, _, err := client.ProvisioningAPI.GetProvisioningKvByName(ctx, name).Execute()
	if err != nil {
		log.Fatalf("read %s: %v", name, err)
	}
	fmt.Printf("read     %s kind=%s host=%s\n", got.GetName(), got.GetKind(), got.GetHost())
}
