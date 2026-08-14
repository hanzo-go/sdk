# Hanzo Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/hanzoai/go-sdk.svg)](https://pkg.go.dev/github.com/hanzoai/go-sdk)

The Go client for the [Hanzo API](https://api.hanzo.ai). Generated from
`hanzoai/cloud`'s `openapi.yaml`, which cloud emits from its own routers and
gates on every release — so a method here is a route the binary serves, and this
SDK and the TypeScript, Python, Java, Kotlin and Rust SDKs all describe one
product at one version. `.spec-lock` names the ref this client is a projection
of.

2477 operations across 192 services, one package.

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

	keys, _, err := client.KeysAPI.GetKeys(context.Background()).Execute()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(keys.Keys), "key(s)")
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
Five are runnable programs:

| Flow | What it does | Operations |
| --- | --- | --- |
| [`hello`](examples/hello) | Prove the key works | `GET /v1/keys` |
| `chat` | One chat completion | `POST /v1/chat/completions` — pinned, not shipped |
| [`money`](examples/money) | Balance and usage | `GET /v1/billing/balance`, `GET /v1/billing/usage` |
| [`store`](examples/store) | Provision a KV store, read it, delete it | `POST /v1/kv`, `GET`/`DELETE /v1/kv/{name}` |
| [`agent`](examples/agent) | Create an agent, run it, poll the run | `POST /v1/agents`, `POST /v1/agents/{ref}/run`, `GET /v1/agents/{ref}/runs` |
| [`tools`](examples/tools) | List the tools this key can reach | `GET /v1/tools` |

```bash
HANZO_API_KEY=sk-... go run ./examples/hello
```

`store` and `agent` are org-scoped — set `HANZO_ORG_ID` as well.

`chat` ships in no language. The document states `POST /v1/chat/completions` and
not its shape — no request body, no response — so the generated method carries
no prompt. The test still pins the address; the example returns the day cloud's
handler declares its types.

## Regenerating

The `*.go` files at the module root are generated. Do not edit them; change the
handler in `hanzoai/cloud` and regenerate:

```bash
./scripts/generate.sh          # the ref .spec-lock names, digest re-checked
./scripts/generate.sh --check  # non-zero if the committed client drifted
```

Reading the document from `git.hanzo.ai` needs `FORGE_TOKEN` (contents:read on
`hanzoai/cloud`). Pass it by value instead with `SPEC=/path/to/openapi.yaml`.
Every generator knob lives in that script and nowhere else, including the eleven
field renames and one model rename the Go generator needs; the renames never
reach the wire.

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
for Go that means the tag itself; `.hanzo/workflows/release.yml` proves it
compiles and warms the proxy so pkg.go.dev picks it up immediately.

## License

Apache-2.0. See [LICENSE](LICENSE).
