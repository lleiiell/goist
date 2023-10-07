package goist

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleSliceIntDiff2Right() {

	s1 := []int{1, 2, 3}
	s2 := []int{2, 4, 5}

	fmt.Println(SliceIntDiff2Right(s1, s2))

	// Output: [1 3]
}

func ExampleToInterfaceSlice() {

	s1 := []int{1, 2, 3}
	is1 := ToInterfaceSlice(s1)
	fmt.Println(reflect.ValueOf(is1).Type(), reflect.ValueOf(s1).Type())

	sl2 := []string{"a", "b", "c"}
	isl2 := ToInterfaceSlice(sl2)
	fmt.Println(reflect.ValueOf(isl2).Type(), reflect.ValueOf(sl2).Type())

	// Output: []interface {} []int
	// []interface {} []string
}

func TestSliceRand(t *testing.T) {

	fmt.Println(sliceRand(5, 22))
}
