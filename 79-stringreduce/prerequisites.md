# Prerequisites for 79-stringreduce

## Before You Start

### 1. Higher-order functions with two-argument function types

You've seen `func(rune) rune` in Map. Now the combining function takes two runes:

```go
func(rune, rune) rune  // takes two runes, returns one
```

Review: [77-stringmap](../77-stringmap/skills.md)

### 2. Converting a string to `[]rune`

To access individual runes by index, convert first:

```go
runes := []rune("hello")
runes[0]  // 'h'
runes[1:]  // ['e','l','l','o']  — all but first
```

### 3. The accumulator pattern

Keeping a running total across loop iterations:

```go
sum := 0
for _, n := range numbers {
    sum += n
}
```

The same concept applies here, but the accumulator is a `rune` and the combining function is passed in.

### 4. Slice indexing and sub-slicing

```go
s := []rune{'a', 'b', 'c'}
s[0]    // 'a'
s[1:]   // ['b', 'c']
```

## You're Ready When You Can...

- [ ] Write a function that accepts `func(rune, rune) rune` as a parameter
- [ ] Initialize an accumulator to the first element of a slice
- [ ] Apply a function in a loop starting from the second element

## Next Steps

- [80-stringformat](../80-stringformat/README.md)
