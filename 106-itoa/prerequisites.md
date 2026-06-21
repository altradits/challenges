# Prerequisites for 106-itoa

## Before You Start

### 1. The Modulo Operator `%`

`n % 10` gives the remainder when `n` is divided by 10. For positive integers, this is always the last (units-place) digit:

```go
fmt.Println(12345 % 10)  // 5
fmt.Println(1234  % 10)  // 4
fmt.Println(100   % 10)  // 0
fmt.Println(7     % 10)  // 7
```

### 2. Integer Division `/`

`n / 10` divides `n` by 10 and discards the remainder. This "removes" the last digit:

```go
fmt.Println(12345 / 10)  // 1234
fmt.Println(1234  / 10)  // 123
fmt.Println(1     / 10)  // 0
```

The loop condition `n > 0` stops when all digits have been extracted.

### 3. Converting a Digit to a Rune Character

`rune('0') + rune(digit)` converts an integer digit (0-9) to its character representation:

```go
digit := 5
char := rune('0') + rune(digit)  // '0' is 48, 48+5 = 53 = '5'
fmt.Println(string(char))        // "5"

// This is WRONG:
fmt.Println(string(5))  // not "5" — it's the character at code point 5
```

### 4. A Slice as an Accumulator

Build up a slice while looping, then convert to string at the end:

```go
var digits []rune
for _, c := range "hello" {
    digits = append(digits, c)
}
fmt.Println(string(digits))  // "hello"
```

### 5. Reversing a Slice

The two-pointer swap pattern:

```go
s := []rune{'5', '4', '3', '2', '1'}
for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
    s[i], s[j] = s[j], s[i]
}
fmt.Println(string(s))  // "12345"
```

### 6. Handling Negative Numbers

```go
n := -42
negative := n < 0
if negative {
    n = -n  // make it positive
}
// ... extract digits from positive n ...
// if negative { result = "-" + result }
```

## Review If Stuck

- [103-split](../103-split/skills.md) — covers building up a result slice incrementally

## You're Ready When You Can...

- [ ] Get the last digit of a number using `n % 10`
- [ ] Remove the last digit using `n / 10`
- [ ] Convert a digit integer to its rune character with `rune('0') + rune(digit)`
- [ ] Reverse a slice using two-pointer swap
- [ ] Handle the special case of `n == 0`

## Next Steps

- [107-thirdchar](../107-thirdchar/README.md)
