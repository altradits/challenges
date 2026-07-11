# Skills for 166-generics-intro

## What You'll Learn

**Previous:** [165-embedding](../165-embedding/skills.md) | **Next:** [167-io-readers-writers](../167-io-readers-writers/skills.md)

**Challenge:** Write generic functions that work with any type using type parameters.

## Generics — Type Parameters (Go 1.18+)

Generics let you write a function once and use it with multiple types — without duplicating code and without losing type safety.

```go
// Before generics — one function per type
func MaxInt(a, b int) int {
    if a > b { return a }
    return b
}

func MaxFloat(a, b float64) float64 {
    if a > b { return a }
    return b
}

// With generics — one function for many types
func Max[T interface{ ~int | ~float64 | ~string }](a, b T) T {
    if a > b { return a }
    return b
}

Max(3, 5)        // 5      — T inferred as int
Max(3.1, 2.7)    // 3.1   — T inferred as float64
Max("b", "a")    // "b"   — T inferred as string
```

## Syntax

```
func Name[TypeParam Constraint](params) returnType { ... }
```

| Part | Example | Meaning |
|------|---------|---------|
| `[T any]` | `func Map[T any]` | T can be anything |
| `[T comparable]` | `func Contains[T comparable]` | T supports == and != |
| `[T, U any]` | `func Map[T, U any]` | Two independent type params |
| `[K comparable, V any]` | `func Keys[K comparable, V any]` | Key must be comparable |

## Constraints

A **constraint** limits what types can be used for a type parameter.

### `any`

The loosest constraint — any type at all. You can only do things every type supports (pass around values, assign, etc.):

```go
func Identity[T any](v T) T { return v }
```

### `comparable`

Types that support `==` and `!=`: integers, strings, booleans, pointers, arrays, structs of comparable fields. Does NOT include slices or maps.

```go
func Contains[T comparable](s []T, target T) bool {
    for _, v := range s {
        if v == target {
            return true
        }
    }
    return false
}
```

### Interface Constraints with `~`

`~int` means "any type whose underlying type is int" — this includes custom types like `type MyInt int`:

```go
type Ordered interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~float32 | ~float64 |
    ~string
}

func Min[T Ordered](s []T) T { ... }
```

The `golang.org/x/exp/constraints` package (and since Go 1.21, `cmp.Ordered`) provides these standard constraints.

### `constraints.Ordered` (standard)

```go
import "cmp"  // Go 1.21+

func Min[T cmp.Ordered](a, b T) T {
    if a < b { return a }
    return b
}
```

## Two Type Parameters

When the input and output types differ, use two type parameters:

```go
func Map[T, U any](s []T, fn func(T) U) []U {
    result := make([]U, len(s))
    for i, v := range s {
        result[i] = fn(v)
    }
    return result
}

// int slice → string slice
Map([]int{1,2,3}, strconv.Itoa)  // ["1","2","3"]
```

## When to Use Generics

| Use generics when... | Use interfaces when... |
|---------------------|----------------------|
| The algorithm is the same for many types | Behaviour varies per type |
| You need type-safe containers | You need runtime polymorphism |
| Writing utility functions (Map, Filter, Reduce) | Designing plugin-style APIs |

The standard library's `slices`, `maps`, and `cmp` packages (Go 1.21) are fully generic.

## Solving the Challenge

```go
package piscine

func Map[T, U any](s []T, fn func(T) U) []U {
    result := make([]U, len(s))
    for i, v := range s {
        result[i] = fn(v)
    }
    return result
}

func Filter[T any](s []T, keep func(T) bool) []T {
    result := []T{}
    for _, v := range s {
        if keep(v) {
            result = append(result, v)
        }
    }
    return result
}

func Reduce[T, U any](s []T, initial U, fn func(U, T) U) U {
    acc := initial
    for _, v := range s {
        acc = fn(acc, v)
    }
    return acc
}

func Contains[T comparable](s []T, target T) bool {
    for _, v := range s {
        if v == target {
            return true
        }
    }
    return false
}

func Keys[K comparable, V any](m map[K]V) []K {
    keys := make([]K, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}

func Min[T interface{ ~int | ~float64 | ~string }](s []T) T {
    m := s[0]
    for _, v := range s[1:] {
        if v < m {
            m = v
        }
    }
    return m
}
```

## Documentation

- [Go Blog — An Introduction to Generics](https://go.dev/blog/intro-generics)
- [A Tour of Go — Generics](https://go.dev/tour/generics/1)
- [Go 1.21 — slices package](https://pkg.go.dev/slices)
- [Go 1.21 — maps package](https://pkg.go.dev/maps)
- [cmp.Ordered constraint](https://pkg.go.dev/cmp#Ordered)

**Next:** [167-io-readers-writers](../167-io-readers-writers/README.md)
