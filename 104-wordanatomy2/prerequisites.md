# Prerequisites for wordanatomy2

## Before You Start

To solve this challenge you need to understand:

### 1. Functions That Take Slice Parameters
A function can accept a `[]string` slice as a parameter and iterate over it.
```go
func checkPrefixes(word string, pref []string) string {
    for _, p := range pref {
        // check each prefix
    }
    return ""
}
```

### 2. First-Match with break
Stop searching as soon as you find the first match by using `break`.
```go
found := ""
for _, p := range pref {
    if len(word) >= len(p) && word[:len(p)] == p {
        found = p
        break  // stop here — first match wins
    }
}
```

### 3. fmt.Sprintf for Formatted Strings
Build a formatted string using `fmt.Sprintf` and format verbs. `%q` wraps the value in double quotes.
```go
result := fmt.Sprintf("prefix: %q, suffix: %q", "un", "able")
// result == `prefix: "un", suffix: "able"`
```

### 4. String Slicing for Prefix/Suffix Matching
```go
word := "understandable"
prefix := "un"
// Check prefix:
if len(word) >= len(prefix) && word[:len(prefix)] == prefix {
    fmt.Println("match!")
}
// Check suffix:
suffix := "able"
if len(word) >= len(suffix) && word[len(word)-len(suffix):] == suffix {
    fmt.Println("suffix match!")
}
```

### 5. Empty String as Default Value
If no match is found, the default value `""` (empty string) should be used in the output.
```go
found := ""  // default: no match
// ... search loop ...
// if nothing matched, found is still ""
```

## Review If Stuck

- [../103-wordanatomy/skills.md](../103-wordanatomy/skills.md) — covers prefix/suffix detection and string slicing; this challenge adds first-match logic and formatted output

## You're Ready When You Can...

- [ ] Accept a `[]string` slice as a function parameter and iterate over it
- [ ] Use `break` to stop a loop on the first match
- [ ] Use `fmt.Sprintf` with `%q` to produce quoted strings in output
- [ ] Return the empty string `""` as a default when no match is found

## Next Steps

- [105-clean-the-list](../105-clean-the-list/README.md)
