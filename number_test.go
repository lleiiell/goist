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
