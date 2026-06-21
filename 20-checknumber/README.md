# 08. Check Number

## What You'll Learn

This exercise teaches you **digit detection and early returns**. By the end, you'll understand:
- How to identify numeric characters using ASCII
- The early return pattern for efficiency
- Boolean logic for "any match" scenarios
- The difference between counting and detecting

## Theory: Numeric Characters

### ASCII Digit Range

Digits in ASCII are contiguous:

```
'0' = 48
'1' = 49
...
'9' = 57
```

**Range check:** `c >= '0' && c <= '9'`

### Early Return Pattern

When you need to find if **ANY** element matches a condition, return immediately upon finding the first match:

```go
func hasMatch(s string) bool {
    for _, c := range s {
        if condition(c) {
            return true  // Found one, stop looking
        }
    }
    return false  // Checked all, none matched
}
```

### Counting vs Detecting

| Goal | Approach | Return Type |
|------|----------|-------------|
| Count all matches | Count every match, return total | `int` |
| Detect any match | Return on first match | `bool` |

## Step-by-Step Approach

1. **Iterate** through each character
2. **Check** if character is a digit (`'0'` to `'9'`)
3. **Return true** immediately if a digit is found
4. **Return false** after checking all characters

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Counting instead of detecting | Returns count, not boolean | Return `true` on first match |
| Checking `c == '0' || c == '1' || ...` | Tedious and error-prone | Use range check |
| Forgetting to return false | Function might not return | Always have a fallback return |

## Practice Tips

- Test with: `"Hello"` (false), `"Hello1"` (true), `"123"` (true)
- The function stops at the FIRST digit found
- Empty string should return false

## The Challenge

Write a function that returns `true` if the string contains any digit.

### Expected Function

```go
func CheckNumber(s string) bool {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `"Hello"` | `false` | No digits |
| `"Hello1"` | `true` | Contains '1' |
| `"123"` | `true` | All digits |
| `""` | `false` | Empty string |
| `"abc!@#"` | `false` | No digits |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Println(CheckNumber("Hello"))   // false
    fmt.Println(CheckNumber("Hello1"))  // true
    fmt.Println(CheckNumber("123"))     // true
    fmt.Println(CheckNumber(""))        // false
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. What ASCII range represents digits?
2. What's the difference between counting and detecting?
3. Why is early return more efficient?

## Next Steps

After completing this, you'll be ready for:
- [21-countvowels](../21-countvowels/README.md) - Countvowels
- [22-reversestring](../22-reversestring/README.md) - Reversestring