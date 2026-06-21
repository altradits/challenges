# 145-time — The time Package

## Challenge

Implement in `package piscine`:

```go
// AgeInYears returns how many full years have passed since birthDate.
func AgeInYears(birthDate time.Time) int

// FormatLog formats a time as "2006-01-02 15:04:05" (the log timestamp format).
func FormatLog(t time.Time) string

// ParseDate parses a date string in "2006-01-02" format.
func ParseDate(s string) (time.Time, error)

// BusinessDays returns the number of weekdays between start and end (exclusive of end).
func BusinessDays(start, end time.Time) int
```

Read `skills.md` before you start.
