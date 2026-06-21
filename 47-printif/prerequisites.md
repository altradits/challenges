# Prerequisites for 47-printif

## Before You Start

### 1. `len()` on strings

`len(s)` returns the number of bytes in a string. For ASCII text this equals character count.

```go
len("abc")  // 3
len("ab")   // 2
len("")     // 0
```

Review: [28-stringlength](../28-stringlength/skills.md)

### 2. `if/else` conditional branching

```go
if condition {
    return "yes"
}
return "no"
```

Review: [85-printif](../85-printif/skills.md)

### 3. The `||` (OR) operator

`A || B` is true when either A or B is true:

```go
if len(s) == 0 || len(s) >= 3 {
    // either empty OR 3+ chars
}
```

Review: [86-printifnot](../86-printifnot/skills.md)

### 4. Returning a string from a function

```go
func Check(s string) string {
    return "result"
}
```

## You're Ready When You Can...

- [ ] Get the length of a string with `len()`
- [ ] Write an `if` with an `||` compound condition
- [ ] Return different string values from a function

## Next Steps

- [48-printifnot](../48-printifnot/README.md)
