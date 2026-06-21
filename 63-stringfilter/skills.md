# Skills for 63-stringfilter

## What You'll Learn

**Previous:** [62-stringmap](../62-stringmap/skills.md) | **Next:** [64-stringreduce](../64-stringreduce/skills.md)

**Challenge:** Implement `Filter(s string, f func(rune) bool) string` — keep only characters where `f` returns true.

## Core Concept: The Filter Pattern

### What Is Filter?

Filter is a higher-order function that **selects** elements from a collection based on a predicate (a function that returns true/false). It's the complement to Map — instead of transforming every element, you keep or discard elements.

```go
func Filter(s string, f func(rune) bool) string {
    var b strings.Builder
    for _, r := range s {
        if f(r) {          // keep only if predicate is true
            b.WriteRune(r)
        }
    }
    return b.String()
}
```

### Calling Filter with Different Predicates

```go
// Keep only digits
Filter("abc123def456", unicode.IsDigit)  // "123456"

// Keep only uppercase letters
Filter("Hello World", unicode.IsUpper)  // "HW"

// Remove spaces
Filter("hello world", func(r rune) bool {
    return r != ' '
})
// "helloworld"
```

### Map vs Filter — Side by Side

```go
// Map: transform EVERY element, output same length
Map("hello", unicode.ToUpper)  // "HELLO"  (5 → 5)

// Filter: KEEP some elements, output may be shorter
Filter("hElLo", unicode.IsUpper)  // "EL"  (5 → 2)
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Writing `b.WriteRune(f(r))` | f returns bool, not rune | `if f(r) { b.WriteRune(r) }` |
| Transforming `r` before writing | Filter shouldn't transform | Write original `r` unchanged |
| Forgetting the condition | Writes all chars = no filtering | The `if f(r)` is the key line |

## Algorithm

1. Create a `strings.Builder`
2. For each rune `r` in `s`: if `f(r)` is true, write `r`
3. Return `builder.String()`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [64-stringreduce](../64-stringreduce/README.md)
