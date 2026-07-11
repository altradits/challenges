# Skills for 168-sync-mutex

## What You'll Learn

**Previous:** [167-io-readers-writers](../167-io-readers-writers/skills.md) | **Next:** [169-testing-advanced](../169-testing-advanced/skills.md)

**Challenge:** Protect shared state with `sync.Mutex` and `sync.RWMutex` to prevent data races.

## The Data Race Problem

When two goroutines access the same variable and at least one writes, you have a **data race** — undefined behaviour:

```go
count := 0
var wg sync.WaitGroup

for i := 0; i < 1000; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        count++  // RACE: read + increment + write is not atomic
    }()
}
wg.Wait()
fmt.Println(count)  // NOT 1000 — random result
```

Detect races:

```bash
go test -race ./...
go run -race main.go
```

## `sync.Mutex` — Mutual Exclusion Lock

A mutex allows only one goroutine to hold the lock at a time:

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

Rules:
- Always use pointer receiver for types with a mutex (`*SafeCounter`, not `SafeCounter`) — copying a mutex is a bug.
- `Lock()` blocks until the mutex is available.
- `defer mu.Unlock()` ensures unlock even if the function panics or returns early.
- Never read without locking — reads are not magically thread-safe.

## `sync.RWMutex` — Read-Write Lock

When reads are much more frequent than writes, use `sync.RWMutex`:

```go
type SafeMap struct {
    mu    sync.RWMutex
    store map[string]int
}

func (m *SafeMap) Set(key string, value int) {
    m.mu.Lock()         // exclusive write lock
    defer m.mu.Unlock()
    m.store[key] = value
}

func (m *SafeMap) Get(key string) (int, bool) {
    m.mu.RLock()         // shared read lock — multiple readers OK
    defer m.mu.RUnlock()
    v, ok := m.store[key]
    return v, ok
}
```

| Operation | Lock method | Other goroutines |
|-----------|-------------|-----------------|
| Write | `Lock()` / `Unlock()` | All blocked |
| Read | `RLock()` / `RUnlock()` | Other readers OK, writers blocked |

## `sync.Map` — Built-In Concurrent Map

For high-concurrency read-heavy maps, Go provides `sync.Map`:

```go
var m sync.Map

m.Store("key", 42)
v, ok := m.Load("key")   // 42, true
m.Delete("key")

m.Range(func(k, v interface{}) bool {
    fmt.Println(k, v)
    return true  // return false to stop iteration
})
```

`sync.Map` is optimised for "many reads, few writes" — not a drop-in replacement for `map` in all cases.

## `sync/atomic` — Lock-Free Atomics

For simple integer operations, atomic is faster than a mutex:

```go
import "sync/atomic"

var count int64

atomic.AddInt64(&count, 1)           // atomic increment
v := atomic.LoadInt64(&count)        // atomic read
atomic.StoreInt64(&count, 0)         // atomic write
old := atomic.SwapInt64(&count, 42)  // atomic swap
```

Use atomics for single-value counters. Use mutexes for protecting multi-field structs.

## Visual: Lock/Unlock Flow

```
Goroutine A          Goroutine B
    |                    |
mu.Lock()           mu.Lock()
    |                    | ← BLOCKED (A holds the lock)
count++                  |
mu.Unlock() ────────────▶ mu.Lock() acquired
                     count++
                     mu.Unlock()
```

## Solving the Challenge

```go
package piscine

import "sync"

type SafeCounter struct {
    mu    sync.Mutex
    count int
}

func NewSafeCounter() *SafeCounter {
    return &SafeCounter{}
}

func (c *SafeCounter) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

func (c *SafeCounter) Add(n int) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count += n
}

func (c *SafeCounter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}

func (c *SafeCounter) Reset() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count = 0
}

type SafeMap struct {
    mu    sync.RWMutex
    store map[string]int
}

func NewSafeMap() *SafeMap {
    return &SafeMap{store: make(map[string]int)}
}

func (m *SafeMap) Set(key string, value int) {
    m.mu.Lock()
    defer m.mu.Unlock()
    m.store[key] = value
}

func (m *SafeMap) Get(key string) (int, bool) {
    m.mu.RLock()
    defer m.mu.RUnlock()
    v, ok := m.store[key]
    return v, ok
}

func (m *SafeMap) Delete(key string) {
    m.mu.Lock()
    defer m.mu.Unlock()
    delete(m.store, key)
}

func (m *SafeMap) Len() int {
    m.mu.RLock()
    defer m.mu.RUnlock()
    return len(m.store)
}
```

## Documentation

- [sync package](https://pkg.go.dev/sync)
- [sync/atomic package](https://pkg.go.dev/sync/atomic)
- [Go Blog — Share memory by communicating](https://go.dev/blog/codelab-share)
- [Go Data Race Detector](https://go.dev/doc/articles/race_detector)

**Next:** [169-testing-advanced](../169-testing-advanced/README.md)
