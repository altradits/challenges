# 175-sync-atomic — Lock-Free Counters with sync/atomic

## Task

In `package piscine`, build a concurrent counter without a mutex, plus a
compare-and-swap helper:

```go
type Counter struct { ... } // your fields, unexported
func NewCounter() *Counter
func (c *Counter) Add(n int64) int64  // adds n, returns the new value
func (c *Counter) Value() int64
func CASAdd(addr *int64, delta int64) int64 // same via CompareAndSwap loop
```

Read `skills.md` before you start.

## Input / Output

- `Add` atomically adds (negative `n` subtracts) and returns the result.
- `CASAdd` does the same for a plain `*int64` using only
  `atomic.CompareAndSwapInt64` in a retry loop — no `atomic.AddInt64`.

## Examples

```
c := NewCounter()
c.Add(5)   →  5
c.Add(-2)  →  3
c.Value()  →  3
var x int64
CASAdd(&x, 10)  →  10
```

## Constraints

- No `sync.Mutex` anywhere — the point is lock-free atomics.
- `Counter` must be safe for concurrent use (tests hammer it with `-race`).
- `CASAdd` must loop until its swap wins (a single try loses updates).

## Time limit

**20 minutes** (round 1). Beat it on every repeat: each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

    go run ./cmd/drill check 175

or run the tests directly:

    go test -race ./175-sync-atomic/

Green means done. Red means read the failure message, fix from memory, check again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, beat your previous time. If you cannot, re-read `skills.md` once,
then close it and go again.
