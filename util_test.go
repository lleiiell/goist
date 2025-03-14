package goist_test

import (
	"errors"
	"fmt"
	"github.com/lleiiell/goist"
	"reflect"
	"testing"
	"time"
)

func TestContains(t *testing.T) {
	sli := []int{1, 2, 3}

	if !goist.Contains(sli, 2) {
		t.Error("should Contains")
		return
	}

	if goist.Contains(sli, 0) {
		t.Error("should Contains")
		return
	}

	if goist.Contains(sli, "2") {
		t.Error("shouldn't Contains")
		return
	}
	if !goist.Contains([]string{"a", "b", "c"}, "c") {
		t.Error("shouldn Contains")
		return
	}
}

/*
first Int63:  4938851740737359089
second Int63:  415379671487512047
*/
func TestRand(t *testing.T) {
	fmt.Println("first Int63: ", goist.Rand().Int63())
	fmt.Println("second Int63: ", goist.Rand().Int63())
}

func TestUnique(t *testing.T) {
	sla := []string{"a", "b", "b", "b", "c"}

	goist.Unique(&sla)

	slaNew := []string{"a", "b", "c"}
	if !reflect.DeepEqual(sla, slaNew) {
		t.Error("sla not equal", slaNew)
		return
	}

	ila := []int{1, 2, 3, 3, 3}
	goist.Unique(&ila)

	ilaNew := []int{1, 2, 3}
	if !reflect.DeepEqual(ila, ilaNew) {
		t.Error("ila not equal", ila, ilaNew)
		return

	}

	ilb := []int{1, 2, 3, 4, 5}
	goist.Unique(&ilb)

	ilbNew := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(ilb, ilbNew) {
		t.Error("ila not equal", ilb, ilbNew)
		return

	}

	itl := []interface{}{1, "1", 2, 2, "a", "b", "b"}
	goist.Unique(&itl)
	itlNew := []interface{}{1, "1", 2, "a", "b"}
	if !reflect.DeepEqual(itl, itlNew) {
		t.Error("itl not equal", itl, itlNew)
		return
	}
	fmt.Println("itl", itl)

}

func TestRetry(t *testing.T) {

	i := 3
	err := goist.Retry(10, 1*time.Second, func() error {
		rd := goist.Rand().Intn(6)
		fmt.Println(fmt.Sprintf("i is %d, rand number is %d,  time is %v", i, rd, time.Now()))
		i--
		if rd == 1 {
			return nil
		} else {
			return errors.New(fmt.Sprintf("%d", rd))
		}
	})

	fmt.Println(err)
}

func ExampleRateLimiter() {
	rl := goist.RateLimiter(10, 1*time.Second)

	for i := 1; i <= 10; i++ {
		if rl() {
			fmt.Printf("Request %d: Allowed\n", i)
		} else {
			fmt.Printf("Request %d: Rate limit exceeded\n", i)
		}
		time.Sleep(2 * time.Second) // Simulate time between requests
	}

	// Output: Request 1: Allowed
	// Request 2: Allowed
	// Request 3: Allowed
	// Request 4: Allowed
	// Request 5: Allowed
	// Request 6: Allowed
	// Request 7: Allowed
	// Request 8: Allowed
	// Request 9: Allowed
	// Request 10: Allowed
}
