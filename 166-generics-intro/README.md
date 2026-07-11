# 166-generics-intro — Generics (Type Parameters)

## Challenge

Implement in `package piscine`:

```go
// Map applies fn to every element of s and returns the results.
func Map[T, U any](s []T, fn func(T) U) []U

// Filter returns elements of s for which keep returns true.
func Filter[T any](s []T, keep func(T) bool) []T

// Reduce folds s into a single value, starting with initial.
func Reduce[T, U any](s []T, initial U, fn func(U, T) U) U

// Contains reports whether target is in s.
func Contains[T comparable](s []T, target T) bool

// Keys returns all keys of m in any order.
func Keys[K comparable, V any](m map[K]V) []K

// Min returns the smallest element in s.
// Constraint: T must be ordered (int, float64, string...).
func Min[T interface{ ~int | ~float64 | ~string }](s []T) T
```

## Examples

```
Map([]int{1, 2, 3}, func(n int) int { return n * 2 })
  →  [2, 4, 6]

Map([]int{1, 2, 3}, func(n int) string { return strconv.Itoa(n) })
  →  ["1", "2", "3"]

Filter([]int{1,2,3,4,5}, func(n int) bool { return n%2==0 })
  →  [2, 4]

Reduce([]int{1,2,3,4}, 0, func(acc, n int) int { return acc + n })
  →  10

Contains([]string{"a","b","c"}, "b")  →  true
Contains([]int{1,2,3}, 9)            →  false

Keys(map[string]int{"a":1,"b":2})  →  ["a","b"] (any order)

Min([]int{3,1,4,1,5,9})     →  1
Min([]string{"c","a","b"})  →  "a"
```

Read `skills.md` before you start.
