#!/usr/bin/env bash
# The call site. Nothing about HOW this SDK is generated lives here.
#
# The invocation is logic and lives once, in `generate.py`; every per-language
# knob is data in `sdks.yaml` beside it. This file says "go, into this checkout"
# and nothing else. What stood here was a SECOND DRIVER — 250 lines re-deriving
# the jar, the locked ref, the digest check, the YAML-to-JSON conversion and the
# drift check, all of which the canonical driver already does. Two copies of one
# contract is one copy too many, and the two had started to differ.
#
#   ./scripts/generate.sh              # regenerate the client at the module root
#   ./scripts/generate.sh --check      # fail if the committed client has drifted
#
# BOTH INPUTS ARRIVE AS VALUES. $SPEC is the document, already fetched at the
# ref `.spec-lock` pins and digest-checked; $OPENAPI is the checkout holding the
# driver. hanzoai/ci's client lane sets both, because it holds the one credential
# that reads this forge.
#
# uv rather than a bare python3: the driver needs PyYAML and the runner image
# promises no interpreter at all, let alone one with it installed.
#
# Requires: java 17+, uv, and gofmt (the row formats what the generator emits).
set -euo pipefail
cd "$(dirname "$0")/.."

: "${OPENAPI:?the generator lives in hanzoai/openapi; the ci client lane sets OPENAPI, or point it at a checkout}"

if [ -n "${SPEC:-}" ]; then set -- --spec "$SPEC" "$@"; fi

exec uv run --with pyyaml python3 "$OPENAPI/generate.py" go --repo "$PWD" "$@"
