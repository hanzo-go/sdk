package hanzoai

// The held answer: a call the platform stopped for a human decision. Hand-
// written like hanzo.go, because the hold is one shape shared by every gated
// operation and the document projects it per operation or not at all.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Approval is a sensitive call the platform stopped for a human decision. It is
// the answer's body verbatim: Status is "held", ID is the handle to poll or
// resolve by, Clause names the policy clause that held the call and Reason says
// why. Resolving a hold answers with these same field names, so there is one
// shape to learn.
type Approval struct {
	Status string `json:"status"`
	ID     string `json:"id"`
	Clause string `json:"clause"`
	Reason string `json:"reason"`
}

// HeldError is what reading a held [Result] gives you instead of a value.
type HeldError struct {
	Approval Approval
}

func (e *HeldError) Error() string {
	return fmt.Sprintf("hanzoai: awaiting approval %s: %s", e.Approval.ID, e.Approval.Clause)
}

// Result is the outcome of an approval-gated call: either a value or the
// approval that holds it.
//
// Go has no sum type, so the value is unexported and reachable only through
// [Result.Value], which returns a *HeldError when the call was held. That is
// the shape chosen over an exported field plus a Held() a caller must remember
// to consult: a struct field can be read silently, an ignored error return
// cannot — every linter in Go flags it and the value you get instead is the
// zero. The wrong path is loud.
type Result[T any] struct {
	value    T
	approval *Approval
	err      error
}

// Value returns the value, or a *HeldError naming the approval that gates it.
func (r Result[T]) Value() (T, error) {
	if r.approval != nil {
		var zero T
		return zero, &HeldError{Approval: *r.approval}
	}
	return r.value, r.err
}

// Held returns the approval and true when the call was held. It never yields
// the value, so it cannot become a second way to read one.
func (r Result[T]) Held() (Approval, bool) {
	if r.approval == nil {
		return Approval{}, false
	}
	return *r.approval, true
}

// Read reads a generated call's answer as an outcome that may be a hold. It
// takes the three values every operation returns, so it goes around the call
// itself:
//
//	agent, err := hanzoai.Read(client.AgentsAPI.PostAgents(ctx).Execute()).Value()
//	var held *hanzoai.HeldError
//	if errors.As(err, &held) {
//		fmt.Println("waiting on", held.Approval.ID)
//	}
//
// A held answer carries the approval and no value. Anything else carries what
// the operation returned, error included — a non-2xx stays the *GenericOpenAPIError
// the client already gives, with the status on the response beside it.
//
// It is a function rather than a method because a Go method cannot take a type
// parameter.
func Read[T any](value T, res *http.Response, err error) Result[T] {
	// The hold is a property of the answer, so it is read before the error: a
	// held body that does not fit the operation's own schema fails to decode,
	// and the caller still needs to learn a person was asked.
	if approval := held(res); approval != nil {
		return Result[T]{approval: approval}
	}
	if err != nil {
		return Result[T]{err: err}
	}
	return Result[T]{value: value}
}

// held returns the approval when the answer's body says the call was held.
//
// THE BODY DECIDES, NOT THE CODE. A 202 alone does not mean held: a dozen
// platform operations answer 202 for "accepted, working on it" and carry a real
// schema — a deployment, a preview, a build. Reading the code alone turns every
// one of those into an approval nobody is waiting on, and the caller then waits
// forever for an answer from someone who was never asked. Only a body that says
// held is a hold; everything else is an ordinary answer and decodes as one.
// This is the discriminator the other clients use, so they stop being one
// contract per language.
//
// The body is put back the way the generated code leaves it, so reading it here
// costs the caller nothing.
func held(res *http.Response) *Approval {
	if res == nil || res.StatusCode != http.StatusAccepted || res.Body == nil {
		return nil
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewReader(body))
	if err != nil || len(body) == 0 {
		return nil
	}
	var approval Approval
	if json.Unmarshal(body, &approval) != nil || approval.Status != "held" {
		return nil
	}
	return &approval
}
