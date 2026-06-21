# Prerequisites for 84-findsubstring

## Before You Start

To solve this challenge you need to understand string slicing, `len()`, and how to write a loop that searches through a string.

### 1. `len(s)` returns the number of bytes in a string

```go
len("Hello")   // 5
len("banana")  // 6
len("")         // 0
```

### 2. String slicing with `s[start:end]`
You can extract a portion of a string using slice notation. The result is a new string from index `start` up to (but not including) `end`.

```go
s := "banana"
s[0:3]   // "ban"
s[1:4]   // "ana"
s[3:]    // "ana"  (from index 3 to end)
```

### 3. String equality with `==`
You can compare two strings directly with `==`.

```go
"ana" == "ana"   // true
"ana" == "ban"   // false
```

### 4. `for` loop with an index counter
A traditional index-based loop (not `for...range`):

```go
for i := 0; i < 10; i++ {
    // i goes 0, 1, 2, ... 9
}
```

### 5. Returning `-1` to signal "not found"
The convention in Go (used by the standard library too) is to return `-1` from a search function when no result is found.

```go
func find(s, sub string) int {
    // ... search ...
    return -1  // not found
}
```

### 6. Early-return pattern from a function
Return as soon as you find the result:

```go
for i := 0; i < len(s); i++ {
    if s[i:i+n] == target {
        return i   // stop and return the position
    }
}
return -1
```

## Review If Stuck

- [81-checknumber](../81-checknumber/skills.md) — covers `for...range` and character comparisons
- [82-count-character](../82-count-character/skills.md) — covers the counter pattern and rune iteration

## You're Ready When You Can...

- [ ] Use `len(s)` to get the length of a string
- [ ] Slice a string with `s[i:j]` and predict the result
- [ ] Write a `for` loop using an integer index (not `range`)
- [ ] Compare two strings with `==`
- [ ] Return a specific integer from a function, including `-1` for "not found"

## Next Steps

- [85-printif](../85-printif/README.md)
