# 171-encoding-binary — Byte Order with encoding/binary

## Task

In `package piscine`, write little-endian and big-endian integer codecs
with `encoding/binary`, plus a packed record type:

```go
func EncodeUint32LE(v uint32) []byte
func DecodeUint32LE(b []byte) uint32
func EncodeUint64BE(v uint64) []byte
func DecodeUint64BE(b []byte) uint64
func PackRecord(id uint32, score uint16) []byte
func UnpackRecord(b []byte) (uint32, uint16, error)
```

Read `skills.md` before you start.

## Input / Output

- `Encode*` return exactly 4 / 8 bytes; `Decode*` read the first 4 / 8.
- `PackRecord` returns 6 bytes: `id` little-endian, then `score`
  little-endian. `UnpackRecord` reverses it and errors when `len(b) != 6`.

## Examples

```
EncodeUint32LE(0x01020304)  →  [0x04 0x03 0x02 0x01]
DecodeUint32LE(...)         →  0x01020304
EncodeUint64BE(1)           →  [0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x01]
PackRecord(0x01020304, 0x0506)  →  [0x04 0x03 0x02 0x01 0x06 0x05]
```

## Constraints

- Use `binary.LittleEndian` / `binary.BigEndian` (`PutUint32`,
  `Uint32`, …) — no manual bit shifting.
- `Decode*` on short input must not panic (define: missing bytes read
  as zero — document it in a comment).

## Time limit

**20 minutes** (round 1). Beat it on every repeat: each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

    go run ./cmd/drill check 171

or run the tests directly:

    go test ./171-encoding-binary/

Green means done. Red means read the failure message, fix from memory, check again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, beat your previous time. If you cannot, re-read `skills.md` once,
then close it and go again.
