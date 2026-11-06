# 01. String Length

## What You'll Learn

This exercise teaches you the **fundamentals of string iteration in Go**. By the end, you'll understand:
- How Go strings work internally (UTF-8 encoded byte sequences)
- The `for...range` loop for iterating over strings
- How to count elements without using built-in functions
- The difference between bytes and runes

## Theory: Strings in Go

### What is a String?

In Go, a `string` is a **read-only slice of bytes**. It represents text encoded in UTF-8 format.

```go
s := "Hello"
// Internally: []byte{'H', 'e', 'l', 'l', 'o'}
```

### Bytes vs Runes

- **Byte**: A single 8-bit value (0-255). One ASCII character = 1 byte.
- **Rune**: A Unicode code point. Can be 1-4 bytes in UTF-8.

```go
s := "Hello"        // 5 bytes, 5 runes
s := "Hello 世界"    // 9 bytes, 7 runes (Chinese chars are 3 bytes each)
```

### The `for...range` Loop

`for...range` is the **safest way** to iterate over strings in Go because it iterates over **runes**, not bytes.

```go
for index, runeValue := range myString {
    // index: byte position
    // runeValue: the actual character (as rune/int32)
}
```

**Example:**
```go
s := "Hi"
for i, c := range s {
    fmt.Println(i, string(c))
}
// Output:
// 0 H
// 1 i
```

### Why Not Use `len()`?

This exercise asks you to count manually to understand:
1. How iteration works
2. What each character represents
3. Building logic from scratch

In real code, you'd use `len()`, but understanding the manual process makes you a better programmer.

## Step-by-Step Approach

1. **Initialize a counter** to 0
2. **Iterate** through the string using `for...range`
3. **Increment** the counter for each character
4. **Return** the counter

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Using `len(s)` | The exercise forbids it | Use `for...range` to count |
| Using byte indexing `s[i]` | Breaks with multi-byte UTF-8 chars | Use `for...range` which handles runes |
| Forgetting to return the count | Function returns 0 always | `return count` at the end |

## Practice Tips

- Draw the string as boxes: `[H][e][l][l][o]`
- Trace through the loop: "First iteration: H, count=1. Second: e, count=2..."
- Test with empty string: `""` should return 0
- Test with single char: `"a"` should return 1

## Task

Write a function that takes a string and returns its length **without using `len()`**.

### Expected Function

```go
func StringLength(s string) int {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `"Hello"` | `5` | 5 characters |
| `""` | `0` | Empty string |
| `"Go"` | `2` | 2 characters |
| `"a"` | `1` | Single character |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Println(StringLength("Hello"))   // 5
    fmt.Println(StringLength(""))        // 0
    fmt.Println(StringLength("Go is fun")) // 9
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. What does `for...range` give you for each iteration?
2. How many times will the loop run for `"Hi"`?
3. What's the difference between `s[0]` and the first value from `for...range`?

## Next Steps

After completing this, you'll be ready for:
- [29-firstchar](../29-firstchar/README.md) - Firstchar
- [30-lastchar](../30-lastchar/README.md) - Lastchar

---

## Time limit

**10 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 28
```

or run the tests directly:

```bash
go test ./28-stringlength/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 28` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
