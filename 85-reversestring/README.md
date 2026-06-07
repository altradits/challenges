# 10. Reverse String

## What You'll Learn

This exercise teaches you **string reversal and rune slice manipulation**. By the end, you'll understand:
- Why strings can't be modified directly in Go
- Converting strings to rune slices for manipulation
- The two-pointer swap technique
- Building strings from rune slices

## Theory: String Immutability and Runes

### Strings Are Immutable

In Go, strings are **read-only**. You cannot change individual characters:

```go
s := "Hello"
s[0] = 'J'  // ERROR: cannot assign to s[0]
```

### The Rune Slice Solution

To modify a string, convert it to a `[]rune`:

```go
runes := []rune("Hello")  // [H e l l o]
runes[0] = 'J'            // Now we can modify!
result := string(runes)   // "Jello"
```

### Why Runes, Not Bytes?

Strings in Go are UTF-8 encoded. A single character might be multiple bytes:

```go
s := "Hello 世界"
len(s)        // 11 bytes (9 ASCII + 6 for 2 Chinese chars)
len([]rune(s)) // 9 runes (each char = 1 rune)
```

**Always use `[]rune` for character-level operations!**

## The Two-Pointer Swap Technique

To reverse a slice, swap elements from both ends:

```
Original:  [H, e, l, l, o]
Step 1:    [o, e, l, l, H]  (swap index 0 and 4)
Step 2:    [o, l, l, e, H]  (swap index 1 and 3)
Step 3:    [o, l, l, e, H]  (index 2 is middle, stop)
Result:    "olleH"
```

### Algorithm

```go
for i := 0; i < len(runes)/2; i++ {
    // Swap elements at i and len-1-i
    runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
}
```

## Step-by-Step Approach

1. **Convert** string to `[]rune`
2. **Loop** from 0 to half the length
3. **Swap** elements at `i` and `len-1-i`
4. **Convert** back to string
5. **Return** result

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Using byte slice `[]byte` | Breaks with UTF-8 chars | Use `[]rune` |
| Loop condition `i <= len/2` | Swaps middle element twice | Use `i < len/2` |
| Forgetting to convert back | Returns rune slice, not string | `string(runes)` |

## Practice Tips

- Draw the swap process on paper
- Test with odd length: `"abc"` → `"cba"` (middle stays)
- Test with even length: `"abcd"` → `"dcba"`
- Remember: `len([]rune(s))` gives character count

## The Challenge

Write a function that reverses a string.

### Expected Function

```go
func ReverseString(s string) string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `"Hello"` | `"olleH"` | Standard reversal |
| `"Go is fun!"` | `"!nuf si oG"` | With spaces and punctuation |
| `""` | `""` | Empty string |
| `"a"` | `"a"` | Single character |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Println(ReverseString("Hello"))     // olleH
    fmt.Println(ReverseString("Go is fun!")) // !nuf si oG
    fmt.Println(ReverseString(""))          // (empty)
    fmt.Println(ReverseString("a"))         // a
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. Why can't you modify a string directly in Go?
2. What's the difference between `[]byte` and `[]rune`?
3. How many swaps are needed for a 5-character string?

## Next Steps

After completing this, you'll be ready for:
- [86-ispalindrome](../86-ispalindrome/README.md) - Ispalindrome
- [87-removespaces](../87-removespaces/README.md) - Removespaces