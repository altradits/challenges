# 41-select-and-sync — select, Mutex, and sync Patterns

## Challenge

Implement in `package piscine`:

```go
// SafeCounter is a goroutine-safe counter.
type SafeCounter struct { /* unexported fields */ }

func NewSafeCounter() *SafeCounter
func (c *SafeCounter) Inc()
func (c *SafeCounter) Value() int

// Timeout runs fn in a goroutine and returns its result,
// or returns ("", ErrTimeout) if fn takes longer than d.
func Timeout(fn func() string, d time.Duration) (string, error)

// Once runs fn only the first time it is called, no matter how many goroutines call it.
// Returns the result of the first call every time.
func NewOnce(fn func() string) func() string
```

Read `skills.md` before you start.
