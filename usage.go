// Usage limits: the refusal a plan or a wallet answers with, and the headers an
// AI answer carries about where its class stands. Hand-written beside the
// generated client, like hanzo.go, under the names the JS and Python SDKs use.
//
//	_, resp, err := client.AiAPI.PostDecisions(ctx).AiDecisionsRequest(req).Execute()
//	if errors.Is(err, hanzoai.ErrFreePlanCap) { ... }
//	var ule *hanzoai.UsageLimitError
//	if errors.As(err, &ule) { fmt.Println(ule.ResetsAt, ule.Actions) }
//	fmt.Println(hanzoai.ReadUsage(resp.Header).Usage)
//
// GET /v1/ai/limits (client.AiAPI.AiLimits) answers where every class stands;
// PUT /v1/ai/limits (AiSetLimits) turns credits after the allowance on or off.
package hanzoai

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

// The refusal codes a plan or a wallet answers with, in the body's error.code.
const (
	// CodePlanAllowanceUsed: the plan's included usage of the model's class is
	// used and nothing else may pay. 402.
	CodePlanAllowanceUsed = "plan_allowance_used"
	// CodePaidPlanRequired: the model needs a paid plan or prepaid balance. 402.
	CodePaidPlanRequired = "paid_plan_required"
	// CodeFreePlanCap: the free plan's daily cap on the model is used. 429,
	// with Retry-After.
	CodeFreePlanCap = "free_plan_cap"
	// CodeModelCap: the model used its share of the plan; Fallback names the
	// model that answers instead. 402.
	CodeModelCap = "model_cap"
	// CodeUsageCapExceeded: a session or day request window is spent. 429.
	CodeUsageCapExceeded = "usage_cap_exceeded"
	// CodeInsufficientBalance: the wallet's known balance cannot cover the
	// request. 402.
	CodeInsufficientBalance = "insufficient_balance"
)

// One sentinel per code. errors.Is(err, ErrModelCap) holds for any error a
// generated method returns whose body carries that code.
var (
	ErrPlanAllowanceUsed   = errors.New("hanzoai: " + CodePlanAllowanceUsed)
	ErrPaidPlanRequired    = errors.New("hanzoai: " + CodePaidPlanRequired)
	ErrFreePlanCap         = errors.New("hanzoai: " + CodeFreePlanCap)
	ErrModelCap            = errors.New("hanzoai: " + CodeModelCap)
	ErrUsageCapExceeded    = errors.New("hanzoai: " + CodeUsageCapExceeded)
	ErrInsufficientBalance = errors.New("hanzoai: " + CodeInsufficientBalance)
)

// limits is each code's sentinel and the status the gateway sends it with.
var limits = map[string]struct {
	err    error
	status int
}{
	CodePlanAllowanceUsed:   {ErrPlanAllowanceUsed, http.StatusPaymentRequired},
	CodePaidPlanRequired:    {ErrPaidPlanRequired, http.StatusPaymentRequired},
	CodeFreePlanCap:         {ErrFreePlanCap, http.StatusTooManyRequests},
	CodeModelCap:            {ErrModelCap, http.StatusPaymentRequired},
	CodeUsageCapExceeded:    {ErrUsageCapExceeded, http.StatusTooManyRequests},
	CodeInsufficientBalance: {ErrInsufficientBalance, http.StatusPaymentRequired},
}

// UsageLimitError is a request a plan or a wallet refused. errors.As reaches it
// through the *GenericOpenAPIError every generated method returns.
//
// Chat from an API key is refused rather than answered by a fallback model
// unless the client asks for the fallback:
//
//	client.GetConfig().AddDefaultHeader("X-Hanzo-Fallback", "allow")
type UsageLimitError struct {
	Status     int           // 402 or 429
	Code       string        // one of the Code constants
	Type       string        // billing_error or rate_limit_error
	Message    string        // the sentence to show a person
	Class      string        // premium, ours or free
	Model      string        // the model refused
	Fallback   string        // the model that answers instead (model_cap)
	ResetsAt   string        // RFC 3339, when it lifts by itself; "" when it does not
	UpgradeURL string        // the plan that raises it
	Actions    []UsageAction // the ways on
}

