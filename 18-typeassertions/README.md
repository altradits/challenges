# 33-typeassertions — Type Assertions and Type Switches

## Challenge

Implement in `package piscine`:

```go
// Describe returns a human-readable description of any value.
// int → "integer: 42"
// string → "string: hello"
// bool → "bool: true"
// []int → "int slice of length 3"
// anything else → "unknown type"
func Describe(v any) string

// Sum adds together any mix of ints and float64s.
// Ignores values of other types.
func Sum(values []any) float64
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
go run ./cmd/drill check 18
```

or run the tests directly:

```bash
go test ./18-typeassertions/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 18` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
