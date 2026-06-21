# Prerequisites for 22-reversestring

## Before You Start

To solve this challenge you need to understand:

### 1. The `for...range` Loop

From [13-stringlength skills.md](../13-stringlength/skills.md): you know how to iterate over characters. This challenge takes that further by working with indices too.

### 2. String Immutability

You have already seen that strings cannot be modified in place. Every challenge so far (ToUpper, ToLower, RemoveSpaces) built a **new** string rather than changing the original. This challenge introduces a better tool for structural rearrangement.

### 3. Slices — A Brief Introduction

A **slice** in Go is a resizable, mutable sequence. `[]rune` is a slice of runes. Unlike a string, you can change individual elements:

```go
nums := []int{10, 20, 30}
nums[0] = 99          // OK: slices are mutable
fmt.Println(nums)     // [99 20 30]
```

You will use `[]rune` to hold the characters while you rearrange them.

### 4. Converting Between `string` and `[]rune`

To go from string to rune slice and back:

```go
runes := []rune("Hello")   // string → []rune
s := string(runes)         // []rune → string
```

### 5. Simultaneous Assignment

Go supports swapping two variables in one line:

```go
a, b := 1, 2
a, b = b, a   // now a=2, b=1 — no temporary variable needed
```

This is used to swap elements during the two-pointer reversal.

## Review If Stuck

- [17-toupper skills.md](../17-toupper/skills.md) — building a new string from characters; immutability
- [15-lastchar skills.md](../15-lastchar/skills.md) — `len(s)` and index arithmetic

## You're Ready When You Can...

- [ ] Explain why you cannot write `s[0] = 'X'` on a string
- [ ] Convert a string to `[]rune` and back to `string`
- [ ] Access an element of a slice by index and change it
- [ ] Swap two variables without a temporary using `a, b = b, a`

## Next Steps

- [23-ispalindrome](../23-ispalindrome/README.md)
- [23-ispalindrome README](../23-ispalindrome/README.md) — next challenge
