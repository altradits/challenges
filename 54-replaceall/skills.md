# Skills for 54-replaceall

**Previous:** [53-findsubstring](../53-findsubstring/skills.md) | **Next:** [55-split](../55-split/skills.md)

**Challenge:** Replace ALL occurrences of substring `old` with `new` in a text string, without using `strings.ReplaceAll`.

## Core Concept: Scan-and-Build Pattern

### What Is It?

The **scan-and-build** pattern walks through a string with a manual index `i`, decides at each position whether to substitute a chunk or copy a single character, and builds the result incrementally. It finds every match — not just the first — because it never stops after the first replacement.

### The Pattern

```go
func ReplaceAll(text, old, new string) string {
    if old == "" {
        return text  // avoid infinite loop on empty pattern
    }
    result := ""
    i := 0
    for i < len(text) {
        // Does `old` start at position i?
        if strings.HasPrefix(text[i:], old) {
            result += new          // substitute `old` with `new`
            i += len(old)          // skip past the matched `old`
        } else {
            result += string(text[i])  // copy one character as-is
            i++
        }
    }
    return result
}
```

### What `strings.HasPrefix(text[i:], old)` Does

`text[i:]` is a slice of the string starting at byte `i`. `strings.HasPrefix` checks whether that slice starts with `old`. This is a safe way to check if `old` appears at position `i` — it won't panic if `old` is longer than the remaining text.

You can also write it manually:

```go
func matchAt(text, pattern string, i int) bool {
    if i+len(pattern) > len(text) {
        return false
    }
    return text[i:i+len(pattern)] == pattern
}
```

### Tracing Through `"banana"` replacing `"na"` with `"NA"`

| i | Match `"na"` at i? | Action | result |
|---|-------------------|--------|--------|
| 0 | "banana"[0:] → "banana": HasPrefix "na"? No | copy 'b', i=1 | `"b"` |
| 1 | "anana": HasPrefix "na"? No | copy 'a', i=2 | `"ba"` |
| 2 | "nana": HasPrefix "na"? Yes | append "NA", i=4 | `"baNA"` |
| 4 | "na": HasPrefix "na"? Yes | append "NA", i=6 | `"baNANA"` |
| 6 | loop ends | | `"baNANA"` |

Result: `"baNANA"` — correct.

### Why Empty `old` Causes an Infinite Loop

If `old == ""`, then `strings.HasPrefix(anything, "")` is always true. The code appends `new` but `i += len("")` means `i += 0` — the index never advances. The loop runs forever. Always guard against this:

```go
if old == "" {
    return text
}
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Forgetting `if old == ""` | Infinite loop | Return early when `old` is empty |
| Using `i++` even after a match | After replacing "na" at position 2, next check is at 3, not 4 — could miss adjacent matches | Use `i += len(old)` on match; `i++` only on non-match |
| Forgetting `string()` wrap | `text[i]` is a byte; `result += text[i]` is a type error | Write `result += string(text[i])` |

## Algorithm

1. If `old == ""`, return `text` unchanged
2. Set `result = ""`, `i = 0`
3. While `i < len(text)`:
   - If `strings.HasPrefix(text[i:], old)`: append `new` to result, advance `i` by `len(old)`
   - Else: append `string(text[i])` to result, advance `i` by 1
4. Return `result`

## The Challenge

See [README.md](README.md).

**Next:** [55-split](../55-split/README.md)
