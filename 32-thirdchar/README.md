# 32. Every Third Character

## What You'll Learn

This exercise teaches you **modular indexing and selective extraction**. By the end, you'll understand:
- Using modulo operator for periodic patterns
- Extracting characters at regular intervals
- Handling edge cases with short strings
- Building strings from selective characters

## Theory: Modulo for Patterns

### The Modulo Operator

`%` gives the remainder of division:

```
index % 3 == 0  // Every 3rd character (0, 3, 6, 9...)
index % 3 == 1  // 1st, 4th, 7th...
index % 3 == 2  // 2nd, 5th, 8th...
```

### Extracting Every Nth Character

```go
for i, c := range s {
    if i % n == n-1 {  // Every nth (1-indexed)
        result += string(c)
    }
}
```

## Step-by-Step Approach

1. **Iterate** with index using `for...range`
2. **Check** if `(index + 1) % 3 == 0` (every 3rd)
3. **Append** matching characters to result
4. **Return** result with newline

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Using `i % 3 == 0` | Gets 0, 3, 6... (1st, 4th, 7th) | Use `(i+1) % 3 == 0` for 3rd, 6th, 9th |
| Forgetting newline | Output missing `\n` | Add `\n` at end |
| Empty string handling | Should return just `\n` | Check length first |

## The Challenge

Write a function that returns every 3rd character of a string.

### Expected Function

```go
func ThirdChar(s string) string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `"123456789"` | `"369\n"` | 3rd, 6th, 9th |
| `""` | `"\n"` | Empty string |
| `"a b c d e f"` | `"b e\n"` | Every 3rd including spaces |
| `"12"` | `"\n"` | No 3rd character |

## Knowledge Check

Before coding, make sure you can answer:
1. What does `i % 3` give you?
2. How do you get the 3rd, 6th, 9th characters?
3. What should empty string return?

## Next Steps

After completing this, you'll be ready for:
- **33-zipstring**: Run-length encoding
- **34-saveandmiss**: Periodic skipping
