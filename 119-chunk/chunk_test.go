package piscine

import (
	"io"
	"os"
	"testing"
)

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w
	fn()
	w.Close()
	os.Stdout = old
	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	return string(out)
}

func TestChunk(t *testing.T) {
	cases := []struct {
		name  string
		slice []int
		size  int
		want  string
	}{
		{"empty slice", []int{}, 10, "[]\n"},
		{"zero size", []int{0, 1, 2, 3, 4, 5, 6, 7}, 0, "\n"},
		{"negative size", []int{0, 1, 2}, -2, "\n"},
		{"three", []int{0, 1, 2, 3, 4, 5, 6, 7}, 3, "[[0 1 2] [3 4 5] [6 7]]\n"},
		{"five", []int{0, 1, 2, 3, 4, 5, 6, 7}, 5, "[[0 1 2 3 4] [5 6 7]]\n"},
		{"four", []int{0, 1, 2, 3, 4, 5, 6, 7}, 4, "[[0 1 2 3] [4 5 6 7]]\n"},
		{"bigger than slice", []int{1, 2}, 9, "[[1 2]]\n"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := captureStdout(t, func() { Chunk(tc.slice, tc.size) })
			if got != tc.want {
				t.Errorf("Chunk(%v, %d) printed %q, want %q", tc.slice, tc.size, got, tc.want)
			}
		})
	}
}
