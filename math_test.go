package goist_test

import (
	"fmt"
	"github.com/lleiiell/goist"
	"testing"
)

func TestIdMask(t *testing.T) {
	fmt.Println(goist.IdMask(1))
	fmt.Println(goist.IdMask(999))
	fmt.Println(goist.IdMask(goist.Rand().Int63()))
}

func ExampleIdUnmask() {
	fmt.Println(goist.IdUnmask("1clb"))
	fmt.Println(goist.IdUnmask("rr5fk"))
	fmt.Println(goist.IdUnmask("1djlwzxe5tj1a4qw"))
	// Output: 1
	// 999
	// 6521156348675224222
}

func ExampleFloatDivide() {
	a := uint(10)
	b := uint(3)

	fmt.Println(float64(a / b))
	fmt.Println(float32(a) / float32(b))
	fmt.Println(float64(a) / float64(b))
	// Output: 3
	// 3.3333333
	// 3.3333333333333335
}
