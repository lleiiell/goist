package goist

import (
	"time"
)

// DayBeginAt unix timestamp of zero o'clock of the day
func DayBeginAt(t time.Time) int64 {
	day := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	return day.Unix()
}

// TimeDiffDays number of days between given timestamps
func TimeDiffDays(start time.Time, end time.Time) int {
	return int(end.Sub(start).Hours() / 24)
}
