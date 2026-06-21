# Skills for 12-recursion

## What You'll Learn

**Previous:** [11-arrays](../11-arrays/skills.md) | **Next:** [13-structs](../13-structs/skills.md)

**Challenge:** Write `Factorial(n uint) uint` that returns `n!` using recursion.

## Core Concept: Recursion

### What Is Recursion?

Recursion is when a function calls itself. Each call works on a smaller version of the problem. The function keeps calling itself until it reaches a condition where it can return an answer directly — that is the **base case**.

Every recursive function needs exactly two parts:

1. **Base case** — the condition that stops the recursion and returns a direct answer
2. **Recursive case** — a call to the same function with a simpler input, moving toward the base case

Without a base case the function calls itself forever, which crashes the program with a **stack overflow**.

### Factorial Example

Factorial is the classic introduction to recursion. The mathematical definition maps directly to code:

```
0! = 1          ← base case
n! = n × (n-1)! ← recursive case
```

```go
func Factorial(n uint) uint {
    if n == 0 {             // base case: stop here
        return 1
    }
    return n * Factorial(n-1)  // recursive case: smaller problem
}
```

### How the Call Stack Works

When Go calls a function it pushes a **stack frame** onto the call stack. The frame holds the function's local variables and where to return when the function finishes. Recursive calls stack up until the base case is reached, then unwind one by one.

Trace of `Factorial(4)`:

```
Factorial(4)
  → 4 * Factorial(3)
           → 3 * Factorial(2)
                    → 2 * Factorial(1)
                             → 1 * Factorial(0)
                                      → 1        (base case — start unwinding)
                             → 1 * 1 = 1
                    → 2 * 1 = 2
           → 3 * 2 = 6
  → 4 * 6 = 24
```

Each indented call is waiting for the result below it. Once `Factorial(0)` returns `1`, every waiting call can finish in reverse order.

### A Second Example: Fibonacci

The Fibonacci sequence — where each number is the sum of the two before it — is another natural fit for recursion:

```
fib(0) = 0
fib(1) = 1
fib(n) = fib(n-1) + fib(n-2)
```

```go
func fib(n uint) uint {
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 1
    }
    return fib(n-1) + fib(n-2)  // two recursive calls
}
```

Notice Fibonacci has **two** base cases and **two** recursive calls — recursion is flexible.

### GCD with Recursion (Preview of the Next Challenge)

The Euclidean algorithm for greatest common divisor is elegant with recursion:

```go
func Gcd(a, b uint) uint {
    if b == 0 {
        return a        // base case
    }
    return Gcd(b, a%b)  // recursive case — Euclidean algorithm
}
```

`a % b` is always smaller than `b`, so `b` shrinks toward `0` on every call.

### Recursion vs Iteration

Both approaches can solve the same problems. Choose based on clarity.

| | Recursion | Iteration |
|--|-----------|-----------|
| **Best for** | Tree traversals, divide-and-conquer, problems defined recursively (factorial, GCD, Fibonacci) | Simple loops over sequences, when performance matters |
| **Reads like** | The mathematical definition | Step-by-step instructions |
| **Risk** | Stack overflow for very deep calls | Accidental infinite loop |
| **Go style** | Used when the structure is naturally recursive | Preferred for most everyday loops |

Iterative factorial for comparison:

```go
func FactorialIterative(n uint) uint {
    result := uint(1)
    for i := uint(2); i <= n; i++ {
        result *= i
    }
    return result
}
```

Both versions produce identical results. The recursive version matches the mathematical definition; the iterative version uses a loop.

### Stack Overflow

If a recursive function never reaches its base case, Go runs out of stack space and crashes:

```go
// This crashes — no base case
func infinite(n uint) uint {
    return n * infinite(n-1)
}
```

```console
runtime: goroutine stack exceeds 1000000000-byte limit
fatal error: stack overflow
```

Always verify your base case catches every terminating condition before calling the function.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| No base case | Stack overflow crash | Add `if n == 0 { return 1 }` |
| Wrong base case value | Returns wrong answers for all inputs | `0! = 1`, not `0` |
| Recursive call moves in wrong direction | Stack overflow — `n` grows instead of shrinking | Call `Factorial(n-1)`, not `Factorial(n+1)` |
| Using `int` instead of `uint` | Type mismatch with function signature | Match the signature: `func Factorial(n uint) uint` |

## Solving This Challenge

### Algorithm

1. If `n == 0`, return `1` (base case — `0! = 1`).
2. Otherwise return `n * Factorial(n-1)` (recursive case).

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [13-structs](../13-structs/README.md)
