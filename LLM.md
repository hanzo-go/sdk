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
| `v0.1.0-alpha.7` | `feb1137` (#17) | version bump only. Branch `next` HEAD. The intended current release. |

So the root-package regeneration was tried, tagged `v1.0.0`, and then abandoned —
but the tag was never (and can never be) unpublished: the module proxy is
immutable. Because `v1.0.0` outranks `v0.1.0-alpha.7` in semver, **every consumer
doing `go get` receives the reverted, abandoned experiment**, not alpha.7.

The only fix is to publish a **higher** version carrying the intended client —
`v1.0.1` or later, from whichever layout wins. Retagging or deleting `v1.0.0`
does nothing; proxy.golang.org has already cached it (2026-07-05).

Do not tag `v0.1.0-alpha.6` or `.7` — both exist and point at other commits.

## Branches

- `main` — Stainless client. **3 commits behind `next`** (it never received
  #15/#16/#17). Its version is `0.1.0-alpha.5`, matching its lineage.
- `next` — `feb1137`, alpha.7, the live release branch. Builds clean.
  release-please's flow is `next` → `main` (see the
  `release-please--branches--main--changes--next` branch), and it has stalled.
- `sdk/openapi-generator`, `generated`, `feat/metering-commerce-client` — other
  in-flight attempts. Check them before starting a fourth.

`main` and `next` have diverged; reconciling them is a decision someone has to
make, not something to paper over. Nothing should be released from `main` until
it is.

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

From hanzoai/cloud `8143fc0e` + hanzoai/openapi `f581a0e`, `go build ./...` fails
with 40 errors. Two causes, both upstream in `hanzoai/openapi`:

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
