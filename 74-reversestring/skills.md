# Skills for 74-reversestring

## What You'll Learn

**Previous:** [73-retainfirsthalf](../73-retainfirsthalf/skills.md) | **Next:** [75-titlecase](../75-titlecase/skills.md)

**Challenge:** Write `ReverseString(s string) string` that returns the string with all characters in reverse order. Do not use any built-in reverse functions.

## Core Concept: Converting a String to a Rune Slice for Mutation

### What Is It?

Go strings are immutable — you cannot change them in place. To reverse a string, you need to either:
1. Iterate backwards and build a new string, or
2. Convert the string to a **rune slice** (`[]rune`), reverse the slice in place, and convert back to string

The rune slice approach is the idiomatic Go method and handles multi-byte Unicode characters correctly.

### Method 1: Iterate Backwards (Using `strings.Builder`)

```go
import "strings"

func ReverseString(s string) string {
    runes := []rune(s)   // convert to slice so we can index by rune position
    var b strings.Builder
    for i := len(runes) - 1; i >= 0; i-- {
        b.WriteRune(runes[i])
    }
    return b.String()
}
```

Step by step:
1. `[]rune(s)` — convert the string to a slice of runes (safe for Unicode)
2. Start `i` at the last index: `len(runes) - 1`
3. Count down: `i--` each iteration
4. Stop when `i < 0`
5. Write each rune to the builder in reverse order

### Method 2: Reverse a Rune Slice In-Place

```go
func ReverseString(s string) string {
    runes := []rune(s)
    left, right := 0, len(runes)-1
    for left < right {
        runes[left], runes[right] = runes[right], runes[left]
        left++
        right--
    }
    return string(runes)
}
```

This uses the **two-pointer swap** technique:
- `left` starts at index 0, `right` at the last index
- Swap the characters at both ends
- Move inward until they meet

### Diagram: Two-Pointer Reverse of "Hello"

```
Initial:   H  e  l  l  o
index:     0  1  2  3  4
           ^              ^
          left           right

Step 1: swap H and o
           o  e  l  l  H
           left moves right, right moves left

Step 2: swap e and l  (indices 1 and 3)
           o  l  l  e  H

Step 3: left(2) == right(2), stop

Result: "olleH"
```

### Why Convert to `[]rune` Instead of `[]byte`?

`[]byte(s)` would work for ASCII but would break multi-byte Unicode characters. For example, the emoji `"👋"` is 4 bytes. Reversing the bytes would corrupt the encoding. `[]rune(s)` decodes the string into Unicode code points first, so each element is one character regardless of byte length.

### Converting Between String and Rune Slice

```go
s := "Hello"
runes := []rune(s)        // string -> rune slice
back := string(runes)     // rune slice -> string
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Reversing `[]byte(s)` | Corrupts multi-byte Unicode characters | Use `[]rune(s)` |
| `for i := len(s) - 1` without converting to runes | `s[i]` gives a byte, not a rune | Convert first: `runes := []rune(s)`, then `runes[i]` |
| Off-by-one: `i >= 0` vs `i > 0` | `i >= 0` is correct — you need to include index 0 | Make sure the loop includes `i == 0` |
| Swapping with temp variable instead of multi-assign | More verbose than needed | Use `a, b = b, a` for clean in-place swaps |

## Solving This Challenge

### Algorithm (backwards iteration)

1. Convert `s` to `[]rune`
2. Create a `strings.Builder`
3. Loop from the last index down to 0
4. Write each rune to the builder
5. Return `b.String()`

### Algorithm (two-pointer in-place)

1. Convert `s` to `[]rune`
2. Set `left = 0`, `right = len-1`
3. While `left < right`: swap `runes[left]` and `runes[right]`, move both inward
4. Return `string(runes)`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [75-titlecase](../75-titlecase/README.md)
