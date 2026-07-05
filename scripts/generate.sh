#!/usr/bin/env bash
# Regenerate the Hanzo Go SDK from the unified OpenAPI spec.
#
# The ONE way: hanzoai/openapi `hanzo.yaml` (built by merge.py from the
# per-service specs) is the single source of truth. This SDK is generated
# from it with openapi-generator — no Stainless, no hand-drift.
#
#   ./scripts/generate.sh                 # pulls spec from hanzoai/openapi@main
#   SPEC=/path/to/hanzo.yaml ./scripts/generate.sh   # local spec override
#
# Requires: java 17+, curl.
set -euo pipefail
cd "$(dirname "$0")/.."

GENERATOR_VERSION="${GENERATOR_VERSION:-7.14.0}"
SPEC_URL="${SPEC_URL:-https://raw.githubusercontent.com/hanzoai/openapi/main/hanzo.yaml}"
SPEC="${SPEC:-}"
JAR="${JAR:-/tmp/openapi-generator-cli-${GENERATOR_VERSION}.jar}"

if [ -z "$SPEC" ]; then
  SPEC="$(mktemp)"; curl -fsSL "$SPEC_URL" -o "$SPEC"
fi
if [ ! -f "$JAR" ]; then
  curl -fsSL -o "$JAR" \
    "https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/${GENERATOR_VERSION}/openapi-generator-cli-${GENERATOR_VERSION}.jar"
fi

# Delete the previous generation so removed operations/models don't linger.
if [ -f .openapi-generator/FILES ]; then
  while IFS= read -r f; do [ -n "$f" ] && rm -f "./$f"; done < .openapi-generator/FILES
fi

java -jar "$JAR" generate \
  -i "$SPEC" -g go \
  --additional-properties=packageName=hanzoai,isGoSubmodule=false,structPrefix=true,enumClassPrefix=true \
  --global-property=apiTests=false,modelTests=false \
  --git-user-id=hanzoai --git-repo-id=go-sdk \
  -o .

# Pin a modern Go toolchain (generator defaults to 1.18).
go mod edit -go=1.23
gofmt -w . >/dev/null 2>&1 || true
echo "generated $(ls api_*.go model_*.go 2>/dev/null | wc -l) source files"
