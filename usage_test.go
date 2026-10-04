package hanzoai

import (
	"context"
	"errors"
	"flag"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"
)

// token is a real IAM bearer for TestLive: go test . -run TestLive -token="$(hanzo auth token)".
var token = flag.String("token", "", "a Hanzo IAM bearer for TestLive; empty skips it")

// The bodies the gateway writes (hanzoai/ai routers/filter_balance.go). The
// free_plan_cap one is byte-for-byte what POST /v1/decisions answered org hanzo
// on 2026-10-04.
const (
	freePlanCap = `{"error":{"message":"Free plan: today's Kai requests are used. Upgrade for more: https://hanzo.ai/pay","type":"rate_limit_error","code":"free_plan_cap","class":"ours","resets_at":"2026-10-05T00:00:00Z","upgrade_url":"https://hanzo.ai/pay/cart?plan=dev","actions":[{"kind":"upgrade","label":"Upgrade your plan","url":"https://hanzo.ai/pay/cart?plan=dev","plan":"dev"},{"kind":"topup","label":"Add prepaid credit","url":"https://hanzo.ai/pay"}]}}`
	modelCap    = `{"error":{"message":"Kai has used its share of your plan for now. https://hanzo.ai/pay","type":"billing_error","code":"model_cap","class":"ours","model":"kai","fallback":"enso","resets_at":"2026-10-11T00:00:00Z","actions":[{"kind":"switch","label":"Try Enso","model":"enso"},{"kind":"topup","label":"Add prepaid credit","url":"https://hanzo.ai/pay"}]}}`
	allowance   = `{"error":{"message":"This request is outside what your plan includes. https://hanzo.ai/pay","type":"billing_error","code":"plan_allowance_used","class":"premium","model":"anthropic/claude-sonnet-4.5","resets_at":"2026-11-01T00:00:00Z","upgrade_url":"https://hanzo.ai/pay/cart?plan=max-5x","actions":[{"kind":"upgrade","label":"Upgrade your plan","url":"https://hanzo.ai/pay/cart?plan=max-5x","plan":"max-5x"},{"kind":"credits","label":"Continue with credits","url":"/v1/ai/limits"}]}}`
	paidPlan    = `{"error":{"message":"This model needs a paid plan or prepaid balance. https://hanzo.ai/pay","type":"billing_error","code":"paid_plan_required","class":"premium","model":"typesafe/jev-1.13","actions":[{"kind":"topup","label":"Add prepaid credit","url":"https://hanzo.ai/pay"}]}}`
	usageCap    = `{"error":{"message":"You've used this session's requests on your plan. They reset at 2026-10-05T03:00:00Z.","type":"rate_limit_error","code":"usage_cap_exceeded","limit":"session","resets_at":"2026-10-05T03:00:00Z"}}`
	balance     = `{"error":{"message":"Insufficient balance for this request.","type":"billing_error","code":"insufficient_balance"}}`
)

var sentinels = []error{
	ErrPlanAllowanceUsed, ErrPaidPlanRequired, ErrFreePlanCap,
	ErrModelCap, ErrUsageCapExceeded, ErrInsufficientBalance,
}

// answering answers every API call with status and body, and the IAM mint.
func answering(t *testing.T, status int, body string, header http.Header) *Client {
	t.Helper()
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/iam/oauth/token", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"access_token":"tok","expires_in":3600}`))
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range header {
			w.Header()[k] = v
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write([]byte(body))
	})
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return New(Options{ID: "cid", Secret: "sec", Base: srv.URL, Issuer: srv.URL, Resource: srv.URL})
}

// Every generated method returns the same *GenericOpenAPIError, but builds it
// three ways: decoded into the operation's typed refusal (decisions), decoded
// into the problem envelope, which a gateway body fails (limits), or not decoded
// at all (chat). The refusal must read the same through each.
var calls = map[string]func(*Client) (*http.Response, error){
	"decisions": func(c *Client) (*http.Response, error) {
		_, resp, err := c.AiAPI.PostDecisions(context.Background()).
			AiDecisionsRequest(AiDecisionsRequest{Model: "kai"}).Execute()
		return resp, err
	},
	"limits": func(c *Client) (*http.Response, error) {
		_, resp, err := c.AiAPI.AiLimits(context.Background()).Execute()
		return resp, err
	},
	"chat": func(c *Client) (*http.Response, error) {
		_, resp, err := c.AiAPI.PostChatCompletions(context.Background()).
			OpenaiChatCompletionRequest(OpenaiChatCompletionRequest{Model: PtrString("kai")}).Execute()
		return resp, err
	},
}

