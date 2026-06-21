# 05. To Upper

## What You'll Learn

This exercise teaches you **ASCII character manipulation and case conversion**. By the end, you'll understand:
- The ASCII table and character codes
- How uppercase and lowercase letters relate mathematically
- Conditional logic for character transformation
- Building new strings from transformed characters

## Theory: ASCII and Character Codes

### What is ASCII?

ASCII (American Standard Code for Information Interchange) assigns a unique number to each character:

```
Character:   A   B   C   ...   Z   a   b   c   ...   z
ASCII Code:  65  66  67       90  97  98  99       122
```

### The Case Conversion Formula

Uppercase and lowercase letters differ by exactly **32** in ASCII:

```
'A' (65) + 32 = 'a' (97)
'a' (97) - 32 = 'A' (65)
```

**General formula:**
- To convert to uppercase: `char - 32` (if it's lowercase)
- To convert to lowercase: `char + 32` (if it's uppercase)

### ASCII Ranges for Letters

| Case | Range | First | Last |
|------|-------|-------|------|
| Uppercase | 'A' to 'Z' | 65 | 90 |
| Lowercase | 'a' to 'z' | 97 | 122 |

### Checking if a Character is Lowercase

```go
if c >= 'a' && c <= 'z' {
    // c is a lowercase letter
    upper := c - 32
}
```

## Step-by-Step Approach

1. **Initialize** an empty result string
2. **Iterate** through each character in the input
3. **Check** if the character is lowercase (`'a'` to `'z'`)
4. **Convert** to uppercase by subtracting 32 if needed
5. **Append** the (possibly converted) character to result
6. **Return** the result

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Converting all characters | Numbers/symbols will break | Only convert `'a'` to `'z'` |
| Using `+ 32` for uppercase | That converts to lowercase | Use `- 32` for uppercase |
| Forgetting non-letters | Spaces/numbers should stay the same | Only modify if in range |

## Practice Tips

- Memorize: `'a'` = 97, `'A'` = 65, difference = 32
- Draw the ASCII range: `[a][b][c]...[x][y][z]`
- Test with mixed input: `"Hello World!"` → `"HELLO WORLD!"`

## The Challenge

Write a function that converts all lowercase letters to uppercase.

### Expected Function

```go
func ToUpper(s string) string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `"Hello"` | `"HELLO"` | All lowercase converted |
| `"Go is fun!"` | `"GO IS FUN!"` | Mixed case + punctuation |
| `"123abc"` | `"123ABC"` | Numbers unchanged, letters converted |
| `""` | `""` | Empty string |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Println(ToUpper("Hello"))       // HELLO
    fmt.Println(ToUpper("Go is fun!"))  // GO IS FUN!
    fmt.Println(ToUpper("123abc"))      // 123ABC
    fmt.Println(ToUpper(""))            // (empty)
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. What is the ASCII code for 'a'? For 'A'?
2. How do you check if a rune is a lowercase letter?
3. What happens if you subtract 32 from 'A'?

## Next Steps

After completing this, you'll be ready for:
- [33-tolower](../33-tolower/README.md) - Tolower
- [34-countalpha](../34-countalpha/README.md) - Countalpha