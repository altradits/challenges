# 173-crypto-sha256 — Hashing with crypto/sha256

## Task

In `package piscine`, write SHA-256 helpers — the hash behind block
hashing, txids, and P2P checksums:

```go
func SHA256Hex(s string) string
func SHA256d(b []byte) [32]byte
func HashReader(r io.Reader) ([32]byte, error)
func Checksum4(b []byte) [4]byte
```

Read `skills.md` before you start.

## Input / Output

- `SHA256Hex` returns the lowercase hex digest of `s`.
- `SHA256d` returns SHA-256 applied twice (Bitcoin's hash function).
- `HashReader` streams `r` into a hasher (for data too big for memory).
- `Checksum4` returns the first 4 bytes of `SHA256d(b)` (P2P checksum).

## Examples

```
SHA256Hex("abc")  →  "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"
SHA256d([]byte("hello"))  →  9595c9df...833d50 (32 bytes)
Checksum4([]byte("hello"))  →  [0x95 0x95 0xc9 0xdf]
```

## Constraints

- Use `crypto/sha256` (`Sum256`, `New`, `Hash.Write`, `Hash.Sum`).
- `HashReader` must stream with `io.Copy`, not `io.ReadAll`.

## Time limit

**20 minutes** (round 1). Beat it on every repeat: each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

    go run ./cmd/drill check 173

or run the tests directly:

    go test ./173-crypto-sha256/

Green means done. Red means read the failure message, fix from memory, check again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, beat your previous time. If you cannot, re-read `skills.md` once,
then close it and go again.
