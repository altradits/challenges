# 34-fileio — File I/O

## Challenge

Implement in `package piscine`:

```go
// ReadLines returns all lines from a file (without newline characters).
func ReadLines(path string) ([]string, error)

// WriteLines writes lines to a file, one per line.
// Creates the file if it doesn't exist, overwrites if it does.
func WriteLines(path string, lines []string) error

// WordCount reads a file and returns a map of word → frequency.
func WordCount(path string) (map[string]int, error)
```

Read `skills.md` before you start.
