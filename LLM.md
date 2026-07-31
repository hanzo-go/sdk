# LLM.md — hanzo-go/sdk

Go module `github.com/hanzoai/go-sdk`, published by **git tag** → proxy.golang.org.

```bash
go build ./... && go vet ./... && go test ./... && go build ./examples/...
```

Those four are the whole gate, and `hanzo.yml` runs exactly them.

## Repo identity

Canonical repo is **`hanzo-go/sdk`**. `hanzoai/go-sdk` is a GitHub rename
redirect — `gh api repos/hanzoai/go-sdk --jq .full_name` answers `hanzo-go/sdk`.

The **module path stays `github.com/hanzoai/go-sdk`** and is not renamed. A
`go.mod` path must match what consumers `require`; `github.com/hanzo-go/sdk` has
never been on the proxy and resolving it fails with a path mismatch. The clone
redirect makes the old path keep working, so there is nothing to fix.

## One client, at the module root

The repo root `*.go` **is** the client, `package hanzoai`, generated from
`hanzoai/openapi` `hanzo.yaml`. There is no second surface: the Stainless client
and the parallel `cloud/` subpackage are both gone.

That resolves the layout question this file used to leave open. `v1.0.0` had
already put generated output at the root and been reverted by `v0.1.0-alpha.6`;
the root layout is now the one that ships, for the reason the revert never
addressed — two client surfaces in one module is two ways to do one thing.

Hand-written, and safe from regeneration:

| file | why |
|---|---|
| `hanzo.go` | `NewConfig` / `NewClient` — the authenticated constructor |
| `hanzo_test.go` | pins the six flows to their routes; asserts the bearer token |
| `examples/` | the six canonical flows |
| `go.mod`, `README.md`, `hanzo.yml`, `.github/`, `.hanzo/` | repo-owned |

`scripts/generate.sh` deletes only what the generator's own
`.openapi-generator/FILES` manifest lists, so the table above survives a
regeneration. Everything not in that table is generated — do not edit it, edit
the per-service spec in `hanzoai/openapi`.

## Why `hanzo.go` exists

`hanzo.yaml` declares `security: [{bearerAuth: []}]` **once at the document
root**, which OpenAPI applies to every operation. openapi-generator 7.14.0's
`go` generator only emits auth code for operations carrying their *own* explicit
`security` block — 65 of 1518 here — so a straight generated client sends no
`Authorization` header on the other 1453 operations.

Confirmed as a generator limitation, not a spec defect, with a 12-line spec
carrying only root-level `security`: the generated Go contained zero references
to `ContextAccessToken`. `NewConfig` sets the bearer as a default header, which
covers every operation.

## SDKs pull; nobody pushes

`scripts/generate.sh` reads `hanzo.yaml` and regenerates in place. This repo owns
its own generation.

`hanzoai/openapi` also carries `generate.py` + `sdks.yaml` with a `go:` entry
(`take: { .: cloud }`) that *pushes* generated output into this repo. **That path
is superseded** — it targets the `cloud/` layout that no longer exists. The
`sdks.yaml` `go:` entry should be dropped; until it is, do not run
`generate.py go` against this repo, it will recreate the second surface.

```bash
SPEC_TOKEN=<token with contents:read on hanzoai/openapi> ./scripts/generate.sh
SPEC=/path/to/hanzo.yaml ./scripts/generate.sh    # local override
```

`hanzoai/openapi` is **private**, so the spec is read via
`api.github.com/repos/.../contents/hanzo.yaml`. `raw.githubusercontent.com`
404s for it — anything still pointing there is broken.

## Release state

`v1.0.0` is published, is what `go get` resolves to, and is the reverted
experiment. The proxy is immutable (cached 2026-07-05), so retagging or deleting
it does nothing. The only fix is to publish **higher** — the next tag is
`v1.0.1`.

Do not tag `v0.1.0-alpha.6`/`.7`/`.8`/`.9`; all exist and point elsewhere.

Publishing a Go module is pushing the tag. `.github/workflows/release.yml` proves
the tag compiles and warms the proxy; there is no registry and no token.

## Generation knobs

```
-g go --skip-validate-spec
--additional-properties=packageName=hanzoai,withGoMod=false,structPrefix=true,enumClassPrefix=true
```

`--skip-validate-spec` is load-bearing: 35 `/v1/platform/*` operations in
`hanzo.yaml` are missing the required `responses` object, which fails validation
outright. It is a spec-side defect, reported upstream; it does not affect the
generated Go, and `go build`/`go vet`/`go test` are the gates that matter.

The multi-tag collision this file used to record as a hard blocker (seven
operations tagged twice, so the generator re-declared their request structs and
the package would not compile) is **fixed upstream** — the spec now has zero
multi-tag operations.

## Testing

`hanzo_test.go` stands up an `httptest` server, points the SDK at it with
`HANZO_BASE_URL`, and asserts for each of the six flows that the client sends the
documented method and path plus `Authorization: Bearer`. If a regeneration moves
an operation, that test fails instead of the examples silently calling a wrong
endpoint.

This replaces the old Stainless suite, whose `ok` was 187 SKIP / 25 PASS against
a disabled mock server and asserted nothing about the API surface.
