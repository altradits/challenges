# 18. Find Last Character

## What You'll Learn

This exercise teaches you **reverse searching and index calculation**. By the end, you'll understand:
- How to search from the end of a string
- Calculating positions from the end
- The difference between first and last occurrence
- When to use reverse iteration

## Theory: Reverse Search

### Finding the Last Occurrence

To find the last occurrence, you can:

**Approach 1: Search forward, keep updating**
```go
lastIndex := -1
for i, c := range s {
    if c == target {
        lastIndex = i  // Keep updating, ends with last
    }
}
return lastIndex
```

**Approach 2: Search backward**
```go
for i := len(s) - 1; i >= 0; i-- {
    if s[i] == target {
        return i
    }
}
return -1
```

### Why Approach 1 is Often Better

- Works with `for...range` (handles UTF-8)
- Simpler logic
- No need to calculate reverse indices

## Step-by-Step Approach

1. **Initialize** `lastIndex` to -1
2. **Iterate** through the string
3. **If match**: update `lastIndex` to current index
4. **Return** `lastIndex` (will be -1 if never found)

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Returning on first match | That's FindChar, not FindLastChar | Keep updating |
| Using byte indexing | Breaks with UTF-8 | Use `for...range` |
| Forgetting -1 initial | Might return 0 incorrectly | Start with -1 |

## Practice Tips

- Test with: `"Hello"`, find `'l'` → 3 (last 'l')
- Test with: `"Hello"`, find `'H'` → 0 (only occurrence)
- Test with: `"Hello"`, find `'z'` → -1

## The Challenge

Write a function that finds the index of the LAST occurrence of a character.

### Expected Function

```go
func FindLastChar(s string, c rune) int {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `("Hello", 'l')` | `3` | Last 'l' at index 3 |
| `("Hello", 'H')` | `0` | Only occurrence |
| `("Hello", 'z')` | `-1` | Not found |
| `("banana", 'a')` | `5` | Last 'a' at index 5 |
| `("", 'a')` | `-1` | Empty string |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Println(FindLastChar("Hello", 'l'))  // 3
    fmt.Println(FindLastChar("Hello", 'H'))  // 0
    fmt.Println(FindLastChar("Hello", 'z'))  // -1
    fmt.Println(FindLastChar("banana", 'a')) // 5
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. How is FindLastChar different from FindChar?
2. Why do we keep updating instead of returning immediately?
3. What should the initial value of lastIndex be?

## Next Steps

After completing this, you'll be ready for:
- **19-replacechar**: Replacing characters
- **20-printif**: Conditional output
