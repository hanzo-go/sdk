# Hanzo Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/hanzoai/go-sdk/v8.svg)](https://pkg.go.dev/github.com/hanzoai/go-sdk/v8)

The Go client for the [Hanzo API](https://api.hanzo.ai), generated from the API's
own OpenAPI document.

## Install

```bash
go get github.com/hanzoai/go-sdk/v8
```

Go 1.26 or newer. **v1.0.2 is the floor** — earlier versions spell the methods
differently (`CloudGetV1Keys` rather than `GetKeys`).

## Quickstart

```go
package main

import (
	"context"
	"fmt"
	"log"

	hanzoai "github.com/hanzoai/go-sdk/v8"
)

func main() {
	client := hanzoai.New(hanzoai.Options{})

	left, err := client.Budget.Left(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d of %d calls left this %s\n", *left.Left, left.Limit, left.Window)
}
```

```bash
HANZO_CLIENT_ID=... HANZO_CLIENT_SECRET=... go run .
```

## The six

Six capabilities hang off the client, one word each:

| | Answers |
| --- | --- |
| `client.Budget` | what you may spend, what is left, what it cost |
| `client.Policy` | may this subject take this action on this object |
| `client.Audit` | the org's own trail, newest first |
| `client.Search` | one ranked result set over everything the org has stored |
| `client.KB` | the corpus you write and the connectors that fill it |
| `client.Graph` | assertions with provenance and time |

**A refusal is an answer, not an exception.** Anything a gate can refuse hands
back a `call.Answer[T]`, and the value is reachable only through `Value()`:

```go
answer := client.Search.Find(ctx, "q3 incident postmortems", search.Opts{})
hits, err := answer.Value()

var denied *hanzoai.Denied
if errors.As(err, &denied) {
	fmt.Println(denied.Code, denied.Reason)   // "insufficient_balance", …
	for _, cure := range denied.Cures {
		fmt.Println(cure.Kind, cure.URL)      // how to clear it, in order
	}
}
```

`*hanzoai.Held` is a call a person was asked about; `*hanzoai.Fault` is an
outcome with no decision in it. `answer.Request` is the `x-request-id` on every
arm, and it is what `audit.Filter{Request: …}` takes.

Reads no gate refuses — `Budget.Left`, `Audit.List`, `Graph.Read`, `KB.Get` and
the rest — answer their value directly.

The whole generated surface is on the same client, under the same credential:

```go
listing, _, err := client.AccountAPI.GetAccountKeys(ctx).Execute()
```

Every generated call has the same shape: pick a service off the client, name the
operation, add parameters by chaining, then `Execute()`. It returns the decoded
value, the raw `*http.Response`, and an error — or, where the document states no
response shape, the response and the error alone.

## Authenticating

IAM mints; the SDK never accepts a bearer. `New` exchanges the client's own IAM
credentials for an access token (`client_credentials`, RFC 8707 `resource`), holds
it until shortly before expiry, and mints again on a 401.

| option | environment | default |
| --- | --- | --- |
| `ID` | `HANZO_CLIENT_ID` | — |
| `Secret` | `HANZO_CLIENT_SECRET` | — |
| `Base` | `HANZO_BASE_URL` | `https://api.hanzo.ai` |
| `Issuer` | `HANZO_ISSUER_URL` | `https://hanzo.id` |
| `Resource` | `HANZO_RESOURCE` | `Base` |

Every option falls back to its environment variable, so `hanzoai.Options{}` is
the normal case. Four operations take no credential: `GET /v1/models`,
`GET /v1/models/providers`, `GET /v1/commands`, `GET /v1/openapi.json`.

No method takes an org: the tenant is the validated principal, everywhere. To
act as one of your tenant's subjects, scope the client:

```go
scoped := client.As("usr_7")
```

The operator credential leaves with the scope — two credentials on one request
are two answers to who is calling.

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

## Examples

Start with `models`. It calls a public operation, so it runs with nothing set up:

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
HANZO_CLIENT_ID=... HANZO_CLIENT_SECRET=... go run ./examples/hello
```

```
the key is accepted; this org holds 1 key(s)
```

| | Does | Calls | Credential |
| --- | --- | --- | --- |
| [`models`](examples/models) | List the catalogue | `GET /v1/models` | no |
| [`hello`](examples/hello) | Prove the credential works | `GET /v1/account/keys` | yes |
| [`chat`](examples/chat) | One completion | `POST /v1/chat/completions` | yes |
| [`money`](examples/money) | The wallet, the allowance, and the charges that moved them | `client.Budget` | yes |
| [`store`](examples/store) | Create a KV store, read it, delete it | `POST /v1/provisioning/kv`, `GET`/`DELETE /v1/provisioning/kv/{name}` | yes |
| [`agent`](examples/agent) | Create an agent, run it, poll the run | `POST /v1/agent`, `POST /v1/agent/{ref}/run`, `GET /v1/agent/{ref}/runs` | yes |
| [`tools`](examples/tools) | List the tools this credential can reach | `GET /v1/tool` | yes |
| [`errors`](examples/errors) | Read a refusal | `GET /v1/account/keys` | no |
| [`six`](examples/six) | Budget, policy, search, graph and the audit trail of what just happened | all six | yes |

`hello`, `chat`, `money`, `store`, `agent` and `tools` are the flows
`hanzoai/openapi`'s `flows.yaml` names for every Hanzo SDK, so they carry the
same names and routes in every language.

Some operations state their address and not their shape, and their generated
methods hand back the raw `*http.Response` with nothing to unmarshal into.
`chat` and `money` show how to read one.

## Reference

Per-service method lists with a runnable snippet each are in [`docs/`](docs) and
on [pkg.go.dev](https://pkg.go.dev/github.com/hanzoai/go-sdk/v8). The API itself is
documented at [docs.hanzo.ai](https://docs.hanzo.ai).

## Regenerating

Everything at the module root except `hanzo.go`, `identity.go`, the tests and
the six capability packages is generated. Change the handler in cloud and
regenerate:

```bash
OPENAPI=../openapi ./scripts/generate.sh          # the ref .spec-lock names
OPENAPI=../openapi ./scripts/generate.sh --check  # non-zero if the client drifted
```

`OPENAPI` points at a [`hanzoai/openapi`](https://git.hanzo.ai/hanzoai/openapi)
checkout, which holds the one driver every Hanzo SDK is generated by; every
generator knob is its `sdks.yaml` `go:` row. Pass a document you already have
with `SPEC=/path/to/openapi.yaml`.

## Development

```bash
go build ./... && go vet ./... && go test -count=1 ./... && go build ./examples/...
```

## License

Apache-2.0. See [LICENSE](LICENSE).
