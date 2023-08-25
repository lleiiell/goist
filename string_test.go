package goist_test

import (
	"fmt"
	"github.com/lleiiell/goist"
)

func ExampleStrIsEmpty() {
	s1 := ""
	s2 := "a bc"
	s3 := "  "
	s4 := "\n \t"
	fmt.Println(goist.IsStrEmpty(s1), goist.IsStrEmpty(s2), goist.IsStrEmpty(s3), goist.IsStrEmpty(s4))
	// Output: true false true true
}

func ExampleMD5() {
	fmt.Println(goist.MD5("hello"))
	// Output: 5d41402abc4b2a76b9719d911017c592
}
