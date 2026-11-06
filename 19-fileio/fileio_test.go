package piscine

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestWriteThenReadLines(t *testing.T) {
	path := filepath.Join(t.TempDir(), "lines.txt")
	lines := []string{"first", "second", "third"}
	if err := WriteLines(path, lines); err != nil {
		t.Fatalf("WriteLines: %v", err)
	}
	got, err := ReadLines(path)
	if err != nil {
		t.Fatalf("ReadLines: %v", err)
	}
	if !reflect.DeepEqual(got, lines) {
		t.Errorf("round trip = %q, want %q", got, lines)
	}
}

func TestWriteLinesOverwrites(t *testing.T) {
	path := filepath.Join(t.TempDir(), "over.txt")
	if err := WriteLines(path, []string{"old"}); err != nil {
		t.Fatal(err)
	}
	if err := WriteLines(path, []string{"new"}); err != nil {
		t.Fatal(err)
	}
	got, err := ReadLines(path)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, []string{"new"}) {
		t.Errorf("after overwrite = %q, want [new]", got)
	}
}

func TestReadLinesMissing(t *testing.T) {
	if _, err := ReadLines(filepath.Join(t.TempDir(), "nope.txt")); err == nil {
		t.Error("ReadLines of a missing file should return an error")
	}
}

func TestWordCountFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "words.txt")
	if err := os.WriteFile(path, []byte("go is go\nfun go\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	got, err := WordCount(path)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]int{"go": 3, "is": 1, "fun": 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("WordCount = %v, want %v", got, want)
	}
	if _, err := WordCount(filepath.Join(t.TempDir(), "nope.txt")); err == nil {
		t.Error("WordCount of a missing file should return an error")
	}
}
