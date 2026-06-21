# Skills for 28-structs

## What You'll Learn

**Previous:** [12-recursion](../12-recursion/skills.md) | **Next:** [14-pointers](../14-pointers/skills.md)

**Challenge:** Define a `Person` struct and add methods to it.

## Core Concept: Structs

A struct groups related fields under one type name.

```go
type Person struct {
    Name string
    Age  int
}
```

### Creating Struct Values

```go
// Struct literal — named fields (preferred)
p := Person{Name: "Alice", Age: 30}

// Positional (fragile — breaks when fields change)
p := Person{"Alice", 30}

// Zero value — all fields are zero/empty
var p Person   // Name: "", Age: 0
```

### Reading and Writing Fields

```go
fmt.Println(p.Name)  // Alice
p.Age = 31
fmt.Println(p.Age)   // 31
```

### Constructor Functions

Go has no special constructor syntax. Use a plain function named `New<Type>`:

```go
func NewPerson(name string, age int) Person {
    return Person{Name: name, Age: age}
}
```

### Methods on Structs

A method is a function with a *receiver* — a special first parameter that names the type it belongs to.

```go
func (p Person) Greet() string {
    return "Hi, I'm " + p.Name
}

func (p Person) IsAdult() bool {
    return p.Age >= 18
}
```

The receiver `(p Person)` means "this method belongs to Person, and inside the method, the value is called p."

```go
alice := NewPerson("Alice", 30)
fmt.Println(alice.Greet())    // Hi, I'm Alice
fmt.Println(alice.IsAdult())  // true
```

### Exported vs Unexported Fields

Fields that start with an uppercase letter are **exported** (visible outside the package).
Fields starting with lowercase are **unexported** (package-private).

```go
type BankAccount struct {
    Owner   string  // exported — anyone can read it
    balance int     // unexported — only this package can access it
}
```

### Structs are Value Types

Assigning a struct copies all its fields:

```go
a := Person{Name: "Alice", Age: 30}
b := a         // b is a full copy
b.Name = "Bob" // does NOT change a
fmt.Println(a.Name) // Alice
fmt.Println(b.Name) // Bob
```

### Comparing Structs

Structs are comparable with `==` if all their fields are comparable:

```go
p1 := Person{"Alice", 30}
p2 := Person{"Alice", 30}
fmt.Println(p1 == p2) // true
```

### Solving the Challenge

```go
package piscine

type Person struct {
    Name string
    Age  int
}

func NewPerson(name string, age int) Person {
    return Person{Name: name, Age: age}
}

func (p Person) Greet() string {
    return "Hi, I'm " + p.Name
}

func (p Person) IsAdult() bool {
    return p.Age >= 18
}
```

**Next:** [14-pointers](../14-pointers/README.md)
