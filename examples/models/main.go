// models — the catalogue, with no credential at all.
//
// Operation: GET /v1/models (get_models). Four of the document's 2479
// operations carry `security: []`; the other 2475 inherit the root
// `security: [{bearer: []}]`. This is one of the four, and the document says
// why in its own description: the catalogue takes no principal, so the route
// reads Authorization only to annotate gated SKUs and never to admit.
//
// That makes it the one example a reader can run before they have anything —
// no key, no org, no signup — which is the point of it being here. It is also
// the honest end of the auth contract: a public route that answers 200 to
// everyone is not evidence a credential works. `hello` is that evidence.
//
// The document states this address and not its shape, so the generated method
// hands back the raw *http.Response and the reply is decoded here. That decode
// is local on purpose: it describes what the route answers today, and it
// disappears the moment cloud declares the handler's type and this client is
// regenerated.
//
//	go run ./examples/models
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	hanzoai "github.com/hanzoai/go-sdk/v8"
)

func main() {
	resp, err := hanzoai.NewClient("").AiAPI.GetModels(context.Background()).Execute()
	if err != nil {
		log.Fatalf("models: %v", err)
	}
	defer resp.Body.Close()

	var catalogue struct {
		Data []struct {
			ID      string `json:"id"`
			OwnedBy string `json:"owned_by"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&catalogue); err != nil {
		log.Fatalf("decode: %v", err)
	}

	fmt.Printf("%s  %d model(s)\n", resp.Status, len(catalogue.Data))
	for _, m := range catalogue.Data[:min(len(catalogue.Data), 8)] {
		fmt.Printf("  %-28s %s\n", m.ID, m.OwnedBy)
	}
}
