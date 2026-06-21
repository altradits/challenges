# Skills for 145-time

## What You'll Learn

**Previous:** [144-grouping](../144-grouping/skills.md) | **Next:** [146-encoding-json](../146-encoding-json/skills.md)

**Challenge:** Parse, format, and compute with dates and times.

## The `time` Package

```go
import "time"
```

### Getting the Current Time

```go
now := time.Now()    // time.Time — local time
utc := time.Now().UTC()  // UTC
```

### `time.Time` Type

`time.Time` is a struct representing an instant in time. Zero value is January 1, year 1.

```go
var t time.Time          // zero value
t.IsZero()               // true
```

### `time.Duration` — Measuring Time Intervals

Duration is `int64` nanoseconds under the hood:

```go
d := 2*time.Hour + 30*time.Minute  // 2h30m
d := 500 * time.Millisecond
d := 100 * time.Nanosecond

// Constants
time.Second       // 1s
time.Minute       // 1m
time.Hour         // 1h
```

Arithmetic:
```go
future := time.Now().Add(24 * time.Hour)   // 24 hours from now
past   := time.Now().Add(-7 * 24 * time.Hour)  // 7 days ago
diff   := time.Since(past)    // Duration since past
diff   := time.Until(future)  // Duration until future
diff   := b.Sub(a)            // b - a as Duration
```

### Formatting — The Reference Time

Go uses a unique trick: format strings use the reference time `Mon Jan 2 15:04:05 MST 2006`.
Each component has a specific value:

```
Year:    2006
Month:   01  (or Jan)
Day:     02
Hour:    15  (24h) or 03 (12h)
Minute:  04
Second:  05
TZ:      MST or -0700
```

```go
t := time.Now()
fmt.Println(t.Format("2006-01-02"))              // 2024-03-15
fmt.Println(t.Format("2006-01-02 15:04:05"))     // 2024-03-15 09:30:00
fmt.Println(t.Format("January 2, 2006"))          // March 15, 2024
fmt.Println(t.Format(time.RFC3339))               // 2024-03-15T09:30:00Z
```

### Parsing

```go
t, err := time.Parse("2006-01-02", "2024-03-15")
if err != nil {
    return err
}
// t is time.Time in UTC

// Parse in a specific timezone
loc, _ := time.LoadLocation("Africa/Nairobi")
t, err = time.ParseInLocation("2006-01-02 15:04", "2024-03-15 09:30", loc)
```

### Extracting Components

```go
t := time.Now()
t.Year()        // int
t.Month()       // time.Month (January = 1)
t.Day()         // int
t.Hour()        // int (0-23)
t.Minute()      // int
t.Second()      // int
t.Weekday()     // time.Weekday (Sunday = 0, Monday = 1, ...)
t.YearDay()     // day of year (1-366)
```

### Comparing Times

```go
a.Before(b)   // true if a < b
a.After(b)    // true if a > b
a.Equal(b)    // true if a == b (prefer over ==)
```

### Timezones

```go
loc, err := time.LoadLocation("Africa/Nairobi")
nairobi := time.Now().In(loc)

utc := time.Now().UTC()
local := utc.Local()
```

### Sleeping

```go
time.Sleep(2 * time.Second)  // pause goroutine for 2 seconds
```

### Solving the Challenge

```go
package piscine

import "time"

func AgeInYears(birthDate time.Time) int {
    now := time.Now()
    years := now.Year() - birthDate.Year()
    if now.Month() < birthDate.Month() ||
        (now.Month() == birthDate.Month() && now.Day() < birthDate.Day()) {
        years--
    }
    return years
}

func FormatLog(t time.Time) string {
    return t.Format("2006-01-02 15:04:05")
}

func ParseDate(s string) (time.Time, error) {
    return time.Parse("2006-01-02", s)
}

func BusinessDays(start, end time.Time) int {
    count := 0
    d := start
    for d.Before(end) {
        wd := d.Weekday()
        if wd != time.Saturday && wd != time.Sunday {
            count++
        }
        d = d.AddDate(0, 0, 1)
    }
    return count
}
```

**Next:** [146-encoding-json](../146-encoding-json/README.md)
