// Package call is the one request path budget, policy, audit, search, kb and
// graph share, and the one answer they hand back.
//
// It exists because those six live in their own packages — the generated client
// at the module root already defines Answer, Page, Hit, Backend, Wrote,
// Allowance and Charge as projections of unrelated operations, and a
// regeneration may add more — so the shared vocabulary needs a package no
// generator writes to. Everything a capability needs is here: how a URL is
// built, how the RFC 9457 problem envelope is read, and which arm an HTTP
// answer is.
//
// The rule that matters: a refusal the caller can act on is a value, never an
// exception. A budget that says no and a policy that says no come back as
// [Denied]; a call a person was asked about comes back as [Held]. Only an
// outcome with no decision in it — an absent principal, a fault, a dead
// socket — is a [Fault].
package call

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Endpoint is where calls go and what carries them: api.hanzo.ai, and an
// http.Client whose transport presents the IAM access token. The generated
// client and the six capabilities share one, so a credential is presented in
// exactly one place.
type Endpoint struct {
	Base string
	HTTP *http.Client
}

// Money is an amount in integer minor units. Money is never a float: a cent
// that rounds is money that disappears.
type Money struct {
	Cents    int64  `json:"cents"`
	Currency string `json:"currency"`
}

// Page is one page of a listing. Total is how many rows the filter matched
// across every page, which is what a pager needs to size itself.
type Page[T any] struct {
	Items []T   `json:"items"`
	Total int64 `json:"total"`
}

// Cure is one way out of a refusal: Kind names what clears it ("subscribe",
// "credit") and URL is where to do it. They arrive in the order to offer them.
type Cure struct {
	Kind string `json:"kind"`
	URL  string `json:"url"`
}

// Denied is a call that was authenticated and refused. Code is the machine
// reason a caller branches on — insufficient_balance and spend_cap_exceeded are
// different facts with different cures and are never collapsed — Reason is the
// sentence the platform wrote, Product names the gated product where the gate
// is product-scoped, and Cures is how to clear it.
//
// Code is a string and not an enum: cloud decides the vocabulary, so a code
// added tomorrow arrives as data rather than as an unparseable answer.
type Denied struct {
	Code    string
	Reason  string
	Product string
	Cures   []Cure
}

func (d *Denied) Error() string {
	scope := ""
	if d.Product != "" {
		scope = " for " + d.Product
	}
	return "hanzoai: denied" + scope + " (" + d.Code + "): " + d.Reason
}

// Held is a call the platform stopped for a human decision. ID is the handle
// the decision is made under, and Clause names the policy clause that stopped
// it — the same clause policy.Check would have refused on.
type Held struct {
	ID     string
	Clause string
	Reason string
}

func (h *Held) Error() string {
	return "hanzoai: held " + h.ID + " on " + h.Clause + ": " + h.Reason
}

// Fault is an answer with no decision in it: an absent principal, an
// unrecognised 4xx, a 5xx, a body that would not decode, a dead socket.
// Nothing in it is for a caller to act on except Request, which is what
// support finds the request by.
type Fault struct {
	Status  int
	Code    string
	Reason  string
	Request string
}

func (f *Fault) Error() string {
	at := strconv.Itoa(f.Status)
	if f.Code != "" {
		at += " " + f.Code
	}
	if f.Request != "" {
		at += " (request " + f.Request + ")"
	}
	return "hanzoai: " + at + ": " + f.Reason
}

// Answer is what a call that can be refused gives back: a value, or the reason
// there is none.
//
// Go has no sum type, so the value is unexported and reachable only through
// [Answer.Value], which hands back the arm as an error. That is the shape
// chosen over an exported field beside a Refused() a caller must remember to
// consult: a struct field can be read silently, an ignored error return cannot —
// every linter flags it and the value you get instead is the zero.
type Answer[T any] struct {
	// Request is the x-request-id the gateway stamped on the response, present
	// on every arm. It is the value audit.Filter.Request takes, and it is what
	// joins a call to the row the server wrote about it.
	Request string

	value T
	err   error
}

// Value returns what the call produced, or what stopped it: a [*Denied] when a
// gate refused, a [*Held] when a person was asked, a [*Fault] when nothing
// decided anything.
//
//	wrote, err := answer.Value()
//	var denied *call.Denied
//	if errors.As(err, &denied) {
//		fmt.Println(denied.Code, denied.Cures)
//	}
func (a Answer[T]) Value() (T, error) { return a.value, a.err }

// Failed is an answer carrying an error instead of a value, for a capability
// that finds the answer inconsistent after it arrives.
func Failed[T any](request string, err error) Answer[T] {
	return Answer[T]{Request: request, err: err}
}

// Do performs a read no gate refuses and decodes the body into out, which may
// be nil where there is nothing to read. It answers the x-request-id either
// way, so a failed read is still traceable.
func Do(ctx context.Context, e *Endpoint, method, path string, query url.Values, body, out any) (string, error) {
	status, raw, request, err := send(ctx, e, method, path, query, body)
	if err != nil {
		return request, err
	}
	if status >= 300 {
		return request, fault(status, raw, request)
	}
	if out == nil || len(raw) == 0 {
		return request, nil
	}
	if err := json.Unmarshal(raw, out); err != nil {
		return request, fmt.Errorf("hanzoai: decode %s %s: %w", method, path, err)
	}
	return request, nil
}

