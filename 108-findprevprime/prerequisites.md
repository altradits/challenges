# Prerequisites for findprevprime

## Before You Start

To solve this challenge you need to understand:

### 1. Modulo for Divisibility Testing
`n % i == 0` means `i` divides `n` evenly (no remainder). Use this to check if a number is composite.
```go
fmt.Println(10 % 2)  // 0 — 2 divides 10
fmt.Println(10 % 3)  // 1 — 3 does not divide 10
fmt.Println(9 % 3)   // 0 — 3 divides 9
```

### 2. For Loop for Divisibility Check
Check all potential factors from 2 up to the square root of the number.
```go
func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false  // found a factor
        }
    }
    return true  // no factor found
}
```

### 3. Downward Iteration
Count down from a starting value toward zero.
```go
n := 10
for n >= 2 {
    fmt.Println(n)
    n--
}
// 10, 9, 8, ..., 2
```

### 4. Returning Early From a Loop
Return a result as soon as a condition is met, without waiting for the loop to finish.
```go
for n >= 2 {
    if isPrime(n) {
        return n  // found it — stop immediately
    }
    n--
}
```

### 5. Helper Functions
Break complex logic into smaller, named functions.
```go
func isPrime(n int) bool { /* ... */ }

func FindPrevPrime(nb int) int {
    for nb >= 2 {
        if isPrime(nb) {
            return nb
        }
        nb--
    }
    return 0
}
```

## Review If Stuck

- [../96-gcd/skills.md](../96-gcd/skills.md) — covers the modulo operator and breaking a problem into smaller steps
- [../94-digitlen/skills.md](../94-digitlen/skills.md) — covers while-style for loops that count down/change `n`

## You're Ready When You Can...

- [ ] Use `n % i == 0` to test whether `i` divides `n`
- [ ] Write an `isPrime` function that checks divisibility up to `sqrt(n)` using `i*i <= n`
- [ ] Count downward from `nb` to 2 in a loop
- [ ] Return the first `nb` for which `isPrime(nb)` is true

## Next Steps

- [109-fromto](../109-fromto/README.md)
