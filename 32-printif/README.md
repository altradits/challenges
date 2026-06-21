# 20. Print If

## What You'll Learn

This exercise teaches you **conditional output based on string length**. By the end, you'll understand:
- Using `len()` to get string length
- Conditional logic with `if/else`
- Returning different values based on conditions
- The ternary-like pattern in Go

## Theory: Conditional Logic

### The `if/else` Pattern

Go uses `if/else` for conditional branching:

```go
if condition {
    // Do this if true
} else {
    // Do this if false
}
```

### String Length Check

```go
len(s) >= 3    // Length is 3 or more
len(s) < 3     // Length is less than 3
len(s) == 0    // Empty string
```

### Multiple Conditions

You can check multiple conditions:

```go
if len(s) == 0 || len(s) >= 3 {
    // Empty OR 3+ characters
}
```

## Step-by-Step Approach

1. **Check** if string length is 0 OR >= 3
2. **If true**: return `"G\n"`
3. **Else**: return `"Invalid Input\n"`

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Using `=` instead of `==` | Assignment, not comparison | Use `==` |
| Forgetting empty case | Empty string should return "G" | Include `len(s) == 0` |
| Wrong comparison | `>` instead of `>=` | Use `>= 3` |

## Practice Tips

- Test with: `"abc"` (len=3) → "G"
- Test with: `"ab"` (len=2) → "Invalid Input"
- Test with: `""` (len=0) → "G"

## The Challenge

Write a function that returns "G" if length is 0 or >= 3, else "Invalid Input".

### Expected Function

```go
func PrintIf(str string) string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `"abcdefz"` | `"G\n"` | Length 7 >= 3 |
| `"abc"` | `"G\n"` | Length 3 >= 3 |
| `""` | `"G\n"` | Length 0 |
| `"14"` | `"Invalid Input\n"` | Length 2 < 3 |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Print(PrintIf("abcdefz"))
    fmt.Print(PrintIf("abc"))
    fmt.Print(PrintIf(""))
    fmt.Print(PrintIf("14"))
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. What's the difference between `=` and `==`?
2. How do you check for empty string?
3. What does `||` mean?

## Next Steps

After completing this, you'll be ready for:
- [33-printifnot](../33-printifnot/README.md) - Printifnot
- [34-longestword](../34-longestword/README.md) - Longestword