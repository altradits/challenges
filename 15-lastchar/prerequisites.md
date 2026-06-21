# Prerequisites for 15-lastchar

## Before You Start

To solve this challenge you need to understand:

### 1. Writing a Go Function That Returns a `string`

You need to write a function with this exact signature:

```go
func LastChar(s string) string {
    // your code here
}
```

If this looks unfamiliar, see [13-stringlength skills.md](../13-stringlength/skills.md).

### 2. Checking for an Empty String

Before doing anything with a string's contents, confirm it is not empty:

```go
if s == "" {
    return ""
}
```

Trying to access `s[0]` (or any index) on an empty string causes a runtime panic. This check prevents that.

### 3. Converting a `byte` to a `string`

`s[i]` gives you a `byte` (type `uint8`). To return it as human-readable text, convert it:

```go
b := s[0]            // b is a byte, e.g., 72
fmt.Println(b)       // 72 (a number)
fmt.Println(string(b))  // "H" (a character)
```

From [14-firstchar skills.md](../14-firstchar/skills.md).

### 4. String Indexing

You already used `s[0]` in the previous challenge. This challenge needs the **last** index instead. Review that any valid index is between `0` and `len(s)-1`:

```go
s := "Hello"
s[0]  // 'H' — first
s[4]  // 'o' — last (len=5, so last index is 4)
```

## Review If Stuck

- [14-firstchar skills.md](../14-firstchar/skills.md) — covers string indexing and `byte`-to-`string` conversion
- [13-stringlength skills.md](../13-stringlength/skills.md) — covers `for...range` as an alternative approach

## You're Ready When You Can...

- [ ] Write a Go function that takes a `string` and returns a `string`
- [ ] Guard against an empty string before accessing any index
- [ ] Explain why `s[len(s)]` panics but `s[len(s)-1]` is safe
- [ ] Convert a `byte` to a `string` using `string(b)`

## Next Steps

- [16-isempty](../16-isempty/README.md)
- [16-isempty README](../16-isempty/README.md) — next challenge
