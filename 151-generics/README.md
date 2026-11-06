# 151-generics — Generics (Go 1.18+)

## Task

Implement in `package piscine`:

```go
// Map applies fn to each element of s and returns a new slice of results.
func Map[T, U any](s []T, fn func(T) U) []U

// Filter returns elements of s for which fn returns true.
func Filter[T any](s []T, fn func(T) bool) []T

// Reduce folds s into a single value using fn, starting from initial.
func Reduce[T, U any](s []T, initial U, fn func(U, T) U) U

// Keys returns all keys of a map in an unspecified order.
func Keys[K comparable, V any](m map[K]V) []K

// Contains reports whether target is in s.
func Contains[T comparable](s []T, target T) bool
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
go run ./cmd/drill check 151
```

or run the tests directly:

```bash
go test ./151-generics/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 151` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
