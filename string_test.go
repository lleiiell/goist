package goist_test

import (
	"fmt"
	"github.com/lleiiell/goist"
	"testing"
)

func TestUuidSimple(t *testing.T) {
	uid := goist.UuidSimple()
	if len(uid) < 22 {
		t.Error(len(uid), uid)
		return
	}
	fmt.Println(uid)
}

func TestRandomStr(t *testing.T) {
	str := goist.RandStr(5)
	if len(str) != 5 {
		t.Error(str)
		return
	}
	str1 := goist.RandStr(5)
	if len(str1) != 5 {
		t.Error(str1)
		return
	}

	fmt.Println(str)
	fmt.Println(str1)
}

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
