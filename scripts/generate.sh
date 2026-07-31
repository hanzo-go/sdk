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
# hanzoai/openapi is PRIVATE today. raw.githubusercontent.com only serves public
# repos, so the plain URL 404s; when that happens this falls back to the GitHub
# API with a token (SPEC_TOKEN, or GH_TOKEN/GITHUB_TOKEN, or `gh auth token`)
# and says so. SPEC_URL overrides the URL, SPEC overrides the file.
#
# Requires: java 17+, curl, go.
set -euo pipefail
cd "$(dirname "$0")/.."

GENERATOR_VERSION="${GENERATOR_VERSION:-7.14.0}"
SPEC_REPO="${SPEC_REPO:-hanzoai/openapi}"
SPEC_REF="${SPEC_REF:-main}"
SPEC_URL="${SPEC_URL:-https://raw.githubusercontent.com/${SPEC_REPO}/${SPEC_REF}/hanzo.yaml}"
SPEC="${SPEC:-}"
JAR="${JAR:-${TMPDIR:-/tmp}/openapi-generator-cli-${GENERATOR_VERSION}.jar}"

if [ -z "$SPEC" ]; then
  SPEC="$(mktemp)"
  # Public fetch first: it is the plain path, needs no credential, and starts
  # working the day hanzoai/openapi opens. While the repo is private GitHub
  # answers 404 rather than 403, so an anonymous miss is indistinguishable from
  # a deleted file — hence the fallback below, which says which case it was.
  # Both paths use curl -f under set -e, so a failed fetch stops the script
  # instead of regenerating from a stale spec.
  if ! curl -fsSL "$SPEC_URL" -o "$SPEC"; then
    TOKEN="${SPEC_TOKEN:-${GH_TOKEN:-${GITHUB_TOKEN:-$(gh auth token 2>/dev/null || true)}}}"
    : "${TOKEN:?$SPEC_URL is not readable anonymously and no SPEC_TOKEN/GH_TOKEN is set. $SPEC_REPO is private; supply a token with contents:read, or pass SPEC=/path/to/hanzo.yaml}"
    echo "note: $SPEC_URL returned no spec ($SPEC_REPO is private) - reading it through the GitHub API instead" >&2
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
# Validation stays ON. hanzo.yaml validates clean, and a malformed document
# should fail here rather than surface as a compile error spread across 1300
# generated files.
java -jar "$JAR" generate \
  -i "$SPEC" -g go \
  --additional-properties=packageName=hanzoai,withGoMod=false,structPrefix=true,enumClassPrefix=true \
  --git-user-id=hanzoai --git-repo-id=go-sdk \
  -o "$OUT"

# The repo root owns go.mod, examples/, scripts/ and the docs. Keep only the
# generated sources. Previous output is removed via the generator's own FILES
# manifest, so a renamed or dropped operation cannot leave a stale file behind.
#
# Only what this script copies back is removed. The generator's manifest also
# lists files the repo owns and this script never restores — README.md,
# .gitignore, .travis.yml, git_push.sh — so deleting everything it names takes
# the hand-written docs with it.
if [ -f .openapi-generator/FILES ]; then
  grep -E '^([A-Za-z0-9_]+\.go|docs/.+\.md)$' .openapi-generator/FILES |
    while IFS= read -r f; do [ -n "$f" ] && rm -f "$f"; done
fi
rm -rf docs .openapi-generator

cp "$OUT"/*.go .
cp -r "$OUT"/docs "$OUT"/.openapi-generator .
cp "$OUT"/.openapi-generator-ignore .

gofmt -w ./*.go
echo "generated $(ls ./*.go | wc -l) Go files at the module root (package hanzoai)"
