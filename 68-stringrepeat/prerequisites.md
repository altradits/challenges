# Prerequisites for 68-stringrepeat

## Before You Start

### 1. `strings.Builder` (from 65-stringbuilder)

The efficient string building tool:

```go
var b strings.Builder
b.WriteString("abc")
b.WriteString("abc")
fmt.Println(b.String())  // "abcabc"
```

### 2. `b.Grow(n)` — Pre-Allocation

When you know the exact output size, tell Builder to allocate it upfront:

```go
var b strings.Builder
b.Grow(9)  // allocate 9 bytes
b.WriteString("abc")  // write 3
b.WriteString("abc")  // write 3
b.WriteString("abc")  // write 3
// Total 9 bytes — no reallocation needed
```

`b.Grow(len(s) * n)` pre-allocates the exact size for Repeat.

### 3. A Simple Loop n Times

```go
n := 4
for i := 0; i < n; i++ {
    fmt.Println("iteration", i)
}
// prints 4 lines
```

For Repeat, you write the same string in each iteration.

### 4. Edge Cases: Zero and Negative n

If `n <= 0`, return an empty string — you can't repeat something zero or negative times:

```go
if n <= 0 {
    return ""
}
```

### 5. Why Pre-Allocating is a Win Here

For `Repeat("x", 1000)`:
- Without Grow: Builder might reallocate at 8, 16, 32, 64, ..., 512, 1024 bytes (10 reallocations)
- With `b.Grow(1000)`: single allocation, no reallocation

You know the size exactly: `len(s) * n`. Use it.

## Review If Stuck

- [65-stringbuilder](../65-stringbuilder/skills.md) — covers `strings.Builder` and O(n) vs O(n²)
- [67-stringjoin](../67-stringjoin/skills.md) — another example of Builder with Grow

## You're Ready When You Can...

- [ ] Use `var b strings.Builder` and `b.WriteString`
- [ ] Call `b.Grow(len(s) * n)` to pre-allocate the right capacity
- [ ] Write a `for i := 0; i < n; i++` loop
- [ ] Return `""` when `n <= 0`

## Next Steps

- [69-stringreplace](../69-stringreplace/README.md)
