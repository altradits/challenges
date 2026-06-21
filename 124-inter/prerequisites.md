# Prerequisites for inter

## Before You Start

To solve this challenge you need to understand:

### 1. Maps in Go — Using a Map as a Set
A map with `bool` values acts as a set. Checking map membership returns `false` (the zero value) for keys that were never inserted — no need to test for existence separately.

```go
seen := make(map[rune]bool)
seen['a'] = true
seen['b'] = true

fmt.Println(seen['a'])  // true
fmt.Println(seen['z'])  // false (zero value, not panic)
```

### 2. `for range` Over a String
Iterates character by character as runes (Unicode-safe).

```go
for _, c := range "hello" {
    fmt.Printf("%c\n", c)  // h, e, l, l, o
}
```

### 3. Building a String Incrementally
Strings in Go are immutable. Build a result by concatenating with `+` or using `strings.Builder`.

```go
result := ""
result += string('x')  // result is now "x"
```

### 4. Two-Pass Algorithm
This problem requires two passes:
1. First pass over `s2` to record what characters exist in it.
2. Second pass over `s1` to output characters present in both, without duplicates.

```go
// Pass 1: record s2
inS2 := make(map[rune]bool)
for _, c := range s2 {
    inS2[c] = true
}

// Pass 2: filter s1
printed := make(map[rune]bool)
for _, c := range s1 {
    if inS2[c] && !printed[c] {
        // output c
        printed[c] = true
    }
}
```

### 5. `os.Args` Validation
Exactly 2 arguments are required (plus the program name at index 0).

```go
if len(os.Args) != 3 {
    fmt.Println()
    return
}
```

## Review If Stuck
- [123-hiddenp](../123-hiddenp/skills.md) — covers iterating strings and `os.Args` validation
- Prior map challenges — covers `make(map[...])`, insertion, and lookup

## You're Ready When You Can...
- [ ] Create a `map[rune]bool` with `make` and insert keys
- [ ] Look up a map key and get `false` (not a panic) when the key is absent
- [ ] Use `for _, c := range s` to iterate a string as runes
- [ ] Build a result string by appending with `+` or `strings.Builder`
- [ ] Use two separate maps to track "in s2" and "already printed"

## Next Steps

- [125-reversestrcap](../125-reversestrcap/README.md)
