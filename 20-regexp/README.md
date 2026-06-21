# 35-regexp — Regular Expressions

## Challenge

Implement in `package piscine`:

```go
// ValidateEmail returns true if s looks like a valid email address.
func ValidateEmail(s string) bool

// ExtractNumbers returns all integer sequences found in s.
// "foo123bar456" → ["123", "456"]
func ExtractNumbers(s string) []string

// MaskPhone replaces digits in a phone number with *, keeping the last 4.
// "0712345678" → "******5678"
func MaskPhone(s string) string
```

Read `skills.md` before you start.
