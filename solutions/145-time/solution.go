package piscine

import "time"

func AgeInYears(birthDate time.Time) int {
	now := time.Now()
	years := now.Year() - birthDate.Year()
	anniversary := birthDate.AddDate(years, 0, 0)
	if now.Before(anniversary) {
		years--
	}
	return years
}

func FormatLog(t time.Time) string {
	return t.UTC().Format("2006-01-02 15:04:05")
}

func ParseDate(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
}

func BusinessDays(start, end time.Time) int {
	if !end.After(start) {
		return 0
	}
	count := 0
	for d := start; d.Before(end); d = d.AddDate(0, 0, 1) {
		if wd := d.Weekday(); wd != time.Saturday && wd != time.Sunday {
			count++
		}
	}
	return count
}
