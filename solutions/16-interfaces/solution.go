package piscine

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

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
		return fmt.Sprintf("Unknown: area=%.2f, perimeter=%.2f", s.Area(), s.Perimeter())
	}
}
