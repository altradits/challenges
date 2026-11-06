# 20-regexp — Regular Expressions

## Task

Implement in `package piscine`:

```go
// ValidateEmail returns true if s looks like a valid email address.
func ValidateEmail(s string) bool

// ExtractNumbers returns all integer sequences found in s.
// "foo123bar456" → ["123", "456"]
func ExtractNumbers(s string) []string

// MaskPhone replaces digits in a phone number with *, keeping the last 4.
// "0712345678" → "******5678"
func MaskPhone(s string) string
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
go run ./cmd/drill check 20
```

or run the tests directly:

```bash
go test ./20-regexp/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 20` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
