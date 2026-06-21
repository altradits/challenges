# Skills for 125-stringmap

## What You'll Learn

**Previous:** [124-stringfield](../124-stringfield/skills.md) | **Next:** [126-stringfilter](../126-stringfilter/skills.md)

**Challenge:** Implement `Map(s string, f func(rune) rune) string` — apply a function to every character.

## Core Concept: Higher-Order Functions (Functions as Values)

### What Is a Higher-Order Function?

A higher-order function either **takes a function as a parameter** or **returns a function**. This is one of Go's functional programming tools.

```go
// f is a parameter that IS a function
func Map(s string, f func(rune) rune) string {
    var b strings.Builder
    for _, r := range s {
        b.WriteRune(f(r))  // call f on each rune
    }
    return b.String()
}
```

The type `func(rune) rune` means: "a function that takes a rune and returns a rune."

### Calling Map with Different Functions

```go
// Pass a named function
result := Map("hello", unicode.ToUpper)  // "HELLO"

// Pass an anonymous function (lambda)
result := Map("abc", func(r rune) rune {
    return r + 1  // shift each char by 1
})
// "bcd"

// Pass a closure that captures a variable
shift := 3
result := Map("abc", func(r rune) rune {
    return r + rune(shift)
})
// "def"
```

### The Built-in `strings.Map`

```go
import "strings"

strings.Map(unicode.ToUpper, "hello")  // "HELLO"
strings.Map(unicode.ToLower, "WORLD")  // "world"
```

Note: `strings.Map` takes the function first, then the string — opposite order from this challenge's signature.

### Function Types in Go

```go
type transformer func(rune) rune   // define a type alias

var f transformer = unicode.ToUpper
f('a')  // 'A'
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Iterating with `s[i]` | Gets bytes not runes | Use `for _, r := range s` |
| Calling `f(r)` but not writing the result | Output is empty | `b.WriteRune(f(r))` |
| Confusing `func(rune) rune` with `func(string) string` | Type mismatch | Match the signature exactly |

## Algorithm

1. Create a `strings.Builder`
2. For each rune `r` in `s`: write `f(r)` to the builder
3. Return `builder.String()`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [126-stringfilter](../126-stringfilter/skills.md)
