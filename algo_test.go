package goist

import "fmt"

func ExampleBinarySearch() {
	nums := []int{22, 11, 33, 100}
	target1 := 1
	target2 := 11

	i1 := BinarySearch(nums, target1)
	i2 := BinarySearch(nums, target2)

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

func ExampleBubbleSort() {
	nums := []int{22, 11, 33, 55, 44, 66}

	nums2 := BubbleSort(nums)
	fmt.Println(nums2)

	// Output: [11 22 33 44 55 66]
}

func ExampleBucketSort() {
	nums := []int{22, 11, 33, 55, 44, 66, 7}

	nums2 := BucketSort(nums, 2)
	fmt.Println(nums2)

	// Output: [7 11 22 33 44 55 66]
}

func ExampleQuickSort() {
	nums := []int{22, 11, 33, 55, 44, 66}

	nums2 := QuickSort(nums)
	fmt.Println(nums2)

	// Output: [11 22 33 44 55 66]
}

func ExampleInsertionSort() {
	nums := []int{22, 11, 33, 55, 44, 66}
	fmt.Println(nums)

	InsertionSort(nums)
	fmt.Println(nums)
	// Output: [22 11 33 55 44 66]
	// [11 22 33 44 55 66]

}
