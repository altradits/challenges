# Skills for 05-onlyz

## What You'll Learn

**Previous:** [04-onlyf](../04-onlyf/skills.md)

**Challenge:** Onlyz

## New Concepts Explained

### 1. String manipulation and processing

In Go, strings are immutable sequences of bytes encoded in UTF-8. You can iterate over them using `for...range` which gives you runes (Unicode code points) rather than bytes.

```go
for _, char := range myString {
    // char is a rune (int32)
}
```

To build new strings, concatenate using `+` or use `strings.Builder` for efficiency.

**Next:** [06-checknumber](../06-checknumber/skills.md) - 06 Checknumber
