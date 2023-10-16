package goist

import "fmt"

func ExampleIsNumber() {
	fmt.Println(IsNumber("a"), IsNumber([]interface{}{"a", 1}))

	fmt.Println(IsNumber(1), IsNumber(-1), IsNumber(1.23),
		IsNumber(complex(1, 2)), IsNumber(1+2i))

	// Output: false false
	// true true true true true
}
