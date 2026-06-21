# Skills for 41-select-and-sync

## What You'll Learn

**Previous:** [25-channels](../25-channels/skills.md) | **Next:** [27-context](../27-context/skills.md)

**Challenge:** Protect shared state with mutexes; coordinate with select and sync primitives.

## `select` Statement

`select` lets a goroutine wait on multiple channel operations simultaneously:

```go
select {
case v := <-ch1:
    fmt.Println("received from ch1:", v)
case v := <-ch2:
    fmt.Println("received from ch2:", v)
case ch3 <- 42:
    fmt.Println("sent to ch3")
}
```

If multiple cases are ready, one is chosen at random.

### Default Case (Non-Blocking)

```go
select {
case v := <-ch:
    fmt.Println("got", v)
default:
    fmt.Println("no value ready")  // runs immediately if ch is not ready
}
```

### Timeout Pattern

```go
import "time"

select {
case result := <-ch:
    fmt.Println("result:", result)
case <-time.After(2 * time.Second):
    fmt.Println("timed out")
}
```

`time.After(d)` returns a channel that sends one value after duration `d`.

## `sync.Mutex` — Mutual Exclusion

Protects a shared variable from concurrent reads/writes:

```go
import "sync"

type SafeCounter struct {
    mu    sync.Mutex
    count int
}

func (c *SafeCounter) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

func (c *SafeCounter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}
```

Always use `defer mu.Unlock()` immediately after `Lock()` — ensures unlock even on panic/early return.

### `sync.RWMutex` — Multiple Readers, One Writer

```go
var rw sync.RWMutex

// Read lock (many goroutines can hold it simultaneously)
rw.RLock()
// ... read shared data ...
rw.RUnlock()

// Write lock (exclusive)
rw.Lock()
// ... modify shared data ...
rw.Unlock()
```

Use when reads are far more frequent than writes.

## `sync.Once` — Run Exactly Once

```go
var once sync.Once
var instance *DB

func GetDB() *DB {
    once.Do(func() {
        instance = &DB{...}  // runs exactly once, regardless of how many goroutines call GetDB
    })
    return instance
}
```

Thread-safe singleton pattern.

## `sync.Map` — Concurrent Map

The built-in `map` is NOT safe for concurrent use. Use `sync.Map` when multiple goroutines read and write:

```go
var m sync.Map

m.Store("key", "value")

v, ok := m.Load("key")

m.Range(func(k, v any) bool {
    fmt.Println(k, v)
    return true  // return false to stop iteration
})
```

Prefer a `map` with `sync.Mutex` for most cases — `sync.Map` has narrower use cases.

## Solving the Challenge

```go
package piscine

import (
    "errors"
    "sync"
    "time"
)

var ErrTimeout = errors.New("timeout")

type SafeCounter struct {
    mu    sync.Mutex
    count int
}

func NewSafeCounter() *SafeCounter { return &SafeCounter{} }

func (c *SafeCounter) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

func (c *SafeCounter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}

func Timeout(fn func() string, d time.Duration) (string, error) {
    ch := make(chan string, 1)
    go func() { ch <- fn() }()
    select {
    case result := <-ch:
        return result, nil
    case <-time.After(d):
        return "", ErrTimeout
    }
}

func NewOnce(fn func() string) func() string {
    var once sync.Once
    var result string
    return func() string {
        once.Do(func() { result = fn() })
        return result
    }
}
```

**Next:** [27-context](../27-context/README.md)
