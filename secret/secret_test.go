package secret

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"
)

// where is the path the keys are read from, and what the environment resolves
// them in. They are different facts and the test says so by spelling them
// differently: a path is a place, an env is which copy of a name lives there.
const (
	where = "hanzo/cloud"
	which = "test"
)

// estate stands up one server answering both IAM's mint and KMS's reads, with a
// ServiceAccount token projected at a real path — the whole boot, watched end
// to end.
type estate struct {
	mints atomic.Int32 // assertions exchanged
	reads atomic.Int32 // secrets read
	file  string       // where the token is projected
}

// answers is what each upstream does, per call, so a test can make the second
// read differ from the first.
type answers struct {
	iam func(n int32, w http.ResponseWriter)
	kms func(n int32, w http.ResponseWriter, r *http.Request)
}

func newEstate(t *testing.T, a answers) *estate {
	t.Helper()
	e := &estate{file: filepath.Join(t.TempDir(), "token")}
	if err := os.WriteFile(e.file, []byte("assertion.jwt\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc(oauth, func(w http.ResponseWriter, r *http.Request) {
		n := e.mints.Add(1)
		if r.Method != http.MethodPost {
			t.Errorf("mint method = %s, want POST", r.Method)
		}
		if err := r.ParseForm(); err != nil {
			t.Errorf("mint carried no form: %v", err)
		}
		if got := r.PostForm.Get("grant_type"); got != grant {
			t.Errorf("grant_type = %q, want %q — the token IS the credential", got, grant)
		}
		// The assertion is the projected token, trailing newline and all removed.
		if got := r.PostForm.Get("assertion"); got != "assertion.jwt" {
			t.Errorf("assertion = %q, want the projected token", got)
		}
		w.Header().Set("Content-Type", "application/json")
		if a.iam != nil {
			a.iam(n, w)
			return
		}
		w.Write([]byte(`{"access_token":"tok","expires_in":3600}`))
	})
	mux.HandleFunc(secrets, func(w http.ResponseWriter, r *http.Request) {
		n := e.reads.Add(1)
		if got := r.Header.Get("Authorization"); got != "Bearer tok" {
			t.Errorf("authorization = %q, want the minted bearer", got)
		}
		if got := r.URL.Query().Get("env"); got != which {
			t.Errorf("env = %q, want %q", got, which)
		}
		if !strings.HasPrefix(r.URL.Path, secrets+where+"/") {
			t.Errorf("kms path = %q, want it under %q", r.URL.Path, secrets+where+"/")
		}
		w.Header().Set("Content-Type", "application/json")
		if a.kms != nil {
			a.kms(n, w, r)
			return
		}
		key := filepath.Base(r.URL.Path)
		json.NewEncoder(w).Encode(map[string]string{"env": which, "name": key, "value": "value-of-" + key})
	})

	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)

	t.Setenv(tokenVar, e.file)
	t.Setenv(iamVar, srv.URL)
	t.Setenv(kmsVar, srv.URL)
	t.Setenv(envVar, which)
	t.Setenv(devVar, "")
	return e
}

// TestBootReadsEveryKey is the ordinary path: one mint, one read per key, the
// values in memory and nowhere else.
func TestBootReadsEveryKey(t *testing.T) {
	e := newEstate(t, answers{})

	values, err := Boot(context.Background(), where, "DB_URL", "SIGNING_KEY")
	if err != nil {
		t.Fatalf("Boot: %v", err)
	}
	for _, key := range []string{"DB_URL", "SIGNING_KEY"} {
		if values[key] != "value-of-"+key {
			t.Errorf("%s = %q, want %q", key, values[key], "value-of-"+key)
		}
	}
	if got := e.mints.Load(); got != 1 {
		t.Errorf("mints = %d, want 1 — one per Boot", got)
	}
	if got := e.reads.Load(); got != 2 {
		t.Errorf("reads = %d, want 2", got)
	}
}

