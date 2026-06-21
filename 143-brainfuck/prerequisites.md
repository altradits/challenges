# Prerequisites for 143-brainfuck

## Before You Start

### 1. Arrays (`[n]byte`) from 11-arrays

Brainfuck needs a fixed-size memory: `[2048]byte{}`. This is a Go ARRAY (fixed size, value type), not a slice.

```go
memory := [2048]byte{}  // 2048 bytes, all zero
memory[0] = 65          // store byte value 65 ('A')
memory[0]++             // increment in-place
```

Review: [11-arrays](../11-arrays/skills.md)

### 2. `switch` statement from 142-rpncalc

```go
switch code[ip] {
case '>': ptr++
case '<': ptr--
case '+': memory[ptr]++
default:  // ignore
}
```

Review: [142-rpncalc](../142-rpncalc/skills.md)

### 3. Index-based for loop (not range)

The instruction pointer `ip` needs to be modified inside the loop (jump forward/backward), so you can't use `for...range`. Use an index-based loop:

```go
for ip := 0; ip < len(code); ip++ {
    // ip can be modified inside the loop body
}
```

Review: [07-forloops](../07-forloops/skills.md)

### 4. Nested loops for bracket matching depth

```go
depth := 1
for depth > 0 {
    ip++
    if code[ip] == '[' { depth++ }
    if code[ip] == ']' { depth-- }
}
```

Review: [141-brackets](../141-brackets/skills.md) — same nesting logic.

### 5. `fmt.Printf("%c", byte)` for character output

```go
fmt.Printf("%c", memory[ptr])  // prints byte as ASCII char
```

## You're Ready When You Can...

- [ ] Declare and use a `[2048]byte` array
- [ ] Modify an index variable inside a for loop
- [ ] Use a depth counter to match opening/closing brackets
- [ ] Use `fmt.Printf("%c", b)` to print a byte as a character

## Next Steps

- [144-grouping](../144-grouping/README.md)
