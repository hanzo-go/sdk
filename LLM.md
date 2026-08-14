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

The repo root `*.go` **is** the client, `package hanzoai` — 2477 operations
across 192 services. There is no second surface: the Stainless client and the
parallel `cloud/` subpackage are both gone.

Hand-written, and safe from regeneration:

| file | why |
|---|---|
| `hanzo.go` | `NewConfig` / `NewClient` — the authenticated constructor |
| `hanzo_test.go` | pins the six flows to their routes; asserts the bearer token |
| `examples/` | the canonical flows |
| `go.mod`, `README.md`, `LLM.md`, `hanzo.yml`, `.hanzo/`, `scripts/` | repo-owned |

`scripts/generate.sh` deletes only what the generator's own
`.openapi-generator/FILES` manifest lists, so the table above survives a
regeneration. Everything not in it is generated — do not edit it, change the
handler in `hanzoai/cloud`.

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
the 3.0 rule that it is required, and 684 operations state their address and not
their shape on purpose. The gate is `go build`, not the validator.

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

**Six are pinned and five are shipped.** `chat` is `POST /v1/chat/completions`,
which the document states as an address with no request body and no responses,
so the generated method carries no prompt and there is no example to write.
python-sdk and js-sdk ship five for the same measurement. `TestFlows` still pins
the address, because that is what has to hold for the example to return the day
cloud's handler declares its `In`/`Out` types.

`hello` is `get_keys` (`GET /v1/keys`), chosen by probing rather than by reading:
it answers 403 with no key and with a bogus one while `/v1/keys-zzq9` answers
404, so the refusal is this route's. It replaced `bot_authMe`, which stopped
resolving when cloud began relaying all of `/v1/bot` through one wildcard.

## Names moved, twice, and both are in the past

Every method used to read `CloudGetV1Tools`: a `cloud_` service prefix and the
default version, both inside the operationId. The document dropped both —
`get_tools` — and this client had 2703 method names carrying `V1` and 3186
carrying `Cloud` until it was regenerated from the locked ref.

Two survivors are correct and must not be "fixed": `GetTeamTransactorApiV1Statistics`
drops only the FIRST `v1` because the second is a path segment, and
`CloudIntegrationId` is a query-parameter setter named after `cloudIntegrationId`.
Never derive a name — read it off the generated client or the document.

## Release state

`v1.0.0` is the reverted experiment, is published, and outranks every
`v0.1.0-alpha.*`. The proxy is immutable (cached 2026-07-05), so retagging or
deleting it does nothing; the fix was to publish higher, and `v1.0.1` is out at
`260d2e2a`. The next tag is `v1.0.2`.

Publishing a Go module is pushing the tag. `.hanzo/workflows/release.yml` proves
the tag compiles and warms the proxy; there is no registry and no token.

## Testing

`hanzo_test.go` stands up an `httptest` server, points the SDK at it with
`HANZO_BASE_URL`, and asserts for each flow that the client sends the documented
method and path plus `Authorization: Bearer`. If a regeneration moves an
operation, that test fails instead of the examples silently calling a wrong
endpoint.

This replaces the old Stainless suite, whose `ok` was 187 SKIP / 25 PASS against
a disabled mock server and asserted nothing about the API surface.
