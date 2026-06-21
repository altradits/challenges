# Prerequisites for printmemory

## Before You Start

To solve this challenge you need to understand:

### 1. Fixed-Size Arrays
A `[N]T` array has a fixed size that is part of its type. Bytes default to 0 for unset positions.
```go
arr := [10]byte{'h', 'e', 'l', 'l', 'o'}
fmt.Println(arr[0]) // 104 ('h')
fmt.Println(arr[5]) // 0   (unset, zero value)
fmt.Println(len(arr)) // 10
```

### 2. fmt.Printf With %02x
Format a byte as two-digit lowercase hexadecimal:
```go
import "fmt"

b := byte(104)
fmt.Printf("%02x\n", b)  // "68"
fmt.Printf("%02x\n", byte(0)) // "00"
fmt.Printf("%02x\n", byte(42)) // "2a"
```

### 3. Modulo for Row Breaks
Use `(i+1) % 4 == 0` to print a newline after every 4 values.
```go
for i := 0; i < 10; i++ {
    fmt.Printf("%02x ", arr[i])
    if (i+1)%4 == 0 {
        fmt.Println()
    }
}
```

### 4. ASCII Printable Range Check
Characters with byte values 32 to 126 are printable ASCII. Outside this range, substitute a dot.
```go
if b >= 32 && b <= 126 {
    fmt.Printf("%c", b)
} else {
    fmt.Print(".")
}
```

### 5. fmt.Printf With %c
The `%c` verb prints the character corresponding to a byte value:
```go
fmt.Printf("%c", byte(104)) // "h"
fmt.Printf("%c", byte(42))  // "*"
```

## Review If Stuck

- [11-arrays](../11-arrays/skills.md) — dedicated lesson on Go arrays: declaration, iteration, arrays vs slices, `[10]byte`
- [../111-itoa-manual/skills.md](../111-itoa-manual/skills.md) — covers digit-by-digit extraction; similar pattern of processing one value at a time
- [../97-hashcode/skills.md](../97-hashcode/skills.md) — covers ASCII values and converting between int and character

## You're Ready When You Can...

- [ ] Declare and iterate over a `[10]byte` array
- [ ] Format bytes as two-digit hex with `fmt.Printf("%02x", b)`
- [ ] Print newlines after every 4 values using `(i+1) % 4 == 0`
- [ ] Check the printable ASCII range `[32, 126]` and print dot for values outside it

## Next Steps

- [113-printrevcomb](../113-printrevcomb/README.md)
