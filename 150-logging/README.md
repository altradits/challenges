# 150-logging — Logging in Go

## Task

Implement in `package piscine`:

```go
// NewLogger returns a *slog.Logger that writes JSON to w at the given level.
// level is one of: "debug", "info", "warn", "error"
func NewLogger(w io.Writer, level string) *slog.Logger

// RequestLogger returns an http.Handler that wraps next,
// logging method, path, and duration for every request.
func RequestLogger(logger *slog.Logger, next http.Handler) http.Handler
```

Read `skills.md` before you start.

---

## Time limit

**25 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 150
```

or run the tests directly:

```bash
go test ./150-logging/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 150` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
