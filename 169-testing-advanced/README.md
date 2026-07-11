# 169-testing-advanced — Table-Driven Tests and Benchmarks

## Challenge

Write the **tests** (not just the implementation) for a `MathOps` package:

```go
// math_ops.go — this is given to you; write tests for it
package piscine

func Abs(n int) int
func Pow(base, exp int) int
func IsPrime(n int) bool
func Fibonacci(n int) int
```

Create `math_ops_test.go` in `package piscine` with:

1. A **table-driven test** for each function covering at least 5 cases each
2. **Edge cases**: negative numbers, zero, large values
3. A **benchmark** for `IsPrime` and `Fibonacci`
4. A test that uses **subtests** (`t.Run`)

## Expected Test File Structure

```go
package piscine

import "testing"

func TestAbs(t *testing.T) {
    tests := []struct {
        name  string
        input int
        want  int
    }{
        // at least 5 cases
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got := Abs(tc.input)
            if got != tc.want {
                t.Errorf("Abs(%d) = %d, want %d", tc.input, got, tc.want)
            }
        })
    }
}

func BenchmarkIsPrime(b *testing.B) {
    for b.Loop() {
        IsPrime(9999991)
    }
}
```

Read `skills.md` before you start.
