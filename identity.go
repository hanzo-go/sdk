package hanzoai

// The credential. IAM mints; this client never accepts a bearer.
//
// Two mints, one shape. [identity] exchanges the client's own credentials for
// an access token (HIP-0111, client_credentials, scoped by RFC 8707 resource) —
// this is hanzoai/visor's egress identity, and it is the whole of auth. [grant]
// exchanges that token for one bound to a tenant subject, which is what
// [Client.As] rides on.
//
// Both are http.RoundTrippers on the shared http.Client, so they sign the
// generated surface and the six capabilities alike and there is one place a
// credential is presented.

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// DefaultIssuerURL is where IAM answers. IAM mints on its own host; the gateway
// at DefaultBaseURL does not answer the mint. HANZO_ISSUER_URL moves it for a
// private estate, the way HANZO_BASE_URL moves the gateway.
const DefaultIssuerURL = "https://hanzo.id"

// oauth is IAM's token endpoint, and act is its act-grant mint. The act grant
// reads the subject off the query and the operator credential off the request,
// which is why the two mints cannot share a call.
const (
	oauth = "/v1/iam/oauth/token"
	act   = "/v1/iam/tokens/issue"
)

// early is how long before expiry a held token stops being offered. A token
// that expires in flight is a 401 the caller cannot tell from a revoked
// identity, so it is replaced while it still works.
const early = 60 * time.Second

// life is the lifetime assumed when IAM states none.
const life = 5 * time.Minute

// held is a token and when it stops being offered. Both mints cache the same
// way, so they cache in the same place.
type held struct {
	mu     sync.Mutex
	token  string
	expiry time.Time
}

// fresh answers the held token, or what mint gives when the cache is empty,
// near expiry, or force says the held one just failed.
func (h *held) fresh(force bool, mint func() (string, time.Duration, error)) (string, error) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if !force && h.token != "" && time.Now().Before(h.expiry.Add(-early)) {
		return h.token, nil
	}
	token, ttl, err := mint()
	if err != nil {
		return "", err
	}
	if token == "" {
		return "", errors.New("hanzoai: IAM answered with no token")
	}
	if ttl <= 0 {
		ttl = life
	}
	h.token, h.expiry = token, time.Now().Add(ttl)
	return token, nil
}

// identity mints the access token this client presents.
//
// It is the client's OWN IAM identity — the clientId and clientSecret it signs
// in with — exchanged for an access token, not a second credential minted for
// the purpose. The gateway verifies it the way every service verifies a caller:
// iss against the issuer, aud against the resource, signature against the
// published JWKS. So there is one authority, one kind of token, and nothing to
// paste into a config file.
type identity struct {
	issuer   string
	id       string
	secret   string
	resource string
	next     http.RoundTripper

	held
}

// RoundTrip presents a live token on every request.
//
// A client holding no credentials presents nothing rather than failing: the
// operations that take none still answer, and a refusal from the rest is the
// truth about a caller that cannot say who it is.
func (i *identity) RoundTrip(req *http.Request) (*http.Response, error) {
	if i.id == "" || i.secret == "" {
		return i.next.RoundTrip(req)
	}
	return signed(req, i.next, i.token)
}

func (i *identity) token(ctx context.Context, force bool) (string, error) {
	return i.fresh(force, func() (string, time.Duration, error) { return i.mint(ctx) })
}

// mint performs the client_credentials exchange, scoped by RFC 8707 resource so
// the token names what it may spend at and is useless anywhere else.
func (i *identity) mint(ctx context.Context) (string, time.Duration, error) {
	form := url.Values{"grant_type": {"client_credentials"}}
	if i.resource != "" {
		form.Set("resource", i.resource)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		strings.TrimRight(i.issuer, "/")+oauth, strings.NewReader(form.Encode()))
	if err != nil {
		return "", 0, err
	}
	req.SetBasicAuth(i.id, i.secret) // client_secret_basic
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	var out struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
		Error       string `json:"error"`
		Description string `json:"error_description"`
	}
	status, err := answered(i.next, req, &out)
	if err != nil {
		return "", 0, err
	}
	if status != http.StatusOK || out.AccessToken == "" {
		// Say which identity was refused. A 401 here reads identically whether
		// the client id is wrong, the secret is stale, or the app may not use
		// this grant, and the reader is holding none of those.
		return "", 0, fmt.Errorf("hanzoai: %s refused client %q: %d %s %s",
			i.issuer, i.id, status, out.Error, out.Description)
	}
	return out.AccessToken, time.Duration(out.ExpiresIn) * time.Second, nil
}

