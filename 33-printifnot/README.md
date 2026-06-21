# 21. Print If Not

## What You'll Learn

This exercise teaches you **inverse conditional logic**. By the end, you'll understand:
- How to invert conditions
- The relationship between `if` and `else` branches
- Logical operators (`<` vs `>=`)
- Writing conditions from different perspectives

## Theory: Inverse Conditions

### Inverting Logic

| Original Condition | Inverse |
|-------------------|---------|
| `len(s) >= 3` | `len(s) < 3` |
| `len(s) == 0` | `len(s) != 0` |
| `c == 'a'` | `c != 'a'` |

### De Morgan's Laws

For compound conditions:
- `!(A && B)` = `!A || !B`
- `!(A || B)` = `!A && !B`

### This Exercise's Logic

Original (PrintIf):
```go
if len(s) == 0 || len(s) >= 3 {
    return "G\n"
} else {
    return "Invalid Input\n"
}
```

Inverse (PrintIfNot):
```go
if len(s) < 3 {
    return "G\n"
} else {
    return "Invalid Input\n"
}
```

## Step-by-Step Approach

1. **Check** if string length is less than 3
2. **If true**: return `"G\n"`
3. **Else**: return `"Invalid Input\n"`

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Using same condition as PrintIf | Would give same results | Use inverse: `< 3` |
| Forgetting empty string | `""` has length 0, which is < 3 | Empty should return "G" |
| Wrong boundary | `<= 2` vs `< 3` | Both work, but `< 3` is clearer |

## Practice Tips

- Test with: `"abcdefz"` (len=7) → "Invalid Input"
- Test with: `"abc"` (len=3) → "Invalid Input"
- Test with: `""` (len=0) → "G"
- Test with: `"14"` (len=2) → "G"

## The Challenge

Write a function that returns "G" if length < 3, else "Invalid Input".

### Expected Function

```go
func PrintIfNot(str string) string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `"abcdefz"` | `"Invalid Input\n"` | Length 7 >= 3 |
| `"abc"` | `"Invalid Input\n"` | Length 3 >= 3 |
| `""` | `"G\n"` | Length 0 < 3 |
| `"14"` | `"G\n"` | Length 2 < 3 |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Print(PrintIfNot("abcdefz"))
    fmt.Print(PrintIfNot("abc"))
    fmt.Print(PrintIfNot(""))
    fmt.Print(PrintIfNot("14"))
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. What's the inverse of `>= 3`?
2. How does PrintIfNot relate to PrintIf?
3. Why does empty string return "G"?

## Next Steps

After completing this, you'll be ready for:
- [34-longestword](../34-longestword/README.md) - Longestword
- [35-searchreplace](../35-searchreplace/README.md) - Searchreplace