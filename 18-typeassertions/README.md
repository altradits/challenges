# 33-typeassertions — Type Assertions and Type Switches

## Challenge

Implement in `package piscine`:

```go
// Describe returns a human-readable description of any value.
// int → "integer: 42"
// string → "string: hello"
// bool → "bool: true"
// []int → "int slice of length 3"
// anything else → "unknown type"
func Describe(v any) string

// Sum adds together any mix of ints and float64s.
// Ignores values of other types.
func Sum(values []any) float64
```

Read `skills.md` before you start.
