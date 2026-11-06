package piscine

import (
	"reflect"
	"testing"
)

func TestCleanList(t *testing.T) {
	cases := []struct {
		name  string
		input []string
		want  []string
	}{
		{"shopping", []string{"tomatoes", "pastas", "milk", "salt"},
			[]string{"1/ Tomatoes", "2/ Pastas", "3/ Milk", "4/ Salt"}},
		{"milk added", []string{"bread", "eggs"},
			[]string{"1/ Bread", "2/ Eggs", "3/ Milk"}},
		{"empty", []string{}, []string{}},
		{"nil", nil, []string{}},
		{"milk case-insensitive", []string{"MILK"}, []string{"1/ Milk"}},
		{"trimmed", []string{"  bread  "}, []string{"1/ Bread", "2/ Milk"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := CleanList(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("CleanList(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
