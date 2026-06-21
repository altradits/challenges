# Prerequisites for addprimesum

## Before You Start

To solve this challenge you need to understand:

### 1. strconv.Atoi
Converts a string to an integer. Returns the value and an error. Always check the error.
```go
import "strconv"

n, err := strconv.Atoi("42")
if err != nil {
    fmt.Println("not a number")
} else {
    fmt.Println(n) // 42
}

_, err2 := strconv.Atoi("abc")
fmt.Println(err2 != nil) // true — conversion failed
```

### 2. Prime Checking
A number is prime if no integer from 2 to sqrt(n) divides it evenly.
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

### 3. Summing Values in a Loop
Accumulate a running total with `sum += value`:
```go
sum := 0
for i := 2; i <= n; i++ {
    if isPrime(i) {
        sum += i
    }
}
```

### 4. os.Args and Input Validation
Get the argument from `os.Args[1]`, validate count, validate numeric conversion, validate positive value.
```go
if len(os.Args) != 2 {
    fmt.Println(0)
    return
}
n, err := strconv.Atoi(os.Args[1])
if err != nil || n <= 0 {
    fmt.Println(0)
    return
}
```

## Review If Stuck

- [../93-findprevprime/skills.md](../93-findprevprime/skills.md) — covers writing an `isPrime` function using `i*i <= n`
- [../86-searchreplace/skills.md](../86-searchreplace/skills.md) — covers `os.Args` and argument count validation

## You're Ready When You Can...

- [ ] Use `strconv.Atoi` to parse a command-line argument to an integer
- [ ] Check and handle the error from `strconv.Atoi`
- [ ] Write an `isPrime` function using the `i*i <= n` optimization
- [ ] Sum all primes from 2 to `n` using a loop and a running total

## Next Steps

- [103-canjump](../103-canjump/README.md)
