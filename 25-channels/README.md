# 40-channels — Channels

## Challenge

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
