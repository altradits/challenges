# Skills for 30-searchreplace

## What You'll Learn

**Previous:** [29-replaceall](../29-replaceall/skills.md)

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** Searchreplace

## New Concepts Explained

### 1. String iteration and character access

In Go, strings are immutable sequences of bytes encoded in UTF-8. You can iterate over them using `for...range` which gives you runes (Unicode code points) rather than bytes.

```go
for _, char := range myString {
    // char is a rune (int32)
}
```

To access individual characters, you can also use indexing, but remember that `s[i]` returns a byte, not a rune. For UTF-8 safety, use `for...range`.

### 2. Conditional logic and boolean returns

Go uses `if/else` for conditional branching. The condition doesn't need parentheses:

```go
if condition {
    // do something
} else if otherCondition {
    // do something else
} else {
    // default case
}
```

Boolean operators: `&&` (AND), `||` (OR), `!` (NOT).

### 3. Command-line argument handling

Access command-line arguments via `os.Args`:

```go
import "os"

func main() {
    args := os.Args[1:]  // Skip program name
    for _, arg := range args {
        fmt.Println(arg)
    }
}
```

Or use the `flag` package for more complex argument parsing.

## The Challenge

See [README.md](README.md) for the full challenge description, expected function, and test cases.

**Next:** [31-splitjoin](../31-splitjoin/skills.md) - Splitjoin