func TestUsageLimitError(t *testing.T) {
	for _, tc := range []struct {
		name   string
		status int
		body   string
		is     error
		want   UsageLimitError
	}{
		{"free_plan_cap", 429, freePlanCap, ErrFreePlanCap, UsageLimitError{
			Status: 429, Code: CodeFreePlanCap, Type: "rate_limit_error",
			Message: "Free plan: today's Kai requests are used. Upgrade for more: https://hanzo.ai/pay",
			Class:   "ours", ResetsAt: "2026-10-05T00:00:00Z", UpgradeURL: "https://hanzo.ai/pay/cart?plan=dev",
			Actions: []UsageAction{
				{Kind: "upgrade", Label: "Upgrade your plan", URL: "https://hanzo.ai/pay/cart?plan=dev", Plan: "dev"},
				{Kind: "topup", Label: "Add prepaid credit", URL: "https://hanzo.ai/pay"},
			},
		}},
		{"model_cap", 402, modelCap, ErrModelCap, UsageLimitError{
			Status: 402, Code: CodeModelCap, Type: "billing_error",
			Message: "Kai has used its share of your plan for now. https://hanzo.ai/pay",
			Class:   "ours", Model: "kai", Fallback: "enso", ResetsAt: "2026-10-11T00:00:00Z",
			Actions: []UsageAction{
				{Kind: "switch", Label: "Try Enso", Model: "enso"},
				{Kind: "topup", Label: "Add prepaid credit", URL: "https://hanzo.ai/pay"},
			},
		}},
		{"plan_allowance_used", 402, allowance, ErrPlanAllowanceUsed, UsageLimitError{
			Status: 402, Code: CodePlanAllowanceUsed, Type: "billing_error",
			Message: "This request is outside what your plan includes. https://hanzo.ai/pay",
			Class:   "premium", Model: "anthropic/claude-sonnet-4.5", ResetsAt: "2026-11-01T00:00:00Z",
			UpgradeURL: "https://hanzo.ai/pay/cart?plan=max-5x",
			Actions: []UsageAction{
				{Kind: "upgrade", Label: "Upgrade your plan", URL: "https://hanzo.ai/pay/cart?plan=max-5x", Plan: "max-5x"},
				{Kind: "credits", Label: "Continue with credits", URL: "/v1/ai/limits"},
			},
		}},
		{"paid_plan_required", 402, paidPlan, ErrPaidPlanRequired, UsageLimitError{
			Status: 402, Code: CodePaidPlanRequired, Type: "billing_error",
			Message: "This model needs a paid plan or prepaid balance. https://hanzo.ai/pay",
			Class:   "premium", Model: "typesafe/jev-1.13",
			Actions: []UsageAction{{Kind: "topup", Label: "Add prepaid credit", URL: "https://hanzo.ai/pay"}},
		}},
		{"usage_cap_exceeded", 429, usageCap, ErrUsageCapExceeded, UsageLimitError{
			Status: 429, Code: CodeUsageCapExceeded, Type: "rate_limit_error",
			Message:  "You've used this session's requests on your plan. They reset at 2026-10-05T03:00:00Z.",
			ResetsAt: "2026-10-05T03:00:00Z",
		}},
		{"insufficient_balance", 402, balance, ErrInsufficientBalance, UsageLimitError{
			Status: 402, Code: CodeInsufficientBalance, Type: "billing_error",
			Message: "Insufficient balance for this request.",
		}},
	} {
		for op, call := range calls {
			t.Run(tc.name+"/"+op, func(t *testing.T) {
				resp, err := call(answering(t, tc.status, tc.body, nil))
				if resp == nil || resp.StatusCode != tc.status {
					t.Fatalf("response %v, want status %d", resp, tc.status)
				}
				var ule *UsageLimitError
				if !errors.As(err, &ule) {
					t.Fatalf("errors.As(%T %v) found no *UsageLimitError", err, err)
				}
				if !equal(*ule, tc.want) {
					t.Errorf("got  %+v\nwant %+v", *ule, tc.want)
				}
				for _, s := range sentinels {
					if got := errors.Is(err, s); got != (s == tc.is) {
						t.Errorf("errors.Is(err, %v) = %v", s, got)
					}
				}
				// The generated error is still there for a caller who reads the body.
				if _, ok := errors.AsType[*GenericOpenAPIError](err); !ok {
					t.Errorf("%T is not a *GenericOpenAPIError", err)
				}
			})
		}
	}
}

