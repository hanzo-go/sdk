# LLM.md - Hanzo Go SDK

## Overview
Go client for the Hanzo Cloud API (`https://api.hanzo.ai/v1`). Covers the
full unified surface — AI inference plus every `/v1/<service>` product.

Go module: `github.com/hanzoai/go-sdk` (package `hanzoai`).

## The ONE way: generated from the OpenAPI spec

This SDK is **generated**, never hand-written. Source of truth is
`hanzoai/openapi` `hanzo.yaml` (built by that repo's `merge.py` from the
per-service specs). Generator is **openapi-generator** (`go`) — no Stainless.

```bash
./scripts/generate.sh                    # regenerate from hanzoai/openapi@main
SPEC=/path/to/hanzo.yaml ./scripts/generate.sh   # from a local spec
```

Never edit `api_*.go` / `model_*.go` / `client.go` by hand — change the spec
in `hanzoai/openapi` and regenerate. CI (`generate.yml`) does this
automatically on every spec change via a `spec-update` repository_dispatch.

## Auth
Bearer token — an IAM JWT or a `hk-` Cloud API key:

```go
cfg := hanzoai.NewConfiguration()
cfg.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("HANZO_API_KEY"))
client := hanzoai.NewAPIClient(cfg)
```

Base URL defaults to `https://api.hanzo.ai`.

## Build
```bash
go build ./...
go vet ./...
```

## Release
Push a semver tag `vX.Y.Z` → `release.yml` warms the Go module proxy so
pkg.go.dev serves it. Semver only, never a sha pin.
