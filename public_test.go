package hanzoai

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

// The four documented open reads reach the gateway with no credential and
// never ask IAM for a token; every other call still needs one.
func TestOpenReadsNeedNoCredential(t *testing.T) {
	t.Setenv("HANZO_CLIENT_ID", "")
	t.Setenv("HANZO_CLIENT_SECRET", "")
	var auth atomic.Value
	var hits atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits.Add(1)
		auth.Store(r.Header.Get("Authorization"))
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"object":"list","data":[]}`))
	}))
	defer srv.Close()

	c := New(Options{Base: srv.URL, Issuer: srv.URL})
	if _, _, err := c.AiAPI.GetModels(context.Background()).Execute(); err != nil {
		t.Fatalf("models with no credential: %v", err)
	}
	if got := auth.Load().(string); got != "" {
		t.Errorf("Authorization = %q, want none", got)
	}
	if err := keys(c); err == nil {
		t.Fatal("a keyed call succeeded with no credential")
	}
	if hits.Load() != 1 {
		t.Errorf("gateway hits = %d, want 1", hits.Load())
	}
}
