package goist_test

import (
	"fmt"
	"github.com/lleiiell/goist"
	"time"
)

// 1693193334 2023-08-28 11:28:54
// 1693152000 2023-08-28 00:00:00
func ExampleDayBeginAt() {
	t := time.Unix(1693193334, 0)
	u := goist.DayBeginAt(t)

	fmt.Println(u)
	// Output: 1693152000

}

func ExampleTimeDiffDays() {
	t1, _ := time.Parse("2006-01-02 15", "2023-09-01 1")
	t2, _ := time.Parse("2006-01-02 15", "2023-09-03 23")
	t3, _ := time.Parse("2006-01-02 15", "2023-08-28 23")

	fmt.Println(goist.TimeDiffDays(t1, t2), goist.TimeDiffDays(t3, t1))
	// Output: 2 3
}

func ExampleTimeIsLeapYear() {
	t1, _ := time.Parse("2006-01-02 15", "2000-09-01 1")
	t2, _ := time.Parse("2006-01-02 15", "1900-09-03 23")
	t3, _ := time.Parse("2006-01-02 15", "2024-08-28 23")

	fmt.Println(goist.TimeIsLeapYear(t1), goist.TimeIsLeapYear(t2), goist.TimeIsLeapYear(t3))
	// Output: true false true
}

func ExampleTime2short() {
	fmt.Println(goist.Time2short(10236 * time.Millisecond))
	fmt.Println(goist.Time2short(1023600 * time.Millisecond))
	fmt.Println(goist.Time2short(10236000 * time.Millisecond))

	// Output: 10.24s
	// 17.06m
	// 2.84h
}
