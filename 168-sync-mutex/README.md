# 168-sync-mutex — Mutexes and Safe Concurrency

## Challenge

Implement in `package piscine`:

```go
// SafeCounter is a goroutine-safe counter.
type SafeCounter struct { /* unexported fields */ }

func NewSafeCounter() *SafeCounter
func (c *SafeCounter) Inc()
func (c *SafeCounter) Add(n int)
func (c *SafeCounter) Value() int
func (c *SafeCounter) Reset()

// SafeMap is a goroutine-safe string→int map.
type SafeMap struct { /* unexported fields */ }

func NewSafeMap() *SafeMap
func (m *SafeMap) Set(key string, value int)
func (m *SafeMap) Get(key string) (int, bool)
func (m *SafeMap) Delete(key string)
func (m *SafeMap) Len() int
```

Both types must be safe to use from multiple goroutines simultaneously.

## Examples

```go
c := NewSafeCounter()

var wg sync.WaitGroup
for i := 0; i < 1000; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        c.Inc()
    }()
}
wg.Wait()

c.Value()  →  1000  // always — no data race

sm := NewSafeMap()
sm.Set("hits", 42)
sm.Get("hits")   →  42, true
sm.Get("miss")   →  0, false
sm.Len()         →  1
sm.Delete("hits")
sm.Len()         →  0
```

Read `skills.md` before you start.
