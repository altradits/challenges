# 36. Union of Characters

## What You'll Learn

This exercise teaches you **set union operations on strings**. By the end, you'll understand:
- How to compute the union of two strings
- Removing duplicate characters while preserving order
- Using maps for membership testing
- The difference between union and intersection

## Theory: Set Union

### What is Union?

Union combines elements from both sets, without duplicates:

```
"abc" ∪ "bcd" = "abcd"
"hello" ∪ "world" = "helowrd"
```

### Preserving Order

The union should preserve the order characters first appear:

```
"zpadinton" ∪ "paqefwtdjetyiytjneytjoeyjnejeyj"
= "zpadintoqefwjy"
```

### Using Maps for Lookup

```go
seen := make(map[rune]bool)
for _, c := s1 + s2 {
    if !seen[c] {
        seen[c] = true
        result += string(c)
    }
}
```

## Step-by-Step Approach

1. **Initialize** empty map and result string
2. **Iterate** through first string, add unique chars
3. **Iterate** through second string, add unique chars
4. **Return** result

## The Challenge

Write a function that returns the union of two strings.

### Expected Function

```go
func Union(s1, s2 string) string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output |
|-------|-----------------|
| `("zpadinton", "paqefwtdjetyiytjneytjoeyjnejeyj")` | `"zpadintoqefwjy"` |
| `("ddf6vewg64f", "gtwthgdwthdwfteewhrtag6h4ffdhsd")` | `"df6vewg4thras"` |

## Next Steps

After completing this, you'll be ready for:
- [112-inter](../112-inter/README.md) - Inter
- [113-stringbuilder](../113-stringbuilder/README.md) - Stringbuilder