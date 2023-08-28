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
