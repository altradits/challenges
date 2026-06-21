# Skills for wordanatomy

## What You'll Learn

**Previous:** [87-splitjoin](../87-splitjoin/skills.md) | **Next:** [89-wordanatomy2](../89-wordanatomy2/skills.md)

**Challenge:** Given a word, detect and return its prefix, root, and suffix from fixed lists, using the longest-match rule for prefixes and suffixes.

## Core Concept: String Prefix and Suffix Matching

### What Is It?
Check whether a string *starts with* a given prefix or *ends with* a given suffix by slicing. Since no imports are allowed, you cannot use `strings.HasPrefix` — you must use string indexing directly.

### Checking a Prefix Manually
```go
func hasPrefix(word, prefix string) bool {
    if len(word) < len(prefix) {
        return false
    }
    return word[:len(prefix)] == prefix
}
```

### Checking a Suffix Manually
```go
func hasSuffix(word, suffix string) bool {
    if len(word) < len(suffix) {
        return false
    }
    return word[len(word)-len(suffix):] == suffix
}
```

### The Longest Match Rule
When multiple prefixes could match, use the longest one. Sort them by length descending, or check all and keep the longest found:

```go
prefixes := []string{"un", "re", "pre", "mis", "dis", "over", "under", "anti", "inter", "sub"}

foundPrefix := ""
for _, p := range prefixes {
    if hasPrefix(word, p) && len(p) > len(foundPrefix) {
        foundPrefix = p
    }
}
```

### Extracting the Root
Once you know the prefix and suffix, the root is what remains in the middle:
```go
root := word[len(foundPrefix) : len(word)-len(foundSuffix)]
```

If `foundSuffix == ""` then `len(word)-0 = len(word)`, so `word[len(prefix):]` is correct.

### Multiple Return Values
Go functions can return multiple values. Use parentheses around the return types:
```go
func WordAnatomy(word string) (string, string, string) {
    // ...
    return prefix, root, suffix
}
```

Call it and capture all three:
```go
prefix, root, suffix := WordAnatomy("unhappy")
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Not checking `len(word) >= len(prefix)` | Slice out of bounds panic | Always check length before slicing |
| Using first-match instead of longest-match | `"underpaid"` → prefix `"un"` instead of `"under"` | Track the longest found, not the first |
| Computing root before finding both prefix and suffix | Root length is wrong | Find prefix and suffix first, then compute root |

## Solving This Challenge

### Algorithm
1. Search all prefixes; keep the longest one that matches the start of `word`.
2. Search all suffixes; keep the longest one that matches the end of `word`.
3. Compute root as `word[len(prefix) : len(word)-len(suffix)]`.
4. Return `prefix, root, suffix`.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [89-wordanatomy2](../89-wordanatomy2/README.md)
