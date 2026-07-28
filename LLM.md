# LLM.md — hanzoai/go-sdk

Go module `github.com/hanzoai/go-sdk`. Published by **git tag** → pkg.go.dev.
There is no release workflow: `release-please-config.json` exists but nothing
consumes it, so a release is `internal/version.go` + `.release-please-manifest.json`
+ an annotated tag, by hand. Patch bumps only.

```bash
go build ./...
go test ./...
```

## Two client surfaces — both real, don't merge them by hand

| surface | package | origin | edit |
|---|---|---|---|
| repo root `*.go` | `hanzoai` | **legacy Stainless** client, 188 endpoints | frozen |
| `cloud/` | `cloud` | **generated** from `hanzo.yaml` by openapi-generator | never by hand |

Stainless was retired 2026-07. The root package is what it left behind; it is
still the surface most consumers import, so it stays. `cloud/` is the current
spec-driven client covering the fused `api.hanzo.ai/v1` surface. Neither package
imports the other. This dual surface is deliberate — do not "unify" it.

Residue already removed: `.stats.yml` (pointed at a dead
`stainless-sdk-openapi-specs` GCS bucket — a *second*, contradicting spec
source), `model1.go.bak` / `model1_test.go.bak`, and the CI `build` job that was
gated `if: github.repository == 'stainless-sdks/hanzo-ai-go'` (never true here)
plus its only caller `scripts/utils/upload-artifact.sh`.

## The pipeline — one spec in, N clients out

```
hanzoai/cloud    emits its own router spec   -> cloud/openapi.yaml + openapi/generated/hanzo.json
hanzoai/openapi  merges 69 per-service specs -> hanzo.yaml (1885 paths)   [the ONE SDK input]
hanzoai/openapi  generate.py + sdks.yaml     -> projects hanzo.yaml into each SDK repo
this repo        owns its test + version bump + release
```

Regenerate — the only way:

```bash
cd ~/work/hanzo/openapi && python3 generate.py go     # rewrites go-sdk/cloud in place
python3 generate.py go --check                        # non-zero if cloud/ drifted
```

`cloud/` is wholly overwritten by that command (`shutil.rmtree` + `copytree`).
Anything hand-written there is destroyed on the next run. Generation knobs live
in `openapi/sdks.yaml` under `go:` (`packageName: cloud`), **not** here — this
repo carries no generation config by design, so a client cannot drift on its own.

SHAs behind the current `cloud/`: hanzoai/cloud `8143fc0e`, hanzoai/openapi `f581a0e`.

Do **not** regenerate from `~/work/hanzo/cloud/openapi.yaml` directly. That file
is 983 paths — an *input* to the merge, not the SDK input. Generating from it
drops ~48% of the surface and forks a second contract.

## Known blocker — `cloud/` does not compile (upstream, not ours)

As generated from openapi `f581a0e`, `go build ./...` fails with 40 errors. Two
root causes, both in `hanzoai/openapi`, neither fixable in this repo:

1. **Seven operations carry multiple tags.** openapi-generator's go generator
   emits one `api_<tag>.go` per tag and re-declares the per-operation
   `Api<OperationId>Request` struct in each, so they collide:
   `commerce_{authorize,capture,charge,refund}Order` `[Orders, Checkout]`,
   `commerce_store{Authorize,Charge}` `[Store, Checkout]`,
   `pricing_getFullPricing` `[Models, Cloud, Infrastructure]`.
   Authored in the service specs (e.g. `commerce/openapi.yaml`). `merge.py`
   namespaces operationIds but does not collapse to one tag, despite
   `openapi/CLAUDE.md` claiming it "collapses to one primary tag".
2. **Enum constant collision.** `CANCELLED` is declared by both
   `model_commerce_payment_status.go` and `model_commerce_order_status.go`.

Fix belongs in `hanzoai/openapi` (tag collapse in `merge.py`, enum prefix in
`sdks.yaml`), then regenerate here. Until then `cloud/` stays uncommitted —
CI `scripts/lint` runs `go build ./...`, and a Go tag is immutable once the
module proxy sees it, so a broken tag can never be fixed, only superseded.

`cloud/` needs two deps the root package does not: `golang.org/x/oauth2`
(`cloud/client.go`) and `gopkg.in/validator.v2` (generated models).

## Testing caveat — green is largely vacuous

`go test ./...` reports `ok` for the root package, but that is **187 SKIP / 25 PASS**:
the Stainless-generated tests all hit a prism mock server that is disabled
("Mock server tests are disabled"). Only 25 real unit tests
(`internal/apijson`, `apiquery`, `apiform`, `option`) actually assert anything.
Do not read `ok` as verified behaviour.
