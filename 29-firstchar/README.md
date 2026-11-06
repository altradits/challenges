# 02. First Character

## What You'll Learn

This exercise teaches you **string indexing and access patterns** in Go. By the end, you'll understand:
- How to access individual characters in a string
- The difference between byte indexing and rune-based access
- How to handle edge cases (empty strings, single characters)
- Converting between `rune` and `string` types

## Theory: Accessing Characters in Go

### String Indexing in Go

Go strings are **immutable sequences of bytes**. You can access individual bytes using index notation:

```go
s := "Hello"
firstByte := s[0]    // 'H' (byte value 72)
```

**Important:** `s[i]` returns a `byte` (uint8), not a `rune` or `string`.

### The Rune Type

A `rune` is Go's alias for `int32` and represents a **Unicode code point**:

```go
var r rune = 'A'    // Same as: var r rune = 65
```

### Converting Between Types

| From | To | How |
|------|----|-----|
| `byte` | `rune` | `rune(b)` |
| `rune` | `string` | `string(r)` |
| `string` | `[]rune` | `[]rune(s)` |
| `[]rune` | `string` | `string(runes)` |

### Why Convert Rune to String?

A `rune` is a number (int32). To get the actual character representation, convert it to a string:

```go
c := 'H'           // c is a rune (value 72)
char := string(c)   // char is "H" (a string)
```

## Step-by-Step Approach

1. **Check if the string is empty** - return empty string if so
2. **Get the first character** using `for...range` (first iteration)
3. **Convert to string** and return

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| `return s[0]` | Returns `byte`, not `string` | Use `string()` conversion |
| Using `len(s)` to check empty | Works but `for...range` is safer for UTF-8 | Check with `for...range` or just return early |
| Forgetting empty string case | Will panic or return wrong value | Always check `s == ""` first |

## Practice Tips

- Remember: `s[0]` gives you a number (byte value), not a character
- To return a character as a string: `string(s[0])` or `string(rune)`
- For UTF-8 safety, use `for...range` to get the first rune

## The Challenge

Write a function that returns the **first character** of a string as a string.

### Expected Function

```go
func FirstChar(s string) string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `"Hello"` | `"H"` | First character |
| `""` | `""` | Empty string |
| `"G"` | `"G"` | Single character |
| `"Go is fun"` | `"G"` | First character of longer string |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Println(FirstChar("Hello"))   // H
    fmt.Println(FirstChar(""))        // (empty)
    fmt.Println(FirstChar("G"))       // G
    fmt.Println(FirstChar("Go is fun")) // G
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. What type does `s[0]` return?
2. How do you convert a `rune` to a `string`?
3. What happens if you try `s[0]` on an empty string?

## Next Steps

After completing this, you'll be ready for:
- [30-lastchar](../30-lastchar/README.md) - Lastchar
- [31-isempty](../31-isempty/README.md) - Isempty

---

## Time limit

**10 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 29
```

or run the tests directly:

```bash
go test ./29-firstchar/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 29` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
