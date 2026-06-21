# Prerequisites for notdecimal

## Before You Start

To solve this challenge you need to understand:

### 1. String Slicing — Splitting at an Index
You can extract a substring using `s[start:end]` syntax.

```go
s := "174.2"
//    01234
dot := 3
intPart  := s[:dot]   // "174"
fracPart := s[dot+1:] // "2"
```

### 2. Finding a Character's Index by Scanning
Scan the string to find the position of `.`.

```go
dotIdx := -1
for i, c := range dec {
    if c == '.' {
        dotIdx = i
        break
    }
}
if dotIdx == -1 {
    // no decimal point found
}
```

### 3. Validating Character by Character
Check each character is a valid digit (`'0'` through `'9'`), allowing `-` at position 0 and `.` at the dot position.

```go
for i, c := range dec {
    if i == 0 && c == '-' { continue }
    if c == '.' { continue }
    if c < '0' || c > '9' {
        return dec + "\n"  // invalid character
    }
}
```

### 4. Trimming Trailing Bytes from a Substring
To remove trailing `'0'` bytes from the fractional part, shrink the end index.

```go
fracPart := "200"
end := len(fracPart)
for end > 0 && fracPart[end-1] == '0' {
    end--
}
fracPart = fracPart[:end]  // "2"
```

### 5. String Concatenation for the Result
Join the parts back into one string.

```go
return intPart + fracPart + "\n"
```

## Review If Stuck
- [60-fifthandskip](../60-fifthandskip/skills.md) — covers string scanning and conditional building
- [53-fprime](../53-fprime/skills.md) — covers `strconv.Atoi` (for contrast — here we avoid float conversion)

## You're Ready When You Can...
- [ ] Use `s[start:end]` to extract a substring
- [ ] Find the index of a character by scanning with `for i, c := range s`
- [ ] Validate each character against a set of allowed values
- [ ] Trim trailing zeros by decrementing an `end` index
- [ ] Return the input unchanged (with `\n`) when it is not a valid decimal

## Next Steps
- [62-revconcatalternate](../62-revconcatalternate/README.md)
