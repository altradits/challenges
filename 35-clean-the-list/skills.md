# Skills for clean-the-list

## What You'll Learn

**Previous:** [../34-slicesintro/skills.md](../34-slicesintro/skills.md) | **Next:** [../36-cleanstr/README.md](../36-cleanstr/README.md)

**Challenge:** Transform each string in a slice: trim spaces, capitalize only the first letter, prepend an index number, ensure "milk" is in the list.

## Core Concept: Transforming and Building a Slice

### What Is It?
This challenge combines several string operations applied to every element of a slice, then rebuilds the output slice. You also learn how to check if a specific item is present and conditionally append it.

### Processing Each Element
For each item in the input slice:
1. `strings.TrimSpace(item)` — removes leading/trailing spaces
2. `strings.ToLower(item)` — lowercases everything first
3. Capitalize only the first rune using `unicode.ToUpper`
4. Prepend the 1-based index with `fmt.Sprintf`

```go
import (
    "fmt"
    "strings"
    "unicode"
)

func CleanList(lst []string) []string {
    if len(lst) == 0 {
        return []string{}
    }
    result := []string{}
    hasMilk := false
    for i, item := range lst {
        trimmed := strings.TrimSpace(item)
        lowered := strings.ToLower(trimmed)
        // Capitalize first letter only
        if len(lowered) == 0 {
            result = append(result, fmt.Sprintf("%d/ ", i+1))
            continue
        }
        runes := []rune(lowered)
        runes[0] = unicode.ToUpper(runes[0])
        formatted := fmt.Sprintf("%d/ %s", i+1, string(runes))
        result = append(result, formatted)
        if strings.ToLower(trimmed) == "milk" {
            hasMilk = true
        }
    }
    if !hasMilk {
        result = append(result, fmt.Sprintf("%d/ Milk", len(result)+1))
    }
    return result
}
```

### strings.TrimSpace
Removes all leading and trailing whitespace:
```go
strings.TrimSpace("  hello  ") // "hello"
strings.TrimSpace("  hi ")     // "hi"
```

### Capitalizing Only the First Letter
Convert to runes, uppercase the first, then convert back:
```go
runes := []rune("hello world")
runes[0] = unicode.ToUpper(runes[0])
result := string(runes) // "Hello world"
```

### fmt.Sprintf for Index Formatting
Build the `"N/ Item"` format:
```go
fmt.Sprintf("%d/ %s", 1, "Tomatoes") // "1/ Tomatoes"
fmt.Sprintf("%d/ %s", 2, "Milk")     // "2/ Milk"
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Capitalizing with `strings.Title` | Capitalizes every word, and it's deprecated | Use `unicode.ToUpper(runes[0])` on the first rune only |
| Not checking for "milk" case-insensitively | `"Milk"`, `"MILK"`, `"milk"` all count | Compare `strings.ToLower(trimmed) == "milk"` |
| Off-by-one in index | Output should be 1-based | Use `i+1` in the format string |

## Solving This Challenge

### Algorithm
1. If input is empty, return empty slice.
2. For each item: trim spaces, lowercase, capitalize first letter, format with 1-based index.
3. Track whether any item equals "milk" (case-insensitive).
4. If not found, append "milk" at the end with the next index.
5. Return the result slice.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [../36-cleanstr/README.md](../36-cleanstr/README.md)
