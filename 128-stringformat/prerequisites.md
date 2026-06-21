# Prerequisites for 128-stringformat

## Before You Start

### 1. `fmt.Println` and `fmt.Printf` — you've been using these since challenge 01

```go
fmt.Println("Hello")          // prints with newline
fmt.Printf("Value: %d\n", 42) // prints formatted
```

Review: [01-only1](../01-only1/skills.md), [04-onlyf](../04-onlyf/skills.md)

### 2. Format verbs — the basics

```go
fmt.Printf("%s", "text")   // string
fmt.Printf("%d", 42)       // integer
fmt.Printf("%f", 3.14)     // float
fmt.Printf("%v", anything) // any type (general)
```

### 3. `fmt.Sprintf` — returns a string instead of printing

```go
s := fmt.Sprintf("Hello %s", "world")
// s = "Hello world"
```

### 4. Variadic parameters with `...`

A function that accepts any number of arguments:

```go
func greetAll(names ...string) {
    for _, name := range names {
        fmt.Println("Hello", name)
    }
}
greetAll("Alice", "Bob", "Carol")
```

And `args...` unpacks a slice into individual arguments:

```go
args := []interface{}{"Alice", 30}
fmt.Sprintf("Name: %s, Age: %d", args...)
```

## You're Ready When You Can...

- [ ] Use `fmt.Printf` with at least three different format verbs
- [ ] Use `fmt.Sprintf` to build a string without printing it
- [ ] Understand the difference between `fmt.Print`, `fmt.Printf`, and `fmt.Sprintf`

## Next Steps

- You've completed the strings series! Explore:
- [129-financial-freedom-api](../129-financial-freedom-api/README.md)
