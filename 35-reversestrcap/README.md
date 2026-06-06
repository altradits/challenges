# 35. Reverse String Capitalization

## What You'll Learn

This exercise teaches you **selective case transformation**. By the end, you'll understand:
- How to transform only specific characters in a string
- Word boundary detection
- Combining multiple string operations
- Processing command-line arguments

## Theory: Selective Case Changes

### The Pattern

Transform each word by:
1. Converting all characters to lowercase
2. Uppercasing only the LAST character

```
"Hello" → "hellO"
"SMALL" → "smalL"
```

### Word Detection

Words are separated by spaces. Use `strings.Fields` or manual splitting.

## Step-by-Step Approach

1. **Split** input into words
2. **For each word**:
   - Convert to lowercase
   - Uppercase last character
3. **Join** words with spaces
4. **Return** result

## The Challenge

Write a function that uppercases the last letter of each word.

### Expected Function

```go
func ReverseStrCap(s string) string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output |
|-------|-----------------|
| `"First SMALL TesT"` | `"firsT smalL tesT"` |
| `"Hello"` | `"hellO"` |
| `""` | `""` |

## Next Steps

After completing this, you'll be ready for:
- **36-printrevcomb**: Combination generation
- **37-union**: Set operations on strings
