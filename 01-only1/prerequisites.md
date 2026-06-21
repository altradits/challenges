# Prerequisites for 01-only1

## Before You Start

This is your very first Go program. There are no prerequisites — you start here.

All you need before attempting this challenge is:
- Go installed on your machine (`go version` should print a version number)
- A text editor or IDE
- A terminal where you can run `go run .`

## What You'll Need to Understand

### 1. `package main`
Every executable Go program declares itself as `package main` on the first line. This is not optional.

```go
package main
```

### 2. `import "fmt"`
To print output, you need the `fmt` package. You declare it with `import`.

```go
import "fmt"
```

### 3. `func main()`
Go starts execution in the `main` function. Every program has exactly one.

```go
func main() {
    // your code goes here
}
```

### 4. `fmt.Println()`
This function prints a value to the screen, followed by a newline.

```go
fmt.Println("hello")  // prints: hello
```

## You're Ready When You Can...

- [ ] Create a file `main.go` with `package main` at the top
- [ ] Add `import "fmt"` before the function
- [ ] Write a `func main()` block with opening and closing braces
- [ ] Call `fmt.Println("something")` inside it
- [ ] Run `go run .` in the terminal and see the output

## Next Steps

After completing this exercise, move to:
- [02-onlya](../02-onlya/README.md)
