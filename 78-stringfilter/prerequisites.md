# Prerequisites for 78-stringfilter

## Before You Start

### 1. Higher-order functions — functions as parameters

You learned this in Map: a function can accept another function as a parameter.

```go
func doSomething(s string, f func(rune) bool) string {
    // f is a function we call inside
}
```

Review: [77-stringmap](../77-stringmap/skills.md)

### 2. `strings.Builder` for building output

```go
var b strings.Builder
b.WriteRune('a')
result := b.String()
```

Review: [65-stringbuilder](../65-stringbuilder/skills.md)

### 3. `for...range` on strings

```go
for _, r := range s {
    // r is a rune
}
```

### 4. Boolean predicate functions

A predicate is a function that returns `true` or `false`:

```go
unicode.IsDigit('5')   // true
unicode.IsLetter('5')  // false
unicode.IsUpper('A')   // true
```

## You're Ready When You Can...

- [ ] Write a function that takes `func(rune) bool` as a parameter
- [ ] Use `strings.Builder` to collect only some characters
- [ ] Apply a boolean condition inside a `for...range` loop

## Next Steps

- [79-stringreduce](../79-stringreduce/README.md)
