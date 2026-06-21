# Skills for gcd

## What You'll Learn

**Previous:** [95-firstword](../95-firstword/skills.md) | **Next:** [97-hashcode](../97-hashcode/skills.md)

**Challenge:** Find the greatest common divisor of two unsigned integers using the Euclidean algorithm.

## Core Concept: Recursion and the Euclidean Algorithm

### What Is Recursion?
A recursive function calls *itself* with a simpler version of the same problem, until it reaches a base case — a condition where it stops and returns directly.

```go
func countdown(n int) {
    if n == 0 {           // base case — stop here
        fmt.Println("done")
        return
    }
    fmt.Println(n)
    countdown(n - 1)      // recursive call — smaller problem
}
```

Every recursive function needs:
1. A base case that stops the recursion
2. A recursive call that moves toward the base case

### The Euclidean Algorithm
The mathematical rule is:
- `gcd(a, 0) = a` — base case
- `gcd(a, b) = gcd(b, a % b)` — recursive step

Each step the second argument gets smaller (it becomes `a % b`, which is always less than `b`). Eventually the second argument reaches 0 and the recursion stops.

```go
func Gcd(a, b uint) uint {
    if a == 0 || b == 0 {
        return 0
    }
    if b == 0 {
        return a
    }
    return Gcd(b, a%b)
}
```

Step-by-step for `Gcd(42, 10)`:
- `Gcd(42, 10)` → `Gcd(10, 42%10)` = `Gcd(10, 2)`
- `Gcd(10, 2)`  → `Gcd(2, 10%2)`  = `Gcd(2, 0)`
- `Gcd(2, 0)`   → b==0, return `2`
- Answer: `2`

### Iterative Version (No Recursion)
The same algorithm without recursion using a loop:
```go
func Gcd(a, b uint) uint {
    if a == 0 || b == 0 {
        return 0
    }
    for b != 0 {
        a, b = b, a%b  // simultaneous assignment
    }
    return a
}
```

### The uint Type
`uint` is an unsigned integer — always zero or positive. Arithmetic works the same as `int`.
```go
var a uint = 42
var b uint = 10
fmt.Println(a % b) // 2
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| No base case | Infinite recursion → stack overflow | Always have `if b == 0 { return a }` |
| Forgetting the zero-input guard | Returns `a` when either input is 0, but challenge expects 0 | Add `if a == 0 \|\| b == 0 { return 0 }` first |
| Using `int` instead of `uint` | Type mismatch, won't compile | Match the signature: `func Gcd(a, b uint) uint` |

## Solving This Challenge

### Algorithm
1. If `a == 0` or `b == 0`, return `0`.
2. If `b == 0`, return `a` (Euclidean base case).
3. Otherwise return `Gcd(b, a%b)`.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [97-hashcode](../97-hashcode/README.md)
