package goist

import "fmt"

func ExampleBinarySearch() {
	arr := []int{22, 11, 33, 100}
	target1 := 1
	target2 := 11

	i1 := BinarySearch(arr, target1)
	i2 := BinarySearch(arr, target2)

	fmt.Println(i1, i2)

	// Output: false true

}

func ExampleClimbStairs3() {
	n := ClimbStairs3(22)

	fmt.Println(n)

	// Output: 410744
}

func ExampleFibonacciRecursion() {
	n1 := FibonacciRecursion(22)
	n2 := FibonacciDP(22)

	fmt.Println(n1, n1 == n2)

	// Output: 17711 true
}

func ExampleQuickSort() {
	arr := []int{22, 11, 33, 55, 44, 66}

	arr2 := QuickSort(arr)
	fmt.Println(arr2)

	// Output: [11 22 33 44 55 66]
}
