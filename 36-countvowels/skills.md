# Skills for 36-countvowels

## What You'll Learn

**Previous:** [35-checknumber](../35-checknumber/skills.md) | **Next:** [37-reversestring](../37-reversestring/skills.md)

**Challenge:** Write a function `CountVowels(s string) int` that counts both uppercase and lowercase vowels (a, e, i, o, u, A, E, I, O, U).

## Core Concept: Matching Against a Fixed Set Using `switch`

### The Vowel Set

Vowels are five specific letters — uppercase and lowercase variants — that do not form a simple numeric range like `'a'`–`'z'`. You need to check equality against each one.

### Approach A — Long `if` with `||`

```go
if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
   c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
    count++
}
```

This works but is verbose. For a small fixed set, a `switch` statement reads more clearly.

### Approach B — `switch` Statement

```go
func CountVowels(s string) int {
    count := 0
    for _, c := range s {
        switch c {
        case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
            count++
        }
    }
    return count
}
```

In Go, a `switch` on a value checks each `case`. Multiple values can share a `case` line, separated by commas. If none match, execution continues past the switch (no action needed for non-vowels).

### Approach C — Normalise Case First

Convert each character to lowercase, then check against only five values:

```go
func CountVowels(s string) int {
    count := 0
    for _, c := range s {
        // convert uppercase to lowercase using ASCII math
        if c >= 'A' && c <= 'Z' {
            c = c + 32
        }
        if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
            count++
        }
    }
    return count
}
```

All three approaches produce the same result. Choose whichever is clearest to you.

### The `switch` Statement

A `switch` compares one value against many cases:

```go
switch variable {
case value1:
    // runs if variable == value1
case value2, value3:
    // runs if variable == value2 OR variable == value3
default:
    // runs if nothing else matched
}
```

Unlike C or Java, Go's `switch` does NOT fall through between cases by default.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Only counting lowercase vowels | Misses `A E I O U` | Add uppercase cases or normalise first |
| Using a range check `c >= 'a' && c <= 'u'` | Includes non-vowels like `b c d` | Check each vowel specifically |
| Returning on first vowel | Would return `1` instead of count | Do not return inside the loop |

## Solving This Challenge

### Algorithm

1. `count := 0`
2. For each rune `c` in `s`:
   - If `c` is one of `a e i o u A E I O U`: `count++`
3. Return `count`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [37-reversestring](../37-reversestring/README.md)
