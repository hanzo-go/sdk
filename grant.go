package hanzoai

// Act-grant scoping: an operator credential acting as one subject. Hand-written
// like hanzo.go, and for the same reason — the document declares the mint as an
// address and not as a shape, so the generated method for it carries neither
// the target nor the token it answers with.

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"maps"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

// DefaultIssuerURL is where IAM answers. IAM mints on its own host; the gateway
// at DefaultBaseURL does not answer the mint. HANZO_ISSUER_URL moves it for a
// private estate, the way HANZO_BASE_URL moves the gateway.
const DefaultIssuerURL = "https://hanzo.id"

// mint is IAM's token endpoint. The target rides as the id query and never as a
// body: IAM reads the act grant off the credential the request already carries.
const mint = "/v1/iam/tokens/issue"

// skew re-mints a little before expiry, so a request never rides an
// about-to-die token.
const skew = 30 * time.Second

// life is the lifetime assumed when IAM states none.
const life = 5 * time.Minute

// As scopes every call to the tenant subject. The returned client mints a
// subject-bound token from IAM's act grant — the operator credential this
// client already carries names the subject — caches it to expiry, and re-mints
// on a 401. No method then takes a user id: the credential is the scope, so a
// caller cannot pass the wrong one and there is nothing to forget.
//
//	scoped := hanzoai.NewClient("").As("usr_7")
//	keys, _, err := scoped.AccountAPI.GetAccountKeys(ctx).Execute()
//
// subject is a subject id or the externalId the operator filed the member under.
//
// The scoping rides as one http.RoundTripper on the Configuration's HTTPClient,
// so it is the same client with the same methods — there is no second surface
// to learn and nothing to remember to call.
func (c *APIClient) As(subject string) *APIClient {
	cfg := *c.cfg
	cfg.DefaultHeader = maps.Clone(c.cfg.DefaultHeader)

	// The minted token is the whole identity of the scoped client, so the
	// operator credential leaves with it. Two credentials on one request is two
	// answers to who is calling.
	operator := cfg.DefaultHeader["Authorization"]
	delete(cfg.DefaultHeader, "Authorization")

	inner := c.cfg.HTTPClient
	if inner == nil {
		inner = http.DefaultClient
	}
	next := inner.Transport
	if next == nil {
		next = http.DefaultTransport
	}

	scoped := *inner
	scoped.Transport = &grant{
		subject:  subject,
		operator: operator,
		issuer:   issuer(),
		next:     next,
	}
	cfg.HTTPClient = &scoped
	return NewAPIClient(&cfg)
}

// issuer is the host that mints, DefaultIssuerURL unless HANZO_ISSUER_URL says
// otherwise.
func issuer() string {
	if u := os.Getenv("HANZO_ISSUER_URL"); u != "" {
		return u
	}
	return DefaultIssuerURL
}

// grant mints and caches the token that lets an operator credential act as one
// subject, and presents it on every request the scoped client makes. It is not
// exported: a caller reaches it only through [APIClient.As], which is the
// single way to scope a client.
type grant struct {
	subject  string
	operator string // the operator credential, already spelled as IAM reads it
	issuer   string
	next     http.RoundTripper

	mu     sync.Mutex
	token  string
	expiry time.Time
}

// issued is IAM's mint response. It is camelCase on the wire.
type issued struct {
	AccessToken string `json:"accessToken"`
	// Seconds from now.
	ExpiresIn int64 `json:"expiresIn"`
}

// RoundTrip signs the request with a live subject-bound token and replays it
// once on a 401. A refused mint comes straight back as the error of the call,
// so a caller learns the grant was denied rather than watching every request
// fail unexplained.
func (g *grant) RoundTrip(req *http.Request) (*http.Response, error) {
	token, err := g.bearer(req.Context(), false)
	if err != nil {
		if req.Body != nil {
			req.Body.Close()
		}
		return nil, err
	}

	res, err := g.next.RoundTrip(bearing(req, token))
	if err != nil || res.StatusCode != http.StatusUnauthorized {
		return res, err
	}

	// One replay only: a second 401 is the server saying no, not a stale token.
	var body io.ReadCloser
	if req.Body != nil {
		if req.GetBody == nil {
			return res, nil // body already spent and unrecoverable
		}
		if body, err = req.GetBody(); err != nil {
			return res, nil
		}
	}
	if token, err = g.bearer(req.Context(), true); err != nil {
		return res, nil // hand back the 401 the server actually sent
	}
	res.Body.Close()
	replay := bearing(req, token)
	if body != nil {
		replay.Body = body
	}
	return g.next.RoundTrip(replay)
}

// bearing copies req with the subject-bound token on it. A RoundTripper is
// handed a request it does not own and must not write to.
func bearing(req *http.Request, token string) *http.Request {
	out := req.Clone(req.Context())
	out.Header.Set("Authorization", "Bearer "+token)
	return out
}

// bearer returns a live token, minting when the cache is empty, near expiry, or
// force says the held one just failed.
func (g *grant) bearer(ctx context.Context, force bool) (string, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if !force && g.token != "" && time.Until(g.expiry) > skew {
		return g.token, nil
	}
	tok, err := g.issue(ctx)
	if err != nil {
		return "", err
	}
	if tok.AccessToken == "" {
		return "", errors.New("hanzoai: IAM issued no token for the act grant")
	}
	g.token = tok.AccessToken
	g.expiry = time.Now().Add(life)
	if tok.ExpiresIn > 0 {
		g.expiry = time.Now().Add(time.Duration(tok.ExpiresIn) * time.Second)
	}
	return g.token, nil
}

// issue asks IAM for a token bound to the subject.
//
// The request is built here rather than taken from the generated client because
// the mint is not a call to the platform: it goes to IAM, on IAM's host, and
// carries only the operator credential IAM reads the act grant off. The
// gateway's default headers stay behind — two spellings of one credential are
// two answers to who is calling — and the subject rides as the id query, which
// is the target the document names.
func (g *grant) issue(ctx context.Context) (issued, error) {
	var out issued

	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		g.issuer+mint+"?"+url.Values{"id": {g.subject}}.Encode(), nil)
	if err != nil {
		return out, err
	}
	req.Header.Set("Accept", "application/json")
	if g.operator != "" {
		req.Header.Set("Authorization", g.operator)
	}

	res, err := g.next.RoundTrip(req)
	if err != nil {
		return out, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return out, err
	}
	if res.StatusCode >= 300 {
		return out, &GenericOpenAPIError{body: body, error: res.Status}
	}
	return out, json.Unmarshal(body, &out)
}
