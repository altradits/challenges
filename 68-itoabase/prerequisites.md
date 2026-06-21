# Prerequisites for itoabase

## Before You Start

To solve this challenge you need to understand:

### 1. Integer Division and Modulo
`value / base` gives the quotient; `value % base` gives the remainder (the next digit).

```go
fmt.Println(255 / 16)   // 15
fmt.Println(255 % 16)   // 15 (the digit)
fmt.Println(15 / 16)    // 0 (loop ends)
```

### 2. Indexing into a Digit String
A string like `"0123456789ABCDEF"` maps an integer 0–15 to its character representation.

```go
digits := "0123456789ABCDEF"
fmt.Println(string(digits[10]))  // "A"
fmt.Println(string(digits[15]))  // "F"
fmt.Println(string(digits[0]))   // "0"
```

### 3. Prepending to a String (Digit Reversal)
The first digit extracted (by `% base`) is the least significant. Prepend each new digit to put them in the correct order.

```go
result := ""
result = "A" + result  // "A"
result = "B" + result  // "BA"  — B is more significant
```

### 4. Handling Negative Numbers
Detect `value < 0`, negate it for the conversion loop, then prepend `"-"` to the result.

```go
negative := value < 0
if negative {
    value = -value
}
// ... convert ...
if negative {
    result = "-" + result
}
```

### 5. Special Case: Zero
The main loop never executes when `value == 0`. Handle it explicitly.

```go
if value == 0 {
    return "0"
}
```

## Review If Stuck
- [53-fprime](../53-fprime/skills.md) — covers repeated division with `%` and `/` (same pattern as base conversion)
- [67-wordflip](../67-wordflip/skills.md) — covers string building and prepending

## You're Ready When You Can...
- [ ] Use `value % base` to extract the least-significant digit
- [ ] Use `value /= base` to shift to the next digit
- [ ] Index into `"0123456789ABCDEF"` to get the correct character
- [ ] Prepend each new character to build the result in the right order
- [ ] Handle `value == 0` and negative values as special cases

## Next Steps
- [69-options](../69-options/README.md)
