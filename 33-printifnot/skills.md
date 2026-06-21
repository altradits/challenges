# Skills for 33-printifnot

## What You'll Learn

**Previous:** [32-printif](../32-printif/skills.md) | **Next:** [34-longestword](../34-longestword/skills.md)

**Challenge:** Return `"G\n"` if string length is < 3, else `"Invalid Input\n"` — the inverse of PrintIf.

## Core Concept: Inverting Conditions and De Morgan's Law

### PrintIf vs PrintIfNot — Side by Side

```go
// PrintIf: return "G" when len == 0 OR len >= 3
if len(str) == 0 || len(str) >= 3 {
    return "G\n"
}
return "Invalid Input\n"

// PrintIfNot: the EXACT INVERSE — swap the return values
if len(str) < 3 {
    return "G\n"
}
return "Invalid Input\n"
```

Notice that `len(str) < 3` is equivalent to `len(str) == 0 || len(str) == 1 || len(str) == 2` — it captures lengths 0, 1, and 2. Empty string (length 0) is included.

### Inverting Compound Conditions — De Morgan's Law

De Morgan's Law tells you how to flip a compound condition:

```
NOT (A OR B)  =  (NOT A) AND (NOT B)
NOT (A AND B) =  (NOT A) OR  (NOT B)
```

Applied here:

```
Original: len == 0 OR len >= 3
Negated:  len != 0 AND len < 3
Simplified: 0 < len < 3  (len is 1 or 2)
```

So `PrintIfNot` returns `"G\n"` when len is 0, 1, or 2 — when the condition for `PrintIf` was false.

### Why This Pattern Matters

Inverse conditions appear everywhere: "allow if condition" vs "deny if condition", "valid if X" vs "invalid if not X". Understanding how to flip logic is a core programming skill.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Copying PrintIf exactly | Same output, wrong result | Use `< 3` not `>= 3` |
| Thinking `< 3` excludes 0 | `len("") = 0`, which is `< 3` | 0 < 3 is true |
| Using `<= 2` | Works but less readable | `< 3` is equivalent and cleaner |

## Algorithm

1. If `len(str) < 3`: return `"G\n"`
2. Else: return `"Invalid Input\n"`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [34-longestword](../34-longestword/README.md)
