# 07. Count Alpha

## What You'll Learn

This exercise teaches you **character classification and counting**. By the end, you'll understand:
- How to identify alphabetic characters using ASCII ranges
- Counting specific types of characters
- Combining multiple conditions
- The difference between letters, digits, and symbols

## Theory: Character Classification

### What is an Alphabetic Character?

In ASCII, alphabetic characters are:

| Type | Range | Characters |
|------|-------|------------|
| Uppercase | 65-90 | A, B, C, ..., Z |
| Lowercase | 97-122 | a, b, c, ..., z |

### Checking if a Character is a Letter

```go
func isAlpha(c rune) bool {
    return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}
```

### Character Type Summary

| Type | Check | Example |
|------|-------|---------|
| Digit | `c >= '0' && c <= '9'` | '5' |
| Uppercase | `c >= 'A' && c <= 'Z'` | 'H' |
| Lowercase | `c >= 'a' && c <= 'z'` | 'e' |
| Alphabetic | uppercase OR lowercase | 'H', 'e' |
| Space | `c == ' '` | ' ' |
| Punctuation | None of the above | '!', '.' |

## Step-by-Step Approach

1. **Initialize** a counter to 0
2. **Iterate** through each character
3. **Check** if character is alphabetic (uppercase OR lowercase)
4. **Increment** counter if true
5. **Return** the count

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Only checking lowercase | Misses uppercase letters | Check BOTH ranges |
| Using `isalpha` from strings | Not allowed in this exercise | Use ASCII comparisons |
| Counting all characters | Should only count letters | Add condition check |

## Practice Tips

- Draw the ASCII table for A-Z and a-z
- Test with: `"Hello"` (5), `"H1e2l3l4o"` (5), `"123"` (0)
- Remember: spaces, numbers, and symbols are NOT alphabetic

## The Challenge

Write a function that counts only alphabetic characters in a string.

### Expected Function

```go
func CountAlpha(s string) int {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `"Hello world"` | `10` | H,e,l,l,o,w,o,r,l,d |
| `"H e l l o"` | `5` | Only letters, spaces ignored |
| `"H1e2l3l4o"` | `5` | Numbers ignored |
| `"123!@#"` | `0` | No letters |
| `""` | `0` | Empty string |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Println(CountAlpha("Hello world"))  // 10
    fmt.Println(CountAlpha("H e l l o"))    // 5
    fmt.Println(CountAlpha("H1e2l3l4o"))    // 5
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. What ASCII range covers uppercase letters?
2. What ASCII range covers lowercase letters?
3. How do you combine two conditions with OR?

## Next Steps

After completing this, you'll be ready for:
- **08-checknumber**: Detecting digits
- **09-countvowels**: Counting specific letters
