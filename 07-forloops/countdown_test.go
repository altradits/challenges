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

func TestCountDown(t *testing.T) {
	cases := []struct {
		name string
		n    int
		want string
	}{
		{"five", 5, "5\n4\n3\n2\n1\n0\n"},
		{"zero", 0, "0\n"},
		{"one", 1, "1\n0\n"},
		{"ten", 10, "10\n9\n8\n7\n6\n5\n4\n3\n2\n1\n0\n"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := captureStdout(t, func() { CountDown(tc.n) })
			if got != tc.want {
				t.Errorf("CountDown(%d) printed %q, want %q", tc.n, got, tc.want)
			}
		})
	}
}
