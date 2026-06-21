# Skills for 128-brainfuck

## What You'll Learn

**Previous:** [127-rpncalc](../127-rpncalc/skills.md) | **Next:** [129-grouping](../129-grouping/skills.md)

**Challenge:** Write a Brainfuck interpreter — a program that executes Brainfuck code passed as a command-line argument.

## Core Concept: Interpreter Design with Array-as-Memory

### What is a Brainfuck Interpreter?

An interpreter is a program that reads instructions in one language and executes them. Brainfuck has 8 commands operating on an array of bytes:

| Command | Meaning |
|---------|---------|
| `>` | Move pointer right (increment pointer) |
| `<` | Move pointer left (decrement pointer) |
| `+` | Increment byte at pointer |
| `-` | Decrement byte at pointer |
| `.` | Print byte at pointer as ASCII character |
| `[` | If byte at pointer is 0, jump PAST matching `]` |
| `]` | If byte at pointer is NOT 0, jump BACK to matching `[` |
| anything else | Comment, ignore |

### The Memory Model

```go
memory := [2048]byte{}  // 2048 bytes, all initialized to 0
ptr := 0                // data pointer, starts at byte 0
```

The data pointer `ptr` moves left/right. The `[` and `]` commands implement loops.

### The Loop Mechanism (`[` and `]`)

The tricky part is implementing jumps. When you hit `[` and the current byte is 0, you must skip forward past the MATCHING `]`. The approach: count nesting depth.

```go
// Skip forward past matching ]
if code[ip] == '[' && memory[ptr] == 0 {
    depth := 1
    for depth > 0 {
        ip++
        if code[ip] == '[' { depth++ }
        if code[ip] == ']' { depth-- }
    }
}

// Jump back to matching [
if code[ip] == ']' && memory[ptr] != 0 {
    depth := 1
    for depth > 0 {
        ip--
        if code[ip] == ']' { depth++ }
        if code[ip] == '[' { depth-- }
    }
}
```

### Complete Implementation

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        return
    }
    code := os.Args[1]
    memory := [2048]byte{}
    ptr := 0

    for ip := 0; ip < len(code); ip++ {
        switch code[ip] {
        case '>':
            ptr++
        case '<':
            ptr--
        case '+':
            memory[ptr]++
        case '-':
            memory[ptr]--
        case '.':
            fmt.Printf("%c", memory[ptr])
        case '[':
            if memory[ptr] == 0 {
                depth := 1
                for depth > 0 {
                    ip++
                    if code[ip] == '[' { depth++ }
                    if code[ip] == ']' { depth-- }
                }
            }
        case ']':
            if memory[ptr] != 0 {
                depth := 1
                for depth > 0 {
                    ip--
                    if code[ip] == ']' { depth++ }
                    if code[ip] == '[' { depth-- }
                }
            }
        }
    }
}
```

### Key Concepts

**Array (`[2048]byte`)** — fixed-size memory model. Unlike slices, arrays have a fixed compile-time size and are value types.

**Instruction pointer (`ip`)** — an index into the code string, moves through instructions one at a time. The `for` loop increments it, but the `[`/`]` handlers can jump it forward or backward.

**Depth counter for bracket matching** — the same technique used in bracket matching (126-brackets), but this time you're jumping the instruction pointer instead of validating.

**`fmt.Printf("%c", byte)`** — prints a byte as its ASCII character representation.

### Debugging Brainfuck

```
"++++++++++[>+++++++<-]>."

Memory starts: [0, 0, 0, ...]
ptr=0

10 x '+': memory[0] = 10
'[': memory[0]=10 != 0, enter loop
  '>': ptr=1
  7 x '+': memory[1] += 7
  '<': ptr=0
  '-': memory[0]--
']': if memory[0] != 0, jump back to [
... repeat 10 times total ...
After loop: memory[0]=0, memory[1]=70 ('F' in ASCII)
'>': ptr=1
'.': print memory[1] = 70 = 'F'
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Using slice instead of `[2048]byte` | Challenge requires array | Use `[2048]byte{}` |
| Not initializing depth to 1 | Exits loop immediately | `depth := 1` before the inner loop |
| Forgetting `ip--` before outer `ip++` | Skips one instruction | Use `ip--` at end of `]` handler so outer loop's `ip++` brings you to the right place — or restructure |

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [129-grouping](../129-grouping/README.md)
