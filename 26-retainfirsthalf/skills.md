# Skills for 26-retainfirsthalf

## What You'll Learn

**Previous:** [25-countrepeats](../25-countrepeats/skills.md) | **Next:** [27-wordcount](../27-wordcount/skills.md)

**Challenge:** Write a function `RetainFirstHalf(s string) string` that returns the first half of the string. For odd-length strings, round down. A single character returns itself; an empty string returns `""`.

## Core Concept: String Slicing and Integer Division

### String Slicing

Go lets you extract a portion of a string using **slice notation**:

```go
s := "Hello World"
s[0:5]    // "Hello"  — characters at indices 0, 1, 2, 3, 4
s[:5]     // "Hello"  — same thing; 0 is the default start
s[6:]     // "World"  — from index 6 to the end
```

The slice `s[start:end]` includes `start` but excludes `end`. So `s[:n]` gives the first `n` characters.

### Integer Division Rounds Down

When you divide two integers in Go, the fractional part is discarded:

```go
6 / 2   // 3
7 / 2   // 3 (not 3.5)
5 / 2   // 2
1 / 2   // 0
```

This is exactly the behaviour you want: the "first half" of an odd-length string should NOT include the middle character.

### Calculating the Half Index

```go
half := len(s) / 2
```

| Input | `len(s)` | `len(s)/2` | First half |
|-------|---------|-----------|------------|
| `"abcd"` | 4 | 2 | `"ab"` |
| `"abcde"` | 5 | 2 | `"ab"` |
| `"Hello World"` | 11 | 5 | `"Hello"` |
| `"A"` | 1 | 0 | `s[:0]` = `""` — but the challenge says return `"A"` |

### The Full Implementation

```go
func RetainFirstHalf(s string) string {
    if len(s) <= 1 {
        return s   // empty or single char: return as-is
    }
    return s[:len(s)/2]
}
```

The guard `len(s) <= 1` handles both `""` and `"A"`. For `"A"`: `len = 1`, so return `"A"` directly.

### Why the Guard Is Needed

For a 1-character string, `len(s)/2 = 0`, making `s[:0]` return `""` — but the test expects `"A"`. The `<= 1` guard short-circuits before the slicing.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `s[:len(s)/2+1]` | Includes the middle character for odd lengths | Use `s[:len(s)/2]` |
| No guard for length 1 | Returns `""` instead of `"A"` | Check `len(s) <= 1` |
| Float division: `len(s) / 2.0` | Compile error — `len` returns `int` | Use integer division: `len(s) / 2` |

## Solving This Challenge

### Algorithm

1. If `len(s) <= 1`, return `s`
2. Return `s[:len(s)/2]`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [27-wordcount](../27-wordcount/README.md)
