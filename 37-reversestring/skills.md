# Skills for 37-reversestring

## What You'll Learn

**Previous:** [36-countvowels](../36-countvowels/skills.md) | **Next:** [38-ispalindrome](../38-ispalindrome/skills.md)

**Challenge:** Write a function `ReverseString(s string) string` that returns the characters of the string in reverse order.

## Core Concept: Converting a String to `[]rune` for In-Place Modification

### Why You Cannot Reverse a String Directly

Go strings are **immutable** — you cannot write to individual positions:

```go
s := "Hello"
s[0] = 'J'   // COMPILE ERROR: cannot assign to s[0]
```

To rearrange characters, you must first convert to a mutable type.

### The `[]rune` Slice

A `[]rune` is a slice of Unicode code points. Unlike a string, you can read AND write individual elements:

```go
runes := []rune("Hello")   // convert string → []rune
runes[0] = 'J'             // fine — slices are mutable
result := string(runes)    // convert back: "Jello"
```

Always use `[]rune` (not `[]byte`) when working with individual characters, because a single rune can be 1–4 bytes. Using `[]byte` would corrupt multi-byte characters like `é` or `中`.

### The Two-Pointer Swap

To reverse a slice, use two indices — one starting at the left, one at the right — and swap elements until they meet in the middle:

```
Original:  [H, e, l, l, o]   indices: 0 1 2 3 4
Step 1: swap 0 and 4 → [o, e, l, l, H]
Step 2: swap 1 and 3 → [o, l, l, e, H]
Step 3: indices meet at 2 → stop
Result:   "olleH"
```

In code:

```go
for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
    runes[i], runes[j] = runes[j], runes[i]
}
```

- `i` starts at the left (0), `j` starts at the right (`len-1`)
- Each iteration swaps the two outer elements and moves inward
- Stop when `i >= j` (they have met or crossed)

### Full Implementation

```go
func ReverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
```

Step by step:
1. `[]rune(s)` — convert string to a mutable rune slice
2. The `for` loop with two indices swaps from the outside in
3. `string(runes)` — convert the reversed slice back to a string

### How Many Swaps?

For a string of length `n`:
- Even length (e.g., 4): 2 swaps
- Odd length (e.g., 5): 2 swaps (middle element never needs to move)

The loop condition `i < j` handles both cases correctly.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `[]byte(s)` instead of `[]rune(s)` | Breaks multi-byte UTF-8 characters | Use `[]rune` |
| Loop condition `i <= j` | Swaps the middle element with itself (harmless but wasteful for odd length) | Use `i < j` |
| Forgetting `string(runes)` at the end | Returns `[]rune`, not `string` | Add `return string(runes)` |
| Building result by prepending: `result = string(c) + result` | Works but is O(n²) — slow | Use two-pointer swap |

## Solving This Challenge

### Algorithm

1. `runes := []rune(s)`
2. `i = 0`, `j = len(runes)-1`
3. While `i < j`: swap `runes[i]` and `runes[j]`, then `i++`, `j--`
4. Return `string(runes)`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [38-ispalindrome](../38-ispalindrome/README.md)
