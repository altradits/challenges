# 26. Find Substring

## What You'll Learn

This exercise teaches you **pattern matching algorithms**. By the end, you'll understand:
- How to search for multi-character patterns
- The sliding window technique
- Nested loop patterns for pattern matching
- Handling edge cases in substring search

## Theory: Substring Search

### What is Substring Search?

Finding if a pattern exists within a larger text:

```
Text:    "Hello World"
Pattern: "World"
Result:  6 (starts at index 6)
```

### The Sliding Window Approach

Check each possible starting position:

```
Text:    "Hello"
          ↑
          Check "Hello" (len=5)
          ↑
          Check "ello" (len=4)
           ↑
          Check "llo" (len=3)
            ↑
          Check "lo" (len=2)
             ↑
          Check "o" (len=1)
```

### Algorithm

```go
for i := 0; i <= len(text)-len(pattern); i++ {
    match := true
    for j := 0; j < len(pattern); j++ {
        if text[i+j] != pattern[j] {
            match = false
            break
        }
    }
    if match {
        return i
    }
}
return -1
```

## Step-by-Step Approach

1. **Check** if pattern is empty → return 0
2. **Check** if pattern longer than text → return -1
3. **Loop** through each possible start position
4. **Compare** pattern with text at that position
5. **If match**: return start index
6. **If no match**: return -1

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Loop condition wrong | Might go out of bounds | `i <= len(text)-len(pattern)` |
| Forgetting empty pattern | Should return 0 | Handle edge case |
| Comparing wrong indices | Off-by-one errors | `text[i+j]` vs `pattern[j]` |

## Practice Tips

- Test with: `"Hello World"`, find `"World"` → 6
- Test with: `"Hello World"`, find `"hello"` → -1 (case sensitive)
- Test with: `"banana"`, find `"ana"` → 1

## The Challenge

Write a function that finds the first occurrence of a substring.

### Expected Function

```go
func FindSubstring(text, substring string) int {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `("Hello World", "World")` | `6` | Found at index 6 |
| `("Hello World", "hello")` | `-1` | Case sensitive, not found |
| `("Hello World", "")` | `0` | Empty substring |
| `("Hello World", "xyz")` | `-1` | Not found |
| `("banana", "ana")` | `1` | Found at index 1 |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Println(FindSubstring("Hello World", "World"))
    fmt.Println(FindSubstring("Hello World", "hello"))
    fmt.Println(FindSubstring("Hello World", ""))
    fmt.Println(FindSubstring("Hello World", "xyz"))
    fmt.Println(FindSubstring("banana", "ana"))
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. What is the sliding window technique?
2. Why do we need nested loops?
3. What's the maximum starting position?

## Next Steps

After completing this, you'll be ready for:
- **27-replaceall**: Global string replacement
- **28-split**: Splitting strings
