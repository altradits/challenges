# Skills for 04-onlyf

## What You'll Learn

**Previous:** [03-onlyb](../03-onlyb/skills.md)

**Challenge:** Onlyf

## New Concepts Explained

### 1. String manipulation and processing

In Go, strings are immutable sequences of bytes encoded in UTF-8. You can iterate over them using `for...range` which gives you runes (Unicode code points) rather than bytes.

```go
for _, char := range myString {
    // char is a rune (int32)
}
```

To build new strings, concatenate using `+` or use `strings.Builder` for efficiency.

**Next:** [05-onlyz](../05-onlyz/skills.md) - 05 Onlyz
