# 188-merkle-root — Tiny Merkle Root

## Task

In `package piscine`, compute a Bitcoin Merkle root from transaction
hashes (internal byte order):

```go
func MerkleRoot(txids [][32]byte) [32]byte
```

Read `skills.md` before you start.

## Input / Output

- Input: transaction hashes in internal (wire) byte order.
- Output: the Merkle root, also in internal order.

## Examples

With `a = DoubleSHA("tx0")`, `b = DoubleSHA("tx1")`:

```
MerkleRoot(nil)     →  32 zero bytes
MerkleRoot([a])     →  a
MerkleRoot([a, b])  →  f7ed2a56db7927143defd32e380fb083f74970e0d6f31c00b8aded86aefdf1d2
```

## Constraints

- Empty input returns the zero hash (no leaves, no root).
- A lone leaf at any level is duplicated before hashing (Bitcoin rule).
- Hash pairs as `DoubleSHA(left || right)` in internal order — never
  reverse bytes mid-tree.

## Time limit

**25 minutes** (round 1). Beat it on every repeat: each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

    go run ./cmd/drill check 188

or run the tests directly:

    go test ./188-merkle-root/

Green means done. Red means read the failure message, fix from memory, check again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, beat your previous time. If you cannot, re-read `skills.md` once,
then close it and go again.
