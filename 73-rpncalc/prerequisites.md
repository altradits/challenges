# Prerequisites for 73-rpncalc

## Before You Start

### 1. Stack (slice as LIFO) from 72-brackets

You built a stack for bracket matching. The same pattern applies here — but the stack holds `int` values, not runes.

```go
stack := []int{}
stack = append(stack, 42)           // push
top := stack[len(stack)-1]          // peek
stack = stack[:len(stack)-1]        // pop
```

Review: [72-brackets](../72-brackets/skills.md)

### 2. `strings.Fields` — split on any whitespace

```go
tokens := strings.Fields("1   3  *  2 -")
// tokens = ["1", "3", "*", "2", "-"]
```

Review: [19-wordcount](../19-wordcount/skills.md)

### 3. `strconv.Atoi` — string to int with error

```go
n, err := strconv.Atoi("42")    // n=42, err=nil
n, err := strconv.Atoi("xyz")   // n=0, err != nil
if err != nil { /* not a number */ }
```

Review: [22-digitlen](../22-digitlen/skills.md)

### 4. `os.Args` — command-line arguments

```go
if len(os.Args) != 2 {
    fmt.Println("Error")
    return
}
expr := os.Args[1]
```

### 5. `switch` statement

```go
switch token {
case "+": result = a + b
case "-": result = a - b
default:  fmt.Println("Error")
}
```

## You're Ready When You Can...

- [ ] Push and pop from a `[]int` slice
- [ ] Tokenize a string with `strings.Fields`
- [ ] Convert a string to int and detect conversion failure
- [ ] Use a switch statement for multiple cases

## Next Steps

- [74-brainfuck](../74-brainfuck/README.md)
