// agent — create it, run it, read it back.
//
// Operations: POST /v1/agents (cloud_post_v1_agents),
// POST /v1/agents/{ref}/run (cloud_post_v1_agents_by_ref_run),
// GET /v1/agents/{ref} (cloud_get_v1_agents_ref).
//
// Agents are org-scoped, so this needs X-Org-Id alongside the API key.
//
// The run operation currently declares only a `default` response with no
// content schema in hanzo.yaml, so its generated method hands back the raw
// *http.Response. Create and read are typed.
//
//	HANZO_API_KEY=sk-... HANZO_ORG_ID=my-org go run ./examples/agent
package main

import (
	"context"
	"encoding/json"
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

	name := "sdk-example"
	instructions := "You answer in exactly one sentence."
	created, _, err := client.AgentsAPI.CloudPostV1Agents(ctx).
		CloudCreateAgentIn(hanzoai.CloudCreateAgentIn{
			Name:         &name,
			Model:        &model,
			Instructions: &instructions,
		}).Execute()
	if err != nil {
		log.Fatalf("create agent: %v", err)
	}
	ref := created.GetId()
	fmt.Printf("created  %s (%s)\n", ref, created.GetName())

	run, err := client.AgentsAPI.CloudPostV1AgentsByRefRun(ctx, ref).Execute()
	if err != nil {
		log.Fatalf("run agent: %v", err)
	}
	defer run.Body.Close()
	var result any
	if err := json.NewDecoder(run.Body).Decode(&result); err != nil {
		log.Fatalf("decode run: %v", err)
	}
	out, err := json.Marshal(result)
	if err != nil {
		log.Fatalf("encode run: %v", err)
	}
	fmt.Printf("ran      %s\n", out)

	agent, _, err := client.AgentsAPI.CloudGetV1AgentsRef(ctx, ref).Execute()
	if err != nil {
		log.Fatalf("get agent: %v", err)
	}
	fmt.Printf("read     %s runs=%d\n", agent.GetName(), agent.GetRuns())
}
