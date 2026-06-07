# 03. Last Character

## What You'll Learn

This exercise teaches you **reverse string access and position calculation**. By the end, you'll understand:
- How to find the last position in a string
- Converting between length and index positions
- Handling edge cases with string boundaries
- The relationship between string length and valid indices

## Theory: String Positions and Indices

### String Length vs Last Index

In Go (and most programming languages), strings are **0-indexed**:

```
String:  "Hello"
Index:    0    1    2    3    4
Char:     H    e    l    l    o
Length:   5
```

**Key Formula:**
```
lastIndex = length - 1
```

### Why `len() - 1`?

If a string has 5 characters (indices 0, 1, 2, 3, 4), the last valid index is 4, which equals `5 - 1`.

### Accessing the Last Character

There are two common approaches:

**Approach 1: Calculate the index**
```go
lastIndex := len(s) - 1
lastChar := s[lastIndex]
```

**Approach 2: Use `for...range` and track**
```go
var lastChar rune
for _, c := range s {
    lastChar = c  // Overwrites each time, ends with last
}
```

## Step-by-Step Approach

1. **Check if the string is empty** - return empty string if so
2. **Calculate the last index**: `len(s) - 1`
3. **Get the last character**: `s[lastIndex]`
4. **Convert to string** and return

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| `s[len(s)]` | Index out of range! | Use `len(s) - 1` |
| Forgetting empty check | Will panic on empty string | Check `s == ""` first |
| Returning `byte` instead of `string` | Type mismatch | Use `string()` conversion |

## Practice Tips

- Draw strings with indices: `[0:H][1:e][2:l][3:l][4:o]`
- Remember: valid indices are `0` to `length-1`
- Test edge cases: empty string, single character

## The Challenge

Write a function that returns the **last character** of a string as a string.

### Expected Function

```go
func LastChar(s string) string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `"Hello"` | `"o"` | Last character |
| `""` | `""` | Empty string |
| `"G"` | `"G"` | Single character (also first and last) |
| `"Go is fun"` | `"n"` | Last character of longer string |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Println(LastChar("Hello"))    // o
    fmt.Println(LastChar(""))         // (empty)
    fmt.Println(LastChar("G"))        // G
    fmt.Println(LastChar("Go is fun")) // n
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. What is the last valid index of a 5-character string?
2. What happens if you access `s[len(s)]`?
3. How do you convert a `byte` to a `string`?

## Next Steps

After completing this, you'll be ready for:
- [79-isempty](../79-isempty/README.md) - Isempty
- [80-toupper](../80-toupper/README.md) - Toupper