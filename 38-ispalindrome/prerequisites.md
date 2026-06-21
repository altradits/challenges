# Prerequisites for 38-ispalindrome

## Before You Start

This challenge combines several skills you have already learned. Make sure you are confident with all of them before starting.

### 1. Reversing a String with `[]rune`

From [37-reversestring skills.md](../37-reversestring/skills.md): convert to `[]rune`, swap with two pointers, convert back:

```go
runes := []rune(s)
for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
    runes[i], runes[j] = runes[j], runes[i]
}
result := string(runes)
```

### 2. Filtering Characters (Building a Cleaned String)

From [34-countalpha skills.md](../34-countalpha/skills.md) and [35-checknumber skills.md](../35-checknumber/skills.md): you can build a new string that keeps only characters matching a condition:

```go
filtered := ""
for _, c := range s {
    if isLetterOrDigit(c) {
        filtered += string(c)
    }
}
```

### 3. Case Normalisation

From [33-tolower skills.md](../33-tolower/skills.md): convert uppercase to lowercase using `c + 32`:

```go
if c >= 'A' && c <= 'Z' {
    c = c + 32   // now lowercase
}
```

You will need this to make the palindrome check case-insensitive.

### 4. Comparing Two Strings with `==`

```go
if s1 == s2 {
    // identical
}
```

### 5. Boolean Return Values

From [31-isempty skills.md](../31-isempty/skills.md): return `true` or `false`.

## Review If Stuck

- [37-reversestring skills.md](../37-reversestring/skills.md) — `[]rune` two-pointer reversal (critical prerequisite)
- [33-tolower skills.md](../33-tolower/skills.md) — case normalisation
- [34-countalpha skills.md](../34-countalpha/skills.md) — character range checks

## You're Ready When You Can...

- [ ] Reverse a string using `[]rune` and the two-pointer technique
- [ ] Build a cleaned string that contains only alphanumeric characters
- [ ] Convert uppercase letters to lowercase using ASCII math
- [ ] Compare two strings for equality with `==`

## Next Steps

- [39-removespaces](../39-removespaces/README.md)
