# Skills for 11-forloops

## What You'll Learn

**Previous:** [10-findsubstring](../10-findsubstring/skills.md) | **Next:** [12-printif](../12-printif/skills.md)

**Challenge:** Write a function `CountDown(n int)` that prints all integers from `n` down to `0`, each on its own line.

## Core Concept: All Four For-Loop Patterns in Go

### What Is a For Loop?

A for loop runs a block of code repeatedly. Go has only one looping keyword — `for` — but it covers every loop pattern: index-based, while-style, range-based, and infinite.

### 1. Index-Based For Loop

This is the classic C-style loop. It has three parts: initialization, condition, and post.

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
// prints: 0 1 2 3 4
```

| Part | Example | Runs |
|------|---------|------|
| Initialization | `i := 0` | Once, before the loop starts |
| Condition | `i < 5` | Before each iteration; loop stops when false |
| Post | `i++` | After each iteration |

Use this pattern when you need the index or need to count a fixed number of times.

Counting down:

```go
for i := 5; i >= 0; i-- {
    fmt.Println(i)
}
// prints: 5 4 3 2 1 0
```

### 2. While-Style For Loop

Go has no `while` keyword. Use `for condition { }` instead:

```go
n := 1
for n < 100 {
    n *= 2
}
fmt.Println(n)  // 128
```

The loop continues as long as the condition is true. This is identical to `while (condition)` in other languages.

### 3. Infinite Loop with Break

`for { }` with no condition runs forever. Use `break` to exit when a condition is met:

```go
for {
    line := readLine()
    if line == "quit" {
        break
    }
    process(line)
}
```

`break` exits the innermost loop immediately. Use it when the exit condition is checked in the middle of the loop body.

### 4. Range Loop on a Slice

`for i, v := range slice` gives you the index and value of each element:

```go
fruits := []string{"apple", "banana", "cherry"}
for i, v := range fruits {
    fmt.Println(i, v)
}
// 0 apple
// 1 banana
// 2 cherry
```

Use `_` to discard the index if you only need the value:

```go
for _, v := range fruits {
    fmt.Println(v)
}
```

### 5. Range Loop on a String

`for i, r := range s` gives you the **byte index** and **rune value** of each character:

```go
for i, r := range "Hi!" {
    fmt.Printf("index %d: %c\n", i, r)
}
// index 0: H
// index 1: i
// index 2: !
```

The rune type (`int32`) holds a Unicode code point. For ASCII strings the byte index equals the character position. For multi-byte characters (emoji, accented letters) the byte index can jump by more than 1.

### 6. Range Loop on a Map

`for k, v := range m` iterates over key-value pairs. Map iteration order is not guaranteed:

```go
scores := map[string]int{"Alice": 90, "Bob": 85}
for k, v := range scores {
    fmt.Println(k, v)
}
```

### 7. `continue` and `break`

- `continue` skips the rest of the current iteration and moves to the next one.
- `break` exits the loop entirely.

```go
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue   // skip even numbers
    }
    if i > 7 {
        break      // stop at 7
    }
    fmt.Println(i)
}
// prints: 1 3 5 7
```

### 8. Nested For Loops

You can put a loop inside another loop. The inner loop runs completely for each iteration of the outer loop:

```go
for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        fmt.Printf("%d*%d=%d  ", i, j, i*j)
    }
    fmt.Println()
}
```

`break` and `continue` in a nested loop only affect the **innermost** loop they appear in.

### Pattern Comparison Table

| Pattern | Syntax | Use when |
|---------|--------|----------|
| Index | `for i := 0; i < n; i++` | Need index, fixed count |
| While | `for condition { }` | Unknown iterations |
| Range | `for _, v := range collection` | Iterating collections |
| Infinite | `for { ... }` | Loop until break |

### Diagram: Index-Based Loop Execution

```
for i := 5; i >= 0; i--

  i=5: 5 >= 0? YES -> print 5 -> i-- -> i=4
  i=4: 4 >= 0? YES -> print 4 -> i-- -> i=3
  i=3: 3 >= 0? YES -> print 3 -> i-- -> i=2
  i=2: 2 >= 0? YES -> print 2 -> i-- -> i=1
  i=1: 1 >= 0? YES -> print 1 -> i-- -> i=0
  i=0: 0 >= 0? YES -> print 0 -> i-- -> i=-1
  i=-1: -1 >= 0? NO -> loop exits
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `for i := n; i > 0; i--` | Skips printing `0` | Use `i >= 0` as the condition |
| Infinite loop: condition never becomes false | Program hangs forever | Make sure the post step (`i--`) moves toward the exit condition |
| `for i := 0; i < n; i++` for countdown | Counts up, not down | Use `i := n` and `i--` |
| Using `break` expecting it to exit an outer loop | Only exits the innermost loop | Use a labeled break or refactor with a boolean flag |

## Solving This Challenge

### Algorithm

1. Use `for i := n; i >= 0; i--`
2. Inside the loop, call `fmt.Println(i)` to print the current value

```go
func CountDown(n int) {
    for i := n; i >= 0; i-- {
        fmt.Println(i)
    }
}
```

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [12-printif](../12-printif/README.md)
