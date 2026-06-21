# Prerequisites for countrepeats

## Before You Start

To solve this challenge you need to understand:

### 1. String Indexing
Strings in Go can be accessed by position. `s[i]` gives the byte at index `i`. Indexing starts at 0.
```go
s := "hello"
fmt.Println(s[0]) // 104 — byte value of 'h'
fmt.Println(s[1]) // 101 — byte value of 'e'
```

### 2. Comparing Adjacent Characters
You can compare bytes at consecutive positions using `==`.
```go
s := "aabb"
// Check if position 1 equals position 0:
if s[1] == s[0] { // 'a' == 'a' → true
    fmt.Println("same")
}
```

### 3. Index-Based For Loop
Use a standard index loop starting at 1 to compare each character with the previous one.
```go
s := "aabb"
for i := 1; i < len(s); i++ {
    if s[i] == s[i-1] {
        fmt.Println("repeat at index", i)
    }
}
```

### 4. Boolean Flags
A boolean variable tracks state across iterations — for example, whether you are currently inside a consecutive group.
```go
inRepeat := false
// set true when a group starts, false when characters change
```

### 5. Empty String Guard
Always check length before indexing to avoid an index-out-of-range panic.
```go
if len(s) < 2 {
    return 0
}
```

## Review If Stuck

- [../20-cameltosnakecase/skills.md](../20-cameltosnakecase/skills.md) — covers string indexing and character-by-character iteration

## You're Ready When You Can...

- [ ] Write a for loop starting at index 1 that compares `s[i]` to `s[i-1]`
- [ ] Use a boolean flag to track whether you are inside a consecutive run
- [ ] Increment a counter only at the *start* of a new repeat group, not on every matching pair
- [ ] Return 0 safely for an empty or single-character string

## Next Steps

- [Next challenge](../22-digitlen/README.md)
