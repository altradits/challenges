# Skills for 32-errorhandling

## What You'll Learn

**Previous:** [16-interfaces](../16-interfaces/skills.md) | **Next:** [18-typeassertions](../18-typeassertions/skills.md)

**Challenge:** Return errors, wrap them, and define custom error types.

## The `error` Interface

`error` is a built-in interface with one method:

```go
type error interface {
    Error() string
}
```

Every Go function that can fail should return `error` as its last return value.

### Creating Errors

```go
import "errors"

// Simple error
err := errors.New("something went wrong")

// Formatted error
err = fmt.Errorf("user %d not found", userID)
```

### Checking Errors

```go
result, err := Divide(10, 0)
if err != nil {
    fmt.Println("error:", err)
    return
}
fmt.Println(result)
```

**Rule:** Always check `err != nil` immediately after the call.

### Wrapping Errors with `%w`

```go
func ParseAge(s string) (int, error) {
    n, err := strconv.Atoi(s)
    if err != nil {
        return 0, fmt.Errorf("ParseAge: %w", err)  // wraps the original error
    }
    if n < 0 {
        return 0, fmt.Errorf("ParseAge: age cannot be negative")
    }
    return n, nil
}
```

The `%w` verb wraps the original error, preserving it for `errors.Is` and `errors.As` unwrapping.

### `errors.Is` — Checking for a Specific Error

```go
var ErrDivByZero = errors.New("division by zero")

func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, ErrDivByZero
    }
    return a / b, nil
}

_, err := Divide(5, 0)
if errors.Is(err, ErrDivByZero) {
    fmt.Println("can't divide by zero")
}
```

`errors.Is` unwraps the chain — it works even if the error was wrapped with `fmt.Errorf("%w", ...)`.

### `errors.As` — Getting the Concrete Error Type

```go
type ValidationError struct {
    Field   string
    Message string
}
func (e *ValidationError) Error() string {
    return e.Field + ": " + e.Message
}

err := &ValidationError{Field: "email", Message: "invalid format"}

var ve *ValidationError
if errors.As(err, &ve) {
    fmt.Println("bad field:", ve.Field)  // bad field: email
}
```

`errors.As` traverses the error chain and type-asserts to the target type.

### Sentinel Errors

Errors declared at package level as variables are called sentinel errors:

```go
var (
    ErrNotFound    = errors.New("not found")
    ErrPermission  = errors.New("permission denied")
)
```

Callers check with `errors.Is(err, ErrNotFound)`.

### `defer` — Guaranteed Cleanup

`defer` schedules a function call to run when the surrounding function returns — no matter how it returns (normal, error, panic).

```go
func readFile(path string) error {
    f, err := os.Open(path)
    if err != nil {
        return err
    }
    defer f.Close()   // runs when readFile returns, even on error

    // read the file...
    return nil
}
```

**Key rules:**
1. Deferred calls run in **LIFO** order (last-in, first-out — like a stack)
2. Arguments are evaluated **immediately** at the `defer` statement, not when the function runs
3. `defer` works with any function call, including anonymous functions

**LIFO order:**
```go
func audit() {
    defer fmt.Println("third")  // runs last
    defer fmt.Println("second")
    defer fmt.Println("first")  // runs first
}
// Output: first, second, third
```

**Common patterns you will see everywhere in Go:**

```go
// Close a file
defer f.Close()

// Close an HTTP response body
defer resp.Body.Close()

// Release a database connection
defer rows.Close()

// Cancel a context
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// Unlock a mutex
mu.Lock()
defer mu.Unlock()

// Rollback a transaction if not committed
tx, _ := db.Begin()
defer tx.Rollback()   // no-op after tx.Commit()
```

**Defer with a named return value** (captures the return value after the `return`):
```go
func divide(a, b float64) (result float64, err error) {
    defer func() {
        if err != nil {
            err = fmt.Errorf("divide: %w", err)  // wraps the error before returning
        }
    }()
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

### panic and recover

`panic` stops normal execution. `recover` catches a panic inside a `defer`.

```go
func safeDiv(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered: %v", r)
        }
    }()
    result = a / b  // panics if b == 0
    return
}
```

**Rule:** Use panic/recover only at package boundaries (like HTTP handlers), not as normal control flow. Return errors instead.

### Solving the Challenge

```go
package piscine

import (
    "errors"
    "fmt"
    "strconv"
)

var ErrDivByZero = errors.New("division by zero")
var ErrNegativeAge = errors.New("age cannot be negative")

func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, ErrDivByZero
    }
    return a / b, nil
}

func ParseAge(s string) (int, error) {
    n, err := strconv.Atoi(s)
    if err != nil {
        return 0, fmt.Errorf("ParseAge %q: %w", s, err)
    }
    if n < 0 {
        return 0, fmt.Errorf("ParseAge %q: %w", s, ErrNegativeAge)
    }
    return n, nil
}

type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return e.Field + ": " + e.Message
}
```

**Next:** [18-typeassertions](../18-typeassertions/README.md)
