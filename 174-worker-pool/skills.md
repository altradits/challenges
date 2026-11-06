# Skills for 174-worker-pool

## What You'll Learn

**Previous:** [173-crypto-sha256](../173-crypto-sha256/skills.md) | **Next:** [175-sync-atomic](../175-sync-atomic/skills.md)

**Challenge:** Fan work out to N goroutines, collect ordered results.

## The Pattern

A worker pool has three parts: a jobs channel, N worker goroutines, and a
way to collect results:

```go
type task struct {
    index int
    job   func() int
}

jobsCh := make(chan task)
results := make([]int, len(jobs))
var wg sync.WaitGroup
for w := 0; w < workers; w++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        for t := range jobsCh { // ends when channel closes
            results[t.index] = t.job()
        }
    }()
}
for i, j := range jobs {
    jobsCh <- task{i, j}
}
close(jobsCh) // workers exit when drained
wg.Wait()
return results
```

Writing `results[i]` from different workers is safe here: each index is
written by exactly one job, and `wg.Wait()` synchronizes before return.
The race detector agrees — shared index, single writer, barrier read.

## Why Each Piece Matters

- **Buffered or not?** An unbuffered jobs channel is fine — the feeder
  blocks until a worker is ready. No deadlock as long as workers run.
- **`close(jobsCh)`** is what lets `range` end. Close from the feeder,
  never from a worker.
- **`wg.Wait()`** before reading `results` — without it you return zeros.
- **`workers < 1`** must mean 1, otherwise zero workers drain nothing and
  the feeder blocks forever.
- **Empty jobs** return an empty non-nil slice; still start (and stop) the
  pool cleanly.

## Leaks and Races

Every `go` statement here must be matched by `wg.Done`, and `Wait` must
run before return — then no goroutine outlives the call. Run the tests
with `-race`: any shared-state mistake shows up immediately.

## Solving the Challenge

Follow the pattern above exactly: indexed tasks through one channel,
results by index, close-then-wait. Order in the output comes from the
index, not from finish order.
