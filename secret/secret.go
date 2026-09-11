// Package secret reads a service's own credentials at boot, into memory.
//
// A service needs a handful of values it cannot mint for itself: a signing key,
// a database URL, the client secret of its own IAM identity. The two usual
// homes are both worse than this one. A Kubernetes Secret is a base64 field in
// the API server that anyone holding get on the namespace reads, mounted into a
// filesystem where anything in the pod reads it again. An environment variable
// is inherited by every child process and printed by every crash dumper.
//
// So the values arrive over the network at boot instead, held in this process's
// memory and nowhere else, and the caller proves who it is with the identity
// the platform already vouches for: the ServiceAccount token projected into the
// pod. IAM exchanges that assertion for a bearer and KMS answers the reads that
// bearer is scoped to. Nothing is issued to a service to let it do this.
//
// The scope is a service's OWN material, read once, before it serves anything.
// A credential for an upstream that a request travels to is a different fact,
// resolved per call against that caller's org, and is not this.
//
//	values, err := secret.Boot(ctx, "hanzo/cloud", "DB_URL", "SIGNING_KEY")
package secret

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// The environment a service is booted with. These are the platform's own names,
// so a pod already carrying them needs nothing added, and every one has a
// default, so a deployment inside the cluster sets none of them.
const (
	// tokenVar names the file holding the projected ServiceAccount token — the
	// assertion this process authenticates with. The default is where the
	// platform projects it.
	tokenVar  = "HANZO_SA_TOKEN"
	tokenFile = "/var/run/secrets/hanzo/iam/token"

	// iamVar is the origin the assertion is exchanged at: the same host
	// HANZO_ISSUER_URL names for the client, under the name the platform
	// projects into a pod. The public issuer is the default because a service
	// outside the cluster has no other address; a deployment inside it points
	// this at the service so the exchange does not leave the cluster to come
	// back.
	iamVar  = "HANZO_IAM_URL"
	iamHost = "https://hanzo.id"

	// kmsVar is where the values live. The default is the in-cluster KMS.
	kmsVar  = "KMS_URL"
	kmsHost = "http://cloud.hanzo.svc:8000"

	// envVar selects the environment a name resolves in: one name in two
	// environments is two secrets, and reading the wrong one answers 404 rather
	// than a wrong value, which is the failure a caller can act on.
	envVar  = "HANZO_ENV"
	envProd = "prod"

	// devVar admits the one environment fallback, for a laptop with no
	// ServiceAccount token to project. See [Boot].
	devVar = "HANZO_DEV"
)

// The two routes, and the grant between them.
//
// oauth is IAM's token endpoint, the same one the client mints at. The grant is
// RFC 7523 §2.1 — a JWT presented as an authorization grant — so the
// ServiceAccount token IS the credential, and there is no second one to
// distribute, rotate or leak.
const (
	oauth   = "/v1/iam/oauth/token"
	secrets = "/v1/kms/secrets/"
	grant   = "urn:ietf:params:oauth:grant-type:jwt-bearer"
)

// exchange carries both upstreams. Boot runs before a service listens, so a
// hung IAM or KMS has to fail the boot rather than hold it open forever.
var exchange = &http.Client{Timeout: 15 * time.Second}

// Boot resolves each key beneath path and answers them by name.
//
// One mint, one read per key, and the values never touch disk. path is the
// coordinate beneath the caller's own org root — the org comes from the bearer,
// so it is neither named here nor nameable — and each key is a bare name.
//
// EVERY key must resolve. A missing one fails the whole Boot and names it: the
// alternative is a map one entry short and a service that starts, serves, and
// fails at the first request that needed the value.
//
// On a laptop there is no ServiceAccount token to project. With HANZO_DEV=1 and
// no token file, Boot reads each key from the environment under its own name
// and says so. That is the only path by which a secret reaches this process
// from the environment; in production an absent token is an error, never a
// fallback.
func Boot(ctx context.Context, path string, keys ...string) (map[string]string, error) {
	path = strings.Trim(strings.TrimSpace(path), "/")
	if path == "" {
		return nil, errors.New("hanzoai: secret.Boot takes the path the keys live under")
	}
	if len(keys) == 0 {
		return nil, errors.New("hanzoai: secret.Boot takes at least one key")
	}
	for _, key := range keys {
		if !bare(key) {
			return nil, fmt.Errorf("hanzoai: %q is not a key: a key is a bare name, and a subpath belongs in the path", key)
		}
	}

	file := or(tokenVar, tokenFile)
	assertion, err := os.ReadFile(file)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) && os.Getenv(devVar) == "1" {
			return dev(keys)
		}
		return nil, fmt.Errorf("hanzoai: read the service account token at %s: %w", file, err)
	}

	bearer, err := mint(ctx, string(assertion))
	if err != nil {
		return nil, err
	}

	values := make(map[string]string, len(keys))
	// A bearer that expired between two reads is worth exactly one more mint. A
	// second 401 is the server saying no — this identity is not admitted — and
	// retrying that is how a boot loop becomes a mint flood against IAM.
	replayed := false
	for _, key := range keys {
		value, status, err := read(ctx, bearer, path, key)
		if status == http.StatusUnauthorized && !replayed {
			replayed = true
			if bearer, err = mint(ctx, string(assertion)); err != nil {
				return nil, err
			}
			value, _, err = read(ctx, bearer, path, key)
		}
		if err != nil {
			return nil, err
		}
		values[key] = value
	}
	return values, nil
}

