# 32-errorhandling — Error Handling in Go

## Challenge

Implement in `package piscine`:

```go
// Divide returns a/b or an error if b is zero.
func Divide(a, b float64) (float64, error)

// ParseAge parses a string to a non-negative int.
// Returns ErrNegativeAge if the number is negative.
// Returns a wrapped strconv error if the string is not a number.
func ParseAge(s string) (int, error)

// Custom error type
type ValidationError struct {
    Field   string
    Message string
}
func (e *ValidationError) Error() string  // "field: message"
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
go run ./cmd/drill check 17
```

or run the tests directly:

```bash
go test ./17-errorhandling/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 17` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
