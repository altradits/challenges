# Skills for printmemory

## What You'll Learn

**Previous:** [../42-arrays/skills.md](../42-arrays/skills.md) | **Next:** [../44-printrevcomb/README.md](../44-printrevcomb/README.md)

**Challenge:** Print a fixed-size `[10]byte` array as hex values (4 per row), then print the printable ASCII characters (dots for non-printable).

## Core Concept: Fixed-Size Arrays and fmt.Printf Format Verbs

### What Is a [10]byte Array?
A `[10]byte` is a fixed-size array of 10 bytes. Unlike slices, the size is part of the type and cannot change. Iterate with a standard index loop or with `range`.

```go
arr := [10]byte{'h', 'e', 'l', 'l', 'o'}
fmt.Println(arr[0]) // 104 — byte value of 'h'
fmt.Println(arr[5]) // 0   — zero value for unset bytes
```

### Printing Hex Values With fmt.Printf
The `%02x` verb formats a byte as lowercase hexadecimal with at least 2 digits:
```go
fmt.Printf("%02x ", byte(104)) // "68 "
fmt.Printf("%02x ", byte(0))   // "00 "
fmt.Printf("%02x ", byte(42))  // "2a "
```

### The Output Format
Print 4 bytes per row as hex, then the printable ASCII on one final row (dots for non-printable):

```
68 65 6c 6c
6f 10 15 2a
00 00
hello..*..$
```

```go
func PrintMemory(arr [10]byte) {
    // Print hex rows of 4
    for i := 0; i < 10; i++ {
        fmt.Printf("%02x ", arr[i])
        if (i+1)%4 == 0 || i == 9 {
            fmt.Println()
        }
    }
    // Print ASCII representation
    for _, b := range arr {
        if b >= 32 && b <= 126 {
            fmt.Printf("%c", b)
        } else {
            fmt.Print(".")
        }
    }
    fmt.Println()
}
```

### Printable ASCII Range
ASCII graphic characters are in the range 32 (space) to 126 (`~`). Outside this range, print a dot:
```go
if b >= 32 && b <= 126 {
    fmt.Printf("%c", b)  // printable — show the character
} else {
    fmt.Print(".")       // non-printable — show dot
}
```

### Controlling Row Breaks
Print a newline after every 4 values. Use `(i+1) % 4 == 0` to detect the end of a group, and add a special case for the last element (index 9):
```go
if (i+1)%4 == 0 || i == 9 {
    fmt.Println()
}
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Using `%x` instead of `%02x` | Single-digit hex values have no leading zero | Always use `%02x` |
| Treating `[10]byte` like a slice | Arrays and slices are different types | The parameter is `[10]byte`, not `[]byte` |
| Wrong printable range | Using `> 32` misses the space character | Use `>= 32 && <= 126` |

## Solving This Challenge

### Algorithm
1. Loop i from 0 to 9: print `arr[i]` as `%02x `. Print newline every 4 values (and at the end).
2. Loop over `arr`: if byte is in [32,126], print the character; else print `.`.
3. Print final newline.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [../44-printrevcomb/README.md](../44-printrevcomb/README.md)
