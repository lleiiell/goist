package goist_test

import (
	"fmt"
	"github.com/lleiiell/goist"
	"testing"
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
