# 11. Is Palindrome

## What You'll Learn

This exercise teaches you **string comparison and cleaning techniques**. By the end, you'll understand:
- What a palindrome is mathematically
- Normalizing strings for comparison
- The two-pointer comparison technique
- Ignoring non-alphanumeric characters

## Theory: Palindromes

### What is a Palindrome?

A palindrome reads the **same forwards and backwards**:

```
"racecar" → forwards: r-a-c-e-c-a-r
            backwards: r-a-c-e-c-a-r  ✓ MATCH

"hello" → forwards: h-e-l-l-o
          backwards: o-l-l-e-h  ✗ NO MATCH
```

### Normalization for Comparison

When checking palindromes, we often need to **normalize** the string:

1. **Convert to same case** (usually lowercase)
2. **Remove non-alphanumeric characters** (spaces, punctuation)
3. **Compare** normalized string with its reverse

### The Two-Pointer Technique

Instead of reversing the whole string, compare from both ends:

```
String:  "racecar"
          ↑           ↑
          left        right

Step 1: Compare 'r' and 'r' → MATCH
Step 2: Compare 'a' and 'a' → MATCH
Step 3: Compare 'c' and 'c' → MATCH
Step 4: Compare 'e' and 'e' → MATCH (middle, stop)
```

### Alphanumeric Check

A character is alphanumeric if it's:
- A letter: `'a'-'z'` or `'A'-'Z'`
- A digit: `'0'-'9'`

```go
func isAlnum(c rune) bool {
    return (c >= 'a' && c <= 'z') ||
           (c >= 'A' && c <= 'Z') ||
           (c >= '0' && c <= '9')
}
```

## Step-by-Step Approach

1. **Normalize** the string (lowercase, remove non-alphanumeric)
2. **Initialize** two pointers: left at start, right at end
3. **Compare** characters at both pointers
4. **Move** pointers inward (left++, right--)
5. **Return** true if all match, false otherwise

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Comparing original string | Case and spaces cause mismatch | Normalize first |
| Using `==` for strings | Works but inefficient for long strings | Use two-pointer technique |
| Forgetting empty string | Should return true | Handle edge case |

## Practice Tips

- Famous palindromes: "racecar", "madam", "A man a plan a canal Panama"
- Test normalization: `"A man, a plan, a canal: Panama"` → `"amanaplanacanalpanama"`
- Empty string IS a palindrome (reads same forwards and backwards)

## The Challenge

Write a function that checks if a string is a palindrome (ignoring case and non-alphanumeric characters).

### Expected Function

```go
func IsPalindrome(s string) bool {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `"racecar"` | `true` | Perfect palindrome |
| `"Hello"` | `false` | Not a palindrome |
| `"A man a plan a canal Panama"` | `true` | After normalization |
| `""` | `true` | Empty string is palindrome |
| `"a"` | `true` | Single character |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Println(IsPalindrome("racecar"))                    // true
    fmt.Println(IsPalindrome("Hello"))                      // false
    fmt.Println(IsPalindrome("A man a plan a canal Panama")) // true
    fmt.Println(IsPalindrome(""))                           // true
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. What does "normalize" mean in string comparison?
2. How does the two-pointer technique work?
3. Why is an empty string a palindrome?

## Next Steps

After completing this, you'll be ready for:
- **12-removespaces**: Filtering characters from strings
- **13-countrepeats**: Detecting patterns
