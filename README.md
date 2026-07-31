# Hanzo Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/hanzoai/go-sdk.svg)](https://pkg.go.dev/github.com/hanzoai/go-sdk)

The Go client for the [Hanzo API](https://api.hanzo.ai). Generated from
[`hanzoai/openapi`](https://github.com/hanzoai/openapi) `hanzo.yaml` — the one
document that defines every Hanzo service — so this SDK and the TypeScript,
Python, and Rust SDKs all describe the same product.

2452 operations across 263 services, one package.

## Install

```bash
go get github.com/hanzoai/go-sdk
```

The module path is `github.com/hanzoai/go-sdk`. The repository moved to
`hanzo-go/sdk` and GitHub redirects it, but a `go.mod` path has to match what
consumers already require, so the import path does not change.

## Authenticate

Every request carries a bearer token. `NewClient("")` reads `HANZO_API_KEY`
from the environment; pass a key explicitly to override it. The base URL
defaults to `https://api.hanzo.ai` and is overridden by `HANZO_BASE_URL`.

```go
package main

import (
	"context"
	"fmt"
	"log"

	hanzoai "github.com/hanzoai/go-sdk"
)

func main() {
	client := hanzoai.NewClient("")

	me, _, err := client.AuthAPI.BotWhoami(context.Background()).Execute()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(me.GetHandle())
}
```

```bash
HANZO_API_KEY=sk-... go run .
```

Some services are org-scoped and need an `X-Org-Id` header. Build the config
yourself for those:

```go
cfg := hanzoai.NewConfig("")
cfg.AddDefaultHeader("X-Org-Id", org)
client := hanzoai.NewAPIClient(cfg)
```

## Flows

Six journeys, the same six in every Hanzo SDK, calling the operations
[`hanzoai/openapi` `flows.yaml`](https://github.com/hanzoai/openapi/blob/main/flows.yaml)
names — that manifest is what makes "the same six" a fact rather than six repos
independently remembering to agree. `hanzo_test.go` pins each one to its route.
Each is a runnable program:

| Flow | What it does | Operations |
| --- | --- | --- |
| [`hello`](examples/hello) | Resolve your key to an identity | `GET /v1/bot/auth/me` |
| [`chat`](examples/chat) | One chat completion | `POST /v1/chat/completions` |
| [`money`](examples/money) | Balance and usage | `GET /v1/billing/balance`, `GET /v1/billing/usage` |
| [`store`](examples/store) | Provision a KV store, read it, delete it | `POST /v1/kv`, `GET`/`DELETE /v1/kv/{name}` |
| [`agent`](examples/agent) | Create an agent, run it, poll the run | `POST /v1/agents`, `POST /v1/agents/{ref}/run`, `GET /v1/agents/{ref}/runs` |
| [`tools`](examples/tools) | List the tools this key can reach | `GET /v1/tools` |

```bash
HANZO_API_KEY=sk-... go run ./examples/hello
```

`store` and `agent` are org-scoped — set `HANZO_ORG_ID` as well.

## Regenerating

The `*.go` files at the module root are generated. Do not edit them; edit the
per-service spec in `hanzoai/openapi` and regenerate:

```bash
./scripts/generate.sh
```

`hanzoai/openapi` is private today, so the public spec URL 404s and the script
falls back to the GitHub API — supply `SPEC_TOKEN` (or `GH_TOKEN`, or be logged
into `gh`). Local override: `SPEC=/path/to/hanzo.yaml`.

Hand-written and safe from regeneration: `hanzo.go` (the client constructor),
`hanzo_test.go`, and `examples/`.

## Development

```bash
go build ./...
go vet ./...
go test ./...
go build ./examples/...
```

CI runs exactly these four (`hanzo.yml`). A tag `vX.Y.Z` publishes the module —
for Go that means the tag itself; `.github/workflows/release.yml` proves it
compiles and warms the proxy so pkg.go.dev picks it up immediately.

## License

Apache-2.0. See [LICENSE](LICENSE).
