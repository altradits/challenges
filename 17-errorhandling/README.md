# 32-errorhandling — Error Handling in Go

## Challenge

Implement in `package piscine`:

```go
// Divide returns a/b or an error if b is zero.
func Divide(a, b float64) (float64, error)

// ParseAge parses a string to a non-negative int.
// Returns ErrNegativeAge if the number is negative.
// Returns a wrapped strconv error if the string is not a number.
func ParseAge(s string) (int, error)

// Custom error type
type ValidationError struct {
    Field   string
    Message string
}
func (e *ValidationError) Error() string  // "field: message"
```

Read `skills.md` before you start.
