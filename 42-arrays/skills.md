# Skills for arrays

## What You'll Learn

**Previous:** [../41-itoa-35/skills.md](../41-itoa-35/skills.md) | **Next:** [../43-printmemory/skills.md](../43-printmemory/skills.md)

**Challenge:** Sum all elements of a `[5]int` array passed to a function.

## Core Concept: Arrays — Fixed-Size, Value-Type Collections

### What Is an Array in Go?

An array is a **fixed-size**, **value-type** sequence of elements. The size is part of the type itself — `[5]int` and `[10]int` are two completely different types. Because the size is known at compile time, arrays are stack-allocated and require no heap allocation.

Arrays are less common in day-to-day Go than slices, but they are essential when you need a **fixed number of elements** — for example, `[10]byte` for a memory buffer, `[4]byte` for an IPv4 address, or `[32]byte` for a SHA-256 hash.

### Declaring and Initializing Arrays

```go
// Zero-initialized (all elements are the zero value for the type)
var arr [5]int           // [0, 0, 0, 0, 0]

// With literal values
arr := [5]int{1, 2, 3, 4, 5}

// Let the compiler count the size
arr := [...]int{1, 2, 3}  // [3]int — size is 3

// Initialize specific indices only (others are zero)
arr := [5]int{0: 10, 4: 50}  // [10, 0, 0, 0, 50]
```

### Accessing and Mutating Elements

```go
arr := [5]int{10, 20, 30, 40, 50}

fmt.Println(arr[0])  // 10
fmt.Println(arr[4])  // 50

arr[2] = 99
fmt.Println(arr)  // [10 20 99 40 50]
```

### Iterating Over an Array

```go
arr := [5]int{1, 2, 3, 4, 5}

// Index and value
for i, v := range arr {
    fmt.Println(i, v)
}

// Index only
for i := range arr {
    fmt.Println(arr[i])
}

// Classic C-style (also works)
for i := 0; i < len(arr); i++ {
    fmt.Println(arr[i])
}
```

### Arrays vs Slices — Key Differences

| Feature | Array `[5]int` | Slice `[]int` |
|---------|---------------|---------------|
| Size | Fixed at compile time | Dynamic |
| Type includes size? | Yes — `[5]int` ≠ `[6]int` | No — all `[]int` slices share the same type |
| Value vs reference | **Value** — assignment copies all elements | Reference — shares underlying array |
| Pass to function | Gets **copied** entirely | Shares the underlying data |
| `len()` | Compile-time constant | Runtime value |
| `append()` | Not available | Core operation |

**Value semantics in action:**
```go
a := [3]int{1, 2, 3}
b := a          // b is a full copy
b[0] = 99
fmt.Println(a)  // [1, 2, 3]  — a is unchanged
fmt.Println(b)  // [99, 2, 3]
```

Compare with slices:
```go
a := []int{1, 2, 3}
b := a          // b points to the same underlying array
b[0] = 99
fmt.Println(a)  // [99, 2, 3]  — a IS changed
```

### The `[10]byte` Type

This is exactly what `43-printmemory` uses. A byte (`uint8`) holds a single ASCII/UTF-8 code point value from 0 to 255.

```go
var mem [10]byte          // 10 bytes, all zero

mem[0] = 'h'             // character literal — stores ASCII value 104
mem[1] = 72              // numeric literal — stores 72 (which is 'H')

fmt.Printf("%02x ", mem[0])  // "68" — lowercase hex, always 2 digits
fmt.Printf("%c", mem[0])     // "h"  — character representation
fmt.Printf("%d", mem[0])     // "104" — decimal value
```

**Printing printable vs non-printable bytes (as in 43-printmemory):**
```go
arr := [10]byte{'h', 'e', 'l', 'l', 'o'}
for i := 0; i < len(arr); i++ {
    // ASCII graphic (printable) characters: 33–126
    if arr[i] >= 33 && arr[i] <= 126 {
        fmt.Printf("%c", arr[i])  // print the character
    } else {
        fmt.Print(".")            // non-printable → dot
    }
}
```

### Converting an Array to a Slice

Slice the whole array with `[:]`. The slice shares the same underlying memory:

```go
arr := [5]int{1, 2, 3, 4, 5}
s := arr[:]    // []int — length 5, points into arr's memory
s[0] = 99
fmt.Println(arr[0])  // 99 — arr was modified through the slice
```

This is useful when a function accepts a `[]int` but you only have a `[5]int`.

### Why Arrays Exist

- **Stack allocation** — no heap allocation, no garbage collector pressure.
- **Exact size guarantees** — `[32]byte` is always exactly 32 bytes, useful for cryptographic hashes, fixed-size network packets, etc.
- **Value semantics** — copying an array copies all its data, which can be desirable when you want a snapshot with no aliasing.

In practice, most Go code uses slices. Arrays appear when the size is fixed by design (e.g., `[10]byte` memory buffers, `[4]byte` IP addresses).

## Solving This Challenge

### Algorithm

Iterate over every element of the `[5]int` array and accumulate the sum:

```go
func SumArray(arr [5]int) int {
    sum := 0
    for _, v := range arr {
        sum += v
    }
    return sum
}
```

Because `arr` is an array (value type), the function receives a copy — the caller's array is never modified.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Writing `arr []int` in the signature | That's a slice, not an array — won't compile with `[5]int` argument | Write `arr [5]int` |
| Using `append(arr, ...)` | Arrays don't support `append` | Use a loop with index access |
| Expecting the function to modify the caller's array | Arrays are value types; the function gets a copy | Return a new value instead |

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [../43-printmemory/skills.md](../43-printmemory/skills.md)
