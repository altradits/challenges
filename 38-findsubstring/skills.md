# Skills for 38-findsubstring

**Previous:** [37-countwords](../37-countwords/skills.md) | **Next:** [39-replaceall](../39-replaceall/skills.md)

**Challenge:** Find the index of the first occurrence of a substring in a string, returning -1 if not found, without using `strings.Index`.

## Core Concept: The Sliding Window Search

### What Is It?

A **sliding window** moves a fixed-size frame across the text one position at a time, checking whether the window's contents match the pattern. For substring search, the window size equals the length of the pattern.

```
Text:    "banana"  (len=6)
Pattern: "ana"     (len=3)

Window at 0: text[0:3] = "ban" — no match
Window at 1: text[1:4] = "ana" — MATCH! return 1
```

### Byte Indexing in Go

Go strings can be indexed with `s[i]` which gives the **byte** at position `i`. For ASCII text, one byte equals one character.

```go
s := "Hello"
fmt.Println(s[0])           // 72  (byte value of 'H')
fmt.Println(string(s[0]))   // "H"
fmt.Println(s[0] == 'H')    // true
```

### The Algorithm

```go
func FindSubstring(text, pattern string) int {
    if len(pattern) == 0 {
        return 0
    }
    if len(pattern) > len(text) {
        return -1
    }

    for i := 0; i <= len(text)-len(pattern); i++ {
        match := true
        for j := 0; j < len(pattern); j++ {
            if text[i+j] != pattern[j] {
                match = false
                break
            }
        }
        if match {
            return i
        }
    }
    return -1
}
```

### Why `i <= len(text) - len(pattern)`?

The last possible window start is `len(text) - len(pattern)`. If you go further, there are not enough characters left to fit the pattern.

Example: text `"Hello"` (len=5), pattern `"lo"` (len=2).
- Last valid start: `5 - 2 = 3` → window is `"lo"` — fits exactly
- Start `4` would give only one character — too short

### Tracing Through `"banana"` for `"ana"`

| i | text[i:i+3] | j=0: text[i+j]==pattern[j]? | Result |
|---|------------|------------------------------|--------|
| 0 | "ban" | 'b'=='a'? No → break | false |
| 1 | "ana" | 'a'=='a'? Yes, 'n'=='n'? Yes, 'a'=='a'? Yes | **true → return 1** |

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Loop condition `i < len(text)` | Checks windows too short to match | Use `i <= len(text)-len(pattern)` |
| Not handling empty pattern | `FindSubstring("x", "")` should return 0 | Check `len(pattern) == 0` first |
| Not handling pattern longer than text | `len(text)-len(pattern)` would be negative, loop never runs | Check `len(pattern) > len(text)` first |
| Forgetting `break` after mismatch | Still correct but wastes work | Add `break` when mismatch found |

## Algorithm

1. If `pattern` is empty, return `0`
2. If `len(pattern) > len(text)`, return `-1`
3. For `i` from `0` to `len(text)-len(pattern)` inclusive:
   a. Set `match = true`
   b. For `j` from `0` to `len(pattern)-1`:
      - If `text[i+j] != pattern[j]`, set `match = false`, `break`
   c. If `match`, return `i`
4. Return `-1`

## The Challenge

See [README.md](README.md).

**Next:** [39-replaceall](../39-replaceall/README.md)
