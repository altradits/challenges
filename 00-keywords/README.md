# 00-keywords — The 25 Words That Are All of Go

> **Read this before challenge 01. You will recognise every keyword in every challenge from the start.**

Go has exactly **25 reserved keywords**. That is the entire language surface. Compare that to Java (50+), C++ (90+), or Python (35+). Because the list is so short, you can memorise it completely on day one and never be surprised by syntax again.

---

## The 25 Keywords

| # | Keyword | Category | One-Line Meaning |
|---|---------|----------|-----------------|
| 1 | `package` | Structure | Declares which package this file belongs to |
| 2 | `import` | Structure | Brings another package into scope |
| 3 | `func` | Functions | Declares a function or method |
| 4 | `return` | Functions | Sends a value back from a function |
| 5 | `var` | Variables | Declares a variable with an explicit type |
| 6 | `const` | Variables | Declares a compile-time constant |
| 7 | `type` | Types | Defines a new named type or type alias |
| 8 | `struct` | Types | Groups fields into a composite type |
| 9 | `interface` | Types | Declares a set of method signatures |
| 10 | `map` | Types | Declares a key→value hash table |
| 11 | `chan` | Concurrency | Declares a channel for goroutine communication |
| 12 | `go` | Concurrency | Starts a new goroutine (lightweight thread) |
| 13 | `select` | Concurrency | Waits on multiple channel operations |
| 14 | `for` | Control | The only loop keyword in Go |
| 15 | `if` | Control | Conditional branching |
| 16 | `else` | Control | Alternative branch after `if` |
| 17 | `switch` | Control | Multi-way branching (no fall-through by default) |
| 18 | `case` | Control | One arm of a `switch` or `select` |
| 19 | `default` | Control | Fallback arm of `switch` or `select` |
| 20 | `break` | Control | Exits the nearest `for`, `switch`, or `select` |
| 21 | `continue` | Control | Skips to the next iteration of a `for` loop |
| 22 | `fallthrough` | Control | Forces execution into the next `switch` case |
| 23 | `defer` | Lifecycle | Schedules a call to run when the function returns |
| 24 | `goto` | Control | Jumps to a labelled statement (rarely used) |
| 25 | `range` | Control | Iterates over slice, map, string, channel, or array |

---

## Category Deep-Dives

### Structure Keywords: `package`, `import`

Every Go file starts with these two:

```go
package main          // this file belongs to the "main" package

import (
    "fmt"             // bring in the fmt package
    "strings"         // bring in the strings package
)
```

Rules:
- `package main` means this file can be run as a program.
- Any other name (`package piscine`, `package auth`) means it is a library.
- Go **refuses to compile** if you import a package you do not use.
- Go **refuses to compile** if you use a package you have not imported.

---

### Function Keywords: `func`, `return`

```go
func Add(a, b int) int {   // func + name + params + return type
    return a + b            // return sends the value back
}

// Multiple return values — unique to Go
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}
```

`return` can send back multiple values at once. This is how Go handles errors without exceptions.

---

### Variable Keywords: `var`, `const`

```go
// var — for package-level or when you need an explicit type
var name string
var count int = 0

// := — short declaration inside functions (type is inferred)
message := "hello"   // Go infers string

// const — value fixed at compile time, never changes
const Pi    = 3.14159
const AppName = "GoSchool"

// iota — auto-incrementing constant, reset per const block
type Day int
const (
    Monday  Day = iota + 1  // 1
    Tuesday                  // 2
    Wednesday                // 3
)
```

---

### Type Keywords: `type`, `struct`, `interface`, `map`, `chan`

**`type`** creates a new named type:

```go
type Celsius float64
type UserID  int
type Handler func(w http.ResponseWriter, r *http.Request)
```

**`struct`** groups fields:

```go
type User struct {
    Name  string
    Email string
    Age   int
}
u := User{Name: "Alice", Email: "alice@example.com", Age: 30}
```

