#!/usr/bin/env bash
set -eu

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
DOC_URL="https://docs.capitalist.net/api/integration-api.html"
OUT_DIR="${ROOT_DIR}/docs/upstream"
HTML_FILE="${OUT_DIR}/integration-api.html"
MD_FILE="${OUT_DIR}/integration-api.md"

if ! command -v pandoc >/dev/null 2>&1; then
  printf 'pandoc is required to convert the upstream HTML snapshot to Markdown\n' >&2
  exit 1
fi

mkdir -p "$OUT_DIR"

curl -L --compressed -sS "$DOC_URL" -o "$HTML_FILE"
perl -0pi -e 's/data-cfemail="[^"]+"/data-cfemail="__cfemail__"/g' "$HTML_FILE"
pandoc -f html -t gfm --wrap=none "$HTML_FILE" -o "$MD_FILE"

printf 'Updated upstream documentation snapshot from %s\n' "$DOC_URL"
printf 'Review changes with: git diff -- docs/upstream\n'
