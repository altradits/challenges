# 187-sha256d-txid — Double-SHA256 and Txid Display Order

## Task

In `package piscine`, implement Bitcoin's hash function and the txid
display convention that trips up every newcomer:

```go
func DoubleSHA(b []byte) [32]byte
func TxidDisplay(hash [32]byte) string
func TxidParse(s string) ([32]byte, error)
```

Read `skills.md` before you start.

## Input / Output

- `DoubleSHA` returns SHA-256(SHA-256(b)).
- `TxidDisplay` renders the internal hash as humans see it: bytes
  reversed, hex-encoded.
- `TxidParse` reverses the process: 64 hex chars → internal byte order.

## Examples

```
DoubleSHA([]byte("hello"))  →  9595c9df...833d50 (internal order)
TxidDisplay(DoubleSHA([]byte("hello")))  →  "503d8319a48348cdc610a582f7bf754b5833df65038606eb48510790dfc99595"
TxidParse("503d83...9595")  →  the 32 internal bytes, nil
TxidParse("zz")             →  error
```

## Constraints

- `TxidParse` requires exactly 64 hex characters — anything else errors.
- Round-trip: `TxidParse(TxidDisplay(h))` must equal `h` for any `h`.

## Time limit

**20 minutes** (round 1). Beat it on every repeat: each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

    go run ./cmd/drill check 187

or run the tests directly:

    go test ./187-sha256d-txid/

Green means done. Red means read the failure message, fix from memory, check again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, beat your previous time. If you cannot, re-read `skills.md` once,
then close it and go again.
