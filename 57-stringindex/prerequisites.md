# Prerequisites for 57-stringindex

## Before You Start

### 1. String slicing `s[i:j]`

Extracts bytes from index `i` up to (not including) `j`.

```go
s := "hello"
s[0:2]  // "he"
s[3:5]  // "lo"
s[1:]   // "ello"
```

Review: [26-retainfirsthalf](../26-retainfirsthalf/skills.md)

### 2. `len()` on strings

```go
len("hello")  // 5
len("")        // 0
```

### 3. A for loop with an index variable

You need a numeric index to slide through positions:

```go
for i := 0; i <= limit; i++ {
    // use i
}
```

### 4. String comparison with `==`

In Go, two strings are equal if they contain the same bytes:

```go
"hello" == "hello"  // true
"hello" == "Hello"  // false
```

## You're Ready When You Can...

- [ ] Slice a string to extract a window of N characters starting at position i
- [ ] Write a loop that iterates i from 0 to a calculated upper bound
- [ ] Return early from a function when a match is found

## Next Steps

- [58-stringcount](../58-stringcount/README.md)
