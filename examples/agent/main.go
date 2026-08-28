// agent — create an agent, run it, read the run back.
//
// Operations: POST /v1/agents (post_agents),
// POST /v1/agents/{ref}/run (post_agents_by_ref_run),
// GET /v1/agents/{ref}/runs (get_agents_by_ref_runs).
//
// `ref` accepts the public id (agent_...) or the org-unique name, so run and
// read both use the name just created without waiting for an id.
//
// A run is asynchronous, so the last step polls the run list until the run
// this program started reaches a terminal status. The run POST states its
// address and not its shape, so its generated method returns the raw
// *http.Response and the run is identified by reading the list rather than
// from the POST.
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
	"time"

	hanzoai "github.com/hanzoai/go-sdk/v8"
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

	// Names are org-unique, so a hardcoded one collides with the last run.
	name := fmt.Sprintf("sdk-example-%d", os.Getpid())
	instructions := "You answer in exactly one sentence."

	created, _, err := client.AgentsAPI.PostAgents(ctx).
		CreateAgentIn(hanzoai.CreateAgentIn{
			Name:         &name,
			Model:        &model,
			Instructions: &instructions,
		}).Execute()
	if err != nil {
		log.Fatalf("create agent: %v", err)
	}
	fmt.Printf("created  %s (%s)\n", created.GetName(), created.GetId())

	if _, err := client.AgentsAPI.PostAgentsByRefRun(ctx, name).Execute(); err != nil {
		log.Fatalf("run agent: %v", err)
	}
	fmt.Printf("started  a run on %s\n", name)

	run, err := poll(ctx, client, name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("status   %s in %dms\n", run.GetStatus(), run.GetDurationMs())
	fmt.Printf("output   %s\n", run.GetOutput())
}

// poll reads the run list until its newest run reaches a terminal status.
func poll(ctx context.Context, client *hanzoai.APIClient, ref string) (*hanzoai.AgentRunView, error) {
	for deadline := time.Now().Add(2 * time.Minute); time.Now().Before(deadline); time.Sleep(2 * time.Second) {
		list, _, err := client.AgentsAPI.GetAgentsByRefRuns(ctx, ref).Limit(1).Execute()
		if err != nil {
			return nil, fmt.Errorf("list runs: %w", err)
		}
		if len(list.Runs) == 0 {
			continue
		}
		// Runs are newest first.
		switch run := list.Runs[0]; run.GetStatus() {
		case "ok", "error", "failed", "cancelled":
			return &run, nil
		}
	}
	return nil, fmt.Errorf("run on %s did not finish within 2m", ref)
}
