#!/bin/bash
# Copy every reference solution into its challenge directory so the hidden
# tests run against a known-good implementation. Used by CI and by
# contributors verifying tests locally. Mapping: solutions/<path>/*.go ->
# <path>/*.go (top-level challenges and 156-api-lesson sub-exercises).
# Pre-existing files are backed up and restored by revert-overlay.sh.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
MANIFEST="${OVERLAY_MANIFEST:-/tmp/drill-overlay-manifest.txt}"
BACKUP="${OVERLAY_BACKUP:-/tmp/drill-overlay-backup}"
: > "$MANIFEST"
rm -rf "$BACKUP"
mkdir -p "$BACKUP"

count=0
backed=0
while IFS= read -r src; do
  rel="${src#$ROOT/solutions/}"
  dest="$ROOT/$rel"
  mkdir -p "$(dirname "$dest")"
  if [ -e "$dest" ]; then
    mkdir -p "$BACKUP/$(dirname "$rel")"
    cp "$dest" "$BACKUP/$rel"
    backed=$((backed + 1))
  fi
  cp "$src" "$dest"
  echo "$dest" >> "$MANIFEST"
  count=$((count + 1))
done < <(find "$ROOT/solutions" -name '*.go' | sort)

echo "overlaid $count solution files ($backed pre-existing files backed up; manifest: $MANIFEST)"
