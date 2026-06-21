# Skills for 65-stringbuilder

**Previous:** [64-inter](../64-inter/skills.md) | **Next:** [66-stringsplit](../66-stringsplit/skills.md)

**Challenge:** Remove all vowels from a string using `strings.Builder` for efficient string construction (without using `strings.ReplaceAll`).

## Core Concept: `strings.Builder` and Why String Concatenation is O(n²)

### The Problem with `+=` in Loops

This is the single most important performance lesson for string manipulation in Go.

In Go, **strings are immutable**. Every time you write `result += string(c)`, Go:
1. Allocates a new string in memory (size = old length + 1)
2. Copies the entire old string into the new memory
3. Appends the new character

For a loop over a 1000-character string:
- Iteration 1: copy 0 chars → 1 allocation
- Iteration 2: copy 1 char → 1 allocation
- Iteration 3: copy 2 chars → 1 allocation
- ...
- Iteration 1000: copy 999 chars → 1 allocation

**Total copying: 0+1+2+...+999 = ~500,000 operations.** This is O(n²).

### `strings.Builder` is O(n)

`strings.Builder` uses an internal byte buffer that **grows by doubling** (like a slice). Characters are appended in place until the buffer needs to grow, at which point it doubles. The total number of copy operations across all doublings is O(n) — amortized constant time per character.

```go
// BAD — O(n²): creates a new string on every iteration
result := ""
for _, c := range s {
    result += string(c)  // allocates new string every time!
}

// GOOD — O(n): appends to internal buffer, rarely copies
var b strings.Builder
for _, c := range s {
    b.WriteRune(c)       // appends in place
}
result := b.String()    // single final conversion
```

### The `strings.Builder` API

```go
import "strings"

var b strings.Builder

// Write methods:
b.WriteRune('H')          // append a single rune
b.WriteByte('e')          // append a single byte
b.WriteString("llo")      // append a string

// Get result:
result := b.String()      // returns the accumulated string

// Optional: pre-allocate capacity if you know the approximate size
b.Grow(100)  // hint: we'll write ~100 bytes

// Check current length:
n := b.Len()
```

### Applying to This Challenge: Remove Vowels

```go
func RemoveVowels(s string) string {
    var b strings.Builder
    for _, c := range s {
        switch c {
        case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
            // skip vowel
        default:
            b.WriteRune(c)  // keep consonant
        }
    }
    return b.String()
}
```

Or using `strings.ContainsRune`:

```go
func isVowel(c rune) bool {
    return strings.ContainsRune("aeiouAEIOU", c)
}

func RemoveVowels(s string) string {
    var b strings.Builder
    for _, c := range s {
        if !isVowel(c) {
            b.WriteRune(c)
        }
    }
    return b.String()
}
```

### When to Use `strings.Builder`

| Situation | Use |
|-----------|-----|
| Building a string in a loop over many characters | `strings.Builder` |
| One or two simple concatenations | `+` or `+=` is fine |
| Joining a small fixed number of strings | `+` or `+=` is fine |
| Joining many strings (like in Join/Repeat) | `strings.Builder` |

The O(n²) cost only matters for large inputs or many iterations. For tiny strings, `+=` is fine.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Using `+=` in a long loop | O(n²) performance on large inputs | Use `strings.Builder` |
| Forgetting `.String()` | `b` itself is not a string | Call `b.String()` to get the result |
| Calling `b.Reset()` and expecting to reuse | Resets the builder (valid but rarely needed) | Just declare a new `var b strings.Builder` |

## Algorithm (for RemoveVowels)

1. Create `var b strings.Builder`
2. For each rune `c` in `s`:
   - If `c` is NOT a vowel: `b.WriteRune(c)`
3. Return `b.String()`

## The Challenge

See [README.md](README.md).

**Next:** [66-stringsplit](../66-stringsplit/README.md)
