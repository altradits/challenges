# Skills for 70-stringtrim

**Previous:** [69-stringreplace](../69-stringreplace/skills.md) | **Next:** [71-stringcontains](../71-stringcontains/skills.md)

**Challenge:** Implement `Trim(s string) string` that removes leading and trailing whitespace (spaces, tabs, newlines) without using `strings.TrimSpace`.

## Core Concept: Find-First-Non-Space + Find-Last-Non-Space + Slice

### What Is It?

Trimming whitespace does NOT rebuild the string character by character. Instead, it:
1. Finds the index of the first non-whitespace character (the left boundary)
2. Finds the index of the last non-whitespace character (the right boundary)
3. Returns the slice between those boundaries

This is an O(n) operation with no string building — just index arithmetic and one slice.

### `unicode.IsSpace` — The Whitespace Check

`unicode.IsSpace(r)` returns true for all whitespace characters including:
- `' '` (space)
- `'\t'` (tab)
- `'\n'` (newline)
- `'\r'` (carriage return)
- `'\v'` (vertical tab)
- `'\f'` (form feed)

```go
import "unicode"

fmt.Println(unicode.IsSpace(' '))   // true
fmt.Println(unicode.IsSpace('\t'))  // true
fmt.Println(unicode.IsSpace('\n'))  // true
fmt.Println(unicode.IsSpace('a'))   // false
fmt.Println(unicode.IsSpace('!'))   // false
```

### The Implementation

```go
func Trim(s string) string {
    runes := []rune(s)
    n := len(runes)

    // Find start: first non-space index
    start := 0
    for start < n && unicode.IsSpace(runes[start]) {
        start++
    }

    // If all spaces, return empty
    if start == n {
        return ""
    }

    // Find end: last non-space index (work from right)
    end := n - 1
    for end > start && unicode.IsSpace(runes[end]) {
        end--
    }

    return string(runes[start : end+1])
}
```

Note: `end+1` because slice `[a:b]` is exclusive of `b`.

### Working Directly on Bytes (for ASCII-only input)

If the string is ASCII-only, you can work directly on bytes:

```go
func Trim(s string) string {
    start := 0
    for start < len(s) && (s[start] == ' ' || s[start] == '\t' || s[start] == '\n') {
        start++
    }
    end := len(s) - 1
    for end > start && (s[end] == ' ' || s[end] == '\t' || s[end] == '\n') {
        end--
    }
    if start > end {
        return ""
    }
    return s[start : end+1]
}
```

### Tracing Through `"  hello  "`

```
runes: [' ', ' ', 'h', 'e', 'l', 'l', 'o', ' ', ' ']
index:   0    1    2    3    4    5    6    7    8

Forward scan:
start=0: runes[0]=' ' → advance
start=1: runes[1]=' ' → advance
start=2: runes[2]='h' → stop. start=2

Backward scan (from end=8):
end=8: runes[8]=' ' → retreat
end=7: runes[7]=' ' → retreat
end=6: runes[6]='o' → stop. end=6

Return runes[2:7] = "hello"
```

### The Standard Library Version

`strings.TrimSpace(s)` does exactly this. The `strings` package also offers:
- `strings.Trim(s, cutset)` — trims any character in `cutset`
- `strings.TrimLeft(s, cutset)` — trims only from the left
- `strings.TrimRight(s, cutset)` — trims only from the right
- `strings.TrimPrefix(s, prefix)` — removes a specific prefix
- `strings.TrimSuffix(s, suffix)` — removes a specific suffix

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Off-by-one in `end+1` | Misses last character | `s[start:end+1]` not `s[start:end]` |
| Not checking `start == n` | Panics when string is all whitespace | Return `""` if `start >= n` after left scan |
| Using `s[i]` for whitespace check | Misses multi-byte whitespace | Use `unicode.IsSpace` on runes |

## Algorithm

1. Convert to `[]rune`
2. Left scan: advance `start` while `unicode.IsSpace(runes[start])`
3. If `start >= len(runes)`, return `""`
4. Right scan: retreat `end` from `len-1` while `unicode.IsSpace(runes[end])`
5. Return `string(runes[start : end+1])`

## The Challenge

See [README.md](README.md).

**Next:** [71-stringcontains](../71-stringcontains/README.md)
