# Prerequisites for 63-stringfilter

## Before You Start

### 1. Higher-order functions — functions as parameters

You learned this in Map: a function can accept another function as a parameter.

```go
func doSomething(s string, f func(rune) bool) string {
    // f is a function we call inside
}
```

Review: [62-stringmap](../62-stringmap/skills.md)

### 2. `strings.Builder` for building output

```go
var b strings.Builder
b.WriteRune('a')
result := b.String()
```

Review: [50-stringbuilder](../50-stringbuilder/skills.md)

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

- [64-stringreduce](../64-stringreduce/README.md)
