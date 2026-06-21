# Skills for 38-testing-basics

## What You'll Learn

**Previous:** [22-modules-and-tooling](../22-modules-and-tooling/skills.md) | **Next:** [24-goroutines](../24-goroutines/skills.md)

**Challenge:** Write and run tests and benchmarks using Go's `testing` package.

## Test Files

Tests live in files ending with `_test.go` in the same package:

```
mypackage/
├── solution.go        ← your implementation
└── solution_test.go   ← your tests
```

## Writing a Basic Test

```go
package piscine

import "testing"

func TestAbs(t *testing.T) {
    got  := Abs(-5)
    want := 5
    if got != want {
        t.Errorf("Abs(-5) = %d, want %d", got, want)
    }
}
```

- Function must start with `Test`, take `*testing.T`, return nothing
- `t.Errorf` marks the test as failed but continues running
- `t.Fatalf` marks as failed and stops the test immediately

## Table-Driven Tests

The standard Go pattern for testing multiple inputs:

```go
func TestAbs(t *testing.T) {
    tests := []struct {
        name string
        input int
        want  int
    }{
        {"positive", 5, 5},
        {"negative", -5, 5},
        {"zero", 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Abs(tt.input)
            if got != tt.want {
                t.Errorf("Abs(%d) = %d, want %d", tt.input, got, tt.want)
            }
        })
    }
}
```

`t.Run` creates a sub-test. Run a specific sub-test with `go test -run TestAbs/negative`.

## Useful `*testing.T` Methods

```go
t.Error("message")             // fail but continue
t.Errorf("fmt %v", val)        // fail but continue, formatted
t.Fatal("message")             // fail and stop this test
t.Fatalf("fmt %v", val)        // fail and stop, formatted
t.Log("message")               // print only if test fails (or -v)
t.Logf("fmt %v", val)
t.Helper()                     // marks this as a helper; errors point to caller
t.Parallel()                   // run this test in parallel with others
t.Skip("reason")               // skip the test
```

## Running Tests

```bash
go test ./...              # run all tests
go test -v ./...           # verbose (show each test name)
go test -run TestAbs       # run tests matching the pattern
go test -run TestAbs/zero  # run specific subtest
go test -count=1 ./...     # disable result caching
```

## Benchmarks

Benchmarks measure how fast a function runs. The framework runs it repeatedly, adjusting `b.N` until results are stable.

```go
func BenchmarkFibonacci(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Fibonacci(30)
    }
}
```

Run with: `go test -bench=. ./...`

```
BenchmarkFibonacci-8    45678    26345 ns/op    0 B/op    0 allocs/op
```

| Column | Meaning |
|--------|---------|
| `45678` | Number of iterations run |
| `26345 ns/op` | Nanoseconds per operation |
| `0 B/op` | Bytes allocated per op (with `-benchmem`) |
| `0 allocs/op` | Heap allocations per op (with `-benchmem`) |

```bash
go test -bench=. -benchmem ./...   # include memory stats
go test -bench=BenchmarkFib ./...  # run a specific benchmark
go test -bench=. -benchtime=5s     # run for 5 seconds instead of 1
go test -bench=. -count=3          # run each benchmark 3 times
```

**Resetting the timer** — exclude setup time from the benchmark:

```go
func BenchmarkSort(b *testing.B) {
    data := make([]int, 10000)
    for i := range data { data[i] = rand.Intn(1000) }

    b.ResetTimer()  // don't count setup time above
    for i := 0; i < b.N; i++ {
        sort.Ints(data)
    }
}
```

**Why benchmarks matter for Go/Bitcoin:** LND runs on nodes that process many Lightning payments per second. Understanding how to measure and reduce `ns/op` and `allocs/op` is the same skill used to optimize the LND payment pathfinding code.

## Test Helpers

Extract repeated setup into a helper:

```go
func checkEqual(t *testing.T, got, want int) {
    t.Helper()  // errors will point to the caller, not here
    if got != want {
        t.Errorf("got %d, want %d", got, want)
    }
}
```

## Solving the Challenge

```go
// solution_test.go
package piscine

import "testing"

func TestAbs(t *testing.T) {
    tests := []struct{ in, want int }{
        {5, 5}, {-5, 5}, {0, 0}, {-100, 100},
    }
    for _, tt := range tests {
        t.Run("", func(t *testing.T) {
            if got := Abs(tt.in); got != tt.want {
                t.Errorf("Abs(%d) = %d, want %d", tt.in, got, tt.want)
            }
        })
    }
}

func TestPalindrome(t *testing.T) {
    tests := []struct {
        s    string
        want bool
    }{
        {"racecar", true},
        {"hello", false},
        {"", true},
        {"a", true},
    }
    for _, tt := range tests {
        t.Run(tt.s, func(t *testing.T) {
            if got := Palindrome(tt.s); got != tt.want {
                t.Errorf("Palindrome(%q) = %v, want %v", tt.s, got, tt.want)
            }
        })
    }
}

func TestFizzBuzz(t *testing.T) {
    cases := []struct{ n int; want string }{
        {1, "1"}, {3, "Fizz"}, {5, "Buzz"}, {15, "FizzBuzz"},
    }
    for _, tt := range cases {
        if got := FizzBuzz(tt.n); got != tt.want {
            t.Errorf("FizzBuzz(%d) = %q, want %q", tt.n, got, tt.want)
        }
    }
}

func BenchmarkFizzBuzz(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FizzBuzz(15)
    }
}
```

**Next:** [24-goroutines](../24-goroutines/README.md)
