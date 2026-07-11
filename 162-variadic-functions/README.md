# 162-variadic-functions — Variadic Functions

## Challenge

Implement in `package piscine`:

```go
// Sum returns the sum of all provided integers.
func Sum(nums ...int) int

// Max returns the largest of all provided integers.
// Panics if no arguments are given.
func Max(nums ...int) int

// Join concatenates strings with sep between each pair.
func Join(sep string, parts ...string) string

// Wrap calls fn for each argument, collecting return values into a slice.
func Wrap(fn func(int) int, nums ...int) []int
```

## Examples

```
Sum(1, 2, 3, 4, 5)          →  15
Sum()                        →  0

Max(3, 1, 4, 1, 5, 9, 2)   →  9

Join(", ", "a", "b", "c")  →  "a, b, c"
Join("-", "go", "is", "fun") →  "go-is-fun"

double := func(n int) int { return n * 2 }
Wrap(double, 1, 2, 3)       →  [2, 4, 6]
```

Read `skills.md` before you start.
