# Skills for 29-pointers

## What You'll Learn

**Previous:** [13-structs](../13-structs/skills.md) | **Next:** [15-methods](../15-methods/skills.md)

**Challenge:** Use pointers to mutate values and build constructors.

## Core Concept: Pointers

A pointer stores the **memory address** of a variable, not the value itself.

```
Variable x:   address 0xc000018050, value 5
Pointer p:    address 0xc000018060, value 0xc000018050  (points to x)
```

### The Two Operators

```go
x := 42
p := &x    // & = "address of"  → p is *int, stores address of x
*p = 100   // * = "dereference" → follow the pointer, set the value
fmt.Println(x)  // 100 (x was changed through p)
```

### Why Pointers?

**1. Mutation across function calls**

In Go, function arguments are *copied*. Without a pointer, the function can't change the caller's variable:

```go
func tryDouble(n int) { n = n * 2 }  // copies n, original unchanged

func double(n *int) { *n = *n * 2 }  // mutates through pointer

x := 5
tryDouble(x)  // x is still 5
double(&x)    // x is now 10
```

**2. Avoiding large copies**

Passing a large struct by pointer is cheaper than copying all its fields.

### Pointers to Structs

```go
type Person struct{ Name string; Age int }

p := &Person{Name: "Alice", Age: 30}  // p is *Person

// Two equivalent ways to access fields through a pointer:
fmt.Println((*p).Name)  // Alice — explicit dereference
fmt.Println(p.Name)     // Alice — Go shorthand (auto-dereference)
```

Go automatically dereferences struct pointers when accessing fields — you almost never write `(*p).Field`.

### Constructor that Returns a Pointer

```go
func NewPerson(name string, age int) *Person {
    return &Person{Name: name, Age: age}
}
```

The struct is allocated on the heap when you return its address. Go's garbage collector manages it.

### nil Pointers

The zero value of any pointer type is `nil` — it points to nothing.

```go
var p *Person  // nil
p.Name         // PANIC: nil pointer dereference
```

Always check before using:
```go
if p != nil {
    fmt.Println(p.Name)
}
```

### Pointer vs Value: Quick Rule

| Use value `T` when... | Use pointer `*T` when... |
|---|---|
| Function only reads data | Function needs to mutate data |
| Struct is small (a few fields) | Struct is large |
| You want copy semantics | You want shared semantics |

### Solving the Challenge

```go
package piscine

func Double(n *int) {
    *n = *n * 2
}

func Swap(a, b *int) {
    *a, *b = *b, *a
}

type Person struct {
    Name string
    Age  int
}

func NewPerson(name string, age int) *Person {
    return &Person{Name: name, Age: age}
}
```

**Next:** [15-methods](../15-methods/README.md)
