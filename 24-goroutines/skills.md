# Skills for 39-goroutines

## What You'll Learn

**Previous:** [23-testing-basics](../23-testing-basics/skills.md) | **Next:** [25-channels](../25-channels/skills.md)

**Challenge:** Run work concurrently with goroutines; coordinate with WaitGroup.

## Goroutines

A goroutine is a lightweight thread managed by the Go runtime. Start one with `go`:

```go
go someFunction()
go func() {
    fmt.Println("running concurrently")
}()
```

Goroutines are cheap — you can have thousands. They start with ~8KB of stack, which grows as needed.

### The Problem: main() Doesn't Wait

```go
func main() {
    go fmt.Println("hello")
    // main exits before the goroutine runs!
}
```

You must synchronize.

## `sync.WaitGroup`

A WaitGroup waits for a collection of goroutines to finish:

```go
import "sync"

var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)  // tell the WaitGroup: one more goroutine to wait for
    go func(n int) {
        defer wg.Done()  // signal "this goroutine is done"
        fmt.Println(n)
    }(i)       // pass i as argument — avoid closure capture bug
}

wg.Wait()  // block until all goroutines call Done()
```

**Critical: capture loop variables explicitly**

```go
// BUG: all goroutines print the same final value of i
for i := 0; i < 5; i++ {
    go func() { fmt.Println(i) }()
}

// FIX: pass i as a parameter to create a copy
for i := 0; i < 5; i++ {
    go func(n int) { fmt.Println(n) }(i)
}
```

## Data Races

When multiple goroutines read and write the same variable without synchronization, you have a data race — undefined behavior:

```go
// DATA RACE
count := 0
var wg sync.WaitGroup
for i := 0; i < 100; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        count++  // race: multiple goroutines write count simultaneously
    }()
}
wg.Wait()
fmt.Println(count)  // probably not 100
```

Detect races at test time:

```bash
go test -race ./...
go run -race main.go
```

Fix with a mutex or channel (see `41-select-and-sync`).

## GOMAXPROCS

By default, Go uses all available CPU cores. Set manually with:

```go
import "runtime"
runtime.GOMAXPROCS(4)  // use 4 OS threads
```

## Solving the Challenge

```go
package piscine

import "sync"

func ConcurrentMap(data []int, fn func(int) int) []int {
    results := make([]int, len(data))  // pre-allocate, indexed writes are safe
    var wg sync.WaitGroup
    for i, v := range data {
        wg.Add(1)
        go func(idx, val int) {
            defer wg.Done()
            results[idx] = fn(val)  // each goroutine writes to a different index
        }(i, v)
    }
    wg.Wait()
    return results
}

func ParallelSum(data []int, numWorkers int) int {
    if numWorkers <= 0 || len(data) == 0 {
        return 0
    }
    chunkSize := (len(data) + numWorkers - 1) / numWorkers
    var mu sync.Mutex
    total := 0
    var wg sync.WaitGroup
    for i := 0; i < numWorkers; i++ {
        start := i * chunkSize
        if start >= len(data) {
            break
        }
        end := start + chunkSize
        if end > len(data) {
            end = len(data)
        }
        wg.Add(1)
        go func(chunk []int) {
            defer wg.Done()
            sum := 0
            for _, v := range chunk {
                sum += v
            }
            mu.Lock()
            total += sum
            mu.Unlock()
        }(data[start:end])
    }
    wg.Wait()
    return total
}
```

**Next:** [25-channels](../25-channels/README.md)
