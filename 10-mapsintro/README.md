# 10. Maps Introduction

## What You'll Learn

This lesson teaches Go's built-in key-value data structure: the **map**.

## The Challenge

Write `WordFrequency(s string) map[string]int` that returns a map showing how many times each word appears in the string.

### Expected Function

```go
func WordFrequency(s string) map[string]int {
    // Your code here
}
```

### Test Cases

| Input | Expected Output |
|-------|-----------------|
| `"go is go"` | `{"go": 2, "is": 1}` |
| `"hello"` | `{"hello": 1}` |
| `""` | `{}` (empty map) |
| `"a b a c a"` | `{"a": 3, "b": 1, "c": 1}` |

### Usage Example

```go
package main

import "fmt"

func main() {
    freq := WordFrequency("go is go")
    fmt.Println(freq["go"])  // 2
    fmt.Println(freq["is"])  // 1
}
```

## Next Steps

After completing this, you'll be ready for:
- [11-arrays](../11-arrays/README.md) - Fixed-size arrays
