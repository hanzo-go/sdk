#!/usr/bin/env bash
# Regenerate the Hanzo Go SDK from the Hanzo Cloud API document.
#
# THE DOCUMENT IS hanzoai/cloud's `openapi.yaml`, emitted from its own routers
# and gated on every release, so it cannot describe a route the binary does not
# serve. Never edit the generated *.go at the module root; change the handler in
# hanzoai/cloud and regenerate.
#
#   ./scripts/generate.sh                              # regenerate in place
#   ./scripts/generate.sh --check                      # non-zero if the tree drifted
#   SPEC=/path/to/openapi.yaml ./scripts/generate.sh   # the document by value
#
# What stood here read a DIFFERENT file: hanzoai/openapi's `hanzo.yaml`, over
# raw.githubusercontent.com, at whatever `main` was that minute, behind a
# four-deep token chain for a private repo. That file is itself a projection of
# this document with codegen rules applied, so reading it made this client a
# projection of a projection — pinned to nothing, one release behind whenever
# the middle step had not run, and 900-odd routes apart from what cloud serves.
#
# Requires: java 17+, curl, go, and FORGE_TOKEN (contents:read on hanzoai/cloud)
# unless SPEC is passed.
set -euo pipefail
cd "$(dirname "$0")/.."

GENERATOR_VERSION="${GENERATOR_VERSION:-7.14.0}"
SPEC="${SPEC:-}"
JAR="${JAR:-${TMPDIR:-/tmp}/openapi-generator-cli-${GENERATOR_VERSION}.jar}"

check=0
[ "${1:-}" = "--check" ] && check=1

if [ -z "$SPEC" ]; then
  SPEC="$(mktemp)"
  # THE REF THIS TREE NAMES, not a branch. hanzoai/ci's client lane passes the
  # document by value in $SPEC, already digest-checked against the release that
  # published it; by hand, `.spec-lock` says which ref this committed client is a
  # projection of and the digest is re-checked here. A branch read would drag the
  # client onto a document no release shipped.
  ref="$(sed -n 's/^ref=//p' .spec-lock)"
  want="$(sed -n 's/^sha256=//p' .spec-lock)"
  : "${ref:?no SPEC and no .spec-lock — this tree names no document}"
  : "${FORGE_TOKEN:?reading hanzoai/cloud from git.hanzo.ai needs FORGE_TOKEN (contents:read), or pass SPEC=/path/to/the/document}"
  curl -fsSL -H "Authorization: token $FORGE_TOKEN" \
    "https://git.hanzo.ai/v1/repos/hanzoai/cloud/raw/openapi.yaml?ref=$ref" -o "$SPEC"
  got="$(sha256sum "$SPEC" | cut -d' ' -f1)"
  [ "$got" = "$want" ] || { echo "hanzoai/cloud@$ref:openapi.yaml hashes to $got, but .spec-lock says $want — the ref moved under this projection" >&2; exit 1; }
fi

if [ ! -f "$JAR" ]; then
  curl -fsSL -o "$JAR" \
    "https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/${GENERATOR_VERSION}/openapi-generator-cli-${GENERATOR_VERSION}.jar"
fi

STAGE="$(mktemp -d)"

# The document as JSON, because YAML has a ceiling and JSON does not.
#
# swagger-parser hands a YAML document to snakeyaml, which refuses anything over
# 3 * 1024 * 1024 = 3145728 code points. The document passed that mark and this
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
# This is the same conversion generate.py already does, for the same reason, and
# it is deliberately NOT written back as a second committed artifact: there is
# one document and it lives in hanzoai/cloud.
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
    [ -f "$OUT/$b" ] || { echo "DRIFTED: $b is committed and the document does not project it"; rc=1; }
  done
  [ "$rc" = 0 ] && echo "clean: the module root is what the document projects"
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
