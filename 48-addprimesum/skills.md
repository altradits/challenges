# Skills for addprimesum

## What You'll Learn

**Previous:** [../47-zipstring/skills.md](../47-zipstring/skills.md) | **Next:** [../49-canjump/README.md](../49-canjump/README.md)

**Challenge:** A program that takes a positive integer `n` from the command line and prints the sum of all primes up to and including `n`.

## Core Concept: Prime Checking Plus strconv.Atoi for CLI Input

### What Is strconv.Atoi?
`strconv.Atoi` converts a string argument (from `os.Args`) to an integer. It returns both the value and an error. If the conversion fails (non-numeric input), the error is non-nil.

```go
import "strconv"

n, err := strconv.Atoi("42")
// n == 42, err == nil

n2, err2 := strconv.Atoi("abc")
// n2 == 0, err2 != nil
```

### The Full Program

```go
import (
    "fmt"
    "os"
    "strconv"
)

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

func main() {
    if len(os.Args) != 2 {
        fmt.Println(0)
        return
    }
    n, err := strconv.Atoi(os.Args[1])
    if err != nil || n <= 0 {
        fmt.Println(0)
        return
    }
    sum := 0
    for i := 2; i <= n; i++ {
        if isPrime(i) {
            sum += i
        }
    }
    fmt.Println(sum)
}
```

### Using strconv.Atoi Safely
Always check the error. If the user passes `"-2"`, `Atoi` succeeds but gives a negative number — check `n <= 0` separately:
```go
n, err := strconv.Atoi(os.Args[1])
if err != nil || n <= 0 {
    fmt.Println(0)
    return
}
```

### Summing Primes Up to n
Loop from 2 to `n` (inclusive), add `i` to sum whenever `isPrime(i)` is true:
```go
sum := 0
for i := 2; i <= n; i++ {
    if isPrime(i) {
        sum += i
    }
}
```

For `n=5`: primes are 2, 3, 5 → sum = 10.
For `n=7`: primes are 2, 3, 5, 7 → sum = 17.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Not checking `err != nil` from Atoi | Non-numeric input causes wrong output | Always check the error |
| Not checking `n <= 0` | Negative numbers or zero should return 0 | Add `\|\| n <= 0` to the error check |
| Starting loop from 0 or 1 | 0 and 1 are not prime, isPrime handles this, but unnecessary iterations | Start from 2 |

## Solving This Challenge

### Algorithm
1. If `len(os.Args) != 2`, print 0 and return.
2. Parse `os.Args[1]` with `strconv.Atoi`. If error or `n <= 0`, print 0 and return.
3. Sum all `i` from 2 to `n` where `isPrime(i)` is true.
4. Print the sum.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [../49-canjump/README.md](../49-canjump/README.md)
