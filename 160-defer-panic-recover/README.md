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
