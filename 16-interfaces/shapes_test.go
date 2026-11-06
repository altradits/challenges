package piscine

import (
	"math"
	"testing"
)

func TestShapesSatisfyInterface(t *testing.T) {
	var _ Shape = Rectangle{}
	var _ Shape = Circle{}
}

func TestShapeAreaPerimeter(t *testing.T) {
	r := Rectangle{Width: 2, Height: 3}
	if got := r.Area(); got != 6 {
		t.Errorf("Rectangle.Area() = %v, want 6", got)
	}
	if got := r.Perimeter(); got != 10 {
		t.Errorf("Rectangle.Perimeter() = %v, want 10", got)
	}
	c := Circle{Radius: 3}
	if got := c.Area(); math.Abs(got-28.27433388) > 1e-6 {
		t.Errorf("Circle(3).Area() = %v, want ~28.274", got)
	}
	if got := c.Perimeter(); math.Abs(got-18.84955592) > 1e-6 {
		t.Errorf("Circle(3).Perimeter() = %v, want ~18.850", got)
	}
}

func TestTotalArea(t *testing.T) {
	shapes := []Shape{Rectangle{2, 3}, Circle{1}}
	want := 6 + math.Pi
	if got := TotalArea(shapes); math.Abs(got-want) > 1e-9 {
		t.Errorf("TotalArea = %v, want %v", got, want)
	}
	if got := TotalArea(nil); got != 0 {
		t.Errorf("TotalArea(nil) = %v, want 0", got)
	}
}

func TestDescribeShape(t *testing.T) {
	cases := []struct {
		name  string
		shape Shape
		want  string
	}{
		{"rectangle", Rectangle{2, 3}, "Rectangle: area=6.00, perimeter=10.00"},
		{"circle", Circle{3}, "Circle: area=28.27, perimeter=18.85"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Describe(tc.shape); got != tc.want {
				t.Errorf("Describe = %q, want %q", got, tc.want)
			}
		})
	}
}
