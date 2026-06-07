# Skills for 02-onlya

## What You'll Learn

**Previous:** [01-only1](../01-only1/skills.md)

**Challenge:** Onlya

## New Concepts Explained

### 1. String manipulation and processing

In Go, strings are immutable sequences of bytes encoded in UTF-8. You can iterate over them using `for...range` which gives you runes (Unicode code points) rather than bytes.

```go
for _, char := range myString {
    // char is a rune (int32)
}
```

To build new strings, concatenate using `+` or use `strings.Builder` for efficiency.

**Next:** [03-onlyb](../03-onlyb/skills.md) - 03 Onlyb
