# Prerequisites for 31-isempty

## Before You Start

To solve this challenge you need to understand:

### 1. Functions That Return `bool`

So far your functions have returned `int` or `string`. This one returns a **boolean** — a true/false value:

```go
func IsEmpty(s string) bool {
    return true   // or false
}
```

`bool` is a distinct type in Go. You cannot return `0`, `1`, or `""` — only `true` or `false`.

### 2. The `==` Comparison Operator

`==` checks equality and produces a `bool`:

```go
s := "Hello"
s == "Hello"   // true
s == ""        // false
s == "World"   // false
```

You can return the result of a comparison directly:

```go
return s == ""   // returns true if s is empty, false otherwise
```

### 3. The `for...range` Loop and Early Return

From [28-stringlength skills.md](../28-stringlength/skills.md): a `for...range` loop on an empty string runs zero times. You can exploit that to detect emptiness:

```go
for range s {
    return false   // we found at least one character
}
return true        // the loop never ran
```

### 4. `len(s)` Returns Zero for Empty Strings

From [30-lastchar skills.md](../30-lastchar/skills.md): `len(s)` counts bytes. For an empty string it returns `0`:

```go
len("")    // 0
len("Hi")  // 2
```

## Review If Stuck

- [28-stringlength skills.md](../28-stringlength/skills.md) — covers `for...range` iteration
- [30-lastchar skills.md](../30-lastchar/skills.md) — covers `len(s)`

## You're Ready When You Can...

- [ ] Write a Go function that returns `bool`
- [ ] Use `==` to compare a string to `""`
- [ ] Explain the difference between `""` (empty) and `" "` (one space)
- [ ] Describe what happens when `for...range` runs on an empty string

## Next Steps

- [32-toupper](../32-toupper/README.md)
