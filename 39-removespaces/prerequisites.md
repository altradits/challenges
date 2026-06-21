# Prerequisites for 39-removespaces

## Before You Start

To solve this challenge you need to understand:

### 1. Building a Result String by Accumulation

From [32-toupper skills.md](../32-toupper/skills.md): start with an empty string and append characters inside a loop:

```go
result := ""
for _, c := range s {
    result += string(c)
}
return result
```

### 2. Conditional Appending

From [34-countalpha skills.md](../34-countalpha/skills.md): you already skipped non-matching characters (by not counting them). Here you skip them by not appending them:

```go
for _, c := range s {
    if someCondition(c) {
        result += string(c)   // only append when condition is true
    }
}
```

Characters that fail the condition are simply ignored.

### 3. Comparing a Rune to a Character Literal

Rune literals use single quotes:

```go
c == ' '    // true when c is a space (rune value 32)
c != ' '    // true when c is anything except a space
```

Double-quoted strings (`" "`) are a different type — do not mix them with rune comparisons.

### 4. The `for...range` Loop

From [28-stringlength skills.md](../28-stringlength/skills.md): the standard way to walk over every character:

```go
for _, c := range s {
    // c is a rune
}
```

## Review If Stuck

- [32-toupper skills.md](../32-toupper/skills.md) — building a result string inside a loop
- [34-countalpha skills.md](../34-countalpha/skills.md) — conditional logic inside a loop

## You're Ready When You Can...

- [ ] Build a new string by appending characters inside a `for...range` loop
- [ ] Write an `if` condition that is true for every character except a space
- [ ] Explain why `' '` (single quotes) and `" "` (double quotes) are different types in Go

## Next Steps

- [40-countrepeats](../40-countrepeats/README.md)
