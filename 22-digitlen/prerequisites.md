# Prerequisites for digitlen

## Before You Start

To solve this challenge you need to understand:

### 1. Integer Division
In Go, dividing two integers with `/` drops the decimal part — the result is always an integer.
```go
fmt.Println(100 / 10) // 10
fmt.Println(7 / 2)    // 3
fmt.Println(1 / 10)   // 0  ← this is what stops the loop
```

### 2. While-Style For Loop
Go uses `for condition { }` as its while loop. It runs until the condition becomes false.
```go
n := 100
for n > 0 {
    n /= 10  // integer division
}
// n is now 0
```

### 3. Counting Loop Iterations
Declare a counter before the loop and increment it each iteration.
```go
count := 0
n := 100
for n > 0 {
    n /= 10
    count++
}
// count == 3
```

### 4. Negating an Integer
To convert a negative number to positive, check the sign and negate it.
```go
n := -42
if n < 0 {
    n = -n  // now n == 42
}
```

### 5. Early Return for Invalid Input
Return a sentinel value immediately when an argument is out of range.
```go
if base < 2 || base > 36 {
    return -1
}
```

## Review If Stuck

- [../21-countrepeats/skills.md](../21-countrepeats/skills.md) — covers loop counters and using a for loop with a terminating condition

## You're Ready When You Can...

- [ ] Write a `for n > 0` loop that divides `n` by `base` each iteration
- [ ] Count how many iterations the loop runs
- [ ] Negate a negative integer before entering the loop
- [ ] Return `-1` when the base is outside the valid range 2–36

## Next Steps

- [Next challenge](../23-firstword/README.md)
