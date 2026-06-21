# 28. Split

## What You'll Learn

This exercise teaches you **tokenizing strings and slice manipulation**. By the end, you'll understand:
- How to split a string into parts
- Working with slices of strings
- Handling edge cases in splitting
- The difference between split and manual parsing

## Theory: String Splitting

### What is Splitting?

Splitting breaks a string into a slice of substrings:

```
"a,b,c" split by "," → ["a", "b", "c"]
"Hello World" split by " " → ["Hello", "World"]
```

### The Split Algorithm

```go
func split(s, sep string) []string {
    var result []string
    start := 0
    for i := 0; i <= len(s); i++ {
        if i == len(s) || s[i:i+len(sep)] == sep {
            result = append(result, s[start:i])
            start = i + len(sep)
            i += len(sep) - 1  // Skip past separator
        }
    }
    return result
}
```

## Step-by-Step Approach

1. **Initialize** empty result slice and start index
2. **Loop** through the string
3. **When separator found**: extract substring, add to result
4. **Update** start index past separator
5. **Return** result

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Off-by-one in slicing | Missing or extra chars | Careful with indices |
| Forgetting last segment | Trailing text missed | Handle end of string |
| Empty separator | Infinite loop | Check for empty sep |

## The Challenge

Write a function that splits a string by a separator.

### Expected Function

```go
func Split(s, sep string) []string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `("a,b,c", ",")` | `["a", "b", "c"]` | Basic split |
| `("Hello World", " ")` | `["Hello", "World"]` | Space separator |
| `("abc", "")` | `["a", "b", "c"]` | Empty sep = split by char |
| `("", ",")` | `[""]` | Empty string |

## Knowledge Check

Before coding, make sure you can answer:
1. How do you extract a substring using slicing?
2. What happens when the separator is at the start or end?
3. How do you handle an empty separator?

## Next Steps

After completing this, you'll be ready for:
- [41-join](../41-join/README.md) - Join
- [42-cameltosnakecase](../42-cameltosnakecase/README.md) - Cameltosnakecase