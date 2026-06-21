# Skills for iscapitalized

## What You'll Learn

**Previous:** [94-fromto](../94-fromto/skills.md) | **Next:** [96-itoa-manual](../96-itoa-manual/skills.md)

**Challenge:** Return true if every word in a string begins with an uppercase letter or a non-alphabetic character.

## Core Concept: unicode.IsUpper and Word-Boundary Detection

### What Is unicode.IsUpper?
The `unicode` package provides character classification functions. `unicode.IsUpper(r)` returns true if the rune is an uppercase letter. `unicode.IsLetter(r)` returns true if it is any letter.

```go
import "unicode"

fmt.Println(unicode.IsUpper('A'))  // true
fmt.Println(unicode.IsUpper('a'))  // false
fmt.Println(unicode.IsUpper('3'))  // false
fmt.Println(unicode.IsLetter('A')) // true
fmt.Println(unicode.IsLetter('3')) // false
```

### The Challenge Logic
A word is capitalized if its first character is uppercase *or* non-alphabetic. A word starts after a space (or at position 0).

```go
import (
    "strings"
    "unicode"
)

func IsCapitalized(s string) bool {
    if s == "" {
        return false
    }
    words := strings.Fields(s)
    if len(words) == 0 {
        return false
    }
    for _, word := range words {
        firstRune := []rune(word)[0]
        if unicode.IsLetter(firstRune) && !unicode.IsUpper(firstRune) {
            return false  // letter and NOT uppercase → fail
        }
    }
    return true
}
```

### The Key Condition
- If the first character is a letter AND not uppercase → return false.
- If the first character is a digit, punctuation, or symbol → it passes (non-alphabetic is fine).
- If the first character is an uppercase letter → it passes.

```go
if unicode.IsLetter(firstRune) && !unicode.IsUpper(firstRune) {
    return false
}
```

### Converting to []rune for First Character
Using `[]rune(word)[0]` safely gets the first Unicode code point:
```go
word := "Hello"
runes := []rune(word)
fmt.Println(runes[0])              // 72 — 'H'
fmt.Println(unicode.IsUpper(runes[0])) // true
```

### Step-by-Step for `IsCapitalized("Whats 4this 100K?")`
- Words: `["Whats", "4this", "100K?"]`
- `"Whats"`: first rune `'W'` — `IsLetter` yes, `IsUpper` yes → pass
- `"4this"`: first rune `'4'` — `IsLetter` no → pass (non-alphabetic)
- `"100K?"`: first rune `'1'` — `IsLetter` no → pass
- All passed → return true

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Only checking `unicode.IsUpper` | Fails on digits/punctuation at word start | Also allow non-letters: `!IsLetter \|\| IsUpper` |
| Using `s[0]` (byte) instead of `rune` | Breaks for Unicode characters | Convert with `[]rune(word)[0]` |
| Not returning false for empty string | Challenge says empty → false | Guard with `if s == "" { return false }` |

## Solving This Challenge

### Algorithm
1. Return false if `s == ""`.
2. Split into words with `strings.Fields`.
3. For each word: get first rune; if it's a letter AND not uppercase, return false.
4. Return true.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [96-itoa-manual](../96-itoa-manual/README.md)
