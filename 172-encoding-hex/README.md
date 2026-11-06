# 172-encoding-hex — Hex Encoding with encoding/hex

## Task

In `package piscine`, write hex helpers with `encoding/hex` — the way
txids, keys, and scripts are printed everywhere in Bitcoin:

```go
func ToHex(b []byte) string
func FromHex(s string) ([]byte, error)
func IsHex(s string) bool
func ReverseHex(s string) (string, error)
```

Read `skills.md` before you start.

## Input / Output

- `ToHex` lowercases hex-encode bytes; `FromHex` decodes, erroring on
  odd length or non-hex digits.
- `IsHex` reports whether `s` is valid even-length hex (empty is valid).
- `ReverseHex` reverses the *byte* order of a hex string
  (`"0102ff"` → `"ff0201"`), erroring on invalid hex.

## Examples

```
ToHex([]byte{0xde, 0xad})  →  "dead"
FromHex("dead")             →  [0xde 0xad], nil
FromHex("zz")               →  nil, error
IsHex("0123")               →  true
ReverseHex("0102ff")        →  "ff0201", nil
```

## Constraints

- Use `hex.EncodeToString` / `hex.DecodeString` — no manual nibble math.
- `ToHex(nil)` → `""`; `FromHex("")` → empty slice, nil error.

## Time limit

**15 minutes** (round 1). Beat it on every repeat: each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

    go run ./cmd/drill check 172

or run the tests directly:

    go test ./172-encoding-hex/

Green means done. Red means read the failure message, fix from memory, check again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, beat your previous time. If you cannot, re-read `skills.md` once,
then close it and go again.
