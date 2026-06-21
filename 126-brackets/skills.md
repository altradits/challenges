# Skills for 126-brackets

## What You'll Learn

**Previous:** [125-romannumbers](../125-romannumbers/skills.md) | **Next:** [127-rpncalc](../127-rpncalc/skills.md)

**Challenge:** Write a program that takes multiple string arguments and prints `OK` or `Error` for each, depending on whether its brackets are correctly matched.

## Core Concept: Stack-Based Bracket Matching

### Why a Stack?

Brackets must close in LIFO (Last-In, First-Out) order. When you see `([{`, the `{` must close first, then `[`, then `(`. A stack naturally tracks this: push each opening bracket, pop and verify on each closing bracket.

### The Algorithm

```
Input: "{[(0 + 0)(1 + 1)](3*(-1)){()}}"

Push {  → stack: ['{']
Push [  → stack: ['{', '[']
Push (  → stack: ['{', '[', '(']
See )   → pop '(' — matches ')' ✓  stack: ['{', '[']
Push (  → stack: ['{', '[', '(']
See )   → pop '(' — matches ')' ✓  stack: ['{', '[']
See ]   → pop '[' — matches ']' ✓  stack: ['{']
...
End: stack empty → OK
```

### Implementation

```go
package main

import (
    "fmt"
    "os"
)

func isOpen(r rune) bool {
    return r == '(' || r == '[' || r == '{'
}

func matches(open, close rune) bool {
    return (open == '(' && close == ')') ||
        (open == '[' && close == ']') ||
        (open == '{' && close == '}')
}

func checkBrackets(s string) bool {
    stack := []rune{}
    for _, r := range s {
        switch r {
        case '(', '[', '{':
            stack = append(stack, r)
        case ')', ']', '}':
            if len(stack) == 0 || !matches(stack[len(stack)-1], r) {
                return false
            }
            stack = stack[:len(stack)-1]  // pop
        }
    }
    return len(stack) == 0  // valid only if stack is empty at end
}

func main() {
    for _, arg := range os.Args[1:] {
        if checkBrackets(arg) {
            fmt.Println("OK")
        } else {
            fmt.Println("Error")
        }
    }
}
```

### Key Go Techniques

**Slice as stack** — push with `append`, pop with `stack[:len-1]`:
```go
stack = append(stack, r)           // push
top := stack[len(stack)-1]         // peek
stack = stack[:len(stack)-1]       // pop
```

**`os.Args[1:]`** — skip `os.Args[0]` (program name), iterate over all arguments:
```go
for _, arg := range os.Args[1:] { ... }
```

**`switch` on rune** — clean multi-case branching:
```go
switch r {
case '(', '[', '{': // push
case ')', ']', '}': // pop and verify
}
```

### Edge Cases

| Input | Stack at end | Result |
|-------|-------------|--------|
| `""` (empty) | empty | OK — no brackets = valid |
| `"hello"` | empty | OK — no brackets |
| `"([)]"` | fails mid-way | Error |
| `"(("` | `['(', '(']` | Error — unclosed |
| `")"` | empty when closing arrives | Error — nothing to match |

### Common Mistakes

| Mistake | Fix |
|---------|-----|
| Checking `len(stack) == 0` AFTER popping | Check BEFORE popping — empty stack means nothing to match |
| Using `stack[len(stack)]` (off by one) | Pop: `stack[:len(stack)-1]` — index `-1` is the last element |
| Forgetting to check stack empty at end | `return len(stack) == 0` — unclosed brackets stay on the stack |

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [127-rpncalc](../127-rpncalc/README.md)
