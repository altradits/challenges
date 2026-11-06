# 189-block-header — Parsing a Bitcoin Block Header

## Task

In `package piscine`, parse, serialize, and hash an 80-byte Bitcoin
block header:

```go
type Header struct {
    Version    int32
    PrevBlock  [32]byte
    MerkleRoot [32]byte
    Timestamp  uint32
    Bits       uint32
    Nonce      uint32
}

func ParseHeader(b []byte) (Header, error)
func SerializeHeader(h Header) []byte
func HeaderHash(h Header) [32]byte
```

Read `skills.md` before you start.

## Input / Output

- `ParseHeader` decodes exactly 80 bytes (error otherwise).
- `SerializeHeader` encodes back to the identical 80 bytes.
- `HeaderHash` returns the double-SHA256 of the serialization
  (internal byte order).

## Examples

The genesis header (version 1, timestamp 1231006505, bits `0x1d00ffff`,
nonce 2083236893) must round-trip byte-identically, and:

```
HeaderHash(genesis)  →  6fe28c0ab6f1b372c1a6a246ae63f74f931e8365e15a089c68d6190000000000
```

(displayed reversed: `000000000019d668...a8ce26f`)

## Constraints

- All integers are little-endian; hashes stay in internal order.
- `ParseHeader` must reject any length other than 80.

## Time limit

**30 minutes** (round 1). Beat it on every repeat: each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

    go run ./cmd/drill check 189

or run the tests directly:

    go test ./189-block-header/

Green means done. Red means read the failure message, fix from memory, check again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, beat your previous time. If you cannot, re-read `skills.md` once,
then close it and go again.
