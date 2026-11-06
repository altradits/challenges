# 174-worker-pool — Worker Pools with Goroutines

## Task

In `package piscine`, run a batch of jobs on a fixed pool of workers,
returning results in job order:

```go
func RunPool(jobs []func() int, workers int) []int
```

Read `skills.md` before you start.

## Input / Output

- Input: a slice of zero-argument jobs and a worker count.
- Output: results in the same order as `jobs`, whatever order they finish.

## Examples

```
jobs := []func() int{func() int { return 1 }, func() int { return 2 }}
RunPool(jobs, 2)  →  [1 2]
RunPool(nil, 4)   →  []
```

## Constraints

- At most `workers` jobs may run at the same time; `workers < 1` means 1.
- Use channels + `sync.WaitGroup`; close the results path cleanly so no
  goroutine leaks (the tests run with `-race`).
- Jobs must actually run concurrently — a sequential loop fails.

## Time limit

**25 minutes** (round 1). Beat it on every repeat: each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

    go run ./cmd/drill check 174

or run the tests directly:

    go test -race ./174-worker-pool/

Green means done. Red means read the failure message, fix from memory, check again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, beat your previous time. If you cannot, re-read `skills.md` once,
then close it and go again.
