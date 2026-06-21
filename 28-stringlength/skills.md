# Skills for 28-stringlength

## What You'll Learn

**Previous:** [27-context](../27-context/skills.md) | **Next:** [29-firstchar](../29-firstchar/skills.md)


**Next:** [29-firstchar](../29-firstchar/README.md)

**Challenge:** Write a function `StringLength(s string) int` that returns the number of characters in a string without using `len()`.

## Core Concept: The `for...range` Loop on Strings

### What Is It?

`for...range` is Go's way to iterate over every character in a string. It gives you each character one at a time, as a **rune** (Unicode code point). A rune is the correct unit for "one character", even when that character takes multiple bytes.

### How It Works

```go
func StringLength(s string) int {
    count := 0
    for _, c := range s {
        _ = c    // c is a rune — we don't need its value, just to count it
        count++
    }
    return count
}
```

Step by step:
1. `count := 0` — start with zero
2. `for _, c := range s` — loop over every rune in `s`; the `_` discards the byte index, `c` is the character
3. `count++` — add 1 for each character
4. `return count` — hand the total back to the caller

### Why Not Just Use `len(s)`?

`len(s)` counts **bytes**, not characters. For plain ASCII text the numbers match, but for strings containing accented letters or emoji, `len` will be larger:

```go
s := "café"
len(s)              // 5 (the é takes 2 bytes)
// for...range count  // 4 (four characters: c a f é)
```

This exercise trains you to think about iteration — a skill you will use for every challenge that follows.

### Functions in Go — a Quick Primer

This is the first challenge that requires you to write a **function**. Every function in Go has this shape:

```go
func FunctionName(paramName paramType) returnType {
    // body
    return value
}
```

For this challenge:

```go
func StringLength(s string) int {
    // s is the input string
    // int is the return type
}
```

The function must be in a file inside a package. The test file calls `StringLength`, so your function name must match exactly (capital S, capital L).

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `return len(s)` | Forbidden by the challenge | Use `for...range` counter |
| `for i := 0; i < len(s); i++` | Uses `len()` and counts bytes | Use `for _, c := range s` |
| Forgetting `return count` | Compile error: missing return | Add `return count` at the end |
| Returning inside the loop | Returns after the first character | Return only after the loop |

## Solving This Challenge

### Algorithm

1. Declare `count := 0`
2. Loop: `for _, c := range s { _ = c; count++ }`
3. Return `count`

### Trace Through an Example

Input: `"Go"`

| Iteration | Character | count after |
|-----------|-----------|-------------|
| 1 | `'G'` | 1 |
| 2 | `'o'` | 2 |
| (done) | — | return 2 |

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [29-firstchar](../29-firstchar/README.md)
