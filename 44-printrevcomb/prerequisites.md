# Prerequisites for printrevcomb

## Before You Start

To solve this challenge you need to understand:

### 1. Nested For Loops
A loop inside a loop. The inner loop completes all its iterations for each single iteration of the outer loop.
```go
for a := 9; a >= 7; a-- {
    for b := a - 1; b >= 6; b-- {
        fmt.Printf("%d%d\n", a, b)
    }
}
// 98, 97, 96, 87, 86, 76
```

### 2. Strictly Decreasing Loop Variables
Start each inner loop one below the outer loop's current value to guarantee all three values are different and in decreasing order.
```go
for a := 9; a >= 2; a-- {
    for b := a - 1; b >= 1; b-- {
        for c := b - 1; c >= 0; c-- {
            // a > b > c, all different
        }
    }
}
```

### 3. fmt.Printf for Multi-Digit Output
Print multiple values without spaces using `%d%d%d`:
```go
fmt.Printf("%d%d%d", 9, 8, 7) // "987"
fmt.Printf("%d%d%d", 2, 1, 0) // "210"
```

### 4. Comma-Separator Pattern With a Boolean Flag
Use a `first` flag to avoid printing a leading comma before the first element.
```go
first := true
// inside loop:
if !first {
    fmt.Print(", ")
}
fmt.Print(value)
first = false
```

### 5. fmt.Print vs fmt.Println
- `fmt.Print` does NOT add a newline.
- `fmt.Println` adds a newline.
- Use `fmt.Print` for values inside the loop; `fmt.Println()` once at the end.
```go
fmt.Print("987")
fmt.Print(", ")
fmt.Print("986")
fmt.Println()   // adds the final newline
```

## Review If Stuck

- [../39-fromto/skills.md](../39-fromto/skills.md) — covers building a list with separators between elements

## You're Ready When You Can...

- [ ] Write three nested `for` loops with strictly decreasing variables
- [ ] Use a `first` boolean flag to control separator printing
- [ ] Print combinations using `fmt.Printf("%d%d%d", a, b, c)`
- [ ] Print the final newline after all combinations with `fmt.Println()`

## Next Steps

- [Next challenge](../45-thirdtimeisacharm/README.md)
