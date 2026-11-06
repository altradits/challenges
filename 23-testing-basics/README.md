# 38-testing-basics — Testing in Go

## Challenge

Write a complete `solution_test.go` for the following functions (which you also implement):

```go
// Abs returns the absolute value of n.
func Abs(n int) int

// Palindrome reports whether s reads the same forwards and backwards.
func Palindrome(s string) bool

// FizzBuzz returns "Fizz" (div by 3), "Buzz" (div by 5), "FizzBuzz" (both), else the number as string.
func FizzBuzz(n int) string
```

Your test file must include:
- At least one table-driven test with `t.Run`
- At least one test for an edge case (empty string, zero, negative)
- At least one benchmark (`func BenchmarkXxx`)

Read `skills.md` before you start.

---

## Time limit

**15 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 23
```

or run the tests directly:

```bash
go test ./23-testing-basics/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 23` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
