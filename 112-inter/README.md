# 37. Intersection of Characters

## What You'll Learn

This exercise teaches you **set intersection operations on strings**. By the end, you'll understand:
- How to find common characters between two strings
- The difference between intersection and union
- Preserving order from the first string
- Using maps for efficient lookups

## Theory: Set Intersection

### What is Intersection?

Intersection finds elements present in BOTH sets:

```
"abc" ∩ "bcd" = "bc"
"hello" ∩ "world" = "lo"
```

### Preserving First String Order

The result should follow the order of the FIRST string:

```
"padinton" ∩ "paqefwtdjetyiytjneytjoeyjnejeyj"
= "padinto"  (chars from s1 that also appear in s2)
```

## Step-by-Step Approach

1. **Build** a set of characters from second string
2. **Iterate** through first string
3. **If** char exists in second string AND not yet added: append
4. **Return** result

## The Challenge

Write a function that returns characters common to both strings.

### Expected Function

```go
func Intersection(s1, s2 string) string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output |
|-------|-----------------|
| `("padinton", "paqefwtdjetyiytjneytjoeyjnejeyj")` | `"padinto"` |
| `("ddf6vewg64f", "gtwthgdwthdwfteewhrtag6h4ffdhsd")` | `"df6ewg4"` |

## Next Steps

After completing this, you'll be ready for:
- [113-stringbuilder](../113-stringbuilder/README.md) - Stringbuilder
- [114-stringsplit](../114-stringsplit/README.md) - Stringsplit