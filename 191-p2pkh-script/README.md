# 191-p2pkh-script — Pay-to-PubKey-Hash Scripts

## Task

In `package piscine`, build and recognize P2PKH locking scripts — the
most common scriptPubKey on chain:

```go
func BuildP2PKH(pubkeyHash []byte) ([]byte, error)
func ParseP2PKH(script []byte) ([]byte, error)
func IsP2PKH(script []byte) bool
```

Read `skills.md` before you start.

## Input / Output

- `BuildP2PKH` wraps a 20-byte hash in the P2PKH template (25 bytes).
- `ParseP2PKH` validates the template and returns the inner hash.
- `IsP2PKH` reports template match without allocating errors.

## Examples

```
h := []byte{0x00,0x11,...} // 20 bytes
BuildP2PKH(h)  →  [0x76 0xa9 0x14 <20 hash bytes> 0x88 0xac]
ParseP2PKH(script)  →  the 20 hash bytes, nil
IsP2PKH(script)     →  true
IsP2PKH([]byte{0x76})  →  false
```

## Constraints

- `BuildP2PKH` errors unless `len(pubkeyHash) == 20`.
- `ParseP2PKH` errors on any template mismatch (length, opcodes, push).
- `ParseP2PKH` must return a copy — mutating it must not corrupt the input.

## Time limit

**20 minutes** (round 1). Beat it on every repeat: each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

    go run ./cmd/drill check 191

or run the tests directly:

    go test ./191-p2pkh-script/

Green means done. Red means read the failure message, fix from memory, check again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, beat your previous time. If you cannot, re-read `skills.md` once,
then close it and go again.
