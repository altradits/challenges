# 34. Save and Miss

## What You'll Learn

This exercise teaches you **periodic character selection and skipping**. By the end, you'll understand:
- How to process strings in chunks
- Alternating between "save" and "miss" modes
- Using counters to track position within chunks
- Building strings from selective segments

## Theory: Periodic Processing

### The Save/Miss Pattern

Process the string in groups of `n` characters:
- Group 1 (0 to n-1): SAVE
- Group 2 (n to 2n-1): MISS (skip)
- Group 3 (2n to 3n-1): SAVE
- And so on...

### Tracking Position Within Chunk

```go
pos := 0  // Position within current chunk
for _, c := range s {
    if pos < n {  // In "save" zone
        result += string(c)
    }
    pos++
    if pos == 2*n {  // Completed save+miss cycle
        pos = 0
    }
}
```

## Step-by-Step Approach

1. **Check** if n <= 0, return original
2. **Initialize** position tracker
3. **Iterate** through characters
4. **If** position < n: save character
5. **Increment** position, reset at 2*n
6. **Return** result

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Wrong boundary | Using `<=` instead of `<` | Use `< n` for save zone |
| Not resetting | Position keeps growing | Reset at `2*n` |
| Forgetting edge case | n=0 causes issues | Return original if n <= 0 |

## The Challenge

Write a function that saves every other chunk of n characters.

### Expected Function

```go
func SaveAndMiss(s string, n int) string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `("123456789", 3)` | `"123789"` | Save 123, miss 456, save 789 |
| `("abcdefghijklmnopqrstuvwyz", 3)` | `"abcghimnostuz"` | Multiple cycles |
| `("", 3)` | `""` | Empty string |
| `("hello", 0)` | `"hello"` | n=0 returns original |

## Knowledge Check

Before coding, make sure you can answer:
1. How do you track position within a chunk?
2. When do you reset the position tracker?
3. What happens when n=0?

## Next Steps

After completing this, you'll be ready for:
- [110-reversestrcap](../110-reversestrcap/README.md) - Reversestrcap
- [111-union](../111-union/README.md) - Union