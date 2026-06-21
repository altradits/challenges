# 42-context — The context Package

## Challenge

Implement in `package piscine`:

```go
// FetchWithTimeout simulates a slow operation.
// It returns "done" after workDuration, or an error if ctx is cancelled first.
func FetchWithTimeout(ctx context.Context, workDuration time.Duration) (string, error)

// ProcessAll processes each item in items using fn.
// Stops early and returns ctx.Err() if the context is cancelled mid-way.
func ProcessAll(ctx context.Context, items []string, fn func(string)) error
```

Read `skills.md` before you start.
