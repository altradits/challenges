# Prerequisites for union

## Before You Start

To solve this challenge you need to understand:

### 1. Map as a "Seen" Set
Use a `map[rune]bool` to record which characters have already been output. Looking up a missing key returns `false` (no panic).

```go
seen := make(map[rune]bool)
seen['a'] = true
fmt.Println(seen['a'])  // true
fmt.Println(seen['z'])  // false — key never set, zero value returned
```

### 2. `for range` Over a String
Iterate over each character as a rune.

```go
for _, c := range "hello" {
    fmt.Printf("%c\n", c)
}
```

### 3. Conditional Append — Only When Not Seen
Only add to the result when the character is new.

```go
if !seen[c] {
    result += string(c)
    seen[c] = true
}
```

### 4. Two Sequential Loops with One Shared Map
To process both strings without duplicates, loop over `s1` completely, then loop over `s2`, using the same `seen` map throughout.

```go
for _, c := range s1 { /* ... */ }
for _, c := range s2 { /* ... */ }
```

### 5. Printing a Bare Newline
When the argument count is wrong, print only a newline (not the content).

```go
fmt.Println()  // just a newline
```

## Review If Stuck
- [55-inter](../55-inter/skills.md) — covers using a map as a set for character deduplication
- [57-saveandmiss](../57-saveandmiss/skills.md) — covers conditional string building

## You're Ready When You Can...
- [ ] Create a `map[rune]bool` and use it to track seen characters
- [ ] Look up a map key and rely on the `false` zero value for missing keys
- [ ] Walk two strings sequentially with `for range`
- [ ] Append to a result string only when a condition is met
- [ ] Print a bare newline with `fmt.Println()`

## Next Steps
- [59-wdmatch](../59-wdmatch/README.md)
