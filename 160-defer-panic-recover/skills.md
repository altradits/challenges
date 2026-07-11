# Skills for 160-defer-panic-recover

## What You'll Learn

**Previous:** [159-closures](../159-closures/skills.md) | **Next:** [161-switch-statements](../161-switch-statements/skills.md)

**Challenge:** Use `defer` for cleanup, catch panics with `recover`, and understand LIFO defer ordering.

## `defer` — Schedule Work for Later

`defer` pushes a function call onto a stack. When the surrounding function returns, Go executes all deferred calls in **LIFO order** (last registered = first executed).

```go
func ReadFile(path string) error {
    f, err := os.Open(path)
    if err != nil {
        return err
    }
    defer f.Close()  // runs when ReadFile returns, no matter what

    // ... read from f
    return nil
}
```

### LIFO Order

```go
func demo() {
    defer fmt.Println("first")   // runs third
    defer fmt.Println("second")  // runs second
    defer fmt.Println("third")   // runs first
    fmt.Println("body")
}
// Output:
// body
// third
// second
// first
```

### Defer Captures Arguments at Registration Time

```go
x := 10
defer fmt.Println(x)  // captures x = 10 NOW
x = 20
// deferred call prints: 10
```

But if you use a closure, it captures by reference:

```go
x := 10
defer func() { fmt.Println(x) }()  // captures x by reference
x = 20
// deferred call prints: 20
```

### Named Return Values + Defer

You can modify named return values inside a deferred function:

```go
func SafeDivide(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            result = 0
            err = fmt.Errorf("recovered: %v", r)
        }
    }()
    result = a / b  // panics if b == 0
    return
}
```

This pattern is the standard way to turn panics into errors.

## `panic` — Halt Execution

`panic` stops normal execution, unwinds the stack running deferred functions, and crashes the program (unless recovered).

```go
func mustPositive(n int) {
    if n <= 0 {
        panic(fmt.Sprintf("expected positive, got %d", n))
    }
}
```

**When to panic:**
- Programmer errors that should never happen in correct code (e.g., nil pointer you know can't be nil, index out of range you guaranteed is in-range)
- During program initialisation, before the server starts

**When NOT to panic:**
- Any error that can happen at runtime (file not found, network failure, bad user input) — return an `error` instead

## `recover` — Catch a Panic

`recover` is only useful inside a `defer`d function. It returns the value passed to `panic`, or `nil` if there was no panic:

```go
func safeCall(fn func()) (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic: %v", r)
        }
    }()
    fn()
    return nil
}
```

### Visual: panic / recover flow

```
fn() starts
  |
  └─ panic("something bad")
       |
       ├─ deferred functions run in LIFO order
       |    └─ recover() returns "something bad"
       |         └─ err = fmt.Errorf("panic: something bad")
       |
       └─ safeCall returns err (non-nil)
```

## Common Patterns

### Pattern 1: Cleanup with defer

```go
func processJob(id int) {
    fmt.Println("start job", id)
    defer fmt.Println("end job", id)  // always runs

    // ... do work ...
}
```

### Pattern 2: Mutex unlock with defer

```go
mu.Lock()
defer mu.Unlock()  // always unlocks, even on panic or early return
```

### Pattern 3: Benchmark timer

```go
func timed(name string) func() {
    start := time.Now()
    return func() {
        fmt.Printf("%s took %v\n", name, time.Since(start))
    }
}

defer timed("loadData")()
```

## Solving the Challenge

```go
package piscine

import "fmt"

func SafeDivide(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            result = 0
            err = fmt.Errorf("cannot divide by zero")
        }
    }()
    result = a / b
    return
}

func WithCleanup(fn func(), cleanup func()) (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic: %v", r)
        }
        cleanup()
    }()
    fn()
    return nil
}

func DeferOrder() []string {
    log := []string{}
    defer func() { log = append(log, "first") }()
    defer func() { log = append(log, "second") }()
    defer func() { log = append(log, "third") }()
    return log
}
```

Wait — `DeferOrder` has a subtlety. Defers run *after* the return statement. If you `return log` before defers run, the slice you returned won't include the deferred appends. Use a named return:

```go
func DeferOrder() (log []string) {
    defer func() { log = append(log, "first") }()
    defer func() { log = append(log, "second") }()
    defer func() { log = append(log, "third") }()
    return
}
// log == ["third", "second", "first"]
```

## Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `recover()` outside a `defer` | Always returns nil — does nothing | Put `recover` inside a `defer`red closure |
| Panicking for expected errors | Panics are expensive and hard to test | Return `error` for anything that can legitimately fail |
| Deferred call uses loop var | Captured by reference; sees final value | Use closure parameter or `i := i` shadowing |

## Documentation

- [Effective Go — Defer](https://go.dev/doc/effective_go#defer)
- [Effective Go — Panic and Recover](https://go.dev/doc/effective_go#panic)
- [Go Blog — Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover)

**Next:** [161-switch-statements](../161-switch-statements/README.md)
