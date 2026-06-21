# 06. To Lower

## What You'll Learn

This exercise teaches you **the reverse of case conversion**. By the end, you'll understand:
- Converting uppercase to lowercase using ASCII math
- The symmetry of case conversion operations
- Building transformation pipelines
- Handling mixed-case input

## Theory: The Reverse Operation

### Uppercase to Lowercase

The formula is the **reverse** of ToUpper:

```
'A' (65) + 32 = 'a' (97)
```

**Key difference:**
- ToUpper: `lowercase - 32` → uppercase
- ToLower: `uppercase + 32` → lowercase

### ASCII Ranges Recap

| Case | Check | Convert |
|------|-------|---------|
| Lowercase | `c >= 'a' && c <= 'z'` | `c - 32` |
| Uppercase | `c >= 'A' && c <= 'Z'` | `c + 32` |

### Why Not Use `strings.ToLower()`?

This exercise builds your understanding of:
1. How case conversion works internally
2. Character-by-character processing
3. Conditional logic patterns

## Step-by-Step Approach

1. **Initialize** an empty result string
2. **Iterate** through each character
3. **Check** if character is uppercase (`'A'` to `'Z'`)
4. **Convert** to lowercase by adding 32 if needed
5. **Append** to result
6. **Return** result

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Using `- 32` | Converts to uppercase, not lowercase | Use `+ 32` for ToLower |
| Converting all chars | Numbers/symbols break | Only convert `'A'` to `'Z'` |
| Forgetting non-letters | Should stay unchanged | Check range first |

## Practice Tips

- ToUpper and ToLower are mirror operations
- Test: `ToLower(ToUpper("Hello"))` should give `"hello"`
- Remember: `'Z'` (90) + 32 = `'z'` (122) ✓

## The Challenge

Write a function that converts all uppercase letters to lowercase.

### Expected Function

```go
func ToLower(s string) string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `"Hello"` | `"hello"` | Uppercase H converted |
| `"Go is fun!"` | `"go is fun!"` | G converted, rest unchanged |
| `"123ABC"` | `"123abc"` | Numbers unchanged, letters converted |
| `""` | `""` | Empty string |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Println(ToLower("Hello"))      // hello
    fmt.Println(ToLower("Go is fun!")) // go is fun!
    fmt.Println(ToLower("123ABC"))     // 123abc
    fmt.Println(ToLower(""))           // (empty)
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. What ASCII operation converts uppercase to lowercase?
2. How is ToLower different from ToUpper?
3. What characters should NOT be modified?

## Next Steps

After completing this, you'll be ready for:
- [19-countalpha](../19-countalpha/README.md) - Countalpha
- [20-checknumber](../20-checknumber/README.md) - Checknumber