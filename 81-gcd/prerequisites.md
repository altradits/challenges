# Prerequisites for gcd

## Before You Start

To solve this challenge you need to understand:

### 1. The Modulo Operator
`%` gives the remainder after integer division. Central to the Euclidean algorithm.
```go
fmt.Println(42 % 10) // 2  (42 = 4×10 + 2)
fmt.Println(10 % 3)  // 1
fmt.Println(9 % 3)   // 0  ← perfectly divisible
```

### 2. Functions with Return Values
A Go function takes parameters and can return a value. The return type must match.
```go
func Multiply(a, b int) int {
    return a * b
}
```

### 3. The uint Type
`uint` is an unsigned (non-negative) integer. Use it when the challenge signature requires it.
```go
var x uint = 10
var y uint = 3
fmt.Println(x % y) // 1
```

### 4. What is Recursion?
A function calls itself to solve a smaller version of the same problem. It must have a condition (base case) that stops the calls.
```go
func sumDown(n uint) uint {
    if n == 0 {
        return 0        // base case
    }
    return n + sumDown(n-1)  // recursive call
}
```

### 5. If/Else for Base Cases
Without a base case, a recursive function loops forever and crashes with a stack overflow.
```go
if b == 0 {
    return a  // stop here
}
return Gcd(b, a%b)  // keep going
```

## Review If Stuck

- [12-recursion](../12-recursion/skills.md) — dedicated lesson on recursion: base cases, call stack, factorial and GCD examples
- [../79-digitlen/skills.md](../79-digitlen/skills.md) — covers the same "loop until zero" logic as an iterative alternative to recursion

## You're Ready When You Can...

- [ ] Use `%` to compute the remainder of integer division
- [ ] Write a function that calls itself (recursion)
- [ ] Identify the base case condition that stops the recursion
- [ ] Use the `uint` type for unsigned integer parameters

## Next Steps

- [82-hashcode](../82-hashcode/README.md)
