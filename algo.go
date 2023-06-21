package goist

import "sort"

func BinarySearch(nums []int, target int) bool {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func BubbleSort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func FibonacciRecursion(n int) int {
	if n < 2 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}

func FibonacciDP(n int) int {
	if n < 2 {
		return n
	}
	f1, f2, f3 := 0, 1, 0
	for i := 2; i <= n; i++ {
		f3 = f1 + f2
		f1 = f2
		f2 = f3
	}
	return f3
}

func QuickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	pivot := nums[0]
	var left, right []int
	for _, num := range nums[1:] {
		if num < pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}
	left = QuickSort(left)
	right = QuickSort(right)
	return append(append(left, pivot), right...)
}

/*
ClimbStairs3
Function: ClimbStairs3
Language: Go
Description: Given a number of stairs, the task is to find the total number of ways to climb the stairs where at each step, the climber can either climb 1, 2, or 3 steps.
Parameters:
x: Total number of stairs
Output: Number of ways to climb the stairs
*/
func ClimbStairs3(x int) int {
	if x <= 2 {
		return x
	} else if x == 3 {
		return 4
	}
	f1, f2, f3, f4 := 1, 2, 4, 0
	for i := 4; i <= x; i++ {
		f4 = f1 + f2 + f3
		f1 = f2
		f2 = f3
		f3 = f4
	}
	return f4
}
