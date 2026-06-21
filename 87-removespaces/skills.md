# Skills for 87-removespaces

## What You'll Learn

**Previous:** [86-ispalindrome](../86-ispalindrome/skills.md) | **Next:** [88-countrepeats](../88-countrepeats/skills.md)

**Challenge:** Write a function `RemoveSpaces(s string) string` that returns the string with all space characters (`' '`) removed.

## Core Concept: The Filter Pattern — Keep Characters That Pass a Test

### What Is Filtering?

Filtering means scanning every element and keeping only those that satisfy a condition. It is the opposite of transformation (which keeps everything but changes values):

| Pattern | What it does | Example |
|---------|-------------|---------|
| Transform | Changes every character | `ToUpper`: `"hello"` → `"HELLO"` |
| Filter | Keeps some, drops others | `RemoveSpaces`: `"a b"` → `"ab"` |

### The Filter Pattern in Go

```go
func RemoveSpaces(s string) string {
    result := ""
    for _, c := range s {
        if c != ' ' {          // condition: keep if NOT a space
            result += string(c)
        }
        // if c IS a space: do nothing — it is dropped
    }
    return result
}
```

The key difference from previous challenges: there is **no `else` branch**. Characters that fail the condition are simply not added.

### Only Space (`' '`), Not All Whitespace

A space (`' '`) has ASCII value 32. Tabs (`'\t'`) and newlines (`'\n'`) are different characters. This challenge removes only the regular space:

```go
c != ' '   // keeps everything except the space character
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `c != " "` (double quotes) | `" "` is a `string`; `' '` is a `rune` — type mismatch | Use single quotes: `c != ' '` |
| Using `strings.Replace` or `strings.ReplaceAll` | Defeats the learning purpose | Manual iteration |
| Removing `'\t'` and `'\n'` too | Over-filters; challenge says only spaces | Check `c != ' '` specifically |

## Solving This Challenge

### Algorithm

1. `result := ""`
2. For each rune `c` in `s`:
   - If `c != ' '`: append `string(c)` to result
3. Return `result`

### Trace Through an Example

Input: `"Go is fun!"`

| Char | Keep? | result so far |
|------|-------|--------------|
| `G` | Yes | `"G"` |
| `o` | Yes | `"Go"` |
| ` ` | No | `"Go"` |
| `i` | Yes | `"Goi"` |
| `s` | Yes | `"Gois"` |
| ` ` | No | `"Gois"` |
| `f` | Yes | `"Goisf"` |
| `u` | Yes | `"Goisfu"` |
| `n` | Yes | `"Goisfun"` |
| `!` | Yes | `"Goisfun!"` |

Result: `"Goisfun!"`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [88-countrepeats](../88-countrepeats/skills.md)
