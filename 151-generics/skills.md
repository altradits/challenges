# Skills for 151-generics

## What You'll Learn

**Previous:** [150-logging](../150-logging/skills.md) | **Next:** [152-financial-freedom-api](../152-financial-freedom-api/skills.md)

**Challenge:** Write generic functions that work with any type.

## Why Generics?

Before Go 1.18, writing a `Map` that worked for both `[]int` and `[]string` required either:
- Duplicating the code for each type
- Using `any` (losing type safety)

Generics let you write it once with full type safety.

## Type Parameters

```go
func Map[T, U any](s []T, fn func(T) U) []U {
    result := make([]U, len(s))
    for i, v := range s {
        result[i] = fn(v)
    }
    return result
}
```

`[T, U any]` declares two type parameters. `any` is the constraint — T and U can be anything.

### Calling Generic Functions

Go infers the type parameters from arguments:

```go
nums := []int{1, 2, 3, 4}
doubled := Map(nums, func(n int) int { return n * 2 })     // [2 4 6 8]
strs   := Map(nums, func(n int) string { return strconv.Itoa(n) }) // ["1" "2" "3" "4"]
```

You can also specify explicitly: `Map[int, string](nums, ...)`.

## Constraints

Constraints restrict which types are allowed. `any` allows everything.

### `comparable` — types that support `==` and `!=`

```go
func Contains[T comparable](s []T, target T) bool {
    for _, v := range s {
        if v == target {
            return true
        }
    }
    return false
}

Contains([]int{1, 2, 3}, 2)        // true
Contains([]string{"a", "b"}, "c") // false
```

### Union Constraints

```go
type Number interface {
    ~int | ~int64 | ~float64  // ~ means "underlying type is"
}

func Sum[T Number](s []T) T {
    var total T
    for _, v := range s {
        total += v
    }
    return total
}

Sum([]int{1, 2, 3})        // 6
Sum([]float64{1.1, 2.2})   // 3.3
```

`~int` matches `int` AND any type whose underlying type is `int` (like `type MyInt int`).

### Interface as Constraint

Any interface can be a constraint:

```go
type Stringer interface {
    String() string
}

func PrintAll[T Stringer](items []T) {
    for _, item := range items {
        fmt.Println(item.String())
    }
}
```

## Generic Types

Not just functions — types can be generic too:

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(v T)        { s.items = append(s.items, v) }
func (s *Stack[T]) Pop() (T, bool)  {
    var zero T
    if len(s.items) == 0 {
        return zero, false
    }
    v := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return v, true
}

s := Stack[int]{}
s.Push(1)
s.Push(2)
v, _ := s.Pop()  // v = 2
```

## When to Use Generics

**Use generics when:**
- Writing collection operations (Map, Filter, Reduce, Contains)
- Building data structures (Stack, Queue, Pair, Result)
- The logic is identical for multiple concrete types

**Don't use generics when:**
- An interface already handles the problem (prefer `io.Reader` over generic reader)
- The function only ever has one concrete type in practice
- Adding type parameters makes the code harder to read

### `golang.org/x/exp/slices` and `golang.org/x/exp/maps`

Go's experimental packages provide ready-made generic utilities:

```go
import "golang.org/x/exp/slices"

slices.Contains([]int{1, 2, 3}, 2)      // true
slices.Index([]string{"a", "b"}, "b")   // 1
slices.Reverse([]int{1, 2, 3})          // modifies in place

import "golang.org/x/exp/maps"
maps.Keys(map[string]int{"a": 1})       // ["a"]
```

Note: `slices` and `maps` are in the standard library as of Go 1.21.

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

func Filter[T any](s []T, fn func(T) bool) []T {
    var result []T
    for _, v := range s {
        if fn(v) {
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

func Keys[K comparable, V any](m map[K]V) []K {
    keys := make([]K, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}

func Contains[T comparable](s []T, target T) bool {
    for _, v := range s {
        if v == target {
            return true
        }
    }
    return false
}
```

**Next:** [152-financial-freedom-api](../152-financial-freedom-api/README.md)
