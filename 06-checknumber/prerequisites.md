# Prerequisites for 06-checknumber

## Before You Start

To solve this challenge you need to understand how to write a Go function, how to loop over a string, and how to compare characters.

### 1. Writing a function with parameters and a return type
This challenge asks you to write a named function, not just `main`.

```go
func CheckNumber(arg string) bool {
    // arg is the input string
    // bool means the function returns true or false
    return false
}
```

### 2. `for...range` on a string
To inspect each character in a string, use `for...range`. It gives you the byte index (`i`) and the rune value (`r`) of each character.

```go
s := "hello"
for i, r := range s {
    // i is the position (0, 1, 2, ...)
    // r is the character as a rune (int32)
    _ = i  // use _ to ignore a value you don't need
}
```

### 3. Rune comparisons with character literals
Single-quote characters like `'0'` and `'9'` are rune constants. You can compare them with `>=` and `<=` just like integers.

```go
r := 'A'
if r >= '0' && r <= '9' {
    // r is a digit character
}
```

### 4. `if/else` and boolean return values
Returning `true` or `false` from a function:

```go
func IsEven(n int) bool {
    if n % 2 == 0 {
        return true
    }
    return false
}
```

### 5. The basic program skeleton
Even though this challenge defines a function, you still need `package main`, `import "fmt"`, and `func main()` to run and test it.

```go
package main

import "fmt"

func CheckNumber(arg string) bool {
    // your implementation
    return false
}

func main() {
    fmt.Println(CheckNumber("Hello"))   // false
    fmt.Println(CheckNumber("Hello1"))  // true
}
```

## Review If Stuck

- [07-functions](../07-functions/skills.md) — dedicated lesson on writing functions with parameters and return types
- [01-only1](../01-only1/skills.md) — covers `package main`, `import "fmt"`, `func main()`

## You're Ready When You Can...

- [ ] Write a function that takes a `string` parameter and returns `bool`
- [ ] Use `for _, r := range s` to visit every character in a string
- [ ] Compare a rune to `'0'` and `'9'` using `>=` and `<=`
- [ ] Return `true` early from inside a loop
- [ ] Return `false` after the loop when no match is found

## Next Steps

After completing this exercise, move to:
- [08-count-character](../08-count-character/README.md)
