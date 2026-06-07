# 22. Longest Word

## What You'll Learn

This exercise teaches you **finding maximum values and word extraction**. By the end, you'll understand:
- How to track the longest element while iterating
- Splitting strings into words
- Comparing string lengths
- Handling edge cases (empty strings, ties)

## Theory: Finding Maximums

### The Max-Finding Pattern

To find the maximum value in a collection:

```go
max := ""           // Initialize with empty/zero value
for _, item := range items {
    if len(item) > len(max) {
        max = item  // Update when we find larger
    }
}
return max
```

### String Length Comparison

```go
len("Hello") > len("Hi")    // 5 > 2 = true
len("Go") > len("Hello")    // 2 > 5 = false
len("") > len("a")          // 0 > 1 = false
```

## Step-by-Step Approach

1. **Split** the string into words (by spaces)
2. **Initialize** `longest` to empty string
3. **Iterate** through words
4. **If** current word is longer than `longest`: update `longest`
5. **Return** `longest`

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Comparing strings directly | `"b" > "aa"` is true (lexicographic) | Compare `len()` |
| Not handling empty | Might return garbage | Initialize to `""` |
| Forgetting ties | First longest wins (usually fine) | Document the behavior |

## Practice Tips

- Test with: `"Hello World"` → `"Hello"` (both 5, first wins)
- Test with: `"Go is fun"` → `"fun"` (3 letters)
- Test with: `""` → `""`

## The Challenge

Write a function that returns the longest word in a string.

### Expected Function

```go
func LongestWord(s string) string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `"Hello World"` | `"Hello"` | Both 5 chars, first wins |
| `"Go is fun"` | `"fun"` | 3 chars, longest |
| `""` | `""` | Empty string |
| `"a bb ccc dddd"` | `"dddd"` | 4 chars |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Println(LongestWord("Hello World"))
    fmt.Println(LongestWord("Go is fun"))
    fmt.Println(LongestWord(""))
    fmt.Println(LongestWord("a bb ccc dddd"))
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. How do you compare string lengths?
2. What should the initial value of `longest` be?
3. What happens with ties (two words same length)?

## Next Steps

After completing this, you'll be ready for:
- [98-searchreplace](../98-searchreplace/README.md) - Searchreplace
- [99-cleanlist](../99-cleanlist/README.md) - Cleanlist