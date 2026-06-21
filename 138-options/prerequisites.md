# Prerequisites for options

## Before You Start

To solve this challenge you need to understand:

### 1. Bitwise OR to Set a Bit (`|=`)
Set a specific bit without changing others. Bit `n` is set by OR-ing with `1 << n`.

```go
var flags int32
flags |= 1 << 0   // set bit 0 (a)
flags |= 1 << 1   // set bit 1 (b)
flags |= 1 << 25  // set bit 25 (z)
```

### 2. Bit Position from a Letter
Subtract `'a'` from a lowercase letter to get its 0-based position.

```go
c := 'c'
bit := uint(c - 'a')  // 'c' - 'a' = 2
fmt.Println(bit)       // 2
```

### 3. Left Shift (`<<`) to Create a Bitmask
`1 << n` creates an integer with only bit `n` set.

```go
fmt.Println(1 << 0)   // 1   = 00000001
fmt.Println(1 << 1)   // 2   = 00000010
fmt.Println(1 << 7)   // 128 = 10000000
```

### 4. Extracting Byte Groups for Printing (`>>` and `& 0xFF`)
To print a 32-bit integer as 4 groups of 8 bits, shift right and mask.

```go
var flags int32 = 0x01020304
b3 := (flags >> 24) & 0xFF  // most significant byte
b2 := (flags >> 16) & 0xFF
b1 := (flags >> 8)  & 0xFF
b0 := flags         & 0xFF  // least significant byte
fmt.Printf("%08b %08b %08b %08b\n", b3, b2, b1, b0)
```

### 5. Iterating Over a String (arg[1:] for flag letters)
Skip the leading `-` and iterate over the rest of the argument.

```go
arg := "-abc"
for _, c := range arg[1:] {
    fmt.Printf("%c\n", c)  // a, b, c
}
```

### 6. `fmt.Printf` with `%08b`
`%08b` prints an integer in binary, zero-padded to 8 digits.

```go
fmt.Printf("%08b\n", 7)   // 00000111
fmt.Printf("%08b\n", 128) // 10000000
```

## Review If Stuck
- [137-itoabase](../137-itoabase/skills.md) — covers integer arithmetic and digit extraction (similar bit-level thinking)
- Prior `os.Args` challenges — covers iterating over command-line arguments

## You're Ready When You Can...
- [ ] Set a bit with `flags |= int32(1) << bit`
- [ ] Convert a letter to its bit position with `uint(c - 'a')`
- [ ] Extract each byte of a 32-bit int with `>> 24`, `>> 16`, `>> 8`, `& 0xFF`
- [ ] Print a byte as 8 binary digits with `fmt.Printf("%08b", b)`
- [ ] Iterate over `arg[1:]` to process option letters

## Next Steps

- [139-piglatin](../139-piglatin/README.md)