// dev is the development path: the keys, read from the process environment.
//
// It holds the same all-or-nothing contract as the real read, so a laptop fails
// where the cluster does — at boot, naming what is missing — rather than at the
// first request that needed the value.
func dev(keys []string) (map[string]string, error) {
	values := make(map[string]string, len(keys))
	for _, key := range keys {
		value, ok := os.LookupEnv(key)
		if !ok {
			return nil, fmt.Errorf("hanzoai: %s is not in the environment (%s=1, and no service account token at %s)",
				key, devVar, or(tokenVar, tokenFile))
		}
		values[key] = value
	}
	slog.Warn("hanzoai: secrets read from the environment: no service account token and "+devVar+"=1", "keys", len(keys))
	return values, nil
}

// mint exchanges the ServiceAccount token for a bearer at IAM.
//
// The assertion is form-encoded like every other grant IAM serves, so this is
// the endpoint and the encoding the client_credentials exchange beside it
// already uses — one token endpoint, a different grant.
func mint(ctx context.Context, assertion string) (string, error) {
	iam := origin(iamVar, iamHost)
	form := url.Values{"grant_type": {grant}, "assertion": {strings.TrimSpace(assertion)}}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, iam+oauth, strings.NewReader(form.Encode()))
	if err != nil {
		return "", fmt.Errorf("hanzoai: %s: %w", iam, err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	res, err := exchange.Do(req)
	if err != nil {
		return "", fmt.Errorf("hanzoai: %s: %w", iam, err)
	}
	defer res.Body.Close()

	var answer struct {
		AccessToken string `json:"access_token"`
		Error       string `json:"error"`
		Description string `json:"error_description"`
	}
	if json.NewDecoder(res.Body).Decode(&answer) != nil {
		return "", fmt.Errorf("hanzoai: %s answered %d with no readable body", iam, res.StatusCode)
	}
	// IAM may answer 200 carrying {"error":…}, so the token decides and not the
	// status: no token is a refusal whatever the code said.
	if res.StatusCode != http.StatusOK || answer.AccessToken == "" {
		return "", fmt.Errorf("hanzoai: %s refused the service account token: %d %s %s",
			iam, res.StatusCode, answer.Error, answer.Description)
	}
	return answer.AccessToken, nil
}

// read GETs one value.
//
// It answers the status beside the error so [Boot] can tell a bearer that
// expired mid-boot (401, worth one more mint) from a secret that is not there
// (404, worth nothing but naming it).
func read(ctx context.Context, bearer, path, key string) (string, int, error) {
	env := or(envVar, envProd)
	address := origin(kmsVar, kmsHost) + secrets + escape(path) + "/" + url.PathEscape(key) +
		"?env=" + url.QueryEscape(env)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, address, nil)
	if err != nil {
		return "", 0, fmt.Errorf("hanzoai: read %s/%s: %w", path, key, err)
	}
	req.Header.Set("Authorization", "Bearer "+bearer)
	req.Header.Set("Accept", "application/json")

	res, err := exchange.Do(req)
	if err != nil {
		return "", 0, fmt.Errorf("hanzoai: read %s/%s: %w", path, key, err)
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusOK:
	case http.StatusNotFound:
		return "", res.StatusCode, fmt.Errorf("hanzoai: %s/%s is not in kms under env %s", path, key, env)
	default:
		return "", res.StatusCode, fmt.Errorf("hanzoai: read %s/%s: %d", path, key, res.StatusCode)
	}

	// KMS answers {env, name, value}; the other two members restate what was
	// asked for.
	var answer struct {
		Value string `json:"value"`
	}
	if json.NewDecoder(res.Body).Decode(&answer) != nil {
		return "", res.StatusCode, fmt.Errorf("hanzoai: kms answered %s/%s with no readable body", path, key)
	}
	return answer.Value, res.StatusCode, nil
}

// bare reports whether key addresses one secret rather than a place.
//
// A key carries no separator and is neither of the two names that mean a
// directory. That is what the address wants anyway — the subpath is the path's
// half — and it is what keeps a key from reaching anything but a secret, here
// or in a caller that writes the values out as one file per name.
func bare(key string) bool {
	switch key {
	case "", ".", "..":
		return false
	}
	return key == strings.TrimSpace(key) && !strings.ContainsAny(key, `/\`)
}

// or answers what the environment holds for key, or def where it holds nothing.
func or(key, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return def
}

// origin answers a base URL from the environment without its trailing slash, so
// a route composes onto it by concatenation and never doubles the separator.
func origin(key, def string) string { return strings.TrimRight(or(key, def), "/") }

// escape percent-encodes a /-separated path one segment at a time, so the
// separators survive as separators and everything else is data.
func escape(path string) string {
	parts := strings.Split(path, "/")
	for i, part := range parts {
		parts[i] = url.PathEscape(part)
	}
	return strings.Join(parts, "/")
}
