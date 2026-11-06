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

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 163
```

or run the tests directly:

```bash
go test ./163-type-conversions/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 163` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
