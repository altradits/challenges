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

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 168
```

or run the tests directly:

```bash
go test ./168-sync-mutex/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 168` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
