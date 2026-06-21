# Skills for 90-titlecase

## What You'll Learn

**Previous:** [89-reversestring](../89-reversestring/skills.md) | **Next:** [91-wordcount](../91-wordcount/skills.md)

**Challenge:** Write `TitleCase(s string) string` that capitalizes the first letter of each word and lowercases all other letters.

## Core Concept: `strings.Fields` to Split on Whitespace + `unicode` for Case Conversion

### What Is It?

Title case requires two new tools:
1. `strings.Fields(s)` — splits a string into words, handling multiple spaces and leading/trailing spaces automatically
2. `unicode.ToUpper(r)` and `unicode.ToLower(r)` — convert individual runes to upper or lower case

### New Tool: `strings.Fields`

`strings.Fields(s)` splits `s` on any whitespace (spaces, tabs, etc.) and returns a `[]string`. It automatically handles:
- Multiple spaces between words
- Leading/trailing spaces
- Empty string (returns an empty slice)

```go
import "strings"

strings.Fields("hello world")       // ["hello", "world"]
strings.Fields("  multiple   spaces")  // ["multiple", "spaces"]
strings.Fields("")                  // []  (empty slice)
strings.Fields("single")            // ["single"]
```

### New Tool: `unicode.ToUpper` and `unicode.ToLower`

```go
import "unicode"

unicode.ToUpper('h')   // 'H'
unicode.ToLower('H')   // 'h'
unicode.ToUpper('A')   // 'A'  (already upper, no change)
unicode.ToLower('!')   // '!'  (non-letter, no change)
```

These work on any Unicode character, not just ASCII.

### The Algorithm

For each word returned by `strings.Fields`:
1. Convert the first character to uppercase
2. Convert all remaining characters to lowercase
3. Join the transformed words back together

```go
import (
    "strings"
    "unicode"
)

func TitleCase(s string) string {
    words := strings.Fields(s)
    for i, word := range words {
        runes := []rune(word)
        runes[0] = unicode.ToUpper(runes[0])
        for j := 1; j < len(runes); j++ {
            runes[j] = unicode.ToLower(runes[j])
        }
        words[i] = string(runes)
    }
    return strings.Join(words, " ")
}
```

### New Tool: `strings.Join`

`strings.Join(slice, sep)` joins a string slice into a single string, placing `sep` between each element:

```go
strings.Join([]string{"Hello", "World"}, " ")   // "Hello World"
strings.Join([]string{"a", "b", "c"}, "-")       // "a-b-c"
strings.Join([]string{}, " ")                    // ""
```

### Important: Preserving Original Spacing

The README says `"multiple   spaces   here"` should return `"Multiple   Spaces   Here"` — with the same spacing. But `strings.Fields` collapses multiple spaces to one. So if the original string has multiple spaces, a simple Fields + Join approach will change the spacing.

One way to handle this: process character by character, tracking whether the previous character was a space.

```go
func TitleCase(s string) string {
    runes := []rune(s)
    capitalizeNext := true
    for i, r := range runes {
        if r == ' ' {
            capitalizeNext = true
        } else if capitalizeNext {
            runes[i] = unicode.ToUpper(r)
            capitalizeNext = false
        } else {
            runes[i] = unicode.ToLower(r)
        }
    }
    return string(runes)
}
```

This approach:
- Sets a flag `capitalizeNext = true` at the start and after every space
- Capitalizes the next non-space character it sees
- Lowercases everything else
- Preserves the original spacing exactly

### Diagram: Processing "hello world"

```
Input runes: h  e  l  l  o     w  o  r  l  d
capitalizeNext starts: true

i=0 'h': capitalizeNext=true -> ToUpper('h')='H', capitalizeNext=false
i=1 'e': capitalizeNext=false -> ToLower('e')='e'
i=2 'l': -> ToLower('l')='l'
i=3 'l': -> ToLower('l')='l'
i=4 'o': -> ToLower('o')='o'
i=5 ' ': capitalizeNext=true
i=6 'w': capitalizeNext=true -> ToUpper('w')='W', capitalizeNext=false
...

Result: "Hello World"
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Using `strings.Fields` + `strings.Join` | Collapses multiple spaces | Use the character-by-character approach if spacing must be preserved |
| Capitalizing only the first letter without lowercasing the rest | `"ALL CAPS"` becomes `"All CAPS"` instead of `"All Caps"` | Also apply `ToLower` to all non-first characters |
| Modifying `s[i]` directly | Strings are immutable | Convert to `[]rune` first |
| `'A' + 32` to lowercase | Works for ASCII only; fails for non-ASCII | Use `unicode.ToLower` |

## Solving This Challenge

### Algorithm (character-by-character, preserves spacing)

1. Convert `s` to `[]rune`
2. Set `capitalizeNext = true`
3. For each rune:
   - If space: set `capitalizeNext = true`
   - Else if `capitalizeNext`: apply `ToUpper`, set `capitalizeNext = false`
   - Else: apply `ToLower`
4. Return `string(runes)`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [91-wordcount](../91-wordcount/README.md)
