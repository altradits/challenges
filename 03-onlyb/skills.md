# Skills for 03-onlyb

## What You'll Learn

**Previous:** [02-onlya](../02-onlya/skills.md)

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** Onlyb

## New Concepts Explained

### 1. String iteration and character access

In Go, strings are immutable sequences of bytes encoded in UTF-8. You can iterate over them using `for...range` which gives you runes (Unicode code points) rather than bytes.

```go
for _, char := range myString {
    // char is a rune (int32)
}
```

To access individual characters, you can also use indexing, but remember that `s[i]` returns a byte, not a rune. For UTF-8 safety, use `for...range`.

## The Challenge

See [README.md](README.md) for the full challenge description, expected function, and test cases.

**Next:** [04-onlyf](../04-onlyf/skills.md) - Onlyf
