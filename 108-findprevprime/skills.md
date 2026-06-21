# Skills for findprevprime

## What You'll Learn

**Previous:** [107-expandstr](../107-expandstr/skills.md) | **Next:** [109-fromto](../109-fromto/skills.md)

**Challenge:** Find the first prime number that is less than or equal to a given integer by iterating downward from it.

## Core Concept: Prime Checking and Backward Iteration

### What Is a Prime Number?
A prime number is an integer greater than 1 that has no divisors other than 1 and itself. The smallest primes are 2, 3, 5, 7, 11, 13...

### Writing an isPrime Helper
Check divisibility from 2 up to the square root of `n`. If no divisor is found, `n` is prime.

```go
func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}
```

Why `i*i <= n` instead of `i <= n`? Because if `n` has a factor larger than its square root, it must also have one smaller than its square root. Stopping at `sqrt(n)` is a significant performance improvement.

### Iterating Downward From n
Start at `nb` and decrease by 1 until you find a prime:
```go
func FindPrevPrime(nb int) int {
    for nb >= 2 {
        if isPrime(nb) {
            return nb
        }
        nb--
    }
    return 0  // no prime found at or below nb
}
```

### Step-by-Step for `FindPrevPrime(4)`
- nb=4: isPrime(4)? 4%2==0 → No
- nb=3: isPrime(3)? No divisor from 2 to sqrt(3)≈1.7 → Yes
- Return 3

### Step-by-Step for `FindPrevPrime(5)`
- nb=5: isPrime(5)? No divisor from 2 to sqrt(5)≈2.2 → Yes
- Return 5

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Checking `i <= n` instead of `i*i <= n` | Correct but very slow for large numbers | Use `i*i <= n` |
| Returning `nb` before checking `isPrime` | Returns non-prime | Always check `isPrime(nb)` before returning |
| Not handling `nb < 2` | Should return 0 (no primes below 2) | The `nb >= 2` loop condition handles this |

## Solving This Challenge

### Algorithm
1. Write `isPrime(n int) bool` that returns true if `n >= 2` and has no factor in `[2, sqrt(n)]`.
2. In `FindPrevPrime(nb int)`: while `nb >= 2`, check `isPrime(nb)`. If yes, return it. Otherwise `nb--`.
3. Return 0 if no prime found.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [109-fromto](../109-fromto/README.md)
