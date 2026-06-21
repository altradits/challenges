# 151-generics — Generics (Go 1.18+)

## Challenge

Implement in `package piscine`:

```go
// Map applies fn to each element of s and returns a new slice of results.
func Map[T, U any](s []T, fn func(T) U) []U

// Filter returns elements of s for which fn returns true.
func Filter[T any](s []T, fn func(T) bool) []T

// Reduce folds s into a single value using fn, starting from initial.
func Reduce[T, U any](s []T, initial U, fn func(U, T) U) U

// Keys returns all keys of a map in an unspecified order.
func Keys[K comparable, V any](m map[K]V) []K

// Contains reports whether target is in s.
func Contains[T comparable](s []T, target T) bool
```

Read `skills.md` before you start.
