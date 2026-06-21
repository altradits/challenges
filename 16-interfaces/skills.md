# Skills for 31-interfaces

## What You'll Learn

**Previous:** [15-methods](../15-methods/skills.md) | **Next:** [17-errorhandling](../17-errorhandling/skills.md)

**Challenge:** Define a `Shape` interface and implement it for two types.

## Core Concept: Interfaces

An interface defines a set of method signatures. Any type that implements all those methods *automatically* satisfies the interface — no explicit declaration needed.

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

### Implicit Satisfaction

```go
type Rectangle struct{ Width, Height float64 }

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

// Rectangle now satisfies Shape — without writing "implements Shape" anywhere
var s Shape = Rectangle{Width: 3, Height: 4}
```

### Why Interfaces?

They let you write one function that works with many types:

```go
func TotalArea(shapes []Shape) float64 {
    total := 0.0
    for _, s := range shapes {
        total += s.Area()
    }
    return total
}
```

This works for `Rectangle`, `Circle`, or any future type you add.

### The `fmt.Stringer` Interface

The most common interface in Go. If your type implements `String() string`, `fmt.Println` will call it automatically:

```go
type Point struct{ X, Y int }

func (p Point) String() string {
    return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

p := Point{3, 4}
fmt.Println(p)  // (3, 4)
```

### The `error` Interface (preview)

```go
type error interface {
    Error() string
}
```

Any type with an `Error() string` method is an `error`. You'll use this in the next lesson.

### Interface Values and nil

An interface value holds two things: a concrete type and a concrete value. A nil interface has neither:

```go
var s Shape          // nil interface — s is nil
s = Rectangle{3, 4} // now s holds type=Rectangle, value={3,4}
```

### The Empty Interface: `any`

`any` (alias for `interface{}`) accepts any type:

```go
func Print(v any) {
    fmt.Println(v)
}
Print(42)       // works
Print("hello")  // works
Print(true)     // works
```

Use `any` sparingly — it sacrifices type safety.

### Interfaces + `math` package

```go
import "math"

type Circle struct{ Radius float64 }

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }
```

### Solving the Challenge

```go
package piscine

import (
    "fmt"
    "math"
)

type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct{ Width, Height float64 }
type Circle struct{ Radius float64 }

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }
func (c Circle) Area() float64         { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64    { return 2 * math.Pi * c.Radius }

func TotalArea(shapes []Shape) float64 {
    total := 0.0
    for _, s := range shapes {
        total += s.Area()
    }
    return total
}

func Describe(s Shape) string {
    switch v := s.(type) {
    case Rectangle:
        return fmt.Sprintf("Rectangle: area=%.2f, perimeter=%.2f", v.Area(), v.Perimeter())
    case Circle:
        return fmt.Sprintf("Circle: area=%.2f, perimeter=%.2f", v.Area(), v.Perimeter())
    default:
        return fmt.Sprintf("Unknown shape: area=%.2f", s.Area())
    }
}
```

**Next:** [17-errorhandling](../17-errorhandling/README.md)
