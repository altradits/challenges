# Skills for 03-onlyb

## What You'll Learn

**Previous:** [02-onlya](../02-onlya/skills.md)

**Challenge:** Onlyb

## New Concepts Explained

### 1. String manipulation and processing

In Go, strings are immutable sequences of bytes encoded in UTF-8. You can iterate over them using `for...range` which gives you runes (Unicode code points) rather than bytes.

```go
for _, char := range myString {
    // char is a rune (int32)
}
```

To build new strings, concatenate using `+` or use `strings.Builder` for efficiency.

**Next:** [04-onlyf](../04-onlyf/skills.md) - 04 Onlyf
