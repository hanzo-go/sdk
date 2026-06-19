# metering — the one way to meter usage to commerce

`github.com/hanzoai/go-sdk/metering` is the single, reusable hook every Hanzo
product uses to charge for usage. It is the DRY replacement for the balance-check
and usage-record logic that was copy-pasted across the LLM/cloud path
(`ai/routers/filter_balance.go`, `ai/controllers/openai_api.go`,
`gateway/auth_middleware.go`). Any product — search, functions, storage, a CLI —
imports this package and gets the proven, fail-closed billing gate plus usage
recording, against **commerce, the single billing source of truth**.

## What it does

Two operations, matching the proven cloud/gateway path:

| Op | Commerce endpoint | Purpose |
|----|-------------------|---------|
| `Authorize` | `GET /v1/billing/balance` (or `GET /v1/billing/tier` when tier-aware) | Pre-request balance gate. **Fail-closed** by default. |
| `Record`    | `POST /v1/billing/usage` | Post-request usage write (debits the user's balance ledger). |

Auth is the commerce service token (admin-scoped S2S):

```
Authorization: Bearer ${COMMERCE_SERVICE_TOKEN}
X-IAM-Org-Id: <tenant org slug>
```

The token is a secret and **must** come from KMS (the operator wires it from a
KMS-backed secret into `COMMERCE_SERVICE_TOKEN`). This package never reads it
from disk.

## Use it (middleware — the common case)

```go
meter, _ := metering.FromEnv() // COMMERCE_URL + COMMERCE_SERVICE_TOKEN (KMS) + COMMERCE_SERVICE_ORG

mux.Use(meter.Middleware(metering.MiddlewareConfig{
    Provider: "search",
    Price: func(r *http.Request, status int, in metering.AuthInput) int64 {
        if status >= 200 && status < 300 { return 5 } // 5¢ per successful request
        return 0
    },
    Skip: func(r *http.Request) bool { return r.URL.Path == "/healthz" },
}))
```

`Middleware` is plain `func(http.Handler) http.Handler`, so it composes with the
standard library, `gorilla/mux` (`.Use`), `chi`, and anything that speaks
`http.Handler` — no per-framework variants.

It reads the caller identity from the **gateway-minted** `X-User-Id` / `X-Org-Id`
headers (the trust boundary). The gateway strips client-supplied copies on
ingress, so these are safe to trust downstream.

## Use it (imperative — per-unit pricing)

```go
in := metering.IdentityFromGatewayHeaders(r)
if err := meter.Authorize(ctx, in); err != nil {
    // ErrInsufficientBalance -> 402 ; other -> 503 (fail-closed)
}
// ... do work, measure cost ...
meter.Record(ctx, metering.Usage{User: in.User, Org: in.Org, AmountCents: cents, Provider: "functions"})
```

## Configuration (env, operator-wired)

| Var | Default | Notes |
|-----|---------|-------|
| `COMMERCE_URL` | `http://commerce.hanzo.svc.cluster.local:8001` | Commerce base (no `/v1` suffix). |
| `COMMERCE_SERVICE_TOKEN` | — | Admin-scoped S2S token. **KMS-sourced.** |
| `COMMERCE_SERVICE_ORG` | `hanzo` | Default tenant org (`X-IAM-Org-Id`). |
| `METERING_TIER_AWARE` | `false` | Gate on `effectiveAvailable` (prepaid + included plan allotment). |
| `METERING_FAIL_OPEN` | `false` | Allow-on-error. Leave false for paid products. |
| `METERING_DISABLED` | `false` | Force not-configured mode for local dev. |

## Fail-closed contract (aligned with the gateway)

`Authorize` returns:

- `nil` → allow.
- `ErrInsufficientBalance` → out of funds → map to **402**.
- any other error → balance unknown → **fail-closed: deny** → map to **503**
  (set `FailOpen` to allow instead).

When no `COMMERCE_URL` is configured the client is in "not configured" mode:
`Authorize` allows and `Record` is a no-op, so a product can ship the wrap
before its tenant billing is wired.

This is the same balance source and the same status mapping the gateway uses —
no divergent logic.