**`interface`** declares behaviour:

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
// Any type with a Write method satisfies this interface automatically.
```

**`map`** creates hash tables:

```go
scores := map[string]int{
    "Alice": 95,
    "Bob":   87,
}
scores["Carol"] = 91
v, ok := scores["Dave"]  // ok is false if key does not exist
```

**`chan`** creates channels for goroutine communication:

```go
ch := make(chan int, 5)   // buffered channel, capacity 5
ch <- 42                  // send
v := <-ch                 // receive
```

---

### Concurrency Keywords: `go`, `select`

**`go`** starts a goroutine — a lightweight function running concurrently:

```go
go fmt.Println("I run in the background")
go func() {
    // anonymous goroutine
}()
```

**`select`** chooses among multiple channel operations (like a switch for channels):

```go
select {
case msg := <-chA:
    fmt.Println("from A:", msg)
case msg := <-chB:
    fmt.Println("from B:", msg)
case <-time.After(1 * time.Second):
    fmt.Println("timeout")
}
```

---

### Loop & Control Keywords: `for`, `range`, `break`, `continue`

Go has **one loop keyword** — `for`. It covers every loop pattern:

```go
// Classic C-style
for i := 0; i < 10; i++ { }

// While-style
for condition { }

// Infinite loop
for { }

// Range over slice
for i, v := range slice { }

// Range over map
for k, v := range myMap { }

// Range over string (yields runes, not bytes)
for i, ch := range "hello" { }

// Range over channel (reads until channel closes)
for v := range ch { }
```

`break` exits the nearest `for`, `switch`, or `select`. `continue` skips to the next iteration.

---

### Branch Keywords: `if`, `else`, `switch`, `case`, `default`, `fallthrough`

**`if`** with an initialiser — unique to Go:

```go
if err := doSomething(); err != nil {
    // err is scoped to this if/else block
    fmt.Println("error:", err)
}
```

**`switch`** without fall-through (each case breaks automatically):

```go
switch day {
case Monday:
    fmt.Println("start of the week")
case Friday:
    fmt.Println("end of the week")
default:
    fmt.Println("midweek")
}
```

**Type switch** — checks the dynamic type of an interface:

```go
switch v := val.(type) {
case int:
    fmt.Println("integer:", v)
case string:
    fmt.Println("string:", v)
default:
    fmt.Println("unknown type")
}
```

`fallthrough` forces execution into the next case (rare — explicit opt-in):

```go
switch x {
case 1:
    fmt.Println("one")
    fallthrough
case 2:
    fmt.Println("one or two")
}
```

---

### Lifecycle Keyword: `defer`

`defer` schedules a function call to run **after the enclosing function returns**, in LIFO order:

```go
func ReadFile(path string) error {
    f, err := os.Open(path)
    if err != nil {
        return err
    }
    defer f.Close()  // guaranteed to run when ReadFile returns

    // ... read from f ...
    return nil
}
```

Defers run even if the function panics. This makes it ideal for cleanup: closing files, releasing locks, logging exit times.

```go
// Multiple defers — LIFO order
defer fmt.Println("third")
defer fmt.Println("second")
defer fmt.Println("first")
// Output: first, second, third
```

---

### `goto` (Know It Exists, Rarely Use It)

```go
loop:
    if x < 10 {
        x++
        goto loop
    }
```

`goto` jumps to a labelled statement. Used occasionally for generated code or tight performance loops. In most code, `for` is cleaner.

---

## Why 25 Keywords Matters

Every other language feature in Go — slices, goroutines, generics, interfaces — is **built from these 25 words**. When you read unfamiliar Go code you will always be able to parse the structure because the keywords are finite and you know them all.

Compare:
- **Python** has 35 keywords plus magic dunder methods you must separately learn.
- **Java** has 50+ keywords plus modifier keywords (`public`, `protected`, `final`, `static`, `synchronized`, `volatile`, `transient`, `abstract`, `native`...).
- **C++** has 90+ keywords.

Go's constraint is intentional. The language designers optimised for **readability over expressiveness**: a team of 10 should be able to read each other's Go without learning each person's preferred idioms.

---

## Quick Reference Card

```
Package/Import:  package  import
Functions:       func     return
Variables:       var      const
Types:           type     struct  interface  map  chan
Concurrency:     go       select
Loop:            for      range   break      continue
Branch:          if       else    switch     case  default  fallthrough
Lifecycle:       defer
Jump:            goto
```

---

## What's Next

Start with [01-only1](../01-only1/README.md) and you will use `package`, `import`, `func`, and `return` in the very first challenge.

Each subsequent challenge introduces more of these 25 keywords in context. By challenge 27 you will have used all of them at least once.

**Full learning path:** [01-only1](../01-only1/README.md) → [02-onlya](../02-onlya/README.md) → ... → [27-context](../27-context/README.md) → [159-closures](../159-closures/README.md)
