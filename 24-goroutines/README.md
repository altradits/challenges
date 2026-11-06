# 39-goroutines — Goroutines and WaitGroups

## Challenge

Implement in `package piscine`:

```go
// ConcurrentMap applies fn to each element of data in parallel
// and returns the results in the original order.
func ConcurrentMap(data []int, fn func(int) int) []int

// ParallelSum divides data into numWorkers equal parts,
// sums each part in a goroutine, and returns the total.
func ParallelSum(data []int, numWorkers int) int
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
go run ./cmd/drill check 24
```

or run the tests directly:

```bash
go test ./24-goroutines/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 24` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
