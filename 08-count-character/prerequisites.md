# Prerequisites for 08-count-character

## Before You Start

To solve this challenge you need to understand how to iterate over a string, how runes work, and how to keep a running counter.

### 1. `for...range` on a string produces runes
Iterating a string with `for...range` gives you each character as a `rune` (type `int32`).

```go
for _, r := range "hello" {
    // r is a rune: 'h', 'e', 'l', 'l', 'o'
}
```

### 2. Rune equality comparison
A rune is just an integer. You compare runes with `==`.

```go
var r rune = 'l'
if r == 'l' {
    // match
}
```

### 3. Counter pattern: initialize, increment, return
The standard way to count things in a loop:

```go
count := 0
for _, r := range str {
    if r == target {
        count++
    }
}
return count
```

### 4. Function signature with a `rune` parameter
The `rune` type is used for single characters.

```go
func CountChar(str string, c rune) int {
    // c is a single character (rune)
    // return the count as an int
}
```

### 5. Empty string / no match returns 0 automatically
If the string is empty, the loop never runs and `count` stays `0`. If the character never appears, no `count++` is ever called. You do not need special cases for these.

## Review If Stuck

- [06-checknumber](../06-checknumber/skills.md) — covers `for...range` on strings and rune comparisons

## You're Ready When You Can...

- [ ] Write a loop that visits every character in a string as a rune
- [ ] Compare a rune variable to a rune literal like `'l'`
- [ ] Maintain a counter variable, increment it with `count++`, and return it
- [ ] Write a function that takes a `string` and a `rune` as parameters

## Next Steps

After completing this exercise, move to:
- [10-findsubstring](../10-findsubstring/README.md)