// TestAMissingKeyFailsTheBootByName: a map one entry short is a service that
// starts and then fails at the first request that needed the value, so a 404
// fails the whole Boot and says which key was not there.
func TestAMissingKeyFailsTheBootByName(t *testing.T) {
	newEstate(t, answers{kms: func(_ int32, w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/SIGNING_KEY") {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"env": which, "name": "DB_URL", "value": "v"})
	}})

	values, err := Boot(context.Background(), where, "DB_URL", "SIGNING_KEY")
	if err == nil {
		t.Fatalf("Boot succeeded with a missing key: %v", values)
	}
	if !strings.Contains(err.Error(), "SIGNING_KEY") {
		t.Errorf("error = %q, want it to name SIGNING_KEY", err)
	}
	if values != nil {
		t.Errorf("Boot answered %v beside the error; it must answer nothing", values)
	}
}

// TestAnExpiredBearerIsWorthOneMoreMint, and no more. A second 401 is the
// server saying no, and retrying it is how a boot loop becomes a mint flood.
func TestAnExpiredBearerIsWorthOneMoreMint(t *testing.T) {
	e := newEstate(t, answers{kms: func(_ int32, w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}})

	if _, err := Boot(context.Background(), where, "DB_URL"); err == nil {
		t.Fatal("Boot succeeded against a KMS that refuses the bearer")
	}
	if got := e.mints.Load(); got != 2 {
		t.Errorf("mints = %d, want 2 — the first, and one after the 401", got)
	}
	if got := e.reads.Load(); got != 2 {
		t.Errorf("reads = %d, want 2 — the read, and one replay", got)
	}
}

// TestTheEnvironmentIsTheDevelopmentPathOnly. With HANZO_DEV=1 and no token to
// project, Boot reads the environment; that is the only path by which a secret
// reaches the process from the environment.
func TestTheEnvironmentIsTheDevelopmentPathOnly(t *testing.T) {
	t.Setenv(tokenVar, filepath.Join(t.TempDir(), "absent"))
	t.Setenv(devVar, "1")
	t.Setenv("DB_URL", "postgres://localhost")

	values, err := Boot(context.Background(), where, "DB_URL")
	if err != nil {
		t.Fatalf("Boot: %v", err)
	}
	if values["DB_URL"] != "postgres://localhost" {
		t.Errorf("DB_URL = %q, want the environment's value", values["DB_URL"])
	}

	// The same all-or-nothing contract, so a laptop fails where the cluster does.
	if _, err := Boot(context.Background(), where, "DB_URL", "SIGNING_KEY"); err == nil {
		t.Fatal("Boot succeeded with SIGNING_KEY absent from the environment")
	} else if !strings.Contains(err.Error(), "SIGNING_KEY") {
		t.Errorf("error = %q, want it to name SIGNING_KEY", err)
	}
}

// TestNoTokenIsAnErrorInProduction: without HANZO_DEV an absent token fails the
// boot, and never falls quietly back to the environment.
func TestNoTokenIsAnErrorInProduction(t *testing.T) {
	e := newEstate(t, answers{kms: func(_ int32, _ http.ResponseWriter, _ *http.Request) {
		t.Error("kms was read with no service account token")
	}})
	absent := filepath.Join(t.TempDir(), "absent")
	t.Setenv(tokenVar, absent)
	t.Setenv("DB_URL", "the environment must not answer this")

	values, err := Boot(context.Background(), where, "DB_URL")
	if err == nil {
		t.Fatalf("Boot succeeded with no service account token: %v", values)
	}
	if !strings.Contains(err.Error(), absent) {
		t.Errorf("error = %q, want it to name the token path %q", err, absent)
	}
	if got := e.mints.Load(); got != 0 {
		t.Errorf("mints = %d, want 0 — there was nothing to authenticate with", got)
	}
}

// TestAKeyIsANameNotAPlace. The subpath belongs in the path, and a key that
// cannot hold a separator cannot address anything but a secret.
func TestAKeyIsANameNotAPlace(t *testing.T) {
	e := newEstate(t, answers{kms: func(_ int32, _ http.ResponseWriter, r *http.Request) {
		t.Errorf("kms was read at %q for a key that is not a name", r.URL.Path)
	}})

	for _, key := range []string{"", ".", "..", "a/b", `a\b`, " DB_URL", "DB_URL "} {
		if _, err := Boot(context.Background(), where, key); err == nil {
			t.Errorf("Boot accepted %q as a key", key)
		}
	}
	if got := e.mints.Load(); got != 0 {
		t.Errorf("mints = %d, want 0 — a key is refused before anything is sent", got)
	}
}
