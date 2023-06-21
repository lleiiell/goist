package goist

import "fmt"

func ExampleErr_ClimbStairs3() {
	n := ClimbStairs3(22)

	fmt.Println(n)

	// Output: 410744
}

func ExampleErr_Fibonacci() {
	n1 := FibonacciRecursion(22)
	n2 := FibonacciDP(22)

	fmt.Println(n1 == n2)

	// Output: true
}
