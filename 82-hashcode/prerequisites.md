# Prerequisites for hashcode

## Before You Start

To solve this challenge you need to understand:

### 1. ASCII Character Values
Each character maps to an integer. `int(c)` gives you the number; `string(rune(n))` converts back.
```go
fmt.Println(int('A'))         // 65
fmt.Println(int('a'))         // 97
fmt.Println(int(' '))         // 32
fmt.Println(string(rune(65))) // "A"
```

### 2. Modulo Arithmetic
`%` gives the remainder and keeps values within a range.
```go
fmt.Println(130 % 127) // 3
fmt.Println(65 % 127)  // 65
fmt.Println(127 % 127) // 0
```

### 3. for/range Over a String
Iterating with `range` gives you rune values, which is correct for character arithmetic.
```go
s := "ABC"
for _, c := range s {
    fmt.Println(int(c)) // 65, 66, 67
}
```

### 4. Building a String in a Loop
Append characters to a result string using `+=`.
```go
result := ""
for _, c := range "ABC" {
    result += string(rune(int(c) + 1))
}
// result == "BCD"
```

### 5. Conditional Adjustment
Check if a computed value falls in an undesirable range and correct it.
```go
hashed := someComputation()
if hashed < 32 {
    hashed += 33  // shift into printable ASCII range
}
```

## Review If Stuck

- [../81-gcd/skills.md](../81-gcd/skills.md) — covers the modulo operator `%`

## You're Ready When You Can...

- [ ] Get the integer value of a character with `int(c)`
- [ ] Use `%` to wrap a value into a range
- [ ] Convert an integer back to a string character with `string(rune(val))`
- [ ] Iterate over a string with `for _, c := range s`
- [ ] Apply a conditional correction when the value falls below 32

## Next Steps

- [83-lastword](../83-lastword/README.md)