// UsageAction is one way past a refusal.
type UsageAction struct {
	Kind  string `json:"kind"` // upgrade, switch, credits or topup
	Label string `json:"label"`
	// URL is the page for upgrade and topup, and /v1/ai/limits for credits:
	// AiSetLimits with creditsAfterAllowance turns them on.
	URL   string `json:"url,omitempty"`
	Plan  string `json:"plan,omitempty"`  // upgrade
	Model string `json:"model,omitempty"` // switch
}

func (e *UsageLimitError) Error() string {
	return "hanzoai: " + strconv.Itoa(e.Status) + " " + e.Code + ": " + e.Message
}

// Unwrap answers the code's sentinel, which is what errors.Is compares.
func (e *UsageLimitError) Unwrap() error { return limits[e.Code].err }

// Unwrap answers the usage refusal the body carries, or nil, so errors.Is and
// errors.As see a UsageLimitError through any generated method's error.
func (e GenericOpenAPIError) Unwrap() error {
	if u := usageLimit(e.error, e.body); u != nil {
		return u
	}
	return nil
}

// usageLimit reads a refusal body. line is the generated error's text, which
// begins with the status line unless decoding the body into the operation's own
// error schema failed; then the code's status stands in for it.
func usageLimit(line string, body []byte) *UsageLimitError {
	var b struct {
		Error *struct {
			Message    string        `json:"message"`
			Type       string        `json:"type"`
			Code       string        `json:"code"`
			Class      string        `json:"class"`
			Model      string        `json:"model"`
			Fallback   string        `json:"fallback"`
			ResetsAt   string        `json:"resets_at"`
			UpgradeURL string        `json:"upgrade_url"`
			Actions    []UsageAction `json:"actions"`
		} `json:"error"`
	}
	if json.Unmarshal(body, &b) != nil || b.Error == nil {
		return nil
	}
	known, ok := limits[b.Error.Code]
	if !ok {
		return nil
	}
	status, err := strconv.Atoi(strings.SplitN(line, " ", 2)[0])
	if err != nil || status < 100 || status > 599 {
		status = known.status
	}
	return &UsageLimitError{
		Status:     status,
		Code:       b.Error.Code,
		Type:       b.Error.Type,
		Message:    b.Error.Message,
		Class:      b.Error.Class,
		Model:      b.Error.Model,
		Fallback:   b.Error.Fallback,
		ResetsAt:   b.Error.ResetsAt,
		UpgradeURL: b.Error.UpgradeURL,
		Actions:    b.Error.Actions,
	}
}

// Usage is where an answer's class stands, read off its headers. An absent
// header reads "".
type Usage struct {
	Usage    string // X-Hanzo-Usage: ok, near or limited
	Class    string // X-Hanzo-Usage-Class: premium, ours or free
	PaidBy   string // X-Hanzo-Paid-By: plan, credits or free
	Fallback string // X-Hanzo-Fallback: the model that answered instead
	Served   string // X-Hanzo-Served: the Hanzo SKU that answered
	Reason   string // X-Hanzo-Usage-Reason: the refusal code that caused the fallback
}

// ReadUsage reads the usage headers off an answer: resp.Header from any
// Execute, refused or not.
func ReadUsage(h http.Header) Usage {
	return Usage{
		Usage:    h.Get("X-Hanzo-Usage"),
		Class:    h.Get("X-Hanzo-Usage-Class"),
		PaidBy:   h.Get("X-Hanzo-Paid-By"),
		Fallback: h.Get("X-Hanzo-Fallback"),
		Served:   h.Get("X-Hanzo-Served"),
		Reason:   h.Get("X-Hanzo-Usage-Reason"),
	}
}
