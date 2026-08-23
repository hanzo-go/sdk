# Hanzo Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/hanzoai/go-sdk.svg)](https://pkg.go.dev/github.com/hanzoai/go-sdk)

The Go client for the [Hanzo API](https://api.hanzo.ai). It is generated from
the document `hanzoai/cloud` emits from its own routers, so a method here is a
route the running binary serves.

## Install

```bash
go get github.com/hanzoai/go-sdk
```

Needs Go 1.26 or newer. The line pins the current tag rather than resolving
`@latest`, which the proxy caches for a while after a push. **v1.0.2 is the
floor**, and the version is worth naming: the operationIds lost their `cloud_`
prefix and their default version, so what `v1.0.1` spelled `CloudGetV1Keys` this
client spells `GetKeys`. Later versions are drop-in; anything earlier answers to
none of the names below.

## Quickstart

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

	listing, _, err := client.KeysAPI.GetKeys(context.Background()).Execute()
	if err != nil {
		log.Fatal(err)
	}
	for _, key := range listing.Keys {
		fmt.Printf("%-12s %s\n", key.GetType(), key.GetPrefix())
	}
}
```

```bash
HANZO_API_KEY=sk-... go run .
```

Every call has the same shape: pick a service off the client, name the
operation, add parameters by chaining, then `Execute()`. It returns the decoded
value, the raw `*http.Response`, and an error — or, where the document states no
response shape, the response and the error alone.

## Authenticating

### bearer

The document declares the credential — one scheme, `bearer`, required at the
root — so the generated client carries the code that sends it, and all but four
operations need it. The four that do not say `security: []`: `GET /v1/models`,
`GET /v1/models/providers`, `GET /v1/commands`, `GET /v1/openapi.json`. Every
page under [`docs/`](docs) names which one it is and links back here.

A token is either an access token minted by Hanzo IAM or an API key (`sk-`/`pk-`).
`NewClient("")` reads it from **`HANZO_API_KEY`**; pass a key as the argument to
override the environment. The base URL is `https://api.hanzo.ai` unless
**`HANZO_BASE_URL`** says otherwise.

The credential belongs to the client, so it is set once, on the configuration,
and every call carries it. Set it in one place only: the generated client also
reads a token off the context (`ContextAccessToken`) and adds its header there,
so doing both sends `Authorization` twice. For a second identity, build a second
client.

Some services are scoped to an org and want an `X-Org-Id` header. Build the
configuration yourself for those:

```go
cfg := hanzoai.NewConfig("")
cfg.AddDefaultHeader("X-Org-Id", org)
client := hanzoai.NewAPIClient(cfg)
```

## Errors

A refusal comes back as a `*hanzoai.GenericOpenAPIError` **and** a response, not
one instead of the other: the status code is on the response, the API's own
message is on the error. A nil response means the request never went out.

```go
listing, resp, err := client.KeysAPI.GetKeys(ctx).Execute()
if err != nil {
	var apiErr *hanzoai.GenericOpenAPIError
	if errors.As(err, &apiErr) && resp != nil {
		log.Fatalf("%s: %s", resp.Status, apiErr.Body())
	}
	log.Fatalf("unsent: %v", err) // DNS, TLS, a cancelled context
}
```

[`examples/errors`](examples/errors) runs this against the live API and needs no
key — it prints `403 Forbidden` and the refusal cloud wrote.

## Examples

Start with `models`. It calls one of the four public operations, so it runs with
nothing set up at all:

```bash
go run ./examples/models
```

```
200 OK  112 model(s)
  all-mini-lm-l6-v2            do-ai
  anthropic-claude-opus-5      do-ai
  best                         hanzo
  ...
```

Then `hello`, the same shape with a credential behind it:

```bash
HANZO_API_KEY=sk-... go run ./examples/hello
```

```
the key is accepted; this org holds 1 key(s)
```

Six of the eight are the canonical flows from `hanzoai/openapi`'s `flows.yaml`,
the manifest every Hanzo SDK draws its examples from — same names, same routes —
so what you learn here transfers to the Python, TypeScript, Java, Kotlin and
Rust clients. `models` and `errors` are Go's own.

| | Does | Calls | Key |
| --- | --- | --- | --- |
| [`models`](examples/models) | List the catalogue | `GET /v1/models` | no |
| [`hello`](examples/hello) | Prove the key works | `GET /v1/keys` | yes |
| [`chat`](examples/chat) | One completion | `POST /v1/chat/completions` | yes |
| [`money`](examples/money) | Balance, and the usage that moved it | `GET /v1/billing/balance`, `GET /v1/billing/usage` | yes |
| [`store`](examples/store) | Create a KV store, read it, delete it | `POST /v1/kv`, `GET`/`DELETE /v1/kv/{name}` | yes |
| [`agent`](examples/agent) | Create an agent, run it, poll the run | `POST /v1/agents`, `POST /v1/agents/{ref}/run`, `GET /v1/agents/{ref}/runs` | yes |
| [`tools`](examples/tools) | List the tools this key can reach | `GET /v1/tools` | yes |
| [`errors`](examples/errors) | Read a refusal | `GET /v1/keys` | no |

`store` and `agent` are org-scoped, so set `HANZO_ORG_ID` for those.

Some operations state their address and not their shape — cloud has not declared
those handlers' types yet — and their generated methods hand back the raw
`*http.Response` with nothing to unmarshal into. `chat` and `money` show how to
read one.

## Reference

Per-service method lists with a runnable snippet each are in
[`docs/`](docs) — [`docs/KeysAPI.md`](docs/KeysAPI.md) is a good first one — and
on [pkg.go.dev](https://pkg.go.dev/github.com/hanzoai/go-sdk). The API itself is
documented at [docs.hanzo.ai](https://docs.hanzo.ai).

## Regenerating

Everything at the module root except `hanzo.go` and `hanzo_test.go` is
generated. Do not edit it — change the handler in `hanzoai/cloud` and
regenerate:

```bash
OPENAPI=../openapi ./scripts/generate.sh          # the ref .spec-lock names, digest re-checked
OPENAPI=../openapi ./scripts/generate.sh --check  # non-zero if the client drifted
```

`OPENAPI` points at a [`hanzoai/openapi`](https://git.hanzo.ai/hanzoai/openapi)
checkout, which holds the one driver every Hanzo SDK is generated by; every
generator knob is its `sdks.yaml` `go:` row and nowhere else. Reading the
document needs no credential — it comes from `api.hanzo.ai/v1/openapi.json`, the
same address hanzoai/ci's lane reads. Pass one by value with
`SPEC=/path/to/openapi.yaml` to generate against a document you already have.

## Development

```bash
go build ./... && go vet ./... && go test -count=1 ./... && go build ./examples/...
```

Those four are the whole gate, and `hanzo.yml` runs exactly them. Publishing is
pushing a semver tag: for a Go module there is no registry, only the proxy.

## License

Apache-2.0. See [LICENSE](LICENSE).
