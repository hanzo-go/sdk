# LLM.md — hanzo-go/sdk

Go module `github.com/hanzoai/go-sdk`, published by **git tag** → proxy.golang.org.

```bash
go build ./... && go vet ./... && go test ./... && go build ./examples/...
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
remotes and cannot resolve `forge/main`. `scripts/generate.sh` re-fetches at that
ref and refuses to run if the bytes hash to anything else.

It is **not** `hanzoai/openapi`'s `hanzo.yaml`. That file is itself a projection
of this document with codegen rules applied, so reading it made this client a
projection of a projection — one release behind whenever the middle step had not
run. `sdks.yaml` has no `go:` row for the same reason it has no `rust:` row: its
`take:` key promises the generator owns a directory, and this client's directory
is the module root, beside `.git`.

## One client, at the module root

The repo root `*.go` **is** the client, `package hanzoai` — 2656 generated files
beside two hand-written ones. There is no second surface: the Stainless client
and the parallel `cloud/` subpackage are both gone.

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
| `hanzo_test.go` | pins the six flows to their routes; asserts the bearer token |
| `examples/` | the six flows, plus `errors` |
| `go.mod`, `README.md`, `LLM.md`, `hanzo.yml`, `.hanzo/`, `scripts/` | repo-owned |

`scripts/generate.sh` deletes the `*.go` and `docs/*.md` entries of the
generator's own `.openapi-generator/FILES` manifest — and **only** those, which
is the load-bearing part. The manifest also names `README.md`, `.gitignore`,
`.travis.yml` and `git_push.sh`; the generator writes those and this repo owns
them, so honouring the manifest wholesale would take the hand-written front door
with it. Everything the two patterns do reach is generated — do not edit it,
change the handler in `hanzoai/cloud`.

## Why `hanzo.go` exists

The document declares **no `security` at all** — not at the root, not on any of
its 2479 operations. openapi-generator emits auth code only from those
declarations, so the generated client contains zero references to
`ContextAccessToken` and would send no `Authorization` header on any call.
`NewConfig` sets the bearer as a default header, which covers every operation
and is why a hand-written file exists in a generated package.

When cloud's emitter declares a scheme, this stays as it is: a default header is
still one place, and per-operation auth code would be 2479.

## Generation knobs

Every knob is in `scripts/generate.sh` and nowhere else.

```
-g go --skip-validate-spec
--additional-properties=packageName=hanzoai,withGoMod=false,structPrefix=true,enumClassPrefix=true
--name-mappings <eleven fields>
--model-name-mappings Config=StreamConfig
```

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
./scripts/generate.sh --check   # non-zero if the committed client drifted
```

It regenerates into a temp dir and compares both directions. The second
direction asks the committed `.openapi-generator/FILES` whether the generator
ever owned a file before calling it stale — without that it reports `hanzo.go`
and `hanzo_test.go`, which the generator never wrote, and can never pass.

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
the README was one a consumer following the install line could not call. The
README says **v1.0.2 or later** for exactly that reason. Do not write a method
name into the README that the newest TAG does not carry — cut the tag instead.

Publishing a Go module is pushing the tag. `.hanzo/workflows/release.yml` proves
the tag compiles and warms the proxy; there is no registry and no token. The
proxy reads github.com/hanzoai/go-sdk, which GitHub redirects to hanzo-go/sdk,
which the forge push-mirrors within seconds — so a tag pushed here is resolvable
publicly without anything else being done.

## Testing

`hanzo_test.go` stands up an `httptest` server, points the SDK at it with
`HANZO_BASE_URL`, and asserts for each flow that the client sends the documented
method and path plus `Authorization: Bearer`. If a regeneration moves an
operation, that test fails instead of the examples silently calling a wrong
endpoint.

`examples/errors` is the one example that runs to completion with no credential,
because its whole subject is the refusal: it prints `403 Forbidden` and the
API's own `{"code":"forbidden"}` body. Use it to check a base URL end to end.

This replaces the old Stainless suite, whose `ok` was 187 SKIP / 25 PASS against
a disabled mock server and asserted nothing about the API surface.
