# Prerequisites for wordanatomy

## Before You Start

To solve this challenge you need to understand:

### 1. String Slicing for Prefix/Suffix Detection
`s[:n]` gives the first `n` bytes; `s[len(s)-n:]` gives the last `n` bytes.
```go
word := "unhappy"
fmt.Println(word[:2])           // "un"  — first 2 chars
fmt.Println(word[len(word)-2:]) // "py"  — last 2 chars
fmt.Println(word[2:])           // "happy" — without prefix
```

### 2. Comparing Slices with ==
String slices compare equal if content matches:
```go
word := "unhappy"
if word[:2] == "un" {
    fmt.Println("has prefix un")
}
```

### 3. Multiple Return Values
A Go function can return more than one value. Declare return types in parentheses.
```go
func analyze(s string) (string, string) {
    return "prefix", "suffix"
}
p, s := analyze("hello")
```

### 4. Iterating Over a Slice of Strings
Loop through a list of candidates and check each one.
```go
prefixes := []string{"un", "re", "pre"}
for _, p := range prefixes {
    if len(word) >= len(p) && word[:len(p)] == p {
        fmt.Println("found prefix:", p)
    }
}
```

### 5. Tracking the Longest Match
Keep a variable holding the best candidate found so far, update it when a longer match is found.
```go
found := ""
for _, p := range prefixes {
    if matches(word, p) && len(p) > len(found) {
        found = p
    }
}
```

## Review If Stuck

- [../100-replaceall/skills.md](../100-replaceall/skills.md) — covers string slicing and bounds checking
- [../99-longestword/skills.md](../99-longestword/skills.md) — covers the running-maximum pattern for tracking the longest match

## You're Ready When You Can...

- [ ] Slice a string with `word[:n]` and `word[len(word)-n:]`
- [ ] Compare a string slice to a literal using `==`
- [ ] Write a function that returns multiple values
- [ ] Iterate over a `[]string` of candidates and track the longest match

## Next Steps

- [104-wordanatomy2](../104-wordanatomy2/README.md)
