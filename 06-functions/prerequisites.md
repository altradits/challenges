# Prerequisites for 06-functions

## Before You Start

To solve this challenge you need to understand Go program structure, string literals, and string concatenation.

### 1. Go program structure from 01-only1

Every Go file starts with a package declaration. The challenges you have solved so far use `package main`. Your function solutions will use `package piscine` instead, but the structure is the same:

```go
package piscine

// your functions go here
```

There is no `func main()` in your solution file — only the function(s) the challenge asks for.

### 2. String literals and `fmt.Println`

A string literal is a sequence of characters wrapped in double quotes:

```go
"Hello, World!"   // a string literal
'H'               // this is a rune (single character), not a string
```

`fmt.Println` prints a string followed by a newline:

```go
fmt.Println("Hello, World!")   // prints: Hello, World!
```

### 3. String concatenation with `+`

You can join strings together with the `+` operator:

```go
first := "Hello, "
name := "Go"
result := first + name + "!"   // "Hello, Go!"
```

Each `+` joins two strings into one. The result has the same type: `string`.

### 4. Function syntax

A Go function has this shape:

```go
func FunctionName(paramName paramType) returnType {
    return value
}
```

For this challenge:

```go
func Greet(name string) string {
    return "Hello, " + name + "!"
}
```

- `name string` — the function receives one string argument and names it `name`
- `string` after the parentheses — the function returns a string
- `return` — sends the value back to the caller

## Review If Stuck

- [01-only1](../01-only1/skills.md) — covers `package main`, `import "fmt"`, `func main()`
- [02-onlya](../02-onlya/skills.md) — covers string literals and double-quote syntax

## You're Ready When You Can...

- [ ] Write `package piscine` at the top of a file (not `package main`)
- [ ] Declare a function with a string parameter and a string return type
- [ ] Concatenate three strings with `+` to form a result
- [ ] Use `return` to send a value back from a function
- [ ] Explain the difference between `"hello"` (string) and `'h'` (rune)

## Next Steps

- [07-forloops](../07-forloops/README.md)
