# 160-defer-panic-recover — Defer, Panic, and Recover

## Challenge

Implement in `package piscine`:

```go
// SafeDivide returns a/b. If b is 0, it recovers from the panic and returns 0, error.
func SafeDivide(a, b int) (result int, err error)

// WithCleanup calls fn(), then always calls cleanup() — even if fn panics.
// If fn panics, WithCleanup returns the panic value as an error; otherwise nil.
func WithCleanup(fn func(), cleanup func()) (err error)

// DeferOrder appends strings to log in the order they will actually print
// given three defers: "first", "second", "third" registered in that order.
// Returns the log slice (demonstrates LIFO ordering).
func DeferOrder() []string
```

## Examples

```
SafeDivide(10, 2)  →  5, nil
SafeDivide(10, 0)  →  0, error("cannot divide by zero")

// WithCleanup always runs cleanup
log := []string{}
WithCleanup(
    func() { log = append(log, "work") },
    func() { log = append(log, "cleaned") },
)
// log == ["work", "cleaned"]

DeferOrder()  →  ["third", "second", "first"]
```

Read `skills.md` before you start.

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 160
```

or run the tests directly:

```bash
go test ./160-defer-panic-recover/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 160` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
