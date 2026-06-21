# Prerequisites for 16-retainfirsthalf

## Before You Start

To solve this challenge you need to understand `len()`, string slicing, and integer division.

### 1. `len(s)` returns the number of bytes in a string

```go
len("Hello")        // 5
len("Hello World")  // 11
len("")              // 0
len("A")             // 1
```

### 2. String slicing with `s[:n]`
Returns the first `n` characters of `s`.

```go
s := "Hello World"
s[:5]    // "Hello"
s[:0]    // ""
s[:11]   // "Hello World"
```

### 3. Integer division truncates (rounds down)
In Go, dividing two integers always gives an integer result — no decimal part.

```go
10 / 2   // 5
11 / 2   // 5  (not 5.5)
7 / 2    // 3  (not 3.5)
1 / 2    // 0  (not 0.5)
```

### 4. Combining len and integer division
You can use these together to find the midpoint of a string:

```go
str := "Hello World"
half := len(str) / 2    // 11 / 2 = 5
str[:half]              // "Hello"
```

### 5. Early-return with a guard condition
Handle edge cases before the main logic:

```go
if len(str) <= 1 {
    return str
}
// normal logic here
```

## Review If Stuck

- [10-findsubstring](../10-findsubstring/skills.md) — covers `len()` and string slicing `s[start:end]`
- [12-printif](../12-printif/skills.md) — covers `if/else` with `len()`

## You're Ready When You Can...

- [ ] Use `len(s)` to get the length of a string
- [ ] Slice a string with `s[:n]` to get the first `n` characters
- [ ] Predict the result of integer division like `11 / 2` (answer: 5, not 5.5)
- [ ] Write a guard `if` at the start of a function for edge cases

## Next Steps

After completing this exercise, move to:
- [17-reversestring](../17-reversestring/README.md)
