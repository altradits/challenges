# Prerequisites for 101-findsubstring

## Before You Start

### 1. String Indexing with `s[i]`

In Go, `s[i]` accesses the **byte** at index `i`. For ASCII strings, one byte equals one character.

```go
s := "Hello"
fmt.Println(s[0])         // 72 — byte value of 'H'
fmt.Println(s[4])         // 111 — byte value of 'o'
fmt.Println(s[0] == 'H')  // true
```

### 2. String Length with `len(s)`

`len(s)` returns the number of bytes in the string (for ASCII, this equals character count).

```go
fmt.Println(len("Hello"))  // 5
fmt.Println(len(""))       // 0
fmt.Println(len("Go!"))    // 3
```

### 3. String Slicing `s[a:b]`

Extract a substring from index `a` up to (but not including) `b`:

```go
s := "Hello World"
fmt.Println(s[6:11])  // "World"
fmt.Println(s[0:5])   // "Hello"
fmt.Println(s[1:])    // "ello World" (from 1 to end)
```

### 4. Nested Loops

Substring search uses an outer loop over start positions and an inner loop comparing characters:

```go
for i := 0; i < 5; i++ {
    for j := 0; j < 3; j++ {
        fmt.Printf("outer=%d inner=%d\n", i, j)
    }
}
```

### 5. `break` to Exit an Inner Loop

When a mismatch is found, `break` exits the inner loop immediately to avoid wasted work:

```go
match := true
for j := 0; j < len(pattern); j++ {
    if text[i+j] != pattern[j] {
        match = false
        break  // no need to keep checking this window
    }
}
```

## Review If Stuck

- [100-countwords](../100-countwords/skills.md) — covers character-by-character scanning

## You're Ready When You Can...

- [ ] Access the byte at index `i` in a string using `s[i]`
- [ ] Get the length of a string with `len(s)`
- [ ] Write a nested loop in Go
- [ ] Use `break` to exit an inner loop early
- [ ] Extract a substring with `s[a:b]`

## Next Steps

- [102-replaceall](../102-replaceall/README.md)
