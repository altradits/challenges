package piscine

import "testing"

func TestDouble(t *testing.T) {
	cases := []struct {
		name  string
		input int
		want  int
	}{
		{"positive", 5, 10},
		{"zero", 0, 0},
		{"negative", -3, -6},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			n := tc.input
			Double(&n)
			if n != tc.want {
				t.Errorf("Double(&%d) left %d, want %d", tc.input, n, tc.want)
			}
		})
	}
}

func TestSwap(t *testing.T) {
	a, b := 3, 7
	Swap(&a, &b)
	if a != 7 || b != 3 {
		t.Errorf("Swap gave (%d, %d), want (7, 3)", a, b)
	}
	x := 1
	Swap(&x, &x)
	if x != 1 {
		t.Errorf("Swap(&x, &x) gave %d, want 1", x)
	}
}

func TestNewPersonPtr(t *testing.T) {
	p := NewPerson("Alice", 30)
	if p == nil {
		t.Fatal("NewPerson returned nil")
	}
	if p.Name != "Alice" || p.Age != 30 {
		t.Errorf("NewPerson = %+v, want {Alice 30}", *p)
	}
	p.Age = 31
	if p.Age != 31 {
		t.Error("returned pointer should allow mutation through it")
	}
}
