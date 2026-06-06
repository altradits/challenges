# 27. Replace All

## What You'll Learn

This exercise teaches you **global string replacement and pattern scanning**. By the end, you'll understand:
- How to replace ALL occurrences of a substring
- Scanning through a string and building results
- Handling overlapping patterns
- The difference between first and global replacement

## Theory: Global Replacement

### First vs. Global Replace

| Operation | Description | Example |
|-----------|-------------|---------|
| Replace first | Only first occurrence | `"aaa"` → `"baa"` |
| Replace all | Every occurrence | `"aaa"` → `"bbb"` |

### The Scan-and-Build Pattern

```go
func replaceAll(text, old, new string) string {
    var result string
    i := 0
    for i < len(text) {
        if strings.HasPrefix(text[i:], old) {
            result += new
            i += len(old)  // Skip past the old substring
        } else {
            result += string(text[i])
            i++
        }
    }
    return result
}
```

## Step-by-Step Approach

1. **Initialize** empty result and index 0
2. **Loop** while index < text length
3. **Check** if `old` matches at current position
4. **If match**: append `new`, advance by `len(old)`
5. **Else**: append current char, advance by 1
6. **Return** result

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Infinite loop | Not advancing index properly | Always increment i |
| Overlapping issues | `"aaa"` with `"aa"` → `"ba"` or `"bb"`? | Document behavior |
| Wrong slice | `text[i:]` vs `text[i:i+len(old)]` | Use `HasPrefix` or manual check |

## Practice Tips

- Test with: `"Hello World"`, replace `"l"` with `"L"` → `"HeLLo WorLd"`
- Test with: `"banana"`, replace `"na"` with `"NA"` → `"baNANA"`
- Test with: `"abcabc"`, replace `"abc"` with `"xyz"` → `"xyzxyz"`

## The Challenge

Write a function that replaces ALL occurrences of `old` with `new`.

### Expected Function

```go
func ReplaceAll(text, old, new string) string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `("Hello World", "l", "L")` | `"HeLLo WorLd"` | All 'l's replaced |
| `("banana", "na", "NA")` | `"baNANA"` | Both "na" replaced |
| `("Hello World", "", "x")` | `"Hello World"` | Empty old, no change |
| `("abcabc", "abc", "xyz")` | `"xyzxyz"` | Both "abc" replaced |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Println(ReplaceAll("Hello World", "l", "L"))
    fmt.Println(ReplaceAll("banana", "na", "NA"))
    fmt.Println(ReplaceAll("Hello World", "", "x"))
    fmt.Println(ReplaceAll("abcabc", "abc", "xyz"))
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. How does ReplaceAll differ from SearchReplace?
2. Why do we need to track the index manually?
3. What happens with overlapping patterns?

## Next Steps

After completing this, you'll be ready for:
- **28-split**: Splitting strings into parts
- **29-join**: Joining strings together
