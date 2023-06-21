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
