# Prerequisites for 83-checknumber

## Before You Start

To solve this challenge you need to understand:

### 1. Functions That Return `bool`

From [79-isempty skills.md](../79-isempty/skills.md): a boolean function answers a yes/no question:

```go
func CheckNumber(s string) bool {
    return true   // or false
}
```

### 2. The `for...range` Loop

From [76-stringlength skills.md](../76-stringlength/skills.md): iterate over every character:

```go
for _, c := range s {
    // c is each rune in turn
}
```

### 3. Range Checks on Runes

From [82-countalpha skills.md](../82-countalpha/skills.md): you can check whether a rune falls in a range using `>=` and `<=`. For digits:

```go
if c >= '0' && c <= '9' {
    // c is a digit
}
```

The digit characters `'0'`–`'9'` have ASCII values 48–57 (consecutive), so this range check works.

### 4. Returning Inside a Loop (Early Return)

You already returned inside a loop in [77-firstchar](../77-firstchar/skills.md):

```go
for _, c := range s {
    return string(c)   // returns immediately on first iteration
}
```

For this challenge, you will `return true` the moment you find a digit — before the loop finishes.

### 5. The Fallback Return

After a loop that may exit early, Go requires a return statement outside the loop too:

```go
for _, c := range s {
    if condition {
        return true   // early exit
    }
}
return false   // reached only if the loop ran to completion
```

## Review If Stuck

- [79-isempty skills.md](../79-isempty/skills.md) — `bool` return type
- [82-countalpha skills.md](../82-countalpha/skills.md) — range checks with `&&`

## You're Ready When You Can...

- [ ] Write a function that returns `bool`
- [ ] Check whether a rune is in the range `'0'` to `'9'`
- [ ] Return `true` from inside a loop to exit early
- [ ] Place `return false` after the loop as a fallback

## Next Steps

- [83-checknumber skills.md](skills.md) — teaches early return and digit detection
- [84-countvowels README](../84-countvowels/README.md) — next challenge
