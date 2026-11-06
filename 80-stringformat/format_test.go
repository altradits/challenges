package piscine

import "testing"

func TestFormat(t *testing.T) {
	cases := []struct {
		name string
		tmpl string
		args []interface{}
		want string
	}{
		{"string", "Hello %s", []interface{}{"World"}, "Hello World"},
		{"ints", "%d + %d = %d", []interface{}{1, 2, 3}, "1 + 2 = 3"},
		{"mixed", "%s is %d years old", []interface{}{"Alice", 30}, "Alice is 30 years old"},
		{"no placeholders", "No placeholders", nil, "No placeholders"},
		{"float", "pi is %f", []interface{}{3.14}, "pi is 3.140000"},
		{"missing arg kept", "Hello %s %s", []interface{}{"A"}, "Hello A %s"},
		{"extra arg ignored", "Hi", []interface{}{1, 2}, "Hi"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Format(tc.tmpl, tc.args...); got != tc.want {
				t.Errorf("Format(%q, %v) = %q, want %q", tc.tmpl, tc.args, got, tc.want)
			}
		})
	}
}
