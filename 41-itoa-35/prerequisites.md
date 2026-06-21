# Prerequisites for itoa-35

## Before You Start

To solve this challenge you need to understand:

### 1. Extracting Digits With Modulo and Division
`n % 10` gives the last (ones) digit. `n / 10` removes it.
```go
n := 12345
fmt.Println(n % 10) // 5  — ones digit
n /= 10
fmt.Println(n % 10) // 4  — next digit
n /= 10
fmt.Println(n)      // 123 — remaining
```

### 2. Converting a Digit to a Character
`'0' + digit` gives the ASCII code of the digit character. Wrap in `string(rune(...))`.
```go
digit := 7
ch := string(rune('0' + digit))
fmt.Println(ch) // "7"
```

### 3. Prepending to a String
To build the number left-to-right while extracting digits right-to-left, prepend each new digit:
```go
result := ""
result = "5" + result  // "5"
result = "4" + result  // "45"
result = "3" + result  // "345"
```

### 4. Handling Zero as a Special Case
When `n == 0`, the `for n > 0` loop never executes. Return `"0"` before entering the loop.
```go
if n == 0 {
    return "0"
}
```

### 5. Handling Negative Numbers
Negate the value, convert the positive part, then prepend `"-"`.
```go
negative := false
if n < 0 {
    negative = true
    n = -n
}
// ... convert n ...
if negative {
    result = "-" + result
}
```

## Review If Stuck

- [../22-digitlen/skills.md](../22-digitlen/skills.md) — covers the `for n > 0 { n /= base }` pattern that is the backbone of this function
- [../26-hashcode/skills.md](../26-hashcode/skills.md) — covers converting between integer values and string characters

## You're Ready When You Can...

- [ ] Extract the last digit of a number with `n % 10`
- [ ] Remove the last digit with `n /= 10`
- [ ] Convert a digit (0–9) to its character with `string(rune('0' + digit))`
- [ ] Prepend digits to build the number in correct order
- [ ] Handle `n == 0` and `n < 0` as special cases

## Next Steps

- [Next challenge](../43-printmemory/README.md)
