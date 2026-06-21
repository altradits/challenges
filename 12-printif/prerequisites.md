# Prerequisites for 12-printif

## Before You Start

To solve this challenge you need to understand `if/else`, the `len()` function, boolean operators, and how to return strings with embedded newline characters.

### 1. `if/else` in Go
Go's `if/else` does not use parentheses around the condition. Braces `{}` are always required.

```go
if condition {
    // true branch
} else {
    // false branch
}
```

### 2. `len(s)` returns the number of bytes (characters for ASCII)

```go
len("abc")  // 3
len("ab")   // 2
len("")      // 0
```

### 3. Comparison operators

```go
len(str) >= 3   // true if length is 3 or more
len(str) < 3    // true if length is less than 3
str == ""        // true if string is empty
```

### 4. Boolean OR operator `||`
The condition is true if either side is true.

```go
if str == "" || len(str) >= 3 {
    // executes when str is empty OR has 3+ characters
}
```

### 5. Newline in string literals
The `\n` escape sequence means newline. It is part of the string value.

```go
return "G\n"             // the string G followed by a newline
return "Invalid Input\n" // the string Invalid Input followed by a newline
```

## Review If Stuck

- [11-forloops](../11-forloops/skills.md) — dedicated lesson on all for loop patterns in Go
- [06-checknumber](../06-checknumber/skills.md) — covers `for...range`, rune comparisons, and boolean returns
- [10-findsubstring](../10-findsubstring/skills.md) — covers `len()` and string slicing

## You're Ready When You Can...

- [ ] Write an `if/else` block in Go without parentheses around the condition
- [ ] Use `len(str)` in a comparison like `len(str) >= 3`
- [ ] Combine two conditions with `||`
- [ ] Return a string that contains `\n` as a newline
- [ ] Write a function that takes a `string` and returns a `string`

## Next Steps

After completing this exercise, move to:
- [13-printifnot](../13-printifnot/README.md)
