# 159-closures — Closures and First-Class Functions

## Challenge

Implement in `package piscine`:

```go
// MakeCounter returns a function that increments and returns an internal count
// each time it is called, starting from 0.
func MakeCounter() func() int

// MakeAdder returns a function that adds n to any argument passed to it.
func MakeAdder(n int) func(int) int

// Filter returns a new slice containing only elements for which keep returns true.
func Filter(nums []int, keep func(int) bool) []int

// Compose returns a new function that applies g first, then f (f(g(x))).
func Compose(f, g func(int) int) func(int) int
```

## Examples

```
counter := MakeCounter()
counter()  →  1
counter()  →  2
counter()  →  3

add5 := MakeAdder(5)
add5(3)   →  8
add5(10)  →  15

Filter([]int{1,2,3,4,5}, func(n int) bool { return n%2==0 })  →  [2 4]

double := func(x int) int { return x * 2 }
inc    := func(x int) int { return x + 1 }
Compose(double, inc)(3)  →  8   // double(inc(3)) = double(4) = 8
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
go run ./cmd/drill check 159
```

or run the tests directly:

```bash
go test ./159-closures/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 159` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
