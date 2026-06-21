# Prerequisites for 91-findchar

## Before You Start

To solve this challenge you need to understand:

### 1. `for...range` with Index and Character

From [76-stringlength skills.md](../76-stringlength/skills.md) you used `for _, c := range s` to discard the index. Now you need both:

```go
for i, c := range s {
    fmt.Println(i, string(c))
}
// "Hello" prints:
// 0 H
// 1 e
// 2 l
// 3 l
// 4 o
```

`i` is the byte position of each character.

### 2. Early Return

From [83-checknumber skills.md](../83-checknumber/skills.md): return inside a loop to stop immediately when a condition is met. Here you return the index rather than `true`:

```go
for i, c := range s {
    if c == target {
        return i   // found — stop here
    }
}
return -1   // checked all — not found
```

### 3. Functions with Two Parameters

You have written functions with one parameter so far. This challenge has two:

```go
func FindChar(s string, c rune) int {
    // s and c are both available here
}
```

Call it by passing two arguments:

```go
FindChar("Hello", 'l')
```

### 4. Rune Literals

A rune literal uses single quotes: `'l'`, `'H'`, `'a'`. It is a value of type `rune`. Compare it directly to the `c` you receive from `for...range`.

### 5. The `-1` Sentinel

Returning `-1` is the convention for "not found" when the return type is `int`. Valid indices are always `>= 0`, so `-1` is unambiguous.

## Review If Stuck

- [76-stringlength skills.md](../76-stringlength/skills.md) — `for...range` basics
- [83-checknumber skills.md](../83-checknumber/skills.md) — early return pattern

## You're Ready When You Can...

- [ ] Write a `for...range` loop that uses both the index `i` and the character `c`
- [ ] Return an integer index from inside a loop (early return)
- [ ] Explain why `-1` is used instead of `0` for "not found"
- [ ] Write a function that takes two parameters

## Next Steps

- [91-findchar skills.md](skills.md) — teaches returning an index from a search loop
- [92-countchar README](../92-countchar/README.md) — next challenge
