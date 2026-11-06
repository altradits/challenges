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

func TestPrintMemory(t *testing.T) {
	got := captureStdout(t, func() {
		PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
	})
	want := "68 65 6c 6c\n6f 10 15 2a\n00 00\nhello..*..\n"
	if got != want {
		t.Errorf("PrintMemory output = %q, want %q (4 hex bytes per row, then ASCII with dots)", got, want)
	}
	got = captureStdout(t, func() {
		PrintMemory([10]byte{})
	})
	want = "00 00 00 00\n00 00 00 00\n00 00\n..........\n"
	if got != want {
		t.Errorf("PrintMemory zeros = %q, want %q", got, want)
	}
}
