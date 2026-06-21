# Prerequisites for 56-stringcontains

## Before You Start

### 1. Returning a `bool`

Functions that check a condition return `bool`:

```go
func isPositive(n int) bool {
    return n > 0
}

fmt.Println(isPositive(5))   // true
fmt.Println(isPositive(-1))  // false
```

### 2. String Slicing for Comparison

To check if a substring matches at position `i`:

```go
s := "hello world"
substr := "world"
i := 6
fmt.Println(s[i:i+len(substr)] == substr)  // true
```

This is equivalent to checking a window of size `len(substr)` at position `i`.

### 3. Loop Boundary for Safe Slicing

The last valid starting position for a window of size `len(substr)` is `len(s)-len(substr)`:

```go
for i := 0; i <= len(s)-len(substr); i++ {
    if s[i:i+len(substr)] == substr {
        return true
    }
}
```

If `len(substr) > len(s)`, this loop runs 0 times (correct — can't find a longer needle in shorter haystack).

### 4. Empty Substring is Always Found

By convention (and matching the standard library), an empty string is considered to be contained in every string:

```go
Contains("hello", "")  // true
Contains("", "")       // true
```

Check this before the loop: `if substr == "" { return true }`.

### 5. The Relationship Between Contains and Index

`Contains(s, substr)` can be thought of as `Index(s, substr) != -1`. If you've already built `Index` (challenge 120), you can use it here. Or build `Contains` first — it's slightly simpler because you return `true` immediately on the first match without caring about the exact index.

## Review If Stuck

- [38-findsubstring](../38-findsubstring/skills.md) — covers the sliding window algorithm this challenge uses

## You're Ready When You Can...

- [ ] Write a function that returns a `bool`
- [ ] Compare `s[i:i+len(substr)]` to `substr` at a given index
- [ ] Set the loop bound correctly with `i <= len(s)-len(substr)`
- [ ] Handle the empty substring case

## Next Steps

- [57-stringindex](../57-stringindex/README.md)