// Ask performs a call a gate may refuse and maps the answer to an arm. One rule
// decides the arm, the same for every capability:
//
//	2xx, body is not a hold                 ok
//	202, body {"status":"held"}             held
//	402, any code                           denied
//	403 with a refusal code                 denied
//	401, other 4xx, 5xx, transport          fault
func Ask[T any](ctx context.Context, e *Endpoint, method, path string, query url.Values, body any) Answer[T] {
	status, raw, request, err := send(ctx, e, method, path, query, body)
	if err != nil {
		return Failed[T](request, err)
	}
	if held := hold(status, raw); held != nil {
		return Failed[T](request, held)
	}
	if status < 300 {
		var value T
		if len(raw) > 0 {
			if err := json.Unmarshal(raw, &value); err != nil {
				return Failed[T](request, fmt.Errorf("hanzoai: decode %s %s: %w", method, path, err))
			}
		}
		return Answer[T]{Request: request, value: value}
	}
	if denied := denial(status, raw); denied != nil {
		return Failed[T](request, denied)
	}
	return Failed[T](request, fault(status, raw, request))
}

// send performs one request and hands back everything an arm is decided from.
//
// A body that is already an io.Reader goes out as bytes — the knowledge import
// takes an archive, not JSON — and anything else is marshalled.
func send(ctx context.Context, e *Endpoint, method, path string, query url.Values, body any) (int, []byte, string, error) {
	payload, stream := body.(io.Reader)
	if body != nil && !stream {
		encoded, err := json.Marshal(body)
		if err != nil {
			return 0, nil, "", fmt.Errorf("hanzoai: encode %s %s: %w", method, path, err)
		}
		payload = bytes.NewReader(encoded)
	}

	address := strings.TrimRight(e.Base, "/") + path
	if len(query) > 0 {
		address += "?" + query.Encode()
	}
	req, err := http.NewRequestWithContext(ctx, method, address, payload)
	if err != nil {
		return 0, nil, "", fmt.Errorf("hanzoai: %s %s: %w", method, path, err)
	}
	req.Header.Set("Accept", "application/json")
	switch {
	case stream:
		req.Header.Set("Content-Type", "application/octet-stream")
	case body != nil:
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := e.HTTP.Do(req)
	if err != nil {
		return 0, nil, "", fmt.Errorf("hanzoai: %s %s: %w", method, path, err)
	}
	defer res.Body.Close()

	request := res.Header.Get("X-Request-Id")
	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, nil, request, fmt.Errorf("hanzoai: read %s %s: %w", method, path, err)
	}
	return res.StatusCode, raw, request, nil
}

// hold reads the held arm.
//
// THE BODY DECIDES, NOT THE CODE. A 202 alone does not mean held: a dozen
// operations answer 202 for "accepted, working on it" and carry a real schema —
// a deployment, a preview, a build. Reading the code alone turns every one of
// those into an approval nobody is waiting on.
func hold(status int, raw []byte) *Held {
	if status != http.StatusAccepted {
		return nil
	}
	var body struct {
		Status string `json:"status"`
		ID     string `json:"id"`
		Clause string `json:"clause"`
		Reason string `json:"reason"`
	}
	if json.Unmarshal(raw, &body) != nil || body.Status != "held" {
		return nil
	}
	return &Held{ID: body.ID, Clause: body.Clause, Reason: body.Reason}
}

// refusals are the 403 codes that carry a decision. Cloud spells "no validated
// principal" as 403 forbidden, so a bare forbidden is the absence of a caller
// and not a refusal of one — reading it as denied would tell an unauthenticated
// caller their budget said no. The day cloud answers 401 for that, this set
// goes and the rule collapses to 402 or 403 means denied.
var refusals = map[string]bool{
	"policy_denied":        true,
	"entitlement_required": true,
	"spend_cap_exceeded":   true,
	"insufficient_balance": true,
}

// denial reads the denied arm out of either 402 body cloud writes: the RFC 9457
// envelope errmap renders, whose `code` is the reason, and cloud.Refuse's
// {error, product, reason, message, cure} shape, whose `error` is.
func denial(status int, raw []byte) *Denied {
	var body struct {
		Code    string `json:"code"`
		Detail  string `json:"detail"`
		Error   string `json:"error"`
		Product string `json:"product"`
		Message string `json:"message"`
		Cure    []Cure `json:"cure"`
	}
	// A body that will not parse leaves every field empty, which the status
	// below still classifies correctly.
	_ = json.Unmarshal(raw, &body)

	code, reason := body.Code, body.Detail
	if code == "" {
		code, reason = body.Error, body.Message
	}
	if status != http.StatusPaymentRequired && !(status == http.StatusForbidden && refusals[code]) {
		return nil
	}
	if code == "" {
		code = "payment_required"
	}
	return &Denied{Code: code, Reason: reason, Product: body.Product, Cures: body.Cure}
}

// faultText bounds what a body contributes to an error message. An HTML error
// page from something in front of the gateway is not a sentence.
const faultText = 400

// fault reads the problem envelope off an answer nobody can act on.
func fault(status int, raw []byte, request string) *Fault {
	var body struct {
		Code   string `json:"code"`
		Detail string `json:"detail"`
		Title  string `json:"title"`
	}
	_ = json.Unmarshal(raw, &body)

	reason := body.Detail
	if reason == "" {
		reason = body.Title
	}
	if reason == "" {
		reason = strings.TrimSpace(string(raw))
	}
	if len(reason) > faultText {
		reason = reason[:faultText] + "…"
	}
	return &Fault{Status: status, Code: body.Code, Reason: reason, Request: request}
}
