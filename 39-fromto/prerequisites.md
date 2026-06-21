# Prerequisites for fromto

## Before You Start

To solve this challenge you need to understand:

### 1. fmt.Sprintf with %02d
Format an integer as a two-digit string with a leading zero when necessary.
```go
import "fmt"

fmt.Println(fmt.Sprintf("%02d", 5))  // "05"
fmt.Println(fmt.Sprintf("%02d", 10)) // "10"
fmt.Println(fmt.Sprintf("%02d", 0))  // "00"
```

### 2. Counting Up and Down in a Loop
Use `i++` to count up, `i--` to count down. Both have the same `for` syntax.
```go
// Counting up: 1, 2, 3, 4, 5
for i := 1; i <= 5; i++ { }

// Counting down: 5, 4, 3, 2, 1
for i := 5; i >= 1; i-- { }
```

### 3. Building a String List With append and strings.Join
Collect formatted numbers into a slice, then join with a separator.
```go
import "strings"

parts := []string{}
parts = append(parts, "01")
parts = append(parts, "02")
parts = append(parts, "03")
fmt.Println(strings.Join(parts, ", ")) // "01, 02, 03"
```

### 4. Input Validation With Early Return
Check both boundary conditions before doing any work.
```go
if from < 0 || from > 99 || to < 0 || to > 99 {
    return "Invalid\n"
}
```

### 5. Conditional Direction
Use `if from <= to` to decide whether to count up or down.
```go
if from <= to {
    // count up
} else {
    // count down
}
```

## Review If Stuck

- [../38-findprevprime/skills.md](../38-findprevprime/skills.md) — covers downward iteration (`i--`)
- [../31-splitjoin/skills.md](../31-splitjoin/skills.md) — covers `strings.Join` and building slices with `append`

## You're Ready When You Can...

- [ ] Format a number with leading zero using `fmt.Sprintf("%02d", n)`
- [ ] Write both an upward and a downward counting loop
- [ ] Collect formatted strings into a slice with `append`
- [ ] Join the slice with `strings.Join(parts, ", ")`
- [ ] Validate input and return `"Invalid\n"` for out-of-range values

## Next Steps

- [Next challenge](../40-iscapitalized/README.md)
