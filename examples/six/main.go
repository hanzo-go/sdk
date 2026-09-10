// six — budget, policy, audit, search, kb and graph in one run.
//
// The six are one client because they share one answer type, one refusal
// vocabulary, one instant and one correlation id. This walks that: what the
// wallet holds decides whether to ask, the wallet's own account names the org a
// policy question is scoped to, a search comes back with its degradation stated
// rather than hidden, a graph write answers an id, and that id is what finds
// the row the server wrote about it.
//
// Nothing here is a refusal you catch. A budget that says no and a policy that
// says no are values: the answer's denied arm carries the code and the cures.
//
//	HANZO_CLIENT_ID=... HANZO_CLIENT_SECRET=... go run ./examples/six
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	hanzoai "github.com/hanzoai/go-sdk/v8"
	"github.com/hanzoai/go-sdk/v8/audit"
	"github.com/hanzoai/go-sdk/v8/graph"
	"github.com/hanzoai/go-sdk/v8/kb"
	"github.com/hanzoai/go-sdk/v8/search"
)

func main() {
	ctx := context.Background()
	client := hanzoai.New(hanzoai.Options{})
	started := time.Now().UTC()

	// BUDGET — a count and a sum, which never stand in for each other.
	left, err := client.Budget.Left(ctx)
	if err != nil {
		log.Fatalf("allowance: %v", err)
	}
	if left.Left == nil {
		fmt.Printf("budget   %s, calls unbounded\n", left.Plan)
	} else {
		fmt.Printf("budget   %s, %d of %d calls left this %s\n",
			left.Plan, *left.Left, left.Limit, left.Window)
	}

	balance, err := client.Budget.Balance(ctx)
	if err != nil {
		log.Fatalf("balance: %v", err)
	}
	fmt.Printf("wallet   %d %s available in %s\n",
		balance.Available.Cents, balance.Available.Currency, balance.Account)

	// POLICY — may this subject write here? The object is org-scoped and the
	// org is not something a client may assert, so it is read off the wallet
	// the answer above already named.
	subject := balance.Account
	org, _, _ := strings.Cut(subject, "/")
	decision, err := client.Policy.Check(ctx, subject, "write", org+"/graph")
	if err != nil {
		log.Fatalf("policy: %v", err)
	}
	fmt.Printf("policy   %s write %s: allow=%v %s\n",
		decision.Sub, decision.Obj, decision.Allow, decision.Reason)

	// SEARCH — one ranked set over the whole corpus, kb included. Partial is
	// the field that says a leg was down; an answer that ignores it reads a
	// truncated corpus as a complete one.
	answer := client.Search.Find(ctx, "runbook", search.Opts{Kinds: []string{"kb." + kb.Page}, Limit: 5})
	hits, err := answer.Value()
	var denied *hanzoai.Denied
	switch {
	case errors.As(err, &denied):
		fmt.Printf("search   denied (%s): %s\n", denied.Code, denied.Reason)
		for _, cure := range denied.Cures {
			fmt.Printf("           %s at %s\n", cure.Kind, cure.URL)
		}
	case err != nil:
		log.Fatalf("search: %v", err)
	default:
		fmt.Printf("search   %d hit(s) in %s mode, partial=%v, took %s\n",
			len(hits.Items), hits.Mode, hits.Partial, hits.Took)
		for _, hit := range hits.Items {
			fmt.Printf("           %-10s %s\n", hit.Kind, hit.Title)
		}
		for _, backend := range hits.Backends {
			if backend.Status != "ok" {
				fmt.Printf("           %s is %s: %s\n", backend.Name, backend.Status, backend.Error)
			}
		}
	}

	// GRAPH — one assertion with provenance and time. Nothing overwrites
	// anything, so running this twice records once and counts the second as a
	// duplicate.
	entity := fmt.Sprintf("example/six/%d", os.Getpid())
	wrote := client.Graph.Assert(ctx, []graph.Fact{{
		Entity:   entity,
		Relation: "title",
		Value:    "the six capabilities, exercised",
		At:       started,
		Source:   "go-sdk/examples/six",
	}})
	written, err := wrote.Value()
	var held *hanzoai.Held
	switch {
	case errors.As(err, &denied):
		fmt.Printf("graph    denied (%s): %s\n", denied.Code, denied.Reason)
	case errors.As(err, &held):
		fmt.Printf("graph    held %s on %s: %s\n", held.ID, held.Clause, held.Reason)
	case err != nil:
		log.Fatalf("graph: %v", err)
	default:
		fmt.Printf("graph    %d recorded, %d duplicate, %d refused %v\n",
			written.Recorded, written.Duplicate, written.Refused, written.Reasons)

		resolved, err := client.Graph.Resolve(ctx, entity, "title", time.Time{})
		if err != nil {
			log.Fatalf("resolve: %v", err)
		}
		if resolved.Known {
			fmt.Printf("graph    in force: %q from %s\n", resolved.Winner.Value, resolved.Winner.Source)
		}
	}

	// AUDIT — the trail of what just happened. Every answer carries the
	// x-request-id, and that is the join: the server wrote the row, the client
	// only reads it.
	//
	// The route accepts no requestId filter yet, so Filter.Request narrows
	// client-side over the window this run opened. It stops being a scan the
	// day GET /v1/audit takes the id every row already carries.
	fmt.Printf("audit    rows for request %s\n", wrote.Request)
	for event, err := range client.Audit.All(ctx, audit.Filter{Request: wrote.Request, Since: started}) {
		if err != nil {
			log.Fatalf("audit: %v", err)
		}
		fmt.Printf("           %s %-24s %-8s %s %s\n",
			event.At.Format(time.RFC3339), event.Action, event.Result, event.Method, event.Path)
	}
}
