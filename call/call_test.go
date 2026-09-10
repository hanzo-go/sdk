package call

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// One rule decides which arm an HTTP answer is, and it is the same rule for
// every capability. This is that rule, measured.

func endpoint(t *testing.T, status int, answer string) *Endpoint {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Request-Id", "req-1")
		w.WriteHeader(status)
		w.Write([]byte(answer))
	}))
	t.Cleanup(srv.Close)
	return &Endpoint{Base: srv.URL, HTTP: srv.Client()}
}

func TestArms(t *testing.T) {
	ctx := context.Background()

	for _, tc := range []struct {
		name   string
		status int
		answer string
		check  func(*testing.T, Answer[map[string]any])
	}{{
		name:   "2xx is the value",
		status: 200,
		answer: `{"ok":true}`,
		check: func(t *testing.T, a Answer[map[string]any]) {
			value, err := a.Value()
			if err != nil {
				t.Fatalf("err = %v, want none", err)
			}
			if !reflect.DeepEqual(value, map[string]any{"ok": true}) {
				t.Errorf("value = %v", value)
			}
		},
	}, {
		// A dozen routes answer 202 for "accepted, working on it" and carry a
		// real schema. The BODY decides, never the code.
		name:   "202 that is not a hold is the value",
		status: 202,
		answer: `{"status":"queued","id":"dpl_1"}`,
		check: func(t *testing.T, a Answer[map[string]any]) {
			value, err := a.Value()
			if err != nil {
				t.Fatalf("a queued 202 read as a hold: %v", err)
			}
			if value["status"] != "queued" {
				t.Errorf("value = %v", value)
			}
		},
	}, {
		name:   "202 that says held is a hold",
		status: 202,
		answer: `{"status":"held","id":"apr_1","clause":"spend.large","reason":"over the cap"}`,
		check: func(t *testing.T, a Answer[map[string]any]) {
			_, err := a.Value()
			var held *Held
			if !errors.As(err, &held) {
				t.Fatalf("err = %v (%T), want *Held", err, err)
			}
			want := &Held{ID: "apr_1", Clause: "spend.large", Reason: "over the cap"}
			if !reflect.DeepEqual(held, want) {
				t.Errorf("held = %+v, want %+v", held, want)
			}
		},
	}, {
		name:   "402 is a denial whatever the code",
		status: 402,
		answer: `{"status":402,"detail":"Spend cap reached","code":"spend_cap_exceeded"}`,
		check: func(t *testing.T, a Answer[map[string]any]) {
			_, err := a.Value()
			var denied *Denied
			if !errors.As(err, &denied) {
				t.Fatalf("err = %v (%T), want *Denied", err, err)
			}
			if denied.Code != "spend_cap_exceeded" || denied.Reason != "Spend cap reached" {
				t.Errorf("denied = %+v", denied)
			}
		},
	}, {
		name:   "402 with no readable body is still a denial",
		status: 402,
		answer: `not json`,
		check: func(t *testing.T, a Answer[map[string]any]) {
			_, err := a.Value()
			var denied *Denied
			if !errors.As(err, &denied) {
				t.Fatalf("err = %v (%T), want *Denied", err, err)
			}
			if denied.Code != "payment_required" {
				t.Errorf("code = %q, want payment_required", denied.Code)
			}
		},
	}, {
		name:   "403 with a refusal code is a denial",
		status: 403,
		answer: `{"detail":"the plan does not include graph","code":"entitlement_required"}`,
		check: func(t *testing.T, a Answer[map[string]any]) {
			_, err := a.Value()
			var denied *Denied
			if !errors.As(err, &denied) {
				t.Fatalf("err = %v (%T), want *Denied", err, err)
			}
			if denied.Code != "entitlement_required" {
				t.Errorf("code = %q", denied.Code)
			}
		},
	}, {
		// Cloud spells "no validated principal" as 403 forbidden. Reading it as
		// a denial would tell an unauthenticated caller their budget said no.
		name:   "403 forbidden is a fault",
		status: 403,
		answer: `{"detail":"a validated principal is required","code":"forbidden"}`,
		check:  wantFault(403, "a validated principal is required"),
	}, {
		name:   "401 is a fault",
		status: 401,
		answer: `{"detail":"sign in to view billing","code":"unauthorized"}`,
		check:  wantFault(401, "sign in to view billing"),
	}, {
		name:   "404 is a fault",
		status: 404,
		answer: `{"detail":"Not Found","title":"Not Found","status":404}`,
		check:  wantFault(404, "Not Found"),
	}, {
		name:   "500 is a fault",
		status: 500,
		answer: `{"detail":"Something went wrong on our side."}`,
		check:  wantFault(500, "Something went wrong on our side."),
	}, {
		// A body with no problem members at all still says what happened.
		name:   "a fault with no envelope keeps the body",
		status: 502,
		answer: `upstream unreachable`,
		check:  wantFault(502, "upstream unreachable"),
	}} {
		t.Run(tc.name, func(t *testing.T) {
			answer := Ask[map[string]any](ctx, endpoint(t, tc.status, tc.answer), "GET", "/v1/thing", nil, nil)
			if answer.Request != "req-1" {
				t.Errorf("request = %q, want req-1 — every arm names the row", answer.Request)
			}
			tc.check(t, answer)
		})
	}
}

func wantFault(status int, reason string) func(*testing.T, Answer[map[string]any]) {
	return func(t *testing.T, a Answer[map[string]any]) {
		t.Helper()
		_, err := a.Value()
		var fault *Fault
		if !errors.As(err, &fault) {
			t.Fatalf("err = %v (%T), want *Fault", err, err)
		}
		if fault.Status != status || fault.Reason != reason {
			t.Errorf("fault = %+v, want %d %q", fault, status, reason)
		}
		if fault.Request != "req-1" {
			t.Errorf("fault request = %q, want req-1", fault.Request)
		}
	}
}

// A read no gate refuses answers the request id whether it decoded or not.
func TestDoAnswersTheRequestID(t *testing.T) {
	var out map[string]any
	request, err := Do(context.Background(), endpoint(t, 200, `{"a":1}`), "GET", "/v1/thing", nil, nil, &out)
	if err != nil || request != "req-1" {
		t.Fatalf("Do = %q, %v", request, err)
	}

	request, err = Do(context.Background(), endpoint(t, 500, `{"detail":"boom"}`), "GET", "/v1/thing", nil, nil, &out)
	if request != "req-1" {
		t.Errorf("request = %q, want req-1 even on a failure", request)
	}
	var fault *Fault
	if !errors.As(err, &fault) {
		t.Fatalf("err = %v (%T), want *Fault", err, err)
	}
}

// A transport that never answered has no request id and no decision in it.
func TestUnsentCallIsAnError(t *testing.T) {
	e := &Endpoint{Base: "http://127.0.0.1:1", HTTP: http.DefaultClient}
	answer := Ask[map[string]any](context.Background(), e, "GET", "/v1/thing", nil, nil)
	if _, err := answer.Value(); err == nil {
		t.Fatal("want an error when the request never went out")
	}
	if answer.Request != "" {
		t.Errorf("request = %q, want empty", answer.Request)
	}
}
