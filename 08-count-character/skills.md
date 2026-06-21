# Skills for 08-count-character

## What You'll Learn

**Previous:** [06-checknumber](../06-checknumber/skills.md) | **Next:** [10-findsubstring](../10-findsubstring/skills.md)

**Challenge:** Write `CountChar(str string, c rune) int` that returns how many times the rune `c` appears in `str`.

## Core Concept: Counting Matches While Iterating with `for...range`

### What Is It?

In 06-checknumber you used `for...range` to find *whether* a character exists (returning `true` at the first match). Now you need to *count* all matches. Instead of returning early, you keep a counter and increment it every time the condition is true.

### New Thing: The `rune` Type as a Function Parameter

The function signature introduces `rune` explicitly:

```go
func CountChar(str string, c rune) int {
```

A `rune` is an alias for `int32`. It represents a single Unicode code point. When you write `'l'` in Go, that is a rune literal.

```go
CountChar("Hello World", 'l')
// 'l' is a rune — its integer value is 108
```

### How It Works: The Counter Pattern

```go
func CountChar(str string, c rune) int {
    count := 0
    for _, r := range str {
        if r == c {
            count++
        }
    }
    return count
}
```

Step by step:
1. `count := 0` — start at zero
2. `for _, r := range str` — visit each rune in the string
3. `if r == c` — compare the current rune to the target rune
4. `count++` — increment if they match
5. `return count` — return the total after scanning everything

### Diagram: Counting `'l'` in `"Hello World"`

```
"Hello World"
 H  e  l  l  o     W  o  r  l  d
 |  |  |  |  |  |  |  |  |  |  |
 ?  ?  Y  Y  ?  ?  ?  ?  ?  Y  ?   (Y = match)

count: 0 -> 0 -> 1 -> 2 -> 2 -> ... -> 3

return 3
```

### Handling the Edge Cases

The README says:
- If the character is not in the string, return 0 — this is automatic: count starts at 0 and nothing increments it
- If the string is empty, return 0 — the loop body never executes, count stays 0

No special code needed for these cases.

### The `count++` Syntax

`count++` is shorthand for `count = count + 1`. In Go, `++` is a statement, not an expression. You cannot write `x = count++` — that is a compile error.

```go
count++       // valid: increment count
count--       // valid: decrement count
x = count++   // INVALID in Go
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `return 0` inside the loop on non-match | Returns immediately on the first non-matching character | Put `return count` only after the loop |
| `str[i] == c` with index-based loop | `str[i]` returns a `byte`, not a `rune`; breaks on multi-byte UTF-8 | Use `for _, r := range str` |
| Forgetting to initialize `count := 0` | Compiler error: undefined variable | Always initialize before the loop |
| Comparing `c == "l"` (string in quotes) | Type mismatch: `c` is a rune, `"l"` is a string | Use `'l'` (single quotes) for rune literals |

## Solving This Challenge

### Algorithm

1. Initialize `count := 0`
2. Loop over every rune `r` in `str` with `for...range`
3. If `r == c`, increment `count`
4. After the loop, return `count`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [10-findsubstring](../10-findsubstring/README.md)
