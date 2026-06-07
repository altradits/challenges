# 24. Clean List

## What You'll Learn

This exercise teaches you **processing collections of strings**. By the end, you'll understand:
- How to iterate over a slice of strings
- Applying string operations to each element
- Building a new slice with cleaned results
- The map-filter pattern

## Theory: Processing Collections

### The Map Pattern

"Map" means applying a function to each element:

```go
func mapStrings(inputs []string, fn func(string) string) []string {
    result := make([]string, len(inputs))
    for i, s := range inputs {
        result[i] = fn(s)  // Apply function to each
    }
    return result
}
```

### Cleaning Operations

Common cleaning operations:
- Trim spaces: `strings.TrimSpace(s)`
- To uppercase: `strings.ToUpper(s)`
- Remove empty: filter out `""`

## Step-by-Step Approach

1. **Create** a result slice
2. **Iterate** through input strings
3. **Clean** each string (remove extra spaces, trim, etc.)
4. **Add** to result
5. **Return** result

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Modifying original | Slices are references | Create new slice |
| Forgetting empty strings | Should be removed | Filter them out |
| Wrong slice capacity | May cause reallocation | Use `make([]string, 0, len(inputs))` |

## Practice Tips

- Test with: `["Hello", " World ", "Go"]` → `["Hello", "World", "Go"]`
- Test with: `["", "a", "", "b"]` → `["a", "b"]`

## The Challenge

Write a function that cleans a list of strings (trim spaces, remove empty).

### Expected Function

```go
func CleanList(inputs []string) []string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `["Hello", " World ", "Go"]` | `["Hello", "World", "Go"]` | Trimmed |
| `["", "a", "", "b"]` | `["a", "b"]` | Empty removed |
| `[]` | `[]` | Empty slice |
| `["  "]` | `[]` | Only spaces |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Println(CleanList([]string{"Hello", " World ", "Go"}))
    fmt.Println(CleanList([]string{"", "a", "", "b"}))
    fmt.Println(CleanList([]string{}))
    fmt.Println(CleanList([]string{"  "}))
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. How do you create a new slice?
2. What's the map-filter pattern?
3. How do you trim spaces from a string?

## Next Steps

After completing this, you'll be ready for:
- [100-countwords](../100-countwords/README.md) - Countwords
- [101-findsubstring](../101-findsubstring/README.md) - Findsubstring