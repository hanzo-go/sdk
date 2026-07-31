// agent — create it, run it, read it back.
//
// Operations: POST /v1/agents (cloud_AgentsController.Create),
// POST /v1/agents/{ref}/run (cloud_AgentsController.Run),
// GET /v1/agents/{ref} (cloud_AgentsController.Get).
//
// Agents are org-scoped, so this needs X-Org-Id alongside the API key.
//
//	HANZO_API_KEY=sk-... HANZO_ORG_ID=my-org go run ./examples/agent
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

	model := os.Getenv("HANZO_MODEL")
	if model == "" {
		model = "zen-1"
	}

	cfg := hanzoai.NewConfig("")
	if org := os.Getenv("HANZO_ORG_ID"); org != "" {
		cfg.AddDefaultHeader("X-Org-Id", org)
	}
	client := hanzoai.NewAPIClient(cfg)

	instructions := "You answer in exactly one sentence."
	created, _, err := client.AgentsAPIAPI.CloudAgentsControllerCreate(ctx).
		CloudAgentsCreateAgentRequest(hanzoai.CloudAgentsCreateAgentRequest{
			Name:         "sdk-example",
			Model:        model,
			Instructions: &instructions,
		}).Execute()
	if err != nil {
		log.Fatalf("create agent: %v", err)
	}
	ref := created.GetId()
	fmt.Printf("created  %s (%s)\n", ref, created.GetName())

	input := "What is Hanzo?"
	run, _, err := client.AgentsAPIAPI.CloudAgentsControllerRun(ctx, ref).
		CloudAgentsRunRequest(hanzoai.CloudAgentsRunRequest{Input: &input}).Execute()
	if err != nil {
		log.Fatalf("run agent: %v", err)
	}
	fmt.Printf("ran      %s -> %s\n", run.GetId(), run.GetStatus())
	fmt.Printf("output   %s\n", run.GetOutput())

	agent, _, err := client.AgentsAPIAPI.CloudAgentsControllerGet(ctx, ref).Execute()
	if err != nil {
		log.Fatalf("get agent: %v", err)
	}
	fmt.Printf("read     %s status=%s\n", agent.GetName(), agent.GetStatus())
}
