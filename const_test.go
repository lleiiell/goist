package goist

import (
	"fmt"
)

func ExampleConstIota() {
	const (
		x = iota
		y
		z = "zz"
		k
		p = iota
		_
		a1 = iota + 10
		a2
	)
	const (
		a3 = iota << 1
		a4
		a5
	)

	fmt.Println(x, y, z, k, p)
	fmt.Println(a1, a2, a3, a4, a5)

	// Output: 0 1 zz zz 4
	// 16 17 0 2 4
}
