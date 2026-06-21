# 04. Is Empty

## What You'll Learn

This exercise teaches you **boolean logic and string validation**. By the end, you'll understand:
- How to check if a string has content
- Early return patterns for edge cases
- Boolean return types in Go
- The difference between empty string and whitespace

## Theory: Empty Strings in Go

### What is an Empty String?

An empty string has **zero length** and contains no characters:

```go
s := ""           // Empty string
fmt.Println(len(s)) // 0
```

### Empty String vs Whitespace

| Value | Description | Length |
|-------|-------------|--------|
| `""` | Empty string | 0 |
| `" "` | Space character | 1 |
| `"\t"` | Tab character | 1 |
| `"\n"` | Newline | 1 |
| `"   "` | Multiple spaces | 3 |

**Important:** This exercise checks for **empty string only**, not whitespace!

### Boolean Return Pattern

Functions that check conditions typically return `bool`:

```go
func IsValid(s string) bool {
    if condition {
        return true
    }
    return false
}
```

### Early Return Pattern

Check edge cases first and return immediately:

```go
func Process(s string) string {
    if s == "" {
        return ""  // Handle empty case early
    }
    // Main logic here
    return result
}
```

## Step-by-Step Approach

1. **Iterate** through the string using `for...range`
2. **If any character is found**, return `false` (not empty)
3. **If loop completes** without finding characters, return `true` (empty)

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Checking `len(s) == 0` | Works but defeats the learning purpose | Use iteration to check |
| Returning opposite values | Confusing logic | Empty = true, not empty = false |
| Forgetting to return in all paths | Compile error | Ensure all code paths return |

## Practice Tips

- An empty string has NO characters to iterate over
- If `for...range` runs 0 times, the string is empty
- If it runs 1+ times, the string is NOT empty

## The Challenge

Write a function that returns `true` if a string is empty, `false` otherwise.

### Expected Function

```go
func IsEmpty(s string) bool {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `"Hello"` | `false` | Has characters |
| `""` | `true` | No characters |
| `" "` | `false` | Has a space character |
| `"G"` | `false` | Has one character |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Println(IsEmpty("Hello"))  // false
    fmt.Println(IsEmpty(""))       // true
    fmt.Println(IsEmpty(" "))      // false
    fmt.Println(IsEmpty("G"))      // false
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. How many times does `for...range` run on an empty string?
2. What's the difference between `""` and `" "`?
3. When should you use early returns?

## Next Steps

After completing this, you'll be ready for:
- [17-toupper](../17-toupper/README.md) - Toupper
- [18-tolower](../18-tolower/README.md) - Tolower