// grant mints the token that lets an operator credential act as one subject,
// and presents it on every request the scoped client makes. It is unexported: a
// caller reaches it only through [Client.As], which is the single way to scope.
type grant struct {
	subject string
	// operator is the credential IAM reads the act grant off.
	operator *identity
	issuer   string
	// next is the bare transport. The subject-bound token is the whole identity
	// of a scoped request, so the operator's own token never rides beside it.
	next http.RoundTripper

	held
}

func (g *grant) RoundTrip(req *http.Request) (*http.Response, error) {
	return signed(req, g.next, g.token)
}

func (g *grant) token(ctx context.Context, force bool) (string, error) {
	return g.fresh(force, func() (string, time.Duration, error) { return g.issue(ctx) })
}

// issue asks IAM for a token bound to the subject.
//
// The subject rides as the id query, which is the target the document names,
// and the operator's own access token is what IAM reads the act grant off.
func (g *grant) issue(ctx context.Context) (string, time.Duration, error) {
	operator, err := g.operator.token(ctx, false)
	if err != nil {
		return "", 0, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		strings.TrimRight(g.issuer, "/")+act+"?"+url.Values{"id": {g.subject}}.Encode(), nil)
	if err != nil {
		return "", 0, err
	}
	req.Header.Set("Authorization", "Bearer "+operator)
	req.Header.Set("Accept", "application/json")

	// IAM spells the act grant's answer in camelCase, where the OAuth mint
	// spells its own in snake_case. Two endpoints, two conventions, both read
	// here so nothing downstream has to know.
	var out struct {
		AccessToken string `json:"accessToken"`
		ExpiresIn   int64  `json:"expiresIn"`
	}
	status, err := answered(g.next, req, &out)
	if err != nil {
		return "", 0, err
	}
	if status != http.StatusOK {
		return "", 0, fmt.Errorf("hanzoai: %s refused the act grant for %q: %d", g.issuer, g.subject, status)
	}
	return out.AccessToken, time.Duration(out.ExpiresIn) * time.Second, nil
}

// answered sends a mint and decodes what came back. A body that will not decode
// is reported as the status it arrived under, because that is what the reader
// can act on.
func answered(next http.RoundTripper, req *http.Request, out any) (int, error) {
	res, err := next.RoundTrip(req)
	if err != nil {
		return 0, fmt.Errorf("hanzoai: %s: %w", req.URL.Host, err)
	}
	defer res.Body.Close()
	if json.NewDecoder(res.Body).Decode(out) != nil {
		return res.StatusCode, fmt.Errorf("hanzoai: %s answered %d with no readable body", req.URL.Host, res.StatusCode)
	}
	return res.StatusCode, nil
}

// signed presents a live token and, on a 401, mints again and sends once more.
// A second 401 is the server saying no, not a stale token.
func signed(req *http.Request, next http.RoundTripper, token func(context.Context, bool) (string, error)) (*http.Response, error) {
	bearer, err := token(req.Context(), false)
	if err != nil {
		if req.Body != nil {
			req.Body.Close()
		}
		return nil, err
	}

	res, err := next.RoundTrip(bearing(req, bearer))
	if err != nil || res.StatusCode != http.StatusUnauthorized {
		return res, err
	}

	var body io.ReadCloser
	if req.Body != nil {
		if req.GetBody == nil {
			return res, nil // the body is spent and unrecoverable
		}
		if body, err = req.GetBody(); err != nil {
			return res, nil
		}
	}
	if bearer, err = token(req.Context(), true); err != nil {
		return res, nil // hand back the 401 the server actually sent
	}
	res.Body.Close()

	replay := bearing(req, bearer)
	if body != nil {
		replay.Body = body
	}
	return next.RoundTrip(replay)
}

// bearing copies req with the token on it. A RoundTripper is handed a request
// it does not own and must not write to.
func bearing(req *http.Request, token string) *http.Request {
	out := req.Clone(req.Context())
	out.Header.Set("Authorization", "Bearer "+token)
	return out
}
