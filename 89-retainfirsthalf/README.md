# 14. Retain First Half

## What You'll Learn

This exercise teaches you **string slicing and length calculations**. By the end, you'll understand:
- How to extract substrings using slicing
- Integer division and rounding
- Handling odd-length strings
- Edge cases with short strings

## Theory: String Slicing

### What is Slicing?

Slicing extracts a portion of a string using indices:

```go
s := "Hello World"
firstFive := s[0:5]   // "Hello"
lastFive := s[6:11]   // "World"
```

### Slice Syntax

```go
s[start:end]   // Characters from start to end-1
s[:end]        // From beginning to end-1
s[start:]      // From start to end of string
s[:]           // The whole string
```

### Integer Division

In Go, integer division **rounds down** (truncates toward zero):

```go
5 / 2  // 2 (not 2.5)
7 / 2  // 3
1 / 2  // 0
```

### Calculating the Half

```go
halfIndex := len(s) / 2  // Integer division rounds down
firstHalf := s[:halfIndex]
```

| Length | Half Index | First Half |
|--------|------------|------------|
| 4 | 2 | `s[0:2]` |
| 5 | 2 | `s[0:2]` (rounds down) |
| 6 | 3 | `s[0:3]` |
| 1 | 0 | `s[0:0]` = `""` |

## Step-by-Step Approach

1. **Check** if length is 0 or 1 → return string as-is
2. **Calculate** half index: `len(s) / 2`
3. **Return** `s[:halfIndex]`

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Using float division | `len(s) / 2.0` gives float | Use integer division |
| Forgetting edge cases | `""` or `"a"` would break | Check length first |
| Using `s[0:len(s)/2+1]` | Would include middle char for odd length | Use exact half |

## Practice Tips

- Test with even: `"abcd"` → `"ab"` (half = 2)
- Test with odd: `"abcde"` → `"ab"` (half = 2, rounds down)
- Test with 1 char: `"a"` → `"a"` (special case)
- Test with empty: `""` → `""` (special case)

## The Challenge

Write a function that returns the first half of a string.

### Expected Function

```go
func RetainFirstHalf(s string) string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `"This is the 1st halfThis is the 2nd half"` | `"This is the 1st half"` | First half of 38 chars |
| `"A"` | `"A"` | Single character |
| `""` | `""` | Empty string |
| `"Hello World"` | `"Hello"` | First 5 of 11 chars |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
    fmt.Println(RetainFirstHalf("A"))
    fmt.Println(RetainFirstHalf(""))
    fmt.Println(RetainFirstHalf("Hello World"))
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. What is `7 / 2` in integer division?
2. How do you slice the first half of a string?
3. What should happen with a 1-character string?

## Next Steps

After completing this, you'll be ready for:
- [90-wordcount](../90-wordcount/README.md) - Wordcount
- [91-findchar](../91-findchar/README.md) - Findchar