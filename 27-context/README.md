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

---

## Time limit

**15 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 27
```

or run the tests directly:

```bash
go test ./27-context/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 27` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
