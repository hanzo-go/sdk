// The Hanzo Go SDK. Generated from hanzoai/cloud's `openapi.yaml` by
// scripts/generate.sh, at the ref .spec-lock names — see README.md.
//
// The module path carries /v8, because Go's semantic import versioning gives a
// path without one exactly two publishable major versions: v0 and v1. This
// client answers to the release line every other Hanzo client does -- 8.5.x --
// so at the bare path its tags were unpublishable, and they were: v8.5.89 and
// v8.5.156 both sat on the repo while proxy.golang.org served v1.0.5 and always
// would. A tag the proxy rejects is not a release, it is a decoration.
//
// It stays github.com/hanzoai/go-sdk/v8 and NOT hanzo-go/sdk. The repo moved
// and GitHub redirects the clone, but a go.mod path must match what consumers
// require, and the hanzoai path is the one already on the proxy and pkg.go.dev.
// Renaming the host half would strand every existing import; adding /v8 does
// not -- v1 consumers keep resolving v1.0.5 at the old path, untouched, which
// is the whole point of the suffix.
module github.com/hanzoai/go-sdk/v8

go 1.26.8

require (
	golang.org/x/oauth2 v0.36.0
	gopkg.in/validator.v2 v2.0.1
)
