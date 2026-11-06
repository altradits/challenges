#!/bin/bash
# Undo overlay-solutions.sh using its manifest: files that existed before the
# overlay are restored from backup, other tracked files are restored with git
# checkout, and untracked copies are removed.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
MANIFEST="${OVERLAY_MANIFEST:-/tmp/drill-overlay-manifest.txt}"
BACKUP="${OVERLAY_BACKUP:-/tmp/drill-overlay-backup}"

if [ ! -f "$MANIFEST" ]; then
  echo "no manifest at $MANIFEST (nothing to revert)" >&2
  exit 1
fi

restored=0
removed=0
while IFS= read -r path; do
  [ -e "$path" ] || continue
  rel="${path#$ROOT/}"
  if [ -e "$BACKUP/$rel" ]; then
    cp "$BACKUP/$rel" "$path"
    restored=$((restored + 1))
  elif git -C "$ROOT" ls-files --error-unmatch "$rel" >/dev/null 2>&1; then
    git -C "$ROOT" checkout -- "$rel"
    restored=$((restored + 1))
  else
    rm "$path"
    removed=$((removed + 1))
  fi
done < "$MANIFEST"

rm -f "$MANIFEST"
rm -rf "$BACKUP"
echo "reverted: $restored restored, $removed removed"
