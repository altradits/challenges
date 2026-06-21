# Skills for 38-ispalindrome

## What You'll Learn

**Previous:** [37-reversestring](../37-reversestring/skills.md) | **Next:** [39-removespaces](../39-removespaces/skills.md)

**Challenge:** Write a function `IsPalindrome(s string) bool` that returns `true` if the string reads the same forwards and backwards, ignoring case and ignoring non-alphanumeric characters. Empty strings return `true`.

## Core Concept: Composing Earlier Skills to Solve a Harder Problem

### What Is a Palindrome?

A palindrome reads the same forwards and backwards:
- `"racecar"` → `true`
- `"A man a plan a canal Panama"` → `true` (after removing spaces and ignoring case)
- `"Hello"` → `false`

The challenge requires ignoring case and ignoring non-alphanumeric characters (spaces, punctuation).

### The Two-Step Strategy

1. **Clean the string**: keep only alphanumeric characters, converted to lowercase
2. **Compare with its reverse**: if they are equal, it is a palindrome

### Step 1 — Cleaning

Build a new string containing only letters and digits, all lowercased:

```go
cleaned := ""
for _, c := range s {
    if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
        cleaned += string(c)
    } else if c >= 'A' && c <= 'Z' {
        cleaned += string(c + 32)   // lowercase the uppercase letter
    }
}
```

This uses:
- Lowercase range from [34-countalpha](../34-countalpha/skills.md)
- Digit range from [35-checknumber](../35-checknumber/skills.md)
- Case conversion from [33-tolower](../33-tolower/skills.md)

### Step 2 — Reversing

Use the `[]rune` two-pointer technique from [37-reversestring](../37-reversestring/skills.md):

```go
func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
```

### Step 3 — Compare

```go
return cleaned == reverse(cleaned)
```

### Full Implementation

```go
func IsPalindrome(s string) bool {
    cleaned := ""
    for _, c := range s {
        if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
            cleaned += string(c)
        } else if c >= 'A' && c <= 'Z' {
            cleaned += string(c + 32)
        }
    }
    // Two-pointer check (avoids building a reversed copy)
    runes := []rune(cleaned)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        if runes[i] != runes[j] {
            return false
        }
    }
    return true
}
```

The two-pointer check is more efficient: as soon as a mismatch is found it returns `false` without building the full reversed string.

### Why Empty Returns `true`

An empty string has no characters, so there is nothing to disagree — it reads the same forwards and backwards vacuously. The loop runs zero times and returns `true`.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Case-sensitive comparison | `"Racecar"` fails | Convert to lowercase first |
| Not removing spaces/punctuation | `"A man..."` fails | Keep only alphanumeric characters |
| Comparing entire original string | Fails when spaces or case differ | Compare `cleaned` to its reverse |

## Solving This Challenge

### Algorithm

1. Build `cleaned`: lowercase alphanumeric characters only
2. Use two-pointer to compare `cleaned[i]` with `cleaned[j]`; return `false` on mismatch
3. Return `true` if no mismatch found

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [39-removespaces](../39-removespaces/README.md)
