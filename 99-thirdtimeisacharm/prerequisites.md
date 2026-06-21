# Prerequisites for thirdtimeisacharm

## Before You Start

To solve this challenge you need to understand:

### 1. Modulo for Selecting Every Nth Item
`(i+1) % 3 == 0` is true when `i+1` is divisible by 3 — i.e., at 1-based positions 3, 6, 9, ...
```go
for i := 0; i < 9; i++ {
    if (i+1)%3 == 0 {
        fmt.Println(i, "is a multiple of 3 (1-based)")
    }
}
// 2, 5, 8 (0-based indices of 3rd, 6th, 9th items)
```

### 2. for/range Over a String
Using `for i, c := range s` gives both the byte index and the rune value.
```go
s := "hello"
for i, c := range s {
    fmt.Printf("index %d: %c\n", i, c)
}
```

### 3. Building a Result String
Append selected characters to a growing result string.
```go
result := ""
for i, c := range str {
    if (i+1)%3 == 0 {
        result += string(c)
    }
}
```

### 4. Returning a Newline-Terminated String
The challenge requires `\n` at the end of the returned string (even when empty).
```go
return result + "\n"
```

### 5. Stepping a Loop by N
Alternatively, start at index 2 and advance by 3:
```go
for i := 2; i < len(str); i += 3 {
    // i is always at the 3rd, 6th, 9th... positions
}
```

## Review If Stuck

- [../81-gcd/skills.md](../81-gcd/skills.md) — covers modulo arithmetic
- [../78-countrepeats/skills.md](../78-countrepeats/skills.md) — covers index-based for loop over a string

## You're Ready When You Can...

- [ ] Use `(i+1) % 3 == 0` to select every third character (1-based)
- [ ] Iterate over a string with both index and rune using `for i, c := range`
- [ ] Append selected characters to a result string
- [ ] Return a newline-terminated string even when no characters are selected

## Next Steps

- [100-weareunique](../100-weareunique/README.md)
