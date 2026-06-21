# Skills for 153-index-fund-tracker

## What You'll Learn

**Previous:** [152-financial-freedom-api](../152-financial-freedom-api/skills.md) | **Next:** [154-savings-runway-calculator](../154-savings-runway-calculator/skills.md)

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** Index Fund Tracker

## New Concepts Explained

### 1. String iteration and character access

In Go, strings are immutable sequences of bytes encoded in UTF-8. You can iterate over them using `for...range` which gives you runes (Unicode code points) rather than bytes.

```go
for _, char := range myString {
    // char is a rune (int32)
}
```

To access individual characters, you can also use indexing, but remember that `s[i]` returns a byte, not a rune. For UTF-8 safety, use `for...range`.

### 2. Go function definition and usage

Functions in Go are defined using the `func` keyword. They can take parameters and return values:

```go
func FunctionName(param1 type1, param2 type2) returnType {
    // function body
    return result
}
```

The `main()` function is special - it's where program execution begins.

### 3. Looping constructs (for, range)

Go has only one looping construct: the `for` loop. It can be used in several ways:

```go
// Traditional for loop
for i := 0; i < 10; i++ { }

// While-style loop
for condition { }

// Range loop (for collections)
for index, value := range collection { }
```

For strings, `for...range` iterates over runes, making it safe for UTF-8.

### 4. Error handling and validation

Go handles errors explicitly. Functions often return `(value, error)`:

```go
result, err := someFunction()
if err != nil {
    // handle error
    return
}
// use result
```

Always check errors - Go doesn't have exceptions!

## The Challenge

See [README.md](README.md) for the full challenge description, expected function, and test cases.

**Next:** [154-savings-runway-calculator](../154-savings-runway-calculator/README.md)
