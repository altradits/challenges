# 162-variadic-functions — Variadic Functions

## Challenge

Implement in `package piscine`:

```go
// Sum returns the sum of all provided integers.
func Sum(nums ...int) int

// Max returns the largest of all provided integers.
// Panics if no arguments are given.
func Max(nums ...int) int

// Join concatenates strings with sep between each pair.
func Join(sep string, parts ...string) string

// Wrap calls fn for each argument, collecting return values into a slice.
func Wrap(fn func(int) int, nums ...int) []int
```

## Examples

```
Sum(1, 2, 3, 4, 5)          →  15
Sum()                        →  0

Max(3, 1, 4, 1, 5, 9, 2)   →  9

Join(", ", "a", "b", "c")  →  "a, b, c"
Join("-", "go", "is", "fun") →  "go-is-fun"

double := func(n int) int { return n * 2 }
Wrap(double, 1, 2, 3)       →  [2, 4, 6]
```

Read `skills.md` before you start.

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 162
```

or run the tests directly:

```bash
go test ./162-variadic-functions/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 162` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
