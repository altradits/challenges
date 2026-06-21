# Skills for 33-typeassertions

## What You'll Learn

**Previous:** [17-errorhandling](../17-errorhandling/skills.md) | **Next:** [19-fileio](../19-fileio/skills.md)

**Challenge:** Extract concrete types from interface values.

## Type Assertion: `v.(T)`

When you have an interface value, use a type assertion to extract the concrete type:

```go
var i any = "hello"

s := i.(string)      // s is "hello" (panics if i is not a string)
fmt.Println(s)       // hello
```

### Safe Form: Comma-ok

```go
var i any = "hello"

s, ok := i.(string)
if ok {
    fmt.Println("string:", s)
} else {
    fmt.Println("not a string")
}
```

Always use the comma-ok form unless you are **certain** of the type.

### Type Switch

A type switch is the clean way to handle multiple possible types:

```go
func Describe(v any) string {
    switch x := v.(type) {
    case int:
        return fmt.Sprintf("integer: %d", x)
    case string:
        return fmt.Sprintf("string: %s", x)
    case bool:
        return fmt.Sprintf("bool: %t", x)
    case []int:
        return fmt.Sprintf("int slice of length %d", len(x))
    default:
        return "unknown type"
    }
}
```

Inside each `case`, `x` has the concrete type — `x` is an `int` in the `case int:` branch.

### `any` vs `interface{}`

`any` is a type alias for `interface{}`, introduced in Go 1.18. They are identical:

```go
func Print(v interface{}) { fmt.Println(v) }
func Print(v any)         { fmt.Println(v) }  // same thing
```

Prefer `any` in new code.

### When to Use

- Unmarshaling JSON into `map[string]any`
- Building generic utility functions before generics were available
- Working with `fmt` (which accepts `any`)

**Avoid** when you know the concrete type in advance — use the concrete type directly.

### Solving the Challenge

```go
package piscine

import "fmt"

func Describe(v any) string {
    switch x := v.(type) {
    case int:
        return fmt.Sprintf("integer: %d", x)
    case string:
        return fmt.Sprintf("string: %s", x)
    case bool:
        return fmt.Sprintf("bool: %t", x)
    case []int:
        return fmt.Sprintf("int slice of length %d", len(x))
    default:
        return "unknown type"
    }
}

func Sum(values []any) float64 {
    total := 0.0
    for _, v := range values {
        switch x := v.(type) {
        case int:
            total += float64(x)
        case float64:
            total += x
        }
    }
    return total
}
```

**Next:** [19-fileio](../19-fileio/README.md)
