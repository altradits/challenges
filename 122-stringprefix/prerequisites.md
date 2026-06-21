# Prerequisites for 122-stringprefix

## Before You Start

### 1. String slicing `s[:n]`

Get the first `n` bytes of a string:

```go
s := "hello"
s[:3]  // "hel"
s[:0]  // ""
```

Review: [89-retainfirsthalf](../89-retainfirsthalf/skills.md)

### 2. `len()` comparisons

```go
len("hello") < len("hello world")  // true
len("hello") < len("hi")           // false
```

### 3. String equality with `==`

```go
"hello" == "hello"  // true
"hello" == "Hello"  // false (case-sensitive)
```

### 4. Boolean return from a function

```go
func check(s string) bool {
    if condition {
        return true
    }
    return false
}
```

## You're Ready When You Can...

- [ ] Slice a string to get its first N characters
- [ ] Compare the lengths of two strings
- [ ] Return a bool from a function based on a condition

## Next Steps

- [123-stringsuffix](../123-stringsuffix/README.md)
