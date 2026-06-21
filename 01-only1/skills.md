# Skills for 01-only1

## What You'll Learn

**Previous:** Start of course | **Next:** [02-onlya](../02-onlya/skills.md)

**Challenge:** Write a program that prints the character `1` and nothing else.

## Core Concept: The Structure of Every Go Program

### What Is It?

Every runnable Go program has the same three-part skeleton: a package declaration, an import block, and a `main` function. Until you understand these three pieces, no Go program will compile. This challenge exists so you can focus on *just* the structure, without worrying about logic.

### How It Works

Step 1 — Declare the package:

```go
package main
```

`package main` tells the Go compiler "this file is a standalone executable, not a library." Every file in an executable program must start with this line.

Step 2 — Import the packages you need:

```go
import "fmt"
```

`fmt` is the standard formatting library. It provides `Println`, `Printf`, and `Print`. You must import a package before you can use it — Go will refuse to compile if you import something you don't use, or use something you haven't imported.

Step 3 — Write the `main` function:

```go
func main() {
    fmt.Println("1")
}
```

`func main()` is the entry point. When you run `go run .`, Go finds the `main` function and executes it top-to-bottom.

### The Complete Program

```go
package main

import "fmt"

func main() {
    fmt.Println("1")
}
```

### What `fmt.Println` Does

```
fmt.Println("1")
      |         |
      |         argument: the value to print
      |
      function from the fmt package
```

`Println` prints its argument followed by a newline character (`\n`). So the output of this program is:

```
1
```

(with a newline at the end — that is correct behavior)

### Visual: Program Execution Flow

```
go run .
    |
    v
Go compiler finds package main
    |
    v
Go compiler resolves imports (fmt)
    |
    v
func main() begins executing
    |
    v
fmt.Println("1")  <-- prints: 1
    |
    v
main() returns, program exits
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `Package main` (capital P) | Go is case-sensitive; package names are lowercase | Write `package main` |
| Forgetting `import "fmt"` | `fmt.Println` is undefined without the import | Add the import block |
| `fmt.println` (lowercase p) | Go is case-sensitive; exported functions start with uppercase | Write `fmt.Println` |
| `fmt.Println(1)` (no quotes) | This prints an integer, which works here, but know the difference | Use `"1"` for a string literal |
| Missing curly braces `{}` | Go requires braces even for single-line functions | Always include `{` and `}` |

## Solving This Challenge

### Algorithm

1. Declare `package main`
2. Import `fmt`
3. Inside `func main()`, call `fmt.Println("1")`

That's it — three lines of real code.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [02-onlya](../02-onlya/README.md)
