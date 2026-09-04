package hanzoai

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"
)

// A dozen platform operations answer 202 for "accepted, working on it" and
// carry a real schema — a deployment, a preview, a build. Discriminating on the
// code alone turns every one of them into an approval nobody is waiting on, and
// the caller then waits forever for a person who was never asked. The body
// decides.
func TestReadAnAcceptedAnswerIsNotAHold(t *testing.T) {
	e := newEstate(t,
		func(n int32, w http.ResponseWriter) { w.Write([]byte(`{"accessToken":"t","expiresIn":3600}`)) },
		func(n int32, w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte(`{"id":"agent_1","status":"queued"}`))
		})

	res := Read(create(e.scoped("usr_7")).Execute())
	if _, held := res.Held(); held {
		t.Fatal("an accepted answer was read as a held approval — the caller now waits on a person who was never asked")
	}
	agent, err := res.Value()
	if err != nil {
		t.Fatalf("Value on an ordinary 202 = %v, want the decoded body", err)
	}
	if agent.GetId() != "agent_1" {
		t.Errorf("value = %+v, want agent_1", agent)
	}
}

func TestReadHeldCannotBeReadAsSuccess(t *testing.T) {
	e := newEstate(t,
		func(n int32, w http.ResponseWriter) { w.Write([]byte(`{"accessToken":"t","expiresIn":3600}`)) },
		func(n int32, w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte(`{"status":"held","id":"apr_9","clause":"agents.create","reason":"unbounded schedule"}`))
		})

	value, answer, err := create(e.scoped("usr_7")).Execute()
	// 202 is a 2xx, so the call itself succeeded. The Result is what refuses.
	if err != nil {
		t.Fatalf("PostAgents: %v", err)
	}
	res := Read(value, answer, err)

	approval, held := res.Held()
	if !held {
		t.Fatal("Held() = false, want true on a body that says held")
	}
	if approval.Status != "held" || approval.ID != "apr_9" ||
		approval.Clause != "agents.create" || approval.Reason != "unbounded schedule" {
		t.Errorf("approval = %+v, want the body verbatim", approval)
	}

	agent, err := res.Value()
	var heldErr *HeldError
	if !errors.As(err, &heldErr) {
		t.Fatalf("Value error = %v (%T), want *HeldError", err, err)
	}
	if heldErr.Approval.ID != "apr_9" {
		t.Errorf("HeldError.Approval = %+v, want apr_9", heldErr.Approval)
	}
	if agent != nil {
		t.Errorf("value = %+v, want nothing — a held call yields no value", agent)
	}

	// Reading the body to decide put it back, so the answer is still whole for
	// a caller who kept it.
	body, _ := io.ReadAll(answer.Body)
	if len(body) == 0 {
		t.Error("the answer's body was spent deciding whether it was a hold")
	}
}

func TestReadDoneYieldsTheValue(t *testing.T) {
	e := newEstate(t,
		func(n int32, w http.ResponseWriter) { w.Write([]byte(`{"accessToken":"t","expiresIn":3600}`)) },
		func(n int32, w http.ResponseWriter, r *http.Request) { w.Write([]byte(`{"id":"agent_3"}`)) })

	res := Read(create(e.scoped("usr_7")).Execute())
	if _, held := res.Held(); held {
		t.Fatal("Held() = true on a 200")
	}
	agent, err := res.Value()
	if err != nil {
		t.Fatalf("Value: %v", err)
	}
	if agent.GetId() != "agent_3" {
		t.Errorf("value = %+v, want agent_3", agent)
	}
}

func TestReadNonSuccessIsTypedError(t *testing.T) {
	e := newEstate(t,
		func(n int32, w http.ResponseWriter) { w.Write([]byte(`{"accessToken":"t","expiresIn":3600}`)) },
		func(n int32, w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"schedule must be a 5-field cron"}`))
		})

	value, answer, err := create(e.scoped("usr_7")).Execute()
	_, err = Read(value, answer, err).Value()

	var apiErr *GenericOpenAPIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("err = %v (%T), want *GenericOpenAPIError", err, err)
	}
	if answer.StatusCode != http.StatusBadRequest {
		t.Errorf("StatusCode = %d, want 400", answer.StatusCode)
	}
	if got, want := string(apiErr.Body()), `{"error":"schedule must be a 5-field cron"}`; got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
}

// The held answer and the answer that resolves it are one shape, so the same
// struct decodes both.
func TestApprovalIsOneShape(t *testing.T) {
	var approval Approval
	if err := json.Unmarshal([]byte(
		`{"status":"approved","id":"apr_9","clause":"agents.create","reason":"unbounded schedule"}`,
	), &approval); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if approval.Status != "approved" || approval.ID != "apr_9" || approval.Clause != "agents.create" {
		t.Errorf("approval = %+v", approval)
	}
}
