# Skills for wordanatomy2

## What You'll Learn

**Previous:** [../32-wordanatomy/skills.md](../32-wordanatomy/skills.md) | **Next:** [../34-slicesintro/README.md](../34-slicesintro/README.md)

**Challenge:** Given a word and caller-provided prefix/suffix arrays, find the *first* matching prefix and first matching suffix, then return a formatted result string.

## Core Concept: First-Match Search and Formatted String Output

### What Changes From wordanatomy?
In wordanatomy (32), the prefix/suffix lists were hard-coded and you used longest-match. Here:
1. The lists are passed as parameters (`pref []string`, `suff []string`).
2. You use **first match** — the first element in the array that matches wins.
3. The return value is a formatted *string* rather than three separate strings.

### Finding the First Match
Stop as soon as you find a match — do not continue scanning:
```go
foundPrefix := ""
for _, p := range pref {
    if len(word) >= len(p) && word[:len(p)] == p {
        foundPrefix = p
        break  // first match wins
    }
}
```

```go
foundSuffix := ""
for _, s := range suff {
    if len(word) >= len(s) && word[len(word)-len(s):] == s {
        foundSuffix = s
        break  // first match wins
    }
}
```

### Returning a Formatted String
Use `fmt.Sprintf` to build the result string with the required format:
```go
import "fmt"

func Wordanatomy(word string, pref []string, suff []string) string {
    // ... find foundPrefix and foundSuffix ...
    return fmt.Sprintf("prefix: %q, suffix: %q", foundPrefix, foundSuffix)
}
```

The `%q` verb wraps the string in double quotes, matching the expected output:
```
prefix: "un", suffix: "able"
```

### Difference Between First-Match and Longest-Match
```
prefixes: ["an", "en", "un"]
word: "understandable"
```
- First-match: scan left to right. `"an"` — does `"understandable"` start with `"an"`? No. `"en"`? No. `"un"`? Yes → use `"un"`.
- Longest-match (challenge 32): check all, keep longest. Same result here but differs in other cases.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Not using `break` after finding a match | Last match wins instead of first | Add `break` immediately after setting `foundPrefix` or `foundSuffix` |
| Using `%v` instead of `%q` for the format string | Missing quotes in output | Use `%q` to get quoted strings |
| Not handling empty input arrays | Loop does nothing, returns empty — this is actually correct behavior | No special handling needed |

## Solving This Challenge

### Algorithm
1. Scan `pref` for the first element that matches the start of `word`; use `break`.
2. Scan `suff` for the first element that matches the end of `word`; use `break`.
3. Return `fmt.Sprintf("prefix: %q, suffix: %q", foundPrefix, foundSuffix)`.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [../34-slicesintro/README.md](../34-slicesintro/README.md)
