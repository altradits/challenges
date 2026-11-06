# 21-sort-and-math — sort and math Packages

## Task

Implement in `package piscine`:

```go
// MedianSorted returns the median of a slice after sorting it.
// Returns 0 for empty slices.
func MedianSorted(nums []float64) float64

// TopN returns the n largest integers from the slice (sorted descending).
func TopN(nums []int, n int) []int

// Hypotenuse returns the length of the hypotenuse given legs a and b.
func Hypotenuse(a, b float64) float64

// RoundToNearest rounds f to the nearest 0.5.
func RoundToNearest(f float64) float64
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
go run ./cmd/drill check 21
```

or run the tests directly:

```bash
go test ./21-sort-and-math/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 21` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
