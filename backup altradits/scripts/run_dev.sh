#!/bin/bash
set -e

echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "  Altradits Dev Server"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

export DB_URL="${DB_URL:-postgres://altradits:password@localhost:5432/altradits?sslmode=disable}"
export PORT="${PORT:-8080}"

echo "  DB:   $DB_URL"
echo "  Port: $PORT"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

go run cmd/server/main.go
