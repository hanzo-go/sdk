// Package policy asks whether a subject may take an action on an object.
//
// Asking is a question with an answer, so [Client.Check] hands back a
// [Decision] carrying a boolean. Being stopped mid-call is the other half of
// the same vocabulary: that arrives as call.Denied with code "policy_denied",
// or as call.Held naming the clause a person was asked about. One clause
// vocabulary across the pre-check, the hold and the denial — otherwise a caller
// cannot tell which of its own checks it should have run, and the pre-check is
// decoration.
package policy

import (
	"context"

	"github.com/hanzoai/go-sdk/v8/call"
)

// Client asks the caller's own org policy set.
type Client struct{ e *call.Endpoint }

// New returns the policy capability over an endpoint.
func New(e *call.Endpoint) *Client { return &Client{e: e} }

// Decision is one policy answer with the question echoed beside it, so a cached
// or logged decision still says what it answered.
type Decision struct {
	Allow  bool
	Sub    string
	Act    string
	Obj    string
	Reason string
}

// Check asks whether sub may act on obj.
//
// The arguments read as the sentence does — subject, verb, object — and the
// body goes out in the members the route binds: {subject, verb, path}. That
// difference is spelled here and nowhere else.
//
// Two things about this route are worth knowing before its answers are trusted.
// It is the one capability that cannot be built over its generated method,
// because the document declares neither a requestBody nor a responses for it,
// so every generator emits a method that takes nothing and returns nothing.
// And the served predicate is STATELESS: it decides against grants carried in
// the request rather than against the org's stored set, and an empty grant set
// authorizes nothing — so until cloud reads the caller's grants from IAM, as
// the route's own description says it does, a three-argument check answers
// false for every question. The shape here is the one that survives that
// change; the answers only become true when it lands.
func (c *Client) Check(ctx context.Context, sub, act, obj string) (Decision, error) {
	in := struct {
		Subject string `json:"subject"`
		Verb    string `json:"verb"`
		Path    string `json:"path"`
	}{sub, act, obj}

	var wire struct {
		Allow   bool   `json:"allow"`
		Subject string `json:"subject"`
		Verb    string `json:"verb"`
		Path    string `json:"path"`
		Reason  string `json:"reason"`
	}
	if _, err := call.Do(ctx, c.e, "POST", "/v1/authz/check", nil, in, &wire); err != nil {
		return Decision{}, err
	}
	return Decision{
		Allow:  wire.Allow,
		Sub:    wire.Subject,
		Act:    wire.Verb,
		Obj:    wire.Path,
		Reason: wire.Reason,
	}, nil
}
