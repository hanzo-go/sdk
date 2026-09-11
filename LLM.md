# LLM.md — hanzo-go/sdk

Go module `github.com/hanzoai/go-sdk/v8`, published by **git tag** → proxy.golang.org.

```bash
go build ./... && go vet ./... && go test -count=1 ./... && go build ./examples/...
```

Those four are the whole gate, and `hanzo.yml` runs exactly them.

## Repo identity

Canonical repo is **`hanzo-go/sdk`**. `hanzoai/go-sdk` is a rename redirect that
GitHub honours and the forge does not: `git.hanzo.ai/hanzoai/go-sdk` answers
`Repository not found`, so a checkout with that remote fetches nothing and
quietly stays on whatever it last saw. Point `origin` at
`https://git.hanzo.ai/hanzo-go/sdk.git`.

The **module path stays `github.com/hanzoai/go-sdk/v8`** and is not renamed. A
`go.mod` path must match what consumers `require`; `github.com/hanzo-go/sdk` has
never been on the proxy and resolving it fails with a path mismatch.

## The document, and the lock that names it

The client is a projection of **`hanzoai/cloud`'s `openapi.yaml`** — the file
cloud emits from its own routers and regenerates-and-diffs on every release, so
it cannot describe a route the binary does not serve. `.spec-lock` names the ref
and sha256 this committed client came from, in the four lines every Hanzo SDK
uses:

```
ref=<commit sha, never a branch>
sha256=<of openapi.yaml at that ref>
repo=hanzoai/cloud
path=openapi.yaml
```

A branch there would read as a moving target in CI, which clones without your
remotes and cannot resolve `forge/main`. The driver re-fetches at that ref and
refuses to run if the bytes hash to anything else.

