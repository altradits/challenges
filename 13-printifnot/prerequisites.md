# Prerequisites for 13-printifnot

## Before You Start

To solve this challenge you need to understand `if/else`, `len()`, and the `!` (NOT) logical operator. You should already be comfortable with 12-printif.

### 1. `if/else` in Go

```go
if condition {
    return "something"
}
return "other"
```

### 2. `len(s)` and comparison operators

```go
len("ab") < 3    // true  (2 < 3)
len("abc") < 3   // false (3 < 3 is false)
len("") < 3      // true  (0 < 3)
```

### 3. The `!` (NOT) operator inverts a boolean

```go
!true   // false
!false  // true

if !(len(str) >= 3) {
    // same as: if len(str) < 3
}
```

### 4. De Morgan's Law for negating compound conditions

```
!(A || B) is the same as (!A && !B)
!(A && B) is the same as (!A || !B)
```

### 5. String function returning a string with `\n`

```go
func PrintIfNot(str string) string {
    if len(str) < 3 {
        return "G\n"
    }
    return "Invalid Input\n"
}
```

## Review If Stuck

- [12-printif](../12-printif/skills.md) — covers `if/else`, `len()`, boolean operators, and returning strings with `\n`

## You're Ready When You Can...

- [ ] Write an `if/else` that compares `len(str)` to a number
- [ ] Correctly predict what `len("") < 3` evaluates to
- [ ] Explain the difference between `<` and `<=`
- [ ] Explain what `!` does to a boolean expression
- [ ] Solve 12-printif first — this challenge is the inverse of it

## Next Steps

After completing this exercise, move to:
- [15-removespaces](../15-removespaces/README.md)
