package piscine

import (
	"reflect"
	"testing"
)

func TestCountLines(t *testing.T) {
	cases := []struct {
		input string
		want  int
	}{
		{"", 0},
		{"a", 1},
		{"a\nb\nc", 3},
		{"a\nb\n", 2},
		{"\n\n", 2},
	}
	for _, tc := range cases {
		if got := CountLines(tc.input); got != tc.want {
			t.Errorf("CountLines(%q) = %d, want %d", tc.input, got, tc.want)
		}
	}
}

func TestWordFreq(t *testing.T) {
	got := WordFreq("go go test")
	want := map[string]int{"go": 2, "test": 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("WordFreq = %v, want %v", got, want)
	}
	got = WordFreq("")
	if got == nil || len(got) != 0 {
		t.Errorf("WordFreq(\"\") = %v, want empty non-nil map", got)
	}
	got = WordFreq("  a\t a\nb  ")
	if !reflect.DeepEqual(got, map[string]int{"a": 2, "b": 1}) {
		t.Errorf("WordFreq whitespace = %v", got)
	}
}

func TestJoinLines(t *testing.T) {
	if got := JoinLines([]string{"a", "b"}); got != "a\nb" {
		t.Errorf("JoinLines = %q", got)
	}
	if got := JoinLines(nil); got != "" {
		t.Errorf("JoinLines(nil) = %q, want \"\"", got)
	}
	if got := JoinLines([]string{"solo"}); got != "solo" {
		t.Errorf("JoinLines single = %q", got)
	}
}

func TestReverseWords(t *testing.T) {
	cases := []struct {
		input, want string
	}{
		{"hello old world", "world old hello"},
		{"a  b", "b a"},
		{"", ""},
		{"solo", "solo"},
		{"\tgo\nis fun ", "fun is go"},
	}
	for _, tc := range cases {
		if got := ReverseWords(tc.input); got != tc.want {
			t.Errorf("ReverseWords(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}