It is **not** `hanzoai/openapi`'s `hanzo.yaml`. That file is itself a projection
of this document with codegen rules applied, so reading it made this client a
projection of a projection — one release behind whenever the middle step had not
run. Every Hanzo SDK reads this same document, at the ref its own `.spec-lock`
names; five of the seven — python, typescript, java, kotlin and this one — are
pinned at this ref and share one digest. Generating
from the projection instead, measured at the locked ref, costs this client **46
methods** — 2456 against 2502 — and stamps every file's header `API version:
8.0.0`, the projection's own release counter, which names no cloud release. It
buys one thing: the projection tags the 50 operations cloud leaves on
`DefaultAPI`. That is cloud's emitter to fix, not a reason to read a second
document.

The one difference that does **not** apply here: the projection's root
`security` requires a `bearerAuth` its own `components.securitySchemes` does not
define, which costs the languages that emit a credential per operation. Go's
generator reads the defined `bearer` scheme instead, so Go's single
`Authorization` site in `client.go` survives either document.

Go is a **row in `sdks.yaml`** like every other language. It was not, for one
reason: `take:` used to promise the generator owned a *directory*, and this
client's directory is the module root beside `go.mod` and `.git`. `take:` now
names the *files* the generator wrote — see `.generated` below — so the row can
say `{.: .}` and the second driver this repo carried is gone.

## One client, at the module root

The repo root `*.go` **is** the client, `package hanzoai` — 2656 generated files
beside four hand-written ones. There is no second surface: the hand-shipped
client that predated it and the parallel `cloud/` subpackage are both gone.

**The shape, measured against the locked document.** 1814 paths carry 2479
operations, and 191 distinct tags plus the 50 untagged operations make 192
services. The client emits one method per (operation, tag) placement, and there
are 2502 of those: 23 operations are tagged both `iam` and `compat` and so
appear under `IamAPI` and `CompatAPI` both. That leaves **2477 distinct method
names, not 2477 operations** — a count of names undercounts the surface twice
over, because two more pairs collide across services on their own
(`GitAPI`/`GitWebhookAPI` both have `PostGitWebhook`, `IamAPI`/`O11yAPI` both
have `DeleteSession`, from four different operationIds). Every one of the 2479
operations is reachable.

Hand-written, and safe from regeneration:

| file | why |
|---|---|
| `hanzo.go` | `New` / `Options` / `Client` / `As` — the constructor and the six accessors |
| `identity.go` | the IAM mints: client credentials, and the act grant `As` rides on |
| `call/` | `Answer`, `Denied`, `Held`, `Fault`, `Money`, `Page`, and the one request path |
| `budget/` `policy/` `audit/` `search/` `kb/` `graph/` | the six capabilities |
| `secret/` | `Boot` — a service reads its own credentials out of KMS at startup |
| `hanzo_test.go`, `identity_test.go`, `capabilities_test.go` | the round trip against an `httptest` estate |
| `examples/` | the six flows, plus `models`, `errors` and `six` |
| `go.mod`, `README.md`, `LLM.md`, `hanzo.yml`, `.hanzo/`, `scripts/` | repo-owned |

**The six live in their own packages, and that is forced.** The generated client
already defines `Answer`, `Page`, `Hit`, `Backend`, `Wrote`, `Allowance`,
`Charge`, `Policy`, `Audit`, `Graph` and `Link` as projections of unrelated
operations, and a regeneration may add more. A hand-written type cannot share a
package with a generated one of the same name, so `call.Answer[T]`,
`audit.Filter` and `graph.Fact` are where they are — which also means a
regeneration can never collide with them.

**`.generated` is what makes that table a fact rather than a promise.** It names
every path the driver wrote — the root `*.go` and `docs/*.md`, nothing else — so
a regeneration removes a file the document stopped projecting and cannot touch a
file it never wrote, whatever directory that file sits in. It is the same record
in every Hanzo SDK, and it is why the client can live at the module root at all:
ownership is a set of files, not a directory. Everything it names is generated —
do not edit it, change the handler in `hanzoai/cloud`.

## Auth: IAM mints, the SDK never accepts a bearer

The document declares it: one `securityScheme`, `bearer` (`type: http`,
`scheme: bearer`), and a root `security: [{bearer: []}]` every operation
inherits. Four opt out with `security: []` — `GET /v1/models`,
`GET /v1/models/providers`, `GET /v1/commands`, `GET /v1/openapi.json` — and
`/v1/models` says why in its own description: the catalogue takes no principal,
so the route reads `Authorization` to annotate gated SKUs and never to admit.

**The client mints its own token and never takes one.** `New(Options{})` reads
`HANZO_CLIENT_ID` and `HANZO_CLIENT_SECRET` and exchanges them at
`POST {issuer}/v1/iam/oauth/token` — `client_credentials` under
`client_secret_basic`, scoped by the RFC 8707 `resource`, which defaults to the
base URL (HIP-0111). The reference is `hanzoai/visor/egress_identity.go`. There
is no `HANZO_API_KEY` and no way to hand the SDK a bearer: a credential an SDK
is given is a credential nobody rotates and one that says nothing about who is
calling, which is the question every refusal in the estate exists to answer.

The token is held until a minute before expiry and minted again on a 401, once —
a second 401 is the server saying no. Both the mint and the 401 replay live on an
`http.RoundTripper`, which is the one place a credential is presented: the
generated surface and the six capabilities share the same `*http.Client`, so
there is one header on a request and never two. `TestFlows` asserts
`len(Header.Values("Authorization")) == 1` rather than just its value.

A client holding no credentials presents nothing rather than failing. The four
operations that take none still answer.

## `secret.Boot`: where a service's own credentials come from

A service needs values before it can serve: a signing key, a database URL, the
client secret it mints with. `secret.Boot` reads them from KMS at startup, over
the identity the pod already carries, and holds them in memory.

```go
import "github.com/hanzoai/go-sdk/v8/secret"

values, err := secret.Boot(ctx, "hanzo/cloud", "HANZO_CLIENT_SECRET", "DB_URL")
client := hanzoai.New(hanzoai.Options{Secret: values["HANZO_CLIENT_SECRET"]})
```

The projected ServiceAccount token is the assertion: IAM exchanges it at
`POST /v1/iam/oauth/token` for a bearer (RFC 7523,
`grant_type=urn:ietf:params:oauth:grant-type:jwt-bearer`, which IAM serves from
v1.34.99), and KMS answers `GET /v1/kms/secrets/<path>/<key>?env=<env>` under
it. One mint per `Boot` and exactly one more on a 401 — a second 401 is a
refusal, and retrying it is how a boot loop becomes a mint flood. Every key must
resolve or the whole `Boot` fails naming the key, because a map one entry short
is a service that starts, serves, and fails at the first request that needed the
value. A key is a bare name — no separator, no `.` or `..` — so it can address a
secret and nothing else.

| environment | is | default |
| --- | --- | --- |
| `HANZO_SA_TOKEN` | the projected ServiceAccount token | `/var/run/secrets/hanzo/iam/token` |
| `HANZO_IAM_URL` | where the assertion is exchanged | `https://hanzo.id` |
| `KMS_URL` | where the values live | `http://cloud.hanzo.svc:8000` |
| `HANZO_ENV` | which environment a name resolves in | `prod` |
| `HANZO_DEV` | `1` with no token file: read the keys from the environment | — |

These are the platform's own variable names, so a pod already carrying them
needs nothing added; `HANZO_IAM_URL` is the same host `HANZO_ISSUER_URL` names
for the client, under the name a pod is given it by. `HANZO_DEV=1` with no token
file is the only path by which a secret reaches the process from the
environment, and it says so in the log; it looks each key up under its own name,
so `DB_URL` is `DB_URL` either way. In production an absent token is an error,
never a fallback.

The alternative this replaces is a Kubernetes Secret, which is a base64 field
anyone holding get on the namespace reads, mounted where anything in the pod
reads it again.

## Acting as a subject

`client.As(subject)` returns the same client — the same 2502 methods — scoped to
one tenant subject. It mints a subject-bound token from IAM's act grant, caches
it to expiry and re-mints once on a 401, so no method takes a user id: the
credential is the scope, and a caller cannot pass the wrong one.

**One extension point, an `http.RoundTripper` on the Configuration's
`HTTPClient`.** Signing and the 401 replay are the same object, because they are
one question — is this request carrying a live token — asked at the one moment
the answer is known.

The act grant is not a call to the platform. It is
`POST https://hanzo.id/v1/iam/tokens/issue?id=<subject>`, on IAM's own host,
carrying the operator's own minted access token; **`HANZO_ISSUER_URL`** moves it
for a private estate, the way `HANZO_BASE_URL` moves the gateway. The generated
`PostIamTokensIssue` cannot serve: the document states that operation as an
address and not as a shape, so the method carries neither the `id` query nor the
camelCase `{accessToken, expiresIn}` it answers with — the same reason
`identity.go` exists.

A scoped client shares the operator's held token, so scoping costs one act grant
and not a second exchange. The operator credential does not travel with the
minted one: the subject-bound token is the whole identity of a scoped call.

**The two mints spell their answers differently** — `{access_token, expires_in}`
from the OAuth endpoint, `{accessToken, expiresIn}` from the act grant — and
`identity.go` is where that ends.

## The six capabilities, and the answer they share

`Budget`, `Policy`, `Audit`, `Search`, `KB` and `Graph` hang off the same client
as the generated surface. One word each, the same word in the Go, Python and
TypeScript SDKs; Go exports it capitalized because that is Go's rule, and `KB`
is all-caps because that is Go's rule for an initialism. Method names carry no
capability prefix: `client.Graph.Read`, never `client.Graph.ReadGraph`.

**A refusal is a value, not an exception.** Every method a gate can refuse
answers `call.Answer[T]`, whose value is reachable only through `Value()` — Go
has no sum type, and an ignored error return is loud where an exported field
read is silent. `Value()` hands back a `*call.Denied` where a budget or a policy
said no (with `Code`, `Reason`, `Product` and the `Cures` that clear it), a
`*call.Held` where a person was asked, and a `*call.Fault` where nothing decided
anything. `Answer.Request` is the `x-request-id` on every arm, and it is the
join to `audit.Event.Request`.

**Pure reads answer their value directly.** `Budget.Left`, `Budget.Balance`,
`Budget.Plan`, `Budget.Spent`, `Policy.Check`, `Audit.List`, `Audit.All`,
`Graph.Read`, `Graph.Find`, `Graph.Resolve`, `Graph.Walk`, `Graph.Extract`,
`Graph.Vocabulary`, `KB.Get`, `KB.List`, `KB.Connectors`, `KB.Connect` and
`KB.Links` are not refused, so wrapping them in an `Answer` would be a branch
with one arm.

**One rule maps an HTTP answer to an arm**, in `call.Ask`, and no capability
varies it. 2xx is the value; a 202 whose body says `held` is a hold; 402 is a
denial whatever the code; a 403 carrying `policy_denied`,
`entitlement_required`, `spend_cap_exceeded` or `insufficient_balance` is a
denial; everything else is a fault. That last allow-list is a workaround for one
cloud defect: cloud spells "no validated principal" as `403 forbidden`, and
reading a bare forbidden as a denial would tell an unauthenticated caller their
budget said no. The day cloud answers 401 for that, the list goes and the rule
collapses to "402 or 403 means denied".

**The body says whether a call was held, not the status code.** A dozen
operations answer 202 for "accepted, working on it" and carry a real schema — a
deployment, a preview, a build — so discriminating on the code turns each of
them into an approval nobody is waiting on, and the caller waits forever for an
answer from someone who was never asked. Only `"status":"held"` is a hold. That
is the discriminator the other Hanzo clients use, so there is one contract
rather than one per language.

`call.Ask` decides the hold before it decodes, because the hold is a property of
the answer: a held body need not fit the operation's own schema, and the caller
still has to learn a person was asked.

## Generation knobs

Every knob is the `go:` row of `hanzoai/openapi`'s `sdks.yaml` and nowhere else.
`scripts/generate.sh` is a call site: it names the language and this checkout.

```
generator: go            take: {.: .}          format: [gofmt, -w]
properties: packageName=hanzoai, withGoMod=false, structPrefix=true, enumClassPrefix=true
flags:      name-mappings <eleven fields>, model-name-mappings Config=StreamConfig
global:     apiDocs=true, modelDocs=true       (--skip-validate-spec is the driver's, for every language)
```

`format` is there because the generator's Go output is not gofmt'd — 2655 of
2656 files — and an unformatted Go file reformats itself in every contributor's
editor. `apiDocs`/`modelDocs` are on where the fleet default has them off: Go is
the one client that publishes `docs/` as its reference.

`--skip-validate-spec` is load-bearing. The document is OpenAPI **3.1**, which
made `responses` optional on an operation; the validator in 7.14.0 still enforces
the 3.0 rule that it is required, and 716 of 2479 operations state their address
and not their shape on purpose. The gate is `go build`, not the validator.

`structPrefix=true` is what lets an operation carry two tags: the request struct
is named `<Service>API<Op>Request`, so 25 operations reachable from two service
groups produce two distinct structs instead of a redeclaration. This used to be
recorded here as a hard blocker.

**The mappings rename what Go sees and never the wire** — the document's key
stays as the field's json tag. Each mapped name occurs in exactly one schema,
measured against the locked document, so a global mapping reaches nothing else.
Three schemas need them, all because Go emits `Get<F>`/`Has<F>`/`Set<F>` beside
every field:

| schema | what collides |
|---|---|
| `o11y.O11yPodOnboarding` | eight `has<Label>Name` flags beside the labels; the FIELD `HasClusterName` and the METHOD `HasClusterName` are one name. All eight take `<label>NamePresent` — one family, one spelling, so the next sibling label cannot break the build |
| `o11y.PostableProfile` | `has_existing_observability_tool` beside `existing_observability_tool` |
| `o11y.GettableAgentCheckIn` | two spellings of two fields (`integration_config` with `integrationConfig`, `removed_at` with `removedAt`), published together so older AWS agents keep working. The snake one is the legacy wire and takes the suffix — the same correction python and kotlin make |

**Spell every mapped value EXPORTED.** The Go generator takes a mapped name
verbatim where python and kotlin re-case theirs, so a camelCase value lands a
lowercase field: unexported, invisible to a caller, and skipped by
`encoding/json` in both directions. It compiles, and it drops the value the
mapping was written to preserve.

`Config=StreamConfig` is a model rename, and the collision is with this repo:
the document's `Config` is the MQ stream configuration (`POST /v1/mq/streams`,
`Stream.config`), whose constructor `NewConfig` is already `hanzo.go`'s. It is
also the one bare `Config` in a document where every other is qualified —
`TLSConfig`, `iam.config`, `o11y.DiscordConfig`.

## The drift check

```bash
OPENAPI=../openapi ./scripts/generate.sh --check   # non-zero if the client drifted
```

It regenerates into a temp dir and compares against the files `.generated`
names. Restricting the comparison to those is the load-bearing part: `hanzo.go`
and `hanzo_test.go` are root `*.go` files the generator never wrote, so a check
that compared the two directories outright would report the hand-written seam as
drift on every run and could never pass.

## The flows come from `flows.yaml`

`hanzoai/openapi` carries a root-level **`flows.yaml`** naming six example flows
and, per flow, the operationIds to call in order. It is the manifest that makes
"the same examples in every SDK" a fact. `examples/` and the `TestFlows` table
both follow it — do not pick a different operation here without changing it
there first.

**All six ship.** `chat` is `POST /v1/chat/completions`, which the document
states as an address with no request body and no responses, so the generated
method carries no prompt and hands back the raw `*http.Response` — the same
shape `money` already reads. That is an example, not a blocker: it calls the
operation the document declares and prints what the route answered, which is
what java-sdk and kotlin-sdk do with the same untyped operation. What it must
not do is hand-roll the missing request body, because a request invented inside
a generated client is the second authority these SDKs exist to remove. Probed
without a key, the route answers 401 rather than 404, so it is mounted.

`hello` is `get_keys` (`GET /v1/keys`), chosen by probing rather than by reading:
it answers 403 with no key and with a bogus one while `/v1/keys-zzq9` answers
404, so the refusal is this route's. It replaced `bot_authMe`, which stopped
resolving when cloud began relaying all of `/v1/bot` through one wildcard.

## Names moved, twice, and both are in the past

Every method used to read `CloudGetV1Tools`: a `cloud_` service prefix and the
default version, both inside the operationId. The document dropped both, and the
method is `GetTools`. Measured on the two clients: `v1.0.1` has 2478 methods over
263 services, of which 1351 carry `V1` and 1502 carry `Cloud`; the current tree
has 2502 over 192, of which **one** carries `V1` and none carry the prefix.

Every apparent survivor is correct and must not be "fixed":

- `GetTeamTransactorApiV1Statistics` drops only the FIRST `v1`, because the
  second is a path segment.
- 11 methods spell `Cloud` because `cloud` is a path segment of the product —
  `GetCloudAccounts`, `GetPricingCloudPlans`. Another 33 are `Cloudflare`, and a
  grep for `Cloud` counts all 44 as leftovers.
- `CloudIntegrationId` is a query-parameter setter named after the parameter
  `cloudIntegrationId`.

Never derive a name — read it off the generated client or the document.

## What cloud has to change, measured

Each of these was probed against `api.hanzo.ai` or read out of `hanzoai/cloud`.
Each one removes a workaround the three SDKs otherwise write identically.

| route | what it does | what it costs the SDK |
|---|---|---|
| every gated route | answers `403 forbidden` where there is no validated principal | `call.refusals` exists only to tell a refusal from an absent caller |
| every 402 | two bodies: `errmap`'s RFC 9457 envelope with `code`, and `cloud.Refuse`'s `{error, product, reason, message, cure[]}` | `call.denial` reads both |
| `POST /v1/authz/check` | binds `{subject, verb, path, grants}` and answers `{allow, subject, verb, path}`, where the document's own description says `{sub, obj, act}` | the SDK sends what the handler binds; the description is what the contract was written from |
| `POST /v1/authz/check` | decides against grants carried IN THE REQUEST — `authz.Can` fails closed, so an empty grant set authorizes nothing | a three-argument check answers false for every question until cloud reads the caller's grants from IAM, as its description says it does |
| `GET /v1/billing/balance`, `/usage` | declare an address and no shape | `budget` models `{balance, holds, available, account}` and the usage envelope by hand |
| `GET /v1/audit` | no `requestId` filter, though every row carries one; `pageSize` and `p` are strings | `audit.Filter.Request` narrows client-side |
| `/v1/framework/kb.*` | the elective middleware answers **404 with the problem envelope** to a caller whose org has not enabled the module | a 404 there is a refusal wearing an absence; only the plain-text 404 means unrouted |
| `POST /v1/graph/ingest` | validates before it authenticates — an unauthenticated `POST {}` answers `400 timestamp "" is not RFC 3339`, measured | an unauthenticated caller learns the validation rules |
| `/v1/approvals/{id}` | is not served | the `held` arm is terminal: a caller learns a person was asked and cannot poll |
| every metered route | emits no per-call cost | spend is unattributable without a second read |

## Release state

`v1.0.0` is the reverted experiment, is published, and outranks every
`v0.1.0-alpha.*`. The proxy is immutable (cached 2026-07-05), so retagging or
deleting it does nothing; the fix was to publish higher.

**`v1.0.1` is the last tag that predates the rename**, and that is a documented
fact rather than trivia: `go get github.com/hanzoai/go-sdk/v8` resolved to a client
whose `KeysAPI` has `CloudGetV1Keys` and no `GetKeys`, so every method name in
the README was one a consumer following the install line could not call. The fix
was the tag, not a footnote: **`v1.0.2` is out at `3ad23f3e`** and is the first
tag carrying the renamed operationIds. Do not write a method name into the README
that the newest TAG does not carry — cut the tag instead.

The README pins the current tag rather than printing a bare `go get`, because
the proxy's `@latest` and `@v/list` are cached for a while after a push while
`@v/<version>.info` is immediate. A pinned line works the second the tag lands;
a bare one silently resolves to the previous tag until the index catches up,
which is the whole failure this repo once had. Move that pin with every tag —
v1.0.2 stays named in the prose as the floor, which is a different fact.

Publishing a Go module is pushing the tag. `.hanzo/workflows/release.yml` proves
the tag compiles and warms the proxy; there is no registry and no token. The
proxy reads github.com/hanzoai/go-sdk/v8, which GitHub redirects to hanzo-go/sdk,
which the forge push-mirrors within seconds — so a tag pushed here is resolvable
publicly without anything else being done.

Confirming that from a Hanzo machine takes care, because two settings route
around the thing under test: `GOPRIVATE=github.com/hanzoai/*` makes the fetch
skip the proxy, and `url.https://git.hanzo.ai/hanzoai/.insteadOf
https://github.com/hanzoai/` aims the direct fetch at the forge. A green
`go get` under both proves the forge has the tag and says nothing about the
registry. Exporting `GOPRIVATE=` does not undo it — Go falls back to its
`go env` file when the variable is empty, so the override has to be a pattern
that matches nothing. `GOVCS=off` is what makes the check honest: it forbids
version control outright, leaving the proxy as the only source. From an empty
module:

```bash
GOMODCACHE=$(mktemp -d) GOFLAGS= GOVCS=off GOPROXY=https://proxy.golang.org \
GOSUMDB=sum.golang.org GOPRIVATE=example.invalid GONOPROXY=example.invalid \
GIT_CONFIG_GLOBAL=/dev/null go get github.com/hanzoai/go-sdk/v8@v1.0.4
```

A `go.sum` line for the version means sum.golang.org vouched for it as well.
Two files are in the tag but not in the artifact: `AGENTS.md` and `CLAUDE.md`
are symlinks to `LLM.md`, and a module zip carries no symlinks. Everything else
is byte-identical, so a diff of the two is a real integrity check.

There is no `CHANGELOG.md`, and one should not come back. The generator that
wrote the old one left with the client it described, so it stopped at
`0.1.0-alpha.7` while three releases shipped past it — a file that says nothing
happened since July reads as fact. The tag list is the record.

## Testing

`hanzo_test.go` stands up an `httptest` server answering both IAM's mint and the
platform API, and asserts for each flow that the client sends the documented
method and path plus exactly one `Authorization: Bearer`. If a regeneration moves
an operation, that test fails instead of the examples silently calling a wrong
endpoint.

`capabilities_test.go` does the same for the six, a row per method, asserting the
exact request — method, path, query and body — beside the decode. That half of a
client contract is the half a decode cannot check: a method that posts the right
shape to the wrong address passes every test that only reads the reply. It also
walks each arm end to end: a 402 in both bodies cloud writes, a `policy_denied`
403, a `held` 202, and a bare `forbidden` that must NOT read as a denial.

`identity_test.go` is the credential: the `client_credentials` exchange with its
`resource`, the held token, the single 401 replay with the body it carried, the
act grant `As` rides on, and the assertion that concurrent first calls mint once
between them.

Two examples run to completion with no credential, and they are the pair that
proves the auth contract from both ends against the live API: `models` is a
public operation answering everyone, and `errors` is a gated route answering 403
to a client that cannot say who it is. `errors` is also the way to check a base
URL end to end.

```
$ go run ./examples/models              # no credentials set
200 OK  112 model(s)
$ go run ./examples/errors
status    403 Forbidden
$ HANZO_CLIENT_ID=... HANZO_CLIENT_SECRET=... go run ./examples/six
```

`examples/six` is the five-step composition: what the wallet holds decides
whether to ask, the wallet's own account names the org a policy question is
scoped to, a search states its degradation, a graph write answers a request id,
and that id finds the row the server wrote about it.

This replaces the suite that shipped with that client, whose `ok` was 187 SKIP /
25 PASS against a disabled mock server and asserted nothing about the API
surface.
