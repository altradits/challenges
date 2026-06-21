# Prerequisites for 44-thirdchar

## Before You Start

### 1. The Modulo Operator `%`

`a % b` gives the remainder when `a` is divided by `b`. It cycles through values `0` to `b-1`:

```go
for i := 0; i < 9; i++ {
    fmt.Printf("i=%d  i%%3=%d\n", i, i%3)
}
// i=0  i%3=0
// i=1  i%3=1
// i=2  i%3=2
// i=3  i%3=0  ← repeats
// i=4  i%3=1
// i=5  i%3=2
// i=6  i%3=0  ← repeats again
```

The pattern `== 0` catches every 3rd index.

### 2. 0-Based vs 1-Based Indexing

Go `for...range` gives 0-based indices. But "every 3rd character" in human terms means positions 1, 2, **3**, 4, 5, **6**, ... which in 0-based indices are 0, 1, **2**, 3, 4, **5**, ...

To convert: `(i+1) % 3 == 0` selects 0-based indices 2, 5, 8 = 1-based positions 3, 6, 9.

### 3. Iterating with Both Index and Value

```go
for i, c := range "abcdef" {
    fmt.Printf("i=%d c=%c\n", i, c)
}
// i=0 c=a
// i=1 c=b
// i=2 c=c
// i=3 c=d
// i=4 c=e
// i=5 c=f
```

### 4. Conditional Character Accumulation

Add a character to a result string only when a condition is met:

```go
result := ""
for i, c := range "abcdef" {
    if i%2 == 1 {  // every 2nd character (b, d, f)
        result += string(c)
    }
}
fmt.Println(result)  // "bdf"
```

### 5. Appending a Newline

Add `"\n"` at the end of the result:

```go
result := "369"
return result + "\n"  // "369\n"
```

## Review If Stuck

- [37-countwords](../37-countwords/skills.md) — covers `for...range` iteration with index

## You're Ready When You Can...

- [ ] Use `%` to check if a number is divisible by 3
- [ ] Explain the difference between 0-based index 2 and 1-based position 3
- [ ] Iterate a string with `for i, c := range s` and conditionally collect characters
- [ ] Append `"\n"` to a string result

## Next Steps

- [45-zipstring](../45-zipstring/README.md)
