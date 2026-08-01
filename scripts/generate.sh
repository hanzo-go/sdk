#!/usr/bin/env bash
# Regenerate the Hanzo Go SDK from the unified OpenAPI spec.
#
# The ONE way: hanzoai/openapi `hanzo.yaml` is the single source of truth. This
# SDK is generated from it with openapi-generator (go) — no Stainless, no
# hand-drift. Never edit the generated *.go at the repo root; edit the
# per-service spec in hanzoai/openapi and regenerate.
#
#   ./scripts/generate.sh                            # pulls spec from hanzoai/openapi@main
#   ./scripts/generate.sh --check                    # diff only; non-zero if the tree drifted
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

check=0
[ "${1:-}" = "--check" ] && check=1

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

STAGE="$(mktemp -d)"

# The document as JSON, because YAML has a ceiling and JSON does not.
#
# swagger-parser hands a YAML document to snakeyaml, which refuses anything over
# 3 * 1024 * 1024 = 3145728 code points. hanzo.yaml passed that mark and this
# script has been unable to generate since: measured 2026-08-01 at 3,686,318
# code points, `YAMLException: The incoming YAML document exceeds the limit:
# 3145728 code points`. The failure does not say so out loud — the parser logs
# SnakeException, falls through to the Swagger 2.0 compat reader, and dies with
# "Issues with the OpenAPI input", which reads like a malformed spec. It is not:
# the document validates at 0 errors.
#
# `-DmaxYamlCodePoints` is NOT the fix — swagger-parser honours it in generator
# 7.24.0 and ignores it in the 7.14.0 pinned here. JSON avoids snakeyaml
# altogether on every version. The generator reads either format from -i, so
# this costs one temp file and removes a ceiling the document keeps growing into.
#
# This is the same conversion hanzoai/openapi's generate.py and cpp-sdk's
# generate.sh already do, for the same reason, and it is deliberately NOT
# written back as a second committed artifact: there is one document, and it is
# hanzo.yaml.
SPEC_JSON="$STAGE/hanzo.json"
python3 -c 'import json,sys,yaml; json.dump(yaml.safe_load(open(sys.argv[1])), open(sys.argv[2],"w"))' \
  "$SPEC" "$SPEC_JSON"

OUT="$STAGE/gen"
# --skip-validate-spec: the document is OpenAPI 3.1, and 3.1 made `responses`
# OPTIONAL on an operation. The validator in generator 7.14.0 still enforces the
# 3.0 rule that it is required, so it REFUSES a document that is valid. Measured
# on hanzoai/cloud's openapi.yaml: 684 of 1636 operations are routes the router
# proves exist and whose response shape no seam can state, and cloud emits those
# with no `responses` key on purpose (openapi/openapi.go — "absent stays valid
# and absent beats invented").
#
# What keeps a bad document out is not the validator, it is `go build ./...` —
# the whole generated package plus the six example flows, in hanzo.yml's test:
# block. A malformed document fails there with a file and a line.
java -jar "$JAR" generate \
  -i "$SPEC_JSON" -g go \
  --skip-validate-spec \
  --additional-properties=packageName=hanzoai,withGoMod=false,structPrefix=true,enumClassPrefix=true \
  --git-user-id=hanzoai --git-repo-id=go-sdk \
  -o "$OUT"

gofmt -w "$OUT"/*.go

if [ "$check" = 1 ]; then
  # The client is generated, so the only thing that can rot is the committed
  # copy. This is what makes "never edit the generated *.go" a fact rather than
  # a convention — and it is what lets the release train gate on this repo.
  #
  # Compare only what this script writes back: the root *.go and docs/. go.mod,
  # examples/, scripts/ and the hand-written docs are the repo's own and the
  # generator never sees them.
  rc=0
  for f in "$OUT"/*.go; do
    b="$(basename "$f")"
    if ! cmp -s "$f" "$b"; then echo "DRIFTED: $b"; rc=1; fi
  done
  for f in ./*.go; do
    b="$(basename "$f")"
    [ -f "$OUT/$b" ] || { echo "DRIFTED: $b is committed and hanzo.yaml does not project it"; rc=1; }
  done
  [ "$rc" = 0 ] && echo "clean: the module root is what hanzo.yaml projects"
  exit "$rc"
fi

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

echo "generated $(ls ./*.go | wc -l) Go files at the module root (package hanzoai)"
