# Prerequisites for arrays

## Before You Start

To solve this challenge you need to understand:

### 1. Slices (and how arrays differ from them)
Before learning arrays, you should know what a slice is — because arrays and slices look similar but behave very differently. Arrays have a fixed size baked into their type and are copied on assignment; slices are dynamic and reference-based.
```go
// Slice — dynamic, reference type
s := []int{1, 2, 3}

// Array — fixed size, value type
arr := [3]int{1, 2, 3}
```

See [../09-slicesintro/skills.md](../09-slicesintro/skills.md) for a full introduction.

### 2. For Loops
You need a `for` loop (or `for range`) to iterate over all elements of an array and accumulate a sum.
```go
arr := [5]int{1, 2, 3, 4, 5}
sum := 0
for _, v := range arr {
    sum += v
}
fmt.Println(sum)  // 15
```

See [../07-forloops/skills.md](../07-forloops/skills.md) for a full introduction.

### 3. `fmt.Printf` with Format Verbs
The next challenge (`97-printmemory`) uses `fmt.Printf` to format byte values in hexadecimal and as characters. Get comfortable with the basic format verbs now.
```go
fmt.Printf("%d\n", 65)    // "65"   — decimal integer
fmt.Printf("%c\n", 65)    // "A"    — character (rune value)
fmt.Printf("%02x\n", 65)  // "41"   — hex, minimum 2 digits, zero-padded
```

## You're Ready When You Can...

- [ ] Declare an array with a literal: `arr := [5]int{1, 2, 3, 4, 5}`
- [ ] Write a function that takes an array parameter: `func SumArray(arr [5]int) int`
- [ ] Iterate over an array with `for _, v := range arr`
- [ ] Explain the difference between arrays (value type) and slices (reference type)
- [ ] Understand that `[5]int` and `[6]int` are completely different types in Go

## Next Steps

- [12-recursion](../12-recursion/README.md)
