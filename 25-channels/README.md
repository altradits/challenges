# 25-channels — Channels

## Task

Implement in `package piscine`:

```go
// Generator returns a channel that emits 0, 1, 2, ... up to n-1, then closes.
func Generator(n int) <-chan int

// Merge merges two channels into one, closing the output when both inputs close.
func Merge(a, b <-chan int) <-chan int

// WorkerPool processes jobs from the jobs channel using numWorkers goroutines.
// Each worker applies fn to the job and sends the result to the returned channel.
func WorkerPool(jobs <-chan int, numWorkers int, fn func(int) int) <-chan int
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
go run ./cmd/drill check 25
```

or run the tests directly:

```bash
go test ./25-channels/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 25` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
