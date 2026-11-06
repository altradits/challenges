# 186-compactsize-varint — Bitcoin CompactSize Varints

## Task

In `package piscine`, encode and decode Bitcoin's CompactSize unsigned
varint — the length prefix on every vector in the protocol:

```go
func EncodeVarint(n uint64) []byte
func DecodeVarint(b []byte) (uint64, int, error)
```

Read `skills.md` before you start.

## Input / Output

- `EncodeVarint` returns the shortest encoding of `n`.
- `DecodeVarint` returns the value, the bytes consumed, and an error on
  short input or non-canonical (non-shortest) encoding.

## Examples

```
EncodeVarint(252)        →  [0xfc]
EncodeVarint(253)        →  [0xfd 0xfd 0x00]
EncodeVarint(0x10000)    →  [0xfe 0x00 0x00 0x01 0x00]
DecodeVarint([0xfd 0x00 0x00], ...)  →  error (must be [0x00])
```

## Constraints

- Canonical only: reject values that could have used a shorter form.
- `DecodeVarint` must never panic on short input — return an error.
- No `encoding/binary` varint helpers (`Uvarint` is a different format);
  use fixed-width `Put`/`Uint` calls if you like.

## Time limit

**25 minutes** (round 1). Beat it on every repeat: each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

    go run ./cmd/drill check 186

or run the tests directly:

    go test ./186-compactsize-varint/

Green means done. Red means read the failure message, fix from memory, check again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, beat your previous time. If you cannot, re-read `skills.md` once,
then close it and go again.
