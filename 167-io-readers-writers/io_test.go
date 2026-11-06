package piscine

import (
	"bytes"
	"strings"
	"testing"
)

func TestReadAll(t *testing.T) {
	got, err := ReadAll(strings.NewReader("hello"))
	if err != nil || string(got) != "hello" {
		t.Errorf("ReadAll = (%q, %v)", got, err)
	}
}

func TestWriteString(t *testing.T) {
	var buf bytes.Buffer
	n, err := WriteString(&buf, "hi")
	if err != nil || n != 2 || buf.String() != "hi" {
		t.Errorf("WriteString = (%d, %v, %q)", n, err, buf.String())
	}
}

func TestCopyIO(t *testing.T) {
	var buf bytes.Buffer
	n, err := Copy(&buf, strings.NewReader("abc"))
	if err != nil || n != 3 || buf.String() != "abc" {
		t.Errorf("Copy = (%d, %v, %q)", n, err, buf.String())
	}
}

func TestCountLines(t *testing.T) {
	got, err := CountLines(strings.NewReader("a\nb\nc\n"))
	if err != nil || got != 3 {
		t.Errorf("CountLines = (%d, %v), want 3", got, err)
	}
	got, err = CountLines(strings.NewReader("nonewline"))
	if err != nil || got != 0 {
		t.Errorf("CountLines no newline = (%d, %v), want 0", got, err)
	}
}

func TestUpperWriter(t *testing.T) {
	var buf bytes.Buffer
	w := UpperWriter(&buf)
	if _, err := w.Write([]byte("Hello, Go!")); err != nil {
		t.Fatal(err)
	}
	if buf.String() != "HELLO, GO!" {
		t.Errorf("UpperWriter wrote %q", buf.String())
	}
}
