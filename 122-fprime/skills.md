# Skills for fprime

## What You'll Learn

**Previous:** [121-concatslice](../121-concatslice/skills.md) | **Next:** [123-hiddenp](../123-hiddenp/skills.md)

**Challenge:** Display the prime factorization of a positive integer, e.g. `225225` → `3*3*5*5*7*11*13`.

## Core Concept: Prime Factorization by Trial Division

### What Is It?

Prime factorization breaks a number into its prime building blocks. For example, 42 = 2 × 3 × 7. Trial division is the simplest algorithm: try dividing by 2, then 3, then 4, etc. — you only need to go up to the square root of the number.

### How It Works

**Step 1 — Parse and validate the input:**

```go
import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    if len(os.Args) != 2 {
        return
    }
    n, err := strconv.Atoi(os.Args[1])
    if err != nil || n <= 1 {
        return
    }
```

**Step 2 — Trial division loop:**

Start with divisor `d = 2`. While `d * d <= n`, try to divide:
- If `n % d == 0`: `d` is a prime factor. Collect it, divide `n` by `d`, and try `d` again (for repeated factors like `3*3`).
- If `n % d != 0`: increment `d` and try the next candidate.

```go
    factors := []int{}
    d := 2
    for d*d <= n {
        for n%d == 0 {
            factors = append(factors, d)
            n /= d
        }
        d++
    }
    // Any remaining n > 1 is itself a prime factor
    if n > 1 {
        factors = append(factors, n)
    }
```

**Why stop at sqrt(n)?**
If `n` has no factor from 2 up to its square root, then `n` itself is prime. Example: checking 97 — no factor up to 9 (sqrt(97) ≈ 9.8), so 97 is prime.

**Step 3 — Print factors separated by `*`:**

```go
    for i, f := range factors {
        if i > 0 {
            fmt.Print("*")
        }
        fmt.Print(f)
    }
    fmt.Println()
}
```

**Trace through `42`:**
- d=2: 42%2==0 → factor 2, n=21; 21%2!=0
- d=3: 21%3==0 → factor 3, n=7; 7%3!=0
- d=4: 4*4=16 > 7 → loop ends
- n=7 > 1 → factor 7
- Output: `2*3*7`

**Trace through `9539` (already prime):**
- No d from 2 to 97 divides 9539
- n=9539 > 1 → factor 9539
- Output: `9539`

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `for d <= n` instead of `for d*d <= n` | Very slow for large primes (O(n) instead of O(sqrt(n))) | Use `d*d <= n` to stop early |
| Not re-trying the same `d` after dividing | Misses repeated factors like 3×3 in 225225 | Use inner `for n%d == 0` loop |
| Printing nothing when `n` is prime | Prime numbers have exactly one factor: themselves | After the loop, if `n > 1` print it |
| Accepting `n = 1` or `n = 0` | 1 and 0 have no prime factors | Return early if `n <= 1` |

## Solving This Challenge

### Algorithm
1. Validate: exactly 1 arg, parse with `strconv.Atoi`, check `n > 1`.
2. `d = 2`; while `d*d <= n`: while `n%d == 0` collect `d` and divide; then `d++`.
3. If `n > 1` after loop, collect `n`.
4. Print factors joined by `*` followed by a newline.

## The Challenge
See [README.md](README.md) for full description.

**Next:** [123-hiddenp](../123-hiddenp/README.md)
