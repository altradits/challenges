package piscine

import "testing"

func TestNewPerson(t *testing.T) {
	p := NewPerson("Alice", 30)
	if p.Name != "Alice" || p.Age != 30 {
		t.Errorf("NewPerson = %+v, want {Alice 30}", p)
	}
}

func TestPersonGreet(t *testing.T) {
	cases := []struct {
		name  string
		input Person
		want  string
	}{
		{"alice", Person{Name: "Alice", Age: 30}, "Hi, I'm Alice"},
		{"empty", Person{}, "Hi, I'm "},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.input.Greet(); got != tc.want {
				t.Errorf("Greet() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestPersonIsAdult(t *testing.T) {
	cases := []struct {
		name string
		age  int
		want bool
	}{
		{"minor", 15, false},
		{"boundary", 18, true},
		{"adult", 30, true},
		{"zero", 0, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := NewPerson("X", tc.age)
			if got := p.IsAdult(); got != tc.want {
				t.Errorf("Age %d: IsAdult() = %v, want %v", tc.age, got, tc.want)
			}
		})
	}
}
