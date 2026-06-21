# Prerequisites for 89-reversestring

## Before You Start

To solve this challenge you need to understand rune slices, backwards iteration, and how to convert between strings and slices.

### 1. Rune slices: converting a string to `[]rune`
A rune slice lets you index into a string by character position (not byte position). This is important for Unicode safety.

```go
s := "Hello"
runes := []rune(s)   // [72 101 108 108 111]
runes[0]             // 72 ('H')
runes[4]             // 111 ('o')
```

Converting back:
```go
string(runes)  // "Hello"
```

### 2. Iterating backwards with a `for` loop
Count down from the last index to 0:

```go
runes := []rune("Hello")
for i := len(runes) - 1; i >= 0; i-- {
    // visits: runes[4], runes[3], runes[2], runes[1], runes[0]
}
```

### 3. `strings.Builder` for accumulating characters

```go
import "strings"

var b strings.Builder
b.WriteRune('o')
b.WriteRune('l')
b.WriteRune('l')
b.WriteRune('e')
b.WriteRune('H')
b.String()  // "olleH"
```

### 4. Multi-assignment swap (for the two-pointer approach)
Go allows swapping two variables in one line:

```go
a, b := 1, 2
a, b = b, a   // now a=2, b=1
```

Applied to a slice:
```go
runes[left], runes[right] = runes[right], runes[left]
```

### 5. `len()` on a slice returns the number of elements

```go
runes := []rune("Hello")
len(runes)   // 5
```

## Review If Stuck

- [87-removespaces](../87-removespaces/skills.md) — covers `strings.Builder` and building strings character by character
- [88-retainfirsthalf](../88-retainfirsthalf/skills.md) — covers `len()` and working with string indices

## You're Ready When You Can...

- [ ] Convert a `string` to `[]rune` with `[]rune(s)`
- [ ] Convert a `[]rune` back to `string` with `string(runes)`
- [ ] Write a `for` loop that counts down from `len-1` to `0`
- [ ] Use `strings.Builder` to build a string character by character
- [ ] Swap two elements in a slice using multi-assignment

## Next Steps

- [90-titlecase](../90-titlecase/README.md)
