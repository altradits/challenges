# 163-type-conversions — Type Conversions and Runes

## Challenge

Implement in `package piscine`:

```go
// IntToFloat converts an int to float64.
func IntToFloat(n int) float64

// FloatToInt truncates a float64 to int (toward zero).
func FloatToInt(f float64) int

// RuneToString converts a single rune to a one-character string.
func RuneToString(r rune) string

// StringToRunes converts a string to a slice of runes (handles multi-byte UTF-8).
func StringToRunes(s string) []rune

// BytesToString converts a byte slice to a string.
func BytesToString(b []byte) string

// CountRunes returns the number of Unicode characters in s (not bytes).
func CountRunes(s string) int
```

## Examples

```
IntToFloat(42)         →  42.0
FloatToInt(3.9)        →  3
FloatToInt(-3.9)       →  -3

RuneToString('A')      →  "A"
RuneToString('€')      →  "€"

StringToRunes("hello")  →  ['h','e','l','l','o']
StringToRunes("日本語")  →  ['日','本','語']   // 3 runes, not 9 bytes

BytesToString([]byte{72,101,108,108,111})  →  "Hello"

CountRunes("hello")   →  5
CountRunes("日本語")   →  3
```

Read `skills.md` before you start.
