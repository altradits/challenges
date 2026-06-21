# Prerequisites for searchreplace

## Before You Start

To solve this challenge you need to understand:

### 1. os.Args
`os.Args` is a slice of strings. Index 0 is the program name. User-provided arguments start at index 1.
```go
import "os"

// Running: go run . "hello" "a" "o"
// os.Args == ["./program", "hello", "a", "o"]
fmt.Println(len(os.Args)) // 4
fmt.Println(os.Args[1])   // "hello"
fmt.Println(os.Args[2])   // "a"
fmt.Println(os.Args[3])   // "o"
```

### 2. Validating Argument Count
Check `len(os.Args)` before accessing any argument. The program name counts as an element.
```go
if len(os.Args) != 4 {  // 1 program name + 3 user args
    return  // exit without printing anything
}
```

### 3. Character-by-Character Replacement
Loop through a string and compare each character to find matches.
```go
result := ""
for _, c := range text {
    if string(c) == old {
        result += newChar
    } else {
        result += string(c)
    }
}
```

### 4. strings.ReplaceAll (allowed here)
Since this challenge doesn't forbid built-in functions, you can use the standard library:
```go
import "strings"

result := strings.ReplaceAll(text, old, newChar)
fmt.Println(result)
```

### 5. Printing Nothing on Bad Input
A bare `return` in `main()` exits the program without printing anything.
```go
if len(os.Args) != 4 {
    return  // no fmt.Println, no output
}
```

## Review If Stuck

- [../100-replaceall/skills.md](../100-replaceall/skills.md) — covers manual string replacement logic
- Prior challenges — cover `os.Args` for reading command-line arguments

## You're Ready When You Can...

- [ ] Access command-line arguments via `os.Args`
- [ ] Check `len(os.Args)` and return silently if the count is wrong
- [ ] Replace characters in a string either manually or with `strings.ReplaceAll`
- [ ] Print the result (or nothing) depending on input validity

## Next Steps

- [102-splitjoin](../102-splitjoin/README.md)
