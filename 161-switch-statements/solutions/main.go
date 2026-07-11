package piscine

import "fmt"

func DayType(day string) string {
	switch day {
	case "Saturday", "Sunday":
		return "weekend"
	default:
		return "weekday"
	}
}

func Grade(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

func Describe(v interface{}) string {
	switch val := v.(type) {
	case int:
		return fmt.Sprintf("integer: %d", val)
	case string:
		return fmt.Sprintf("string: %s", val)
	case bool:
		return fmt.Sprintf("boolean: %v", val)
	default:
		return "unknown type"
	}
}

func FizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return fmt.Sprintf("%d", n)
	}
}
