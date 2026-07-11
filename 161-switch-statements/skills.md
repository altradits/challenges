# Skills for 161-switch-statements

## What You'll Learn

**Previous:** [160-defer-panic-recover](../160-defer-panic-recover/skills.md) | **Next:** [162-variadic-functions](../162-variadic-functions/skills.md)

**Challenge:** Master Go's switch statement — value switch, expression switch, and type switch.

## The `switch` Statement

Go's `switch` is cleaner than C/Java: **no fall-through by default**, cases can be expressions, and you can switch on types.

### Basic Switch (Value)

```go
switch day {
case "Monday":
    fmt.Println("start of the week")
case "Friday":
    fmt.Println("end of the week")
case "Saturday", "Sunday":   // multiple values in one case
    fmt.Println("weekend")
default:
    fmt.Println("midweek")
}
```

Rules:
- Cases are evaluated top to bottom, first match wins.
- No `break` needed — each case exits automatically.
- `default` runs if no case matched (optional).

### Expression Switch (No Tag)

When you omit the switch expression, each case is a boolean expression — essentially a cleaner `if/else if` chain:

```go
switch {
case score >= 90:
    return "A"
case score >= 80:
    return "B"
case score >= 70:
    return "C"
case score >= 60:
    return "D"
default:
    return "F"
}
```

This is far more readable than a chain of `else if`.

### Type Switch

A **type switch** inspects the dynamic type of an `interface{}` (or `any`) value:

```go
func Describe(v interface{}) string {
    switch val := v.(type) {
    case int:
        return fmt.Sprintf("integer: %d", val)
    case string:
        return fmt.Sprintf("string: %s", val)
    case bool:
        return fmt.Sprintf("boolean: %v", val)
    default:
        return "unknown type"
    }
}
```

The syntax `val := v.(type)` gives you `val` with the **concrete type** inside each case — you can use `val` as an `int` in the `int` case, as a `string` in the `string` case, etc.

### `fallthrough` — Explicit Fall-Through

If you do want fall-through behaviour, use `fallthrough`:

```go
switch x {
case 1:
    fmt.Println("one")
    fallthrough
case 2:
    fmt.Println("one or two")
case 3:
    fmt.Println("three")
}
// if x == 1: prints "one" and "one or two"
// if x == 2: prints "one or two"
```

`fallthrough` is rarely used in idiomatic Go. When you reach for it, consider whether separate cases with shared code would be clearer.

### Switch with Initialiser

Like `if`, `switch` can include a short statement before the condition:

```go
switch n := computeValue(); {
case n < 0:
    fmt.Println("negative")
case n == 0:
    fmt.Println("zero")
default:
    fmt.Println("positive")
}
```

## Switch vs if/else Performance

Both compile to roughly the same bytecode for small numbers of cases. For large numbers of cases over integers, the compiler may generate a jump table — but you rarely need to think about this. Prefer `switch` for readability whenever you have three or more branches.

## Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Adding `break` manually | Redundant in Go; each case already breaks | Remove the `break` |
| Expecting Java fall-through | Each case exits automatically | Use `fallthrough` only if you need the old behavior |
| Forgetting `default` | Program may silently do nothing | Always add a `default` for exhaustive handling |
| Using `switch` on floats | Float equality is unreliable | Use expression switch with ranges for floats |

## FizzBuzz with Switch

A neat Go idiom: switch on a struct literal of booleans:

```go
func FizzBuzz(n int) string {
    switch {
    case n%15 == 0:
        return "FizzBuzz"
    case n%3 == 0:
        return "Fizz"
    case n%5 == 0:
        return "Buzz"
    default:
        return fmt.Sprintf("%d", n)
    }
}
```

Note: check `n%15 == 0` *first*, otherwise `n%3 == 0` would match for multiples of 15.

## Solving the Challenge

```go
package piscine

import "fmt"

func DayType(day string) string {
    switch day {
    case "Saturday", "Sunday":
        return "weekend"
    default:
        return "weekday"
    }
}

func Grade(score int) string {
    switch {
    case score >= 90:
        return "A"
    case score >= 80:
        return "B"
    case score >= 70:
        return "C"
    case score >= 60:
        return "D"
    default:
        return "F"
    }
}

func Describe(v interface{}) string {
    switch val := v.(type) {
    case int:
        return fmt.Sprintf("integer: %d", val)
    case string:
        return fmt.Sprintf("string: %s", val)
    case bool:
        return fmt.Sprintf("boolean: %v", val)
    default:
        return "unknown type"
    }
}

func FizzBuzz(n int) string {
    switch {
    case n%15 == 0:
        return "FizzBuzz"
    case n%3 == 0:
        return "Fizz"
    case n%5 == 0:
        return "Buzz"
    default:
        return fmt.Sprintf("%d", n)
    }
}
```

## Documentation

- [A Tour of Go — Switch](https://go.dev/tour/flowcontrol/9)
- [A Tour of of Go — Type switches](https://go.dev/tour/methods/16)
- [Go Spec — Switch statements](https://go.dev/ref/spec#Switch_statements)

**Next:** [162-variadic-functions](../162-variadic-functions/README.md)
