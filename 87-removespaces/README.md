# 12. Remove Spaces

## What You'll Learn

This exercise teaches you **string filtering and selective character removal**. By the end, you'll understand:
- How to filter characters based on conditions
- Building new strings by excluding certain characters
- The difference between removing and replacing
- Efficient string construction patterns

## Theory: String Filtering

### What is Filtering?

Filtering means **keeping only elements that match a condition**:

```
Input:  "Hello World"
Filter: Keep everything EXCEPT spaces
Output: "HelloWorld"
```

### The Build-New-String Pattern

Since strings are immutable, filtering requires building a new string:

```go
func filterSpaces(s string) string {
    var result string
    for _, c := range s {
        if c != ' ' {  // Keep non-spaces
            result += string(c)
        }
    }
    return result
}
```

### Space vs Whitespace

| Character | Description | Code |
|-----------|-------------|------|
| Space | Regular space | `' '` (32) |
| Tab | Horizontal tab | `'\t'` (9) |
| Newline | Line break | `'\n'` (10) |

**This exercise removes only spaces (`' '`), not other whitespace!**

## Step-by-Step Approach

1. **Initialize** an empty result string
2. **Iterate** through each character
3. **Check** if character is NOT a space
4. **Append** non-space characters to result
5. **Return** the filtered string

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Removing all whitespace | Should only remove spaces | Check `c != ' '` specifically |
| Modifying original string | Strings are immutable | Build new string |
| Using `strings.Replace` | Defeats learning purpose | Manual iteration |

## Practice Tips

- Test with: `"Hello World"` → `"HelloWorld"`
- Test with multiple spaces: `"a  b   c"` → `"abc"`
- Remember: tabs and newlines should stay

## The Challenge

Write a function that removes all space characters from a string.

### Expected Function

```go
func RemoveSpaces(s string) string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `"Hello World"` | `"HelloWorld"` | Space removed |
| `"Go is fun!"` | `"Goisfun!"` | All spaces removed |
| `""` | `""` | Empty string |
| `"NoSpacesHere"` | `"NoSpacesHere"` | No spaces to remove |
| `"a  b   c"` | `"abc"` | Multiple spaces removed |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Println(RemoveSpaces("Hello World"))  // HelloWorld
    fmt.Println(RemoveSpaces("Go is fun!"))   // Goisfun!
    fmt.Println(RemoveSpaces(""))             // (empty)
    fmt.Println(RemoveSpaces("NoSpacesHere")) // NoSpacesHere
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. What's the difference between filtering and replacing?
2. Why do we build a new string instead of modifying the original?
3. What character code represents a space?

## Next Steps

After completing this, you'll be ready for:
- [88-countrepeats](../88-countrepeats/README.md) - Countrepeats
- [89-retainfirsthalf](../89-retainfirsthalf/README.md) - Retainfirsthalf