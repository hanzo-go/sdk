// limits — where the plan stands, and what a refusal says.
//
// Operations: GET /v1/ai/limits (aiLimits), then POST /v1/decisions
// (post_decisions) on kai. Limits are shares, never amounts: each class reads a
// percent and a state, and actions name the ways on.
//
// A refused request comes back as the generated *GenericOpenAPIError, and
// errors.Is / errors.As reach the UsageLimitError inside it when the body's
// error.code is a plan or wallet refusal. ReadUsage reads the X-Hanzo-Usage
// headers off any response, refused or answered.
//
//	HANZO_CLIENT_ID=... HANZO_CLIENT_SECRET=... go run ./examples/limits
package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	hanzoai "github.com/hanzoai/go-sdk/v8"
)

func main() {
	ctx := context.Background()
	client := hanzoai.New(hanzoai.Options{})

	limits, _, err := client.AiAPI.AiLimits(ctx).Execute()
	if err != nil {
		log.Fatalf("limits: %v", err)
	}
	fmt.Printf("plan     %s, %s\n", limits.GetPlan(), limits.GetState())
	for name, class := range limits.GetClasses() {
		fmt.Printf("class    %s %d%% %s, paid by %s\n", name, class.GetPercent(), class.GetState(), class.GetPaying())
	}

	state := "The user asked to delete their account."
	_, resp, err := client.AiAPI.PostDecisions(ctx).AiDecisionsRequest(hanzoai.AiDecisionsRequest{
		Model: "kai",
		State: &hanzoai.AiDecisionSidesFalse{String: &state},
		Questions: map[string]hanzoai.AiDecisionsQuestion{"act": {AiDecisionsChoice: &hanzoai.AiDecisionsChoice{
			Type: "choice",
			Criteria: map[string]hanzoai.AiDecisionsChoiceCriteriaValue{
				"yes": {String: hanzoai.PtrString("delete now")},
				"no":  {String: hanzoai.PtrString("ask to confirm")},
			},
		}}},
	}).Execute()
	if resp != nil {
		fmt.Printf("usage    %+v\n", hanzoai.ReadUsage(resp.Header))
	}

	var limit *hanzoai.UsageLimitError
	switch {
	case err == nil:
		fmt.Println("decided")
	case errors.Is(err, hanzoai.ErrFreePlanCap), errors.Is(err, hanzoai.ErrModelCap):
		errors.As(err, &limit)
		fmt.Printf("refused  %s until %s\n", limit.Code, limit.ResetsAt)
		for _, a := range limit.Actions {
			fmt.Printf("  %-8s %s %s%s\n", a.Kind, a.Label, a.URL, a.Model)
		}
	case errors.As(err, &limit):
		fmt.Printf("refused  %d %s: %s\n", limit.Status, limit.Code, limit.Message)
	default:
		log.Fatalf("decisions: %v", err)
	}
}
