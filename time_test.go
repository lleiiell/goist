package goist

import (
	"fmt"
	"testing"
	"time"
)

// 1693193334 2023-08-28 11:28:54
// 1693152000 2023-08-28 00:00:00
func ExampleDayBeginAt() {
	t := time.Unix(1693193334, 0)
	u := DayBeginAt(t)

	fmt.Println(time.Unix(u, 0).Format(time.TimeOnly))
	// Output: 00:00:00

}

func TestDayEndedAt(t *testing.T) {
	tm := time.Unix(1693193334, 0)
	tme := dayEndedAt(tm)
	fmt.Println(time.Unix(tme, 0).Format(time.TimeOnly))
	// Output: 23:59:59
}

func ExampleTimeDiffDays() {
	t1, _ := time.Parse("2006-01-02 15", "2023-09-01 1")
	t2, _ := time.Parse("2006-01-02 15", "2023-09-03 23")
	t3, _ := time.Parse("2006-01-02 15", "2023-08-28 23")

	fmt.Println(TimeDiffDays(t1, t2), TimeDiffDays(t3, t1))
	// Output: 2 3
}

func ExampleTimeIsLeapYear() {
	t1, _ := time.Parse("2006-01-02 15", "2000-09-01 1")
	t2, _ := time.Parse("2006-01-02 15", "1900-09-03 23")
	t3, _ := time.Parse("2006-01-02 15", "2024-08-28 23")

	fmt.Println(TimeIsLeapYear(t1), TimeIsLeapYear(t2), TimeIsLeapYear(t3))
	// Output: true false true
}

func ExampleTime2short() {
	fmt.Println(Time2short(10236 * time.Millisecond))
	fmt.Println(Time2short(1023600 * time.Millisecond))
	fmt.Println(Time2short(10236000 * time.Millisecond))

	// Output: 10.24s
	// 17.06m
	// 2.84h
}

func ExampleTime2now() {
	fmt.Println(Time2now(time.Unix(time.Now().Unix()-5, 0)))
	fmt.Println(Time2now(time.Unix(time.Now().Unix()-60*2-5, 0)))
	fmt.Println(Time2now(time.Unix(time.Now().Unix()-3600*2-5, 0)))
	fmt.Println(Time2now(time.Unix(time.Now().Unix()-86400*2-5, 0)))
	// Output: 刚刚
	// 2分钟
	// 2小时
	// 2天
}
