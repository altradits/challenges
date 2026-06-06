# 09. Count Vowels

## What You'll Learn

This exercise teaches you **counting specific characters and multi-condition logic**. By the end, you'll understand:
- What vowels are and why they matter
- Checking multiple conditions with OR logic
- Case-insensitive character matching
- Counting vs. detecting patterns

## Theory: Vowels in English

### The Five Vowels

English has 5 vowels: **a, e, i, o, u**

In programming, we check both cases:
- Lowercase: `a, e, i, o, u`
- Uppercase: `A, E, I, O, U`

### Case-Insensitive Matching

There are two approaches:

**Approach 1: Check both cases**
```go
if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
   c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
    // It's a vowel
}
```

**Approach 2: Convert first, then check**
```go
c = ToLower(c)  // Convert to lowercase first
if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
    // It's a vowel
}
```

### Using a Helper Function

For cleaner code, create a helper:

```go
func isVowel(c rune) bool {
    c = ToLower(c)  // Or check both cases
    return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}
```

## Step-by-Step Approach

1. **Initialize** counter to 0
2. **Iterate** through each character
3. **Check** if character is a vowel (any case)
4. **Increment** counter if true
5. **Return** the count

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Only checking lowercase | Misses uppercase vowels | Check both cases or convert first |
| Forgetting 'y' | Sometimes a vowel, but not here | Only a, e, i, o, u |
| Using `==` for multiple values | Can't chain `==` in Go | Use `||` (OR) operator |

## Practice Tips

- Memorize the 5 vowels: a, e, i, o, u
- Test with: `"Hello"` (2: e, o), `"AEIOU"` (5), `"rhythm"` (0)
- Remember: 'y' is NOT counted as a vowel here

## The Challenge

Write a function that counts vowels (a, e, i, o, u) in a string, case-insensitively.

### Expected Function

```go
func CountVowels(s string) int {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `"Hello World"` | `3` | e, o, o |
| `"Go is fun!"` | `3` | o, i, u |
| `""` | `0` | Empty string |
| `"AEIOUaeiou"` | `10` | All vowels, both cases |
| `"rhythm"` | `0` | No vowels |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Println(CountVowels("Hello World"))  // 3
    fmt.Println(CountVowels("Go is fun!"))   // 3
    fmt.Println(CountVowels(""))             // 0
    fmt.Println(CountVowels("AEIOUaeiou"))   // 10
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. What are the 5 vowels in English?
2. How do you check multiple values with OR logic?
3. What's the difference between counting and detecting?

## Next Steps

After completing this, you'll be ready for:
- **10-reversestring**: Building strings backwards
- **11-ispalindrome**: Advanced string comparison
