# Skills for 73-rpncalc

## What You'll Learn

**Previous:** [72-brackets](../72-brackets/skills.md) | **Next:** [74-brainfuck](../74-brainfuck/skills.md)

**Challenge:** Evaluate a Reverse Polish Notation (RPN) expression given as a command-line argument.

## Core Concept: Stack-Based Expression Evaluation

### What is Reverse Polish Notation?

In normal infix notation: `3 + 4 * 2` (operator between operands, needs parentheses for precedence)

In RPN (postfix notation): `3 4 2 * +` (operator comes AFTER its operands — no parentheses needed!)

**Rule**: When you see an operator, pop the last two numbers from the stack, apply the operation, push the result back.

### The Algorithm

```
Input: "1 2 * 3 * 4 +"

Stack starts empty: []

"1" → number → push: [1]
"2" → number → push: [1, 2]
"*" → operator → pop 2, pop 1, push 1*2=2: [2]
"3" → number → push: [2, 3]
"*" → operator → pop 3, pop 2, push 2*3=6: [6]
"4" → number → push: [6, 4]
"+" → operator → pop 4, pop 6, push 6+4=10: [10]

End: one item on stack → result = 10
```

### Implementation in Go

```go
package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Error")
        return
    }

    tokens := strings.Fields(os.Args[1])  // split on any whitespace
    stack := []int{}                       // stack of integers

    for _, token := range tokens {
        switch token {
        case "+", "-", "*", "/", "%":
            if len(stack) < 2 {  // not enough operands
                fmt.Println("Error")
                return
            }
            b := stack[len(stack)-1]  // top of stack
            a := stack[len(stack)-2]  // second from top
            stack = stack[:len(stack)-2]  // pop both
            var result int
            switch token {
            case "+": result = a + b
            case "-": result = a - b
            case "*": result = a * b
            case "/": result = a / b
            case "%": result = a % b
            }
            stack = append(stack, result)
        default:
            n, err := strconv.Atoi(token)
            if err != nil {  // not a number
                fmt.Println("Error")
                return
            }
            stack = append(stack, n)
        }
    }

    if len(stack) != 1 {  // valid expression leaves exactly 1 result
        fmt.Println("Error")
        return
    }
    fmt.Println(stack[0])
}
```

### Key Go Techniques Used

**`strings.Fields`** — splits on any whitespace, handles extra spaces:
```go
strings.Fields("  1   3  *  ")  // ["1", "3", "*"]
```

**`strconv.Atoi`** — converts string to int, returns error if not a number:
```go
n, err := strconv.Atoi("42")    // n=42, err=nil
n, err := strconv.Atoi("abc")   // n=0, err=non-nil
```

**`switch` without condition** — cleaner than if/else chains:
```go
switch token {
case "+": ...
case "-": ...
default: ...
}
```

### Error Conditions

| Condition | Example | Response |
|-----------|---------|----------|
| Wrong arg count | `go run .` | "Error" |
| Unknown token | `"1 x +"` | "Error" |
| Not enough operands | `"1 2 3 +"` | "Error" at end (extra on stack) |
| Stack has != 1 at end | incomplete expression | "Error" |

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [74-brainfuck](../74-brainfuck/skills.md)
