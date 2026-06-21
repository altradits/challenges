# 31-interfaces — Interfaces in Go

## Challenge

Define and implement in `package piscine`:

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct{ Width, Height float64 }
type Circle struct{ Radius float64 }

// Both Rectangle and Circle must satisfy the Shape interface.

func TotalArea(shapes []Shape) float64  // sum of all areas
func Describe(s Shape) string           // "Rectangle: area=6.00, perimeter=10.00"
                                        // "Circle: area=28.27, perimeter=12.57"
```

Read `skills.md` before you start.
