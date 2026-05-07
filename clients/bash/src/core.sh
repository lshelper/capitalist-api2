#!/usr/bin/env bash

CAPITALIST_API2_BASE_URL="${CAPITALIST_API2_BASE_URL:-https://api2.capitalist.net}"

capitalist_api2_now_ms() {
  if command -v gdate >/dev/null 2>&1; then
    gdate +%s%3N
  else
    printf '%s000' "$(date +%s)"
  fi
}

capitalist_api2_sha256_hex() {
  if command -v openssl >/dev/null 2>&1; then
    printf '%s' "$1" | openssl dgst -sha256 -binary | xxd -p -c 256
  elif command -v shasum >/dev/null 2>&1; then
    printf '%s' "$1" | shasum -a 256 | awk '{print $1}'
  else
    printf 'capitalist-api2: openssl or shasum is required\n' >&2
    return 1
  fi
}

capitalist_api2_signature() {
  local timestamp="$1"
  local body="${2:-}"
  local secret="$3"

  capitalist_api2_sha256_hex "${timestamp}${body}${secret}"
}

capitalist_api2_require_credentials() {
  if [ -z "${CAPITALIST_API2_KEY:-}" ] || [ -z "${CAPITALIST_API2_SECRET:-}" ]; then
    printf 'capitalist-api2: set CAPITALIST_API2_KEY and CAPITALIST_API2_SECRET\n' >&2
    return 1
  fi
}

capitalist_api2_urlencode() {
  local value="$1"
  local length="${#value}"
  local i char

  for ((i = 0; i < length; i++)); do
    char="${value:i:1}"
    case "$char" in
      [a-zA-Z0-9.~_-]) printf '%s' "$char" ;;
      *) printf '%%%02X' "'$char" ;;
    esac
  done
}

capitalist_api2_request() {
  local method="$1"
  local path="$2"
  local body="${3:-}"
  local content_type="${4:-application/json}"
  local timestamp signature url

  capitalist_api2_require_credentials || return 1

  timestamp="$(capitalist_api2_now_ms)"
  signature="$(capitalist_api2_signature "$timestamp" "$body" "$CAPITALIST_API2_SECRET")" || return 1
  url="${CAPITALIST_API2_BASE_URL%/}${path}"

  if [ -n "$body" ]; then
    curl -sS -X "$method" "$url" \
      -H "API-Key: ${CAPITALIST_API2_KEY}" \
      -H "X-Request-Timestamp: ${timestamp}" \
      -H "Signature: ${signature}" \
      -H "Content-Type: ${content_type}" \
      --data-binary "$body"
  else
    curl -sS -X "$method" "$url" \
      -H "API-Key: ${CAPITALIST_API2_KEY}" \
      -H "X-Request-Timestamp: ${timestamp}" \
      -H "Signature: ${signature}"
  fi
}

capitalist_api2_request_file() {
  local method="$1"
  local path="$2"
  local file="$3"
  local content_type="${4:-application/octet-stream}"
  local timestamp signature url

  capitalist_api2_require_credentials || return 1

  timestamp="$(capitalist_api2_now_ms)"
  signature="$({ printf '%s' "$timestamp"; cat "$file"; printf '%s' "$CAPITALIST_API2_SECRET"; } | openssl dgst -sha256 -binary | xxd -p -c 256)" || return 1
  url="${CAPITALIST_API2_BASE_URL%/}${path}"

  curl -sS -X "$method" "$url" \
    -H "API-Key: ${CAPITALIST_API2_KEY}" \
    -H "X-Request-Timestamp: ${timestamp}" \
    -H "Signature: ${signature}" \
    -H "Content-Type: ${content_type}" \
    --data-binary "@${file}"
}
