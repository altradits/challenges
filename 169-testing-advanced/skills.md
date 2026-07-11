# Skills for 169-testing-advanced

## What You'll Learn

**Previous:** [168-sync-mutex](../168-sync-mutex/skills.md) | **Next:** [176-linked-list](../176-linked-list/skills.md)

**Challenge:** Write production-quality tests: table-driven, subtests, benchmarks, and helpers.

## Table-Driven Tests

The canonical Go testing pattern. Define cases as a slice of structs, then loop:

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"positive", 2, 3, 5},
        {"negative", -1, -2, -3},
        {"zero", 0, 0, 0},
        {"mixed", -5, 10, 5},
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got := Add(tc.a, tc.b)
            if got != tc.want {
                t.Errorf("Add(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
            }
        })
    }
}
```

### Why Table-Driven?

- Adding a case is one struct literal — no new function, no copy-paste
- Each subtest has a name (`-run TestAdd/positive`) so failures are pinpointed
- Keeps test logic in one place

## Subtests — `t.Run`

```go
t.Run("subtest name", func(t *testing.T) {
    // runs as an independent test
    // t.Fail(), t.Fatal(), t.Skip() work independently
})
```

Run a specific subtest:

```bash
go test -run TestAdd/positive
go test -run TestAdd       # runs all subtests
```

## Test Helpers — `t.Helper()`

```go
func assertEqual(t *testing.T, got, want int) {
    t.Helper()  // marks this as a helper — error line points to caller
    if got != want {
        t.Errorf("got %d, want %d", got, want)
    }
}
```

Without `t.Helper()`, the error line points inside `assertEqual` instead of the calling test.

## Common Test Methods

| Method | Behaviour |
|--------|-----------|
| `t.Error(msg)` | Mark failure, continue test |
| `t.Errorf(fmt, args)` | Mark failure with format, continue |
| `t.Fatal(msg)` | Mark failure, stop test immediately |
| `t.Fatalf(fmt, args)` | Mark failure with format, stop |
| `t.Log(msg)` | Print only when test fails |
| `t.Logf(fmt, args)` | Print formatted, only when test fails |
| `t.Skip(msg)` | Skip this test |

## Benchmarks

```go
func BenchmarkFibonacci(b *testing.B) {
    for b.Loop() {   // Go 1.24+ style
        Fibonacci(30)
    }
}
```

Run benchmarks:

```bash
go test -bench=.              # all benchmarks
go test -bench=BenchmarkFib   # specific benchmark
go test -bench=. -benchmem    # include memory allocations
go test -bench=. -count=5     # run 5 times for stability
```

Output:

```
BenchmarkFibonacci-8    500000    2341 ns/op
```

- `-8` = GOMAXPROCS
- `500000` = iterations run
- `2341 ns/op` = nanoseconds per call

## The `testing.T` and `testing.B` Types

Both embed `testing.TB` which has the common methods. Benchmarks add:

```go
b.N          // number of iterations (set by the framework)
b.ResetTimer()  // reset timer after expensive setup
b.StopTimer()   // pause timer during setup/teardown
b.StartTimer()  // restart timer
b.ReportAllocs() // force memory allocation reporting
```

## Testing With `testify` (Popular Library)

```go
import "github.com/stretchr/testify/assert"

func TestSomething(t *testing.T) {
    assert.Equal(t, 5, Add(2, 3))
    assert.True(t, IsPrime(7))
    assert.NoError(t, err)
}
```

Not in the standard library, but very widely used. Stick to standard `testing` to master the fundamentals first.

## Solving the Challenge

```go
// math_ops.go
package piscine

func Abs(n int) int {
    if n < 0 { return -n }
    return n
}

func Pow(base, exp int) int {
    result := 1
    for exp > 0 {
        result *= base
        exp--
    }
    return result
}

func IsPrime(n int) bool {
    if n < 2 { return false }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 { return false }
    }
    return true
}

func Fibonacci(n int) int {
    if n <= 1 { return n }
    a, b := 0, 1
    for i := 2; i <= n; i++ {
        a, b = b, a+b
    }
    return b
}
```

```go
// math_ops_test.go
package piscine

import "testing"

func TestAbs(t *testing.T) {
    tests := []struct {
        name  string
        input int
        want  int
    }{
        {"positive", 5, 5},
        {"negative", -5, 5},
        {"zero", 0, 0},
        {"large positive", 1000000, 1000000},
        {"large negative", -1000000, 1000000},
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            if got := Abs(tc.input); got != tc.want {
                t.Errorf("Abs(%d) = %d, want %d", tc.input, got, tc.want)
            }
        })
    }
}

func TestIsPrime(t *testing.T) {
    tests := []struct {
        name  string
        input int
        want  bool
    }{
        {"two", 2, true},
        {"three", 3, true},
        {"four", 4, false},
        {"one", 1, false},
        {"zero", 0, false},
        {"negative", -7, false},
        {"large prime", 9999991, true},
        {"large composite", 9999994, false},
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            if got := IsPrime(tc.input); got != tc.want {
                t.Errorf("IsPrime(%d) = %v, want %v", tc.input, got, tc.want)
            }
        })
    }
}

func BenchmarkIsPrime(b *testing.B) {
    for b.Loop() {
        IsPrime(9999991)
    }
}

func BenchmarkFibonacci(b *testing.B) {
    for b.Loop() {
        Fibonacci(40)
    }
}
```

## Documentation

- [testing package](https://pkg.go.dev/testing)
- [Go Blog — Table-Driven Tests](https://go.dev/wiki/TableDrivenTests)
- [Go Blog — Writing Benchmarks](https://go.dev/blog/pprof)
- [testify library](https://github.com/stretchr/testify)

**Next:** [176-linked-list](../176-linked-list/README.md)
