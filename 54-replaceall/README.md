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

## Task

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
- [55-split](../55-split/README.md) - Split
- [56-join](../56-join/README.md) - Join

---

## Time limit

**10 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 54
```

or run the tests directly:

```bash
go test ./54-replaceall/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 54` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