// A refusal that is not a usage limit stays a plain generated error.
func TestNotAUsageLimit(t *testing.T) {
	for _, tc := range []struct {
		name   string
		status int
		body   string
	}{
		{"cloud refusal", 403, `{"error":"forbidden","reason":"no validated principal"}`},
		{"free lane", 429, `{"error":{"message":"The free pool is busy.","type":"rate_limit_error","code":"pool_busy"}}`},
		{"retryable balance", 503, `{"error":{"message":"Balance unavailable.","type":"api_error","code":"balance_unavailable"}}`},
		{"problem envelope", 402, `{"type":"about:blank","title":"Payment Required","status":402,"detail":"no balance","code":"insufficient_balance"}`},
		{"numeric code", 402, `{"error":{"message":"x","code":402}}`},
		{"not json", 502, `upstream unavailable`},
	} {
		for op, call := range calls {
			t.Run(tc.name+"/"+op, func(t *testing.T) {
				_, err := call(answering(t, tc.status, tc.body, nil))
				if err == nil {
					t.Fatal("no error")
				}
				if _, ok := errors.AsType[*UsageLimitError](err); ok {
					t.Errorf("%s read as a usage limit", tc.body)
				}
				for _, s := range sentinels {
					if errors.Is(err, s) {
						t.Errorf("errors.Is(err, %v) for %s", s, tc.body)
					}
				}
			})
		}
	}
}

func TestReadUsage(t *testing.T) {
	for _, tc := range []struct {
		name   string
		header map[string]string
		want   Usage
	}{
		// Measured on api.hanzo.ai.
		{"premium on credits", map[string]string{
			"x-hanzo-usage": "ok", "x-hanzo-usage-class": "premium", "x-hanzo-paid-by": "credits",
			"x-hanzo-served": "anthropic/claude-sonnet-4.5",
		}, Usage{Usage: "ok", Class: "premium", PaidBy: "credits", Served: "anthropic/claude-sonnet-4.5"}},
		{"free class", map[string]string{"x-hanzo-served": "enso"}, Usage{Served: "enso"}},
		{"refused", map[string]string{"x-hanzo-usage": "limited", "x-hanzo-usage-class": "ours"},
			Usage{Usage: "limited", Class: "ours"}},
		{"answered by the fallback", map[string]string{
			"X-Hanzo-Usage": "limited", "X-Hanzo-Usage-Class": "ours", "X-Hanzo-Paid-By": "free",
			"X-Hanzo-Fallback": "enso", "X-Hanzo-Served": "enso", "X-Hanzo-Usage-Reason": "model_cap",
		}, Usage{Usage: "limited", Class: "ours", PaidBy: "free", Fallback: "enso", Served: "enso", Reason: "model_cap"}},
		{"no headers", nil, Usage{}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			h := http.Header{}
			for k, v := range tc.header {
				h.Set(k, v)
			}
			if got := ReadUsage(h); got != tc.want {
				t.Errorf("direct: got %+v, want %+v", got, tc.want)
			}
			// And off a generated method's response, refused or answered.
			for _, status := range []int{200, 429} {
				resp, _ := calls["chat"](answering(t, status, freePlanCap, h))
				if got := ReadUsage(resp.Header); got != tc.want {
					t.Errorf("%d: got %+v, want %+v", status, got, tc.want)
				}
			}
		})
	}
}

