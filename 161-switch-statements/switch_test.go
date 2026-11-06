package piscine

import "testing"

func TestDayType(t *testing.T) {
	cases := []struct {
		day  string
		want string
	}{
		{"Monday", "weekday"},
		{"Friday", "weekday"},
		{"Saturday", "weekend"},
		{"Sunday", "weekend"},
	}
	for _, tc := range cases {
		if got := DayType(tc.day); got != tc.want {
			t.Errorf("DayType(%q) = %q, want %q", tc.day, got, tc.want)
		}
	}
}

func TestGrade(t *testing.T) {
	cases := []struct {
		score int
		want  string
	}{
		{100, "A"}, {90, "A"}, {89, "B"}, {80, "B"}, {79, "C"},
		{70, "C"}, {69, "D"}, {60, "D"}, {59, "F"}, {0, "F"},
	}
	for _, tc := range cases {
		if got := Grade(tc.score); got != tc.want {
			t.Errorf("Grade(%d) = %q, want %q", tc.score, got, tc.want)
		}
	}
}

func TestDescribeSwitch(t *testing.T) {
	cases := []struct {
		name  string
		input interface{}
		want  string
	}{
		{"int", 42, "integer: 42"},
		{"string", "hi", "string: hi"},
		{"bool", true, "boolean: true"},
		{"float", 1.5, "unknown type"},
		{"nil", nil, "unknown type"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Describe(tc.input); got != tc.want {
				t.Errorf("Describe(%v) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

func TestFizzBuzzSwitch(t *testing.T) {
	cases := map[int]string{3: "Fizz", 5: "Buzz", 15: "FizzBuzz", 7: "7", 1: "1"}
	for in, want := range cases {
		if got := FizzBuzz(in); got != want {
			t.Errorf("FizzBuzz(%d) = %q, want %q", in, got, want)
		}
	}
}
