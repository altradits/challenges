# Skills for 40-channels

## What You'll Learn

**Previous:** [24-goroutines](../24-goroutines/skills.md) | **Next:** [26-select-and-sync](../26-select-and-sync/skills.md)

**Challenge:** Pass data between goroutines safely using channels.

## Channels

A channel is a typed pipe for communication between goroutines.

```go
ch := make(chan int)       // unbuffered
ch := make(chan int, 10)   // buffered (capacity 10)
```

### Sending and Receiving

```go
ch <- 42      // send (blocks if channel is full)
v := <-ch     // receive (blocks until a value is available)
v, ok := <-ch // ok = false when channel is closed and empty
```

### Unbuffered vs Buffered

```go
// Unbuffered: sender blocks until receiver is ready (synchronous rendezvous)
ch := make(chan int)
go func() { ch <- 42 }()  // blocks until main receives
fmt.Println(<-ch)           // 42

// Buffered: sender only blocks when buffer is full
ch := make(chan int, 3)
ch <- 1
ch <- 2
ch <- 3
// ch <- 4  would block — buffer full
```

### Closing a Channel

```go
close(ch)
```

- Only the **sender** should close
- Sending to a closed channel panics
- Receiving from a closed, empty channel returns the zero value + `ok=false`

### Ranging Over a Channel

```go
ch := make(chan int, 5)
for i := 0; i < 5; i++ { ch <- i }
close(ch)

for v := range ch {  // stops when ch is closed and drained
    fmt.Println(v)
}
```

### Directional Channel Types

```go
func producer(out chan<- int) { out <- 42 }   // can only send
func consumer(in <-chan int)  { fmt.Println(<-in) }  // can only receive
```

Using directional types in function signatures makes intent clear and catches bugs at compile time.

### Pipeline Pattern

```go
func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

for v := range square(generate(2, 3, 4)) {
    fmt.Println(v)  // 4, 9, 16
}
```

### Solving the Challenge

```go
package piscine

import "sync"

func Generator(n int) <-chan int {
    ch := make(chan int)
    go func() {
        for i := 0; i < n; i++ {
            ch <- i
        }
        close(ch)
    }()
    return ch
}

func Merge(a, b <-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    forward := func(ch <-chan int) {
        defer wg.Done()
        for v := range ch {
            out <- v
        }
    }
    wg.Add(2)
    go forward(a)
    go forward(b)
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}

func WorkerPool(jobs <-chan int, numWorkers int, fn func(int) int) <-chan int {
    results := make(chan int)
    var wg sync.WaitGroup
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := range jobs {
                results <- fn(j)
            }
        }()
    }
    go func() {
        wg.Wait()
        close(results)
    }()
    return results
}
```

**Next:** [26-select-and-sync](../26-select-and-sync/README.md)
