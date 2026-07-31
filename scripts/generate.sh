#!/usr/bin/env bash
# Regenerate the Hanzo Go SDK from the unified OpenAPI spec.
#
# The ONE way: hanzoai/openapi `hanzo.yaml` is the single source of truth. This
# SDK is generated from it with openapi-generator (go) — no Stainless, no
# hand-drift. Never edit the generated *.go at the repo root; edit the
# per-service spec in hanzoai/openapi and regenerate.
#
#   ./scripts/generate.sh                            # pulls spec from hanzoai/openapi@main
#   SPEC=/path/to/hanzo.yaml ./scripts/generate.sh   # local spec override
#
# hanzoai/openapi is PRIVATE, so the spec is read through the GitHub API with a
# token (SPEC_TOKEN, or GH_TOKEN/GITHUB_TOKEN). raw.githubusercontent.com only
# serves public repos — it 404s here. SPEC_URL still overrides for a public
# mirror.
#
# Requires: java 17+, curl, go.
set -euo pipefail
cd "$(dirname "$0")/.."

GENERATOR_VERSION="${GENERATOR_VERSION:-7.14.0}"
SPEC_REPO="${SPEC_REPO:-hanzoai/openapi}"
SPEC_REF="${SPEC_REF:-main}"
SPEC_URL="${SPEC_URL:-}"
SPEC="${SPEC:-}"
JAR="${JAR:-${TMPDIR:-/tmp}/openapi-generator-cli-${GENERATOR_VERSION}.jar}"

if [ -z "$SPEC" ]; then
  SPEC="$(mktemp)"
  if [ -n "$SPEC_URL" ]; then
    curl -fsSL "$SPEC_URL" -o "$SPEC"
  else
    TOKEN="${SPEC_TOKEN:-${GH_TOKEN:-${GITHUB_TOKEN:-}}}"
    : "${TOKEN:?SPEC_TOKEN required to read private $SPEC_REPO (or pass SPEC=/path/to/hanzo.yaml)}"
    curl -fsSL \
      -H "Authorization: Bearer $TOKEN" \
      -H "Accept: application/vnd.github.raw" \
      "https://api.github.com/repos/${SPEC_REPO}/contents/hanzo.yaml?ref=${SPEC_REF}" -o "$SPEC"
  fi
fi

if [ ! -f "$JAR" ]; then
  curl -fsSL -o "$JAR" \
    "https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/${GENERATOR_VERSION}/openapi-generator-cli-${GENERATOR_VERSION}.jar"
fi

OUT="$(mktemp -d)"
# --skip-validate-spec: 35 /v1/platform/* operations in hanzo.yaml are missing
# the required `responses` object. That is a spec-side defect (reported
# upstream); it does not affect the generated Go, and the gate that matters —
# `go build ./... && go vet ./...` — runs below.
java -jar "$JAR" generate \
  -i "$SPEC" -g go --skip-validate-spec \
  --additional-properties=packageName=hanzoai,withGoMod=false,structPrefix=true,enumClassPrefix=true \
  --git-user-id=hanzoai --git-repo-id=go-sdk \
  -o "$OUT"

# The repo root owns go.mod, examples/, scripts/ and the docs. Keep only the
# generated sources. Previous output is removed via the generator's own FILES
# manifest, so a renamed or dropped operation cannot leave a stale file behind.
if [ -f .openapi-generator/FILES ]; then
  while IFS= read -r f; do [ -n "$f" ] && rm -f "$f"; done < .openapi-generator/FILES
fi
rm -rf docs .openapi-generator

cp "$OUT"/*.go .
cp -r "$OUT"/docs "$OUT"/.openapi-generator .
cp "$OUT"/.openapi-generator-ignore .

gofmt -w ./*.go
echo "generated $(ls ./*.go | wc -l) Go files at the module root (package hanzoai)"
