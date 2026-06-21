# 16. Find Character

## What You'll Learn

This exercise teaches you **searching for specific characters and index tracking**. By the end, you'll understand:
- How to search for a character in a string
- Returning the position (index) of a match
- Handling "not found" cases
- The difference between first occurrence and all occurrences

## Theory: Character Search

### What is a Search?

A search looks for a **specific value** in a collection and returns its position:

```
String:  "Hello World"
Search:  'o'
Result:  4 (first 'o' is at index 4)
```

### Index vs Position

| Term | Meaning | Example |
|------|---------|---------|
| Index | 0-based position | `"Hello"[4]` = `'o'` |
| Position | 1-based (human) | 5th character |

**In programming, we use 0-based indices.**

### Not Found Convention

When a search fails, we need a way to indicate "not found":

| Approach | Value | Pro | Con |
|----------|-------|-----|-----|
| Return -1 | `-1` | Common convention | Must remember to check |
| Return error | `error` | Explicit | More complex |
| Return bool + index | Two values | Clear | Changes function signature |

**This exercise uses `-1` for "not found".**

## Step-by-Step Approach

1. **Iterate** through the string with index
2. **Compare** each character with target
3. **If match**: return the index immediately
4. **If no match after loop**: return -1

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Returning the character | Should return position | Return index (int) |
| Forgetting -1 case | Function might not return | Always have fallback |
| Using `strings.Index` | Defeats learning purpose | Manual search |

## Practice Tips

- Test with: `"Hello"`, find `'l'` → index 2
- Test with: `"Hello"`, find `'z'` → -1
- Test with: `"Hello"`, find `'H'` → index 0
- Remember: first occurrence only!

## The Challenge

Write a function that finds the index of the first occurrence of a character.

### Expected Function

```go
func FindChar(s string, c rune) int {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `("Hello", 'l')` | `2` | First 'l' at index 2 |
| `("Hello", 'H')` | `0` | First character |
| `("Hello", 'z')` | `-1` | Not found |
| `("", 'a')` | `-1` | Empty string |
| `("banana", 'a')` | `1` | First 'a' at index 1 |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Println(FindChar("Hello", 'l'))  // 2
    fmt.Println(FindChar("Hello", 'H'))  // 0
    fmt.Println(FindChar("Hello", 'z'))  // -1
    fmt.Println(FindChar("", 'a'))       // -1
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. What does "first occurrence" mean?
2. Why return -1 instead of 0 for "not found"?
3. How do you get both index and character in `for...range`?

## Next Steps

After completing this, you'll be ready for:
- [44-countchar](../44-countchar/README.md) - Countchar
- [45-findlastchar](../45-findlastchar/README.md) - Findlastchar