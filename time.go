package goist

import (
	"fmt"
	"math"
	"time"
)

// DayBeginAt unix timestamp of zero o'clock of the day
func DayBeginAt(t time.Time) int64 {
	day := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	return day.Unix()
}

// DayEndedAt unix timestamp of 23:59:59 o'clock of the day
func DayEndedAt(t time.Time) int64 {
	return DayBeginAt(t) + 86400 - 1
}

// TimeDiffDays number of days between given time
func TimeDiffDays(first time.Time, second time.Time) int {

	d1 := time.Date(first.Year(), first.Month(), first.Day(), 0, 0, 0, 0, time.Local)
	d2 := time.Date(second.Year(), second.Month(), second.Day(), 0, 0, 0, 0, time.Local)

	return int(math.Abs(d2.Sub(d1).Hours() / 24))
}

func TimeIsSameDay(first time.Time, second time.Time) bool {
	return first.YearDay() == second.YearDay() && first.Year() == second.Year()
}

func TimeIsLeapYear(t time.Time) bool {
	y := t.Year()
	return y%4 == 0 && (y%100 != 0 || y%400 == 0)
}

// Time2short short string format
func Time2short(d time.Duration) string {
	// u := uint64(d)
	if d < time.Second {
		switch {
		case d == 0:
			return "0"
		case d < time.Microsecond:
			return fmt.Sprintf("%.2fns", float64(d))
		case d < time.Millisecond:
			return fmt.Sprintf("%.2fus", float64(d)/1000)
		default:
			return fmt.Sprintf("%.2fms", float64(d)/1000/1000)
		}
	} else {
		switch {
		case d < (time.Minute):
			return fmt.Sprintf("%.2fs", float64(d)/1000/1000/1000)
		case d < (time.Hour):
			return fmt.Sprintf("%.2fm", float64(d)/1000/1000/1000/60)
		default:
			return fmt.Sprintf("%.2fh", float64(d)/1000/1000/1000/60/60)
		}
	}
}

func Time2now(t time.Time) string {
	d := time.Since(t)

	if d < time.Minute {
		return "刚刚"
	} else if d < time.Hour {
		return fmt.Sprintf("%d分钟", uint(d)/1000/1000/1000/60)
	} else if d < time.Duration(24)*time.Hour {
		return fmt.Sprintf("%d小时", uint(d)/1000/1000/1000/60/60)
	} else {
		return fmt.Sprintf("%d天", uint(d)/1000/1000/1000/60/60/24)
	}

}

func CountWeekday(start, end time.Time, day time.Weekday) int {
	total := 0

	for start.Before(end) {
		if start.Weekday() == day {
			total = total + 1
		}
		start = start.AddDate(0, 0, 1)
	}

	return total
}

func tickerRun(f func() error, duration time.Duration) (err error) {

	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	ch := make(chan struct{}, 1)
	ch <- struct{}{}

	for {
		select {
		case <-ch:
			err = f()
			if err != nil {
				return err
			}
		case <-ticker.C:
			err = f()
			if err != nil {
				return err
			}
		}
	}

}
