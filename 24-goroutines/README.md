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