// TestLive reads the live estate as the bearer -token names.
func TestLive(t *testing.T) {
	if *token == "" {
		t.Skip("no -token")
	}
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: DefaultBaseURL}}
	cfg.AddDefaultHeader("Authorization", "Bearer "+*token)
	c := NewAPIClient(cfg)
	ctx := context.Background()

	lim, resp, err := c.AiAPI.AiLimits(ctx).Execute()
	if err != nil {
		t.Fatalf("GET /v1/ai/limits: %v", err)
	}
	if !slices.Contains([]string{"ok", "near", "limited"}, lim.GetState()) || lim.GetPlan() == "" {
		t.Errorf("limits plan=%q state=%q", lim.GetPlan(), lim.GetState())
	}
	t.Logf("limits   %s plan=%s state=%s classes=%d actions=%d credits_after_allowance=%v",
		resp.Header.Get("X-Api-Version"), lim.GetPlan(), lim.GetState(), len(lim.GetClasses()), len(lim.GetActions()), lim.GetCreditsAfterAllowance())

	models, _, err := c.AiAPI.GetModels(ctx).Execute()
	if err != nil {
		t.Fatalf("GET /v1/models: %v", err)
	}
	classes, variable := map[string]int{}, 0
	for _, m := range models.GetData() {
		classes[m.GetClass()]++
		if m.Pricing.GetVariable() {
			variable++
		}
		if m.GetId() == "kai" && (m.GetClass() != "ours" || m.GetFamily() != "kai") {
			t.Errorf("kai class=%q family=%q, want ours/kai", m.GetClass(), m.GetFamily())
		}
	}
	if classes["premium"] == 0 || classes["ours"] == 0 || classes["free"] == 0 || variable == 0 {
		t.Errorf("models classes=%v variable=%d", classes, variable)
	}
	t.Logf("models   %d: classes=%v variable=%d", len(models.GetData()), classes, variable)

	syncs, _, err := c.SyncAPI.GetSync(ctx).Execute()
	if err != nil {
		t.Fatalf("GET /v1/sync: %v", err)
	}
	t.Logf("sync     %d link(s)", len(syncs.GetData()))

	state := "The user asked to delete their account."
	q := AiDecisionsQuestion{AiDecisionsChoice: &AiDecisionsChoice{Type: "choice",
		Criteria: map[string]AiDecisionsChoiceCriteriaValue{
			"yes": {String: PtrString("delete now")}, "no": {String: PtrString("ask to confirm")},
		}}}
	decision, resp, err := c.AiAPI.PostDecisions(ctx).AiDecisionsRequest(AiDecisionsRequest{
		Model: "kai", State: &AiDecisionSidesFalse{String: &state},
		Questions: map[string]AiDecisionsQuestion{"act": q},
	}).Execute()
	if resp == nil {
		t.Fatalf("POST /v1/decisions: %v", err)
	}
	u := ReadUsage(resp.Header)
	var ule *UsageLimitError
	switch {
	case err == nil:
		t.Logf("decide   %s answered %d question(s); usage=%q class=%q", decision.Model, len(decision.Answers), u.Usage, u.Class)
	case errors.As(err, &ule):
		if u.Usage != "limited" {
			t.Errorf("refused with X-Hanzo-Usage %q, want limited", u.Usage)
		}
		t.Logf("decide   %d %s class=%s resets=%s actions=%d is(ErrFreePlanCap)=%v usage=%q",
			ule.Status, ule.Code, ule.Class, ule.ResetsAt, len(ule.Actions), errors.Is(err, ErrFreePlanCap), u.Usage)
	default:
		t.Fatalf("POST /v1/decisions: %v %s", err, err.(*GenericOpenAPIError).Body())
	}

	_, resp, err = c.AiAPI.PostChatCompletions(ctx).OpenaiChatCompletionRequest(OpenaiChatCompletionRequest{
		Model: PtrString("enso"), MaxTokens: PtrInt32(16),
		Messages: []OpenaiChatCompletionMessage{{Role: PtrString("user"), Content: PtrString("Say hi.")}},
	}).Execute()
	if err != nil {
		t.Fatalf("POST /v1/chat/completions: %v", err)
	}
	if u := ReadUsage(resp.Header); u.Served != "enso" {
		t.Errorf("chat served %+v, want enso", u)
	}
	t.Logf("chat     %+v", ReadUsage(resp.Header))
}

func equal(a, b UsageLimitError) bool {
	return slices.Equal(a.Actions, b.Actions) &&
		a.Status == b.Status && a.Code == b.Code && a.Type == b.Type && a.Message == b.Message &&
		a.Class == b.Class && a.Model == b.Model && a.Fallback == b.Fallback &&
		a.ResetsAt == b.ResetsAt && a.UpgradeURL == b.UpgradeURL
}
