# Skills for 72-removespaces

## What You'll Learn

**Previous:** [71-printifnot](../71-printifnot/skills.md) | **Next:** [73-retainfirsthalf](../73-retainfirsthalf/skills.md)

**Challenge:** Write `RemoveSpaces(s string) string` that returns a new string with all space characters (`' '`) removed.

## Core Concept: Building a New String by Filtering Characters

### What Is It?

You cannot remove characters from a string in place — Go strings are immutable. To build a modified string, you iterate over the original and collect only the characters you want to keep. This is the **filter** pattern.

### Method 1: `strings.ReplaceAll` (Simple)

The standard library makes this a one-liner:

```go
import "strings"

func RemoveSpaces(s string) string {
    return strings.ReplaceAll(s, " ", "")
}
```

`strings.ReplaceAll(s, old, new)` returns a copy of `s` with every occurrence of `old` replaced by `new`. Replacing `" "` with `""` (empty string) removes all spaces.

### Method 2: Manual Iteration with `strings.Builder` (Teaches the Concept)

For any custom filtering, you need to build the result yourself:

```go
import "strings"

func RemoveSpaces(s string) string {
    var b strings.Builder
    for _, r := range s {
        if r != ' ' {
            b.WriteRune(r)
        }
    }
    return b.String()
}
```

Step by step:
1. `var b strings.Builder` — declare an efficient string accumulator
2. `for _, r := range s` — visit every rune
3. `if r != ' '` — skip spaces; keep everything else
4. `b.WriteRune(r)` — append the character to the builder
5. `return b.String()` — convert the builder to a final string

### Why Not Just Use `+=` String Concatenation?

You could write `result += string(r)`, but this creates a new string in memory on every iteration. For large strings it is very slow. `strings.Builder` uses a growing internal buffer and is the idiomatic Go way.

```go
// Works but slow for long strings:
result := ""
for _, r := range s {
    if r != ' ' {
        result += string(r)   // allocates a new string each time
    }
}
return result

// Better:
var b strings.Builder
for _, r := range s {
    if r != ' ' {
        b.WriteRune(r)
    }
}
return b.String()
```

### `strings.Builder` API

| Method | What it does |
|--------|--------------|
| `b.WriteRune(r)` | Appends a single rune |
| `b.WriteString(s)` | Appends a whole string |
| `b.WriteByte(c)` | Appends a single byte |
| `b.String()` | Returns the accumulated string |

### Diagram: Filtering "Hello World"

```
Input:  H e l l o   W o r l d
        | | | | | | | | | | |
Keep?   Y Y Y Y Y N Y Y Y Y Y   (N = space, skip it)
        | | | | |   | | | | |
Output: H e l l o   W o r l d
        =
        "HelloWorld"
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `strings.Replace` instead of `strings.ReplaceAll` | `Replace` takes a count argument; you must pass `-1` to replace all | Use `strings.ReplaceAll` |
| `if r == " "` (double quotes) | Type mismatch: `r` is a rune, `" "` is a string | Use `r == ' '` (single quotes for rune) |
| Not returning `b.String()` | Builder accumulated the result but you returned the builder itself | Always call `.String()` at the end |
| `b.WriteString(string(r))` | Works but unnecessarily converts rune to string | Use `b.WriteRune(r)` directly |

## Solving This Challenge

### Algorithm

**Simple approach:** `return strings.ReplaceAll(s, " ", "")`

**Manual approach:**
1. Declare a `strings.Builder`
2. For each rune in `s`, if it is not a space, write it to the builder
3. Return `b.String()`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [73-retainfirsthalf](../73-retainfirsthalf/README.md)
