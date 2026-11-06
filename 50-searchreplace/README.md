# 23. Search and Replace

## What You'll Learn

This exercise teaches you **pattern matching and string substitution**. By the end, you'll understand:
- How to find and replace substrings (not just single characters)
- Building strings with dynamic content insertion
- Handling cases where the pattern isn't found
- The difference between character and substring operations

## Theory: Substring Search and Replace

### What is a Substring?

A substring is a **sequence of characters within a string**:

```
"Hello World"
  ^^^^^
  "Hello" is a substring

"Hello World"
       ^^^^^
       "World" is a substring
```

### Search and Replace Pattern

```go
func searchReplace(text, old, new string) string {
    // 1. Find where "old" appears in "text"
    // 2. Build new string: text before old + new + text after old
}
```

### Building with Slices

```go
index := FindSubstring(text, old)  // Position of old
if index == -1 {
    return text  // Not found, return original
}
before := text[:index]              // Part before old
after := text[index+len(old):]      // Part after old
return before + new + after
```

## Step-by-Step Approach

1. **Find** the index of `old` in `text`
2. **If not found** (-1): return original `text`
3. **Extract** parts before and after `old`
4. **Combine**: `before + new + after`
5. **Return** result

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Using character replace | Works for single chars only | Use substring slicing |
| Wrong slice indices | Off-by-one errors | `[:index]` and `[index+len(old):]` |
| Forgetting not-found case | Would panic or return wrong | Check for -1 first |

## Practice Tips

- Test with: `"Hello World"`, replace `"World"` with `"Go"` → `"Hello Go"`
- Test with: `"Hello World"`, replace `"xyz"` with `"Go"` → `"Hello World"` (no change)
- Remember: `old` can be multiple characters!

## Task

Write a function that replaces the first occurrence of `old` with `new` in `text`.

### Expected Function

```go
func SearchReplace(text, old, new string) string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `("Hello World", "World", "Go")` | `"Hello Go"` | Replaced |
| `("Hello World", "xyz", "Go")` | `"Hello World"` | Not found |
| `("Hello", "l", "L")` | `"HeLlo"` | First 'l' replaced |
| `("", "a", "b")` | `""` | Empty text |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Println(SearchReplace("Hello World", "World", "Go"))
    fmt.Println(SearchReplace("Hello World", "xyz", "Go"))
    fmt.Println(SearchReplace("Hello", "l", "L"))
    fmt.Println(SearchReplace("", "a", "b"))
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. How do you extract parts of a string using slicing?
2. What happens if `old` is not found?
3. How is this different from ReplaceChar?

## Next Steps

After completing this, you'll be ready for:
- [51-cleanlist](../51-cleanlist/README.md) - Cleanlist
- [52-countwords](../52-countwords/README.md) - Countwords

---

## Time limit

**10 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 50
```

or run the tests directly:

```bash
go test ./50-searchreplace/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 50` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
