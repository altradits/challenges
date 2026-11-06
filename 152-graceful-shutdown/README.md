# 152-graceful-shutdown — Graceful HTTP Shutdown

## Challenge

Implement in `package piscine`:

```go
// NewServer builds an http.Server with the given handler and addr.
func NewServer(addr string, handler http.Handler) *http.Server

// RunWithGracefulShutdown starts srv and blocks until SIGINT or SIGTERM is received.
// It then calls srv.Shutdown with a 10-second deadline, waiting for in-flight requests
// to complete before returning.
func RunWithGracefulShutdown(srv *http.Server) error
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
go run ./cmd/drill check 152
```

or run the tests directly:

```bash
go test ./152-graceful-shutdown/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 152` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
