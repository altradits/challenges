# Prerequisites for splitjoin

## Before You Start

To solve this challenge you need to understand:

### 1. Slices and append
A slice is a dynamic list. `append` adds an element and returns the updated slice.
```go
parts := []string{}
parts = append(parts, "hello")
parts = append(parts, "world")
// parts == []string{"hello", "world"}
fmt.Println(len(parts)) // 2
```

### 2. String Slicing for Substring Matching
Extract a substring with `s[i:j]` and compare it with `==`.
```go
s := "a,b,c"
sep := ","
if s[1:2] == sep {
    fmt.Println("comma at index 1")
}
```

### 3. Index-Based For Loop
Use `for i := 0; i < len(s)` when you need to advance by variable amounts (skip past separators).
```go
i := 0
for i < len(s) {
    // do work
    i++ // or i += len(sep)
}
```

### 4. String Concatenation
Build strings by appending with `+=`.
```go
result := ""
result += "hello"
result += ", "
result += "world"
// result == "hello, world"
```

### 5. Joining Elements With a Separator
Add the separator *between* elements — not before the first or after the last.
```go
arr := []string{"a", "b", "c"}
result := arr[0]
for i := 1; i < len(arr); i++ {
    result += "," + arr[i]
}
// result == "a,b,c"
```

## Review If Stuck

- [../85-replaceall/skills.md](../85-replaceall/skills.md) — covers index-based scanning and substring comparison with `s[i:j]`

## You're Ready When You Can...

- [ ] Create a `[]string{}` slice and append to it with `append`
- [ ] Detect a separator in a string using `s[i:i+len(sep)] == sep`
- [ ] Write a scanning loop that advances by `len(sep)` on a match
- [ ] Join a slice with a separator by starting with `arr[0]` and looping from index 1

## Next Steps

- [88-wordanatomy](../88-wordanatomy/README.md)
