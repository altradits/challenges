# Skills for 04-onlyf

## What You'll Learn

**Previous:** [03-onlyb](../03-onlyb/skills.md)

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** Onlyf

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

**Next:** [05-onlyz](../05-onlyz/skills.md) - Onlyz
