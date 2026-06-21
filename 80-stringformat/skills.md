# Skills for 80-stringformat

## What You'll Learn

**Previous:** [79-stringreduce](../79-stringreduce/skills.md) | **Next:** [81-checknumber](../81-checknumber/skills.md)

**Challenge:** Implement a `Format(template string, args ...interface{}) string` function using `fmt.Sprintf` format verbs.

## Core Concept: `fmt.Sprintf` — Formatted String Construction

### What Is It?

`fmt.Sprintf` (Sprintf = "string print formatted") builds a string by substituting format verbs (`%s`, `%d`, etc.) with argument values. It returns the result as a string instead of printing it.

```go
import "fmt"

s := fmt.Sprintf("Hello %s, you are %d years old", "Alice", 30)
// s = "Hello Alice, you are 30 years old"
```

### Format Verbs Reference

| Verb | Type | Example | Output |
|------|------|---------|--------|
| `%s` | string | `fmt.Sprintf("%s", "go")` | `go` |
| `%d` | integer | `fmt.Sprintf("%d", 42)` | `42` |
| `%f` | float | `fmt.Sprintf("%f", 3.14)` | `3.140000` |
| `%.2f` | float, 2 decimals | `fmt.Sprintf("%.2f", 3.14159)` | `3.14` |
| `%v` | any value | `fmt.Sprintf("%v", true)` | `true` |
| `%T` | type name | `fmt.Sprintf("%T", 42)` | `int` |
| `%q` | quoted string | `fmt.Sprintf("%q", "hi")` | `"hi"` |
| `%x` | hex | `fmt.Sprintf("%x", 255)` | `ff` |
| `%X` | hex uppercase | `fmt.Sprintf("%X", 255)` | `FF` |
| `%b` | binary | `fmt.Sprintf("%b", 10)` | `1010` |
| `%%` | literal `%` | `fmt.Sprintf("100%%")` | `100%` |

### Width and Padding

```go
fmt.Sprintf("%5d", 42)    // "   42"  — right-aligned, width 5
fmt.Sprintf("%-5d", 42)   // "42   "  — left-aligned, width 5
fmt.Sprintf("%05d", 42)   // "00042"  — zero-padded, width 5
fmt.Sprintf("%8.2f", 3.14) // "    3.14"
```

### `fmt.Sprintf` vs `fmt.Printf` vs `fmt.Println`

```go
fmt.Sprintf("Hello %s", name)  // returns a string
fmt.Printf("Hello %s\n", name) // prints to stdout, no newline added
fmt.Println("Hello", name)     // prints to stdout, space-separated, newline added
```

### Variadic Arguments (`...interface{}`)

The `...` means "zero or more arguments of any type":

```go
func Format(template string, args ...interface{}) string {
    return fmt.Sprintf(template, args...)
}
```

The `args...` in the call to `Sprintf` "unpacks" the slice back into individual arguments.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `fmt.Printf` instead of `Sprintf` | Prints, doesn't return | Use `fmt.Sprintf` to get a string |
| Wrong number of args | Runtime panic or `%!(EXTRA ...)` | Match verb count to arg count |
| `%d` for a float | `%!d(float64=3.14)` | Use `%f` for floats |
| Forgetting `args...` when passing to Sprintf | Passes the slice as one arg | Use `args...` to unpack |

## Algorithm

Implementing `Format` is simple — it delegates to `fmt.Sprintf`:

```go
func Format(template string, args ...interface{}) string {
    return fmt.Sprintf(template, args...)
}
```

The learning goal here is knowing WHAT format verbs are available and WHEN to use each one.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [81-checknumber](../81-checknumber/README.md)
