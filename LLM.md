# LLM.md — hanzo-go/sdk

Go module `github.com/hanzoai/go-sdk`, published by **git tag** → proxy.golang.org.

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

The **module path stays `github.com/hanzoai/go-sdk`** and is not renamed. A
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
run. All five Hanzo SDKs read this one document at this one digest. Generating
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
beside two hand-written ones. There is no second surface: the hand-shipped
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
| `hanzo.go` | `NewConfig` / `NewClient` — the authenticated constructor |
| `hanzo_test.go` | pins the six flows to their routes; asserts one bearer header |
| `examples/` | the six flows, plus `models` and `errors` |
| `go.mod`, `README.md`, `LLM.md`, `hanzo.yml`, `.hanzo/`, `scripts/` | repo-owned |

**`.generated` is what makes that table a fact rather than a promise.** It names
every path the driver wrote — the root `*.go` and `docs/*.md`, nothing else — so
a regeneration removes a file the document stopped projecting and cannot touch a
file it never wrote, whatever directory that file sits in. It is the same record
in every Hanzo SDK, and it is why the client can live at the module root at all:
ownership is a set of files, not a directory. Everything it names is generated —
do not edit it, change the handler in `hanzoai/cloud`.

## Auth, and why `hanzo.go` exists

The document declares it: one `securityScheme`, `bearer` (`type: http`,
`scheme: bearer`), and a root `security: [{bearer: []}]` every operation
inherits. Four opt out with `security: []` — `GET /v1/models`,
`GET /v1/models/providers`, `GET /v1/commands`, `GET /v1/openapi.json` — and
`/v1/models` says why in its own description: the catalogue takes no principal,
so the route reads `Authorization` to annotate gated SKUs and never to admit.

openapi-generator emits auth code from that declaration and nowhere else, so
what it emits is now real: `configuration.go` defines
`ContextAccessToken = contextKey("accesstoken")` and `client.go` writes
`Authorization: Bearer <token>` from it in `prepareRequest`. Before the
declaration landed the generated client held zero references to either.

That reader takes the token **per call**, off the context. `NewConfig` puts it
on the Configuration instead, once, because the credential belongs to the client
and not to a call — and that is the whole reason a hand-written file exists in a
generated package.

**One place, and the reason is mechanical.** `prepareRequest` reads the context
key and then `Add`s every default header, both under `Authorization`, so a token
set in both places is sent twice. `TestFlows` asserts
`len(Header.Values("Authorization")) == 1` rather than just its value, so the
day someone wires the context path as well the test says so. A second identity
is a second client.

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

## Release state

`v1.0.0` is the reverted experiment, is published, and outranks every
`v0.1.0-alpha.*`. The proxy is immutable (cached 2026-07-05), so retagging or
deleting it does nothing; the fix was to publish higher.

**`v1.0.1` is the last tag that predates the rename**, and that is a documented
fact rather than trivia: `go get github.com/hanzoai/go-sdk` resolved to a client
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
proxy reads github.com/hanzoai/go-sdk, which GitHub redirects to hanzo-go/sdk,
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
GIT_CONFIG_GLOBAL=/dev/null go get github.com/hanzoai/go-sdk@v1.0.4
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

`hanzo_test.go` stands up an `httptest` server, points the SDK at it with
`HANZO_BASE_URL`, and asserts for each flow that the client sends the documented
method and path plus `Authorization: Bearer`. If a regeneration moves an
operation, that test fails instead of the examples silently calling a wrong
endpoint.

Two examples run to completion with no credential, and they are the pair that
proves the auth contract from both ends against the live API:

```
$ go run ./examples/models              # HANZO_API_KEY unset
200 OK  112 model(s)
  all-mini-lm-l6-v2            do-ai
  anthropic-claude-opus-5      do-ai
  best                         hanzo
  ...
$ go run ./examples/errors              # a key the API refuses
status    403 Forbidden
refused   {"status":403,"code":"forbidden","error":"sign in to manage API keys"}
$ HANZO_API_KEY=sk-... go run ./examples/hello
the key is accepted; this org holds 1 key(s)
```

`models` is a public operation answering everyone; `errors`/`hello` are the same
route answering 403 without a credential and 200 with one, which is what makes
the credential load-bearing rather than decorative. `errors` is also the way to
check a base URL end to end.

This replaces the suite that shipped with that client, whose `ok` was 187 SKIP /
25 PASS against a disabled mock server and asserted nothing about the API
surface.
