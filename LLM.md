# LLM.md — hanzoai/go-sdk

Go module `github.com/hanzoai/go-sdk`, published by **git tag** → proxy.golang.org.
Read the release-state section before you tag anything.

```bash
go build ./...
go test ./...
```

## Release state is broken — read this first

`v1.0.0` is published and is what `go get github.com/hanzoai/go-sdk` resolves to.
It should not be. The history:

| tag | commit | what it is |
|---|---|---|
| `v0.1.0-alpha.5` | `6cb36eb` | Stainless client. Last good release on `main`'s lineage. |
| `v1.0.0` | `dbd4195` (#15) | "regenerate SDK from unified OpenAPI spec (retire Stainless)" — deleted the Stainless client, put openapi-generator output at the repo **root** as `package hanzoai`, added `generate.yml` + `release.yml`. |
| `v0.1.0-alpha.6` | `8145bd3` (#16) | **a full revert of #15.** Removed `.openapi-generator/FILES` (4412 lines), `.openapi-generator-ignore`, `generate.yml`, `release.yml`; restored the Stainless client. |
| `v0.1.0-alpha.7` | `feb1137` (#17) | version bump only. Was branch `next` HEAD. |
| `v0.1.0-alpha.8` | this line | `next` merged into `main`, plus the Stainless-residue removal below. |

So the root-package regeneration was tried, tagged `v1.0.0`, and then abandoned —
but the tag was never (and can never be) unpublished: the module proxy is
immutable. Because `v1.0.0` outranks `v0.1.0-alpha.7` in semver, **every consumer
doing `go get` receives the reverted, abandoned experiment**, not alpha.7.

The only fix is to publish a **higher** version carrying the intended client —
`v1.0.1` or later, from whichever layout wins. Retagging or deleting `v1.0.0`
does nothing; proxy.golang.org has already cached it (2026-07-05).

Do not tag `v0.1.0-alpha.6` or `.7` — both exist and point at other commits.

## Branches

- `main` — Stainless client, and now the release branch. `next` was merged in
  (it had stalled 3 commits ahead, holding #15/#16/#17).
- `next` — `feb1137`, alpha.7. Merged into `main`; land new work on `main`.
  release-please's flow is `next` → `main` (see the
  `release-please--branches--main--changes--next` branch), and it has stalled.
- `sdk/openapi-generator`, `generated`, `feat/metering-commerce-client` — other
  in-flight attempts. Check them before starting a fourth.

The `next` → `main` reconciliation is done, so `main` is releasable again. The
`v1.0.0` problem above is not fixed by it and still needs a decision.

## Two client surfaces

| surface | package | origin |
|---|---|---|
| repo root `*.go` | `hanzoai` | legacy **Stainless** client, 188 endpoints. Stainless retired 2026-07. |
| `cloud/` | `cloud` | **generated** from `hanzo.yaml` by openapi-generator. Not committed — see blocker. |

Two *layouts* have been attempted for the generated client and they conflict:
`v1.0.0` put it at the repo root as `package hanzoai` (replacing Stainless);
`openapi/sdks.yaml` currently says `take: { .: cloud }`, i.e. a `cloud/`
subpackage coexisting with Stainless. Pick one deliberately.

Residue removed from `main`: `.stats.yml` (it recorded an `openapi_spec_url`
pointing at a `stainless-sdk-openapi-specs` GCS bucket — a second, dead spec
source competing with `hanzo.yaml`), `model1.go.bak` / `model1_test.go.bak`, and
the CI `build` job gated `github.repository == 'stainless-sdks/hanzo-ai-go'`
(never true here) plus its only caller `scripts/utils/upload-artifact.sh`. All of
these are still present on `next` and will need removing there too.

## The pipeline — one spec in, N clients out

```
hanzoai/cloud    emits its own router spec   -> cloud/openapi.yaml (983 paths)
hanzoai/openapi  merges 69 per-service specs -> hanzo.yaml (1885 paths)  [the ONE SDK input]
hanzoai/openapi  generate.py + sdks.yaml     -> projects hanzo.yaml into each SDK repo
this repo        owns its test, version bump and release
```

Regenerate — the only way:

```bash
cd ~/work/hanzo/openapi && python3 generate.py go     # rewrites go-sdk/cloud in place
python3 generate.py go --check                        # non-zero if it drifted
```

`cloud/` is wholly overwritten by that command (`shutil.rmtree` + `copytree`);
anything hand-written there is destroyed on the next run. Generation knobs live
in `openapi/sdks.yaml` under `go:`, **not** here — this repo carries no
generation config by design, so a client cannot drift on its own.

Do **not** generate from `~/work/hanzo/cloud/openapi.yaml` directly. That file is
an *input* to the merge, not the SDK input; generating from it drops ~48% of the
surface and forks a second contract.

## Known blocker — the generated `cloud/` does not compile

From hanzoai/cloud `8143fc0e` + hanzoai/openapi `2861089`, `go build ./...` fails
with 30 errors. (Re-verified after regenerating at `2861089`, which picked up the
KMS contract change in `3300cda` — `/v1/kms/orgs/{org}/secrets` became
`/v1/kms/secrets`, the org now coming from the token. That regeneration fixed the
routes and changed nothing about the compile failure.) Two causes, both upstream
in `hanzoai/openapi`:

1. **Seven operations carry multiple tags.** The go generator emits one
   `api_<tag>.go` per tag and re-declares that operation's
   `Api<OperationId>Request` struct in each, so they collide:
   `commerce_{authorize,capture,charge,refund}Order` `[Orders, Checkout]`,
   `commerce_store{Authorize,Charge}` `[Store, Checkout]`,
   `pricing_getFullPricing` `[Models, Cloud, Infrastructure]`. Authored in the
   service specs (e.g. `commerce/openapi.yaml`). `merge.py` namespaces
   operationIds but does not collapse to one tag, despite `openapi/CLAUDE.md`
   claiming it "collapses to one primary tag".
2. **Enum constant collision** — `CANCELLED` is declared by both
   `model_commerce_payment_status.go` and `model_commerce_order_status.go`.

The generated client *does* carry the `/v1/admin/plugins` operator surface
(`PluginAdminPlugins`, `PluginAdminEnablePlugin`, `PluginAdminDisablePlugin`,
`PluginAdminReloadPlugin` in `cloud/api_admin.go`). The published `v1.0.0` does
**not** — it predates that spec change. That surface reaches consumers only once
the compile blocker is fixed and a `v1.0.1`+ is published.

`cloud/` needs two deps the root package does not: `golang.org/x/oauth2`
(`cloud/client.go`) and `gopkg.in/validator.v2` (generated models). They are kept
out of `go.mod` while `cloud/` is uncommitted, so `go mod tidy` stays a no-op.

## Testing caveat — green is largely vacuous

`go test ./...` prints `ok` for the root package, but that is **187 SKIP / 25 PASS**:
the Stainless-generated tests all target a prism mock server that is disabled
("Mock server tests are disabled"). Only 25 real unit tests
(`internal/apijson`, `apiquery`, `apiform`, `option`) assert anything, and none
of them touch the API surface. `v1.0.0` has **no test files at all**. Do not read
`ok` as verified behaviour.
