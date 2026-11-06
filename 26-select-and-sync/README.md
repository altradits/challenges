# 26-select-and-sync — select, Mutex, and sync Patterns

## Task

Implement in `package piscine`:

```go
// SafeCounter is a goroutine-safe counter.
type SafeCounter struct { /* unexported fields */ }

func NewSafeCounter() *SafeCounter
func (c *SafeCounter) Inc()
func (c *SafeCounter) Value() int

// Timeout runs fn in a goroutine and returns its result,
// or returns ("", ErrTimeout) if fn takes longer than d.
func Timeout(fn func() string, d time.Duration) (string, error)

// Once runs fn only the first time it is called, no matter how many goroutines call it.
// Returns the result of the first call every time.
func NewOnce(fn func() string) func() string
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
go run ./cmd/drill check 26
```

or run the tests directly:

```bash
go test ./26-select-and-sync/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 26` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
