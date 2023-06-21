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

/*
BubbleSort
冒泡排序的时间复杂度为 O(n^2)，效率较低
*/
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

func BucketSort(nums []int, bucketSize int) []int {
	if len(nums) < 2 {
		return nums
	}

	// 计算最小值和最大值
	minValue, maxValue := nums[0], nums[0]
	for _, num := range nums {
		if num < minValue {
			minValue = num
		} else if num > maxValue {
			maxValue = num
		}
	}

	// 计算桶的数量
	bucketCount := (maxValue-minValue)/bucketSize + 1

	// 初始化桶
	buckets := make([][]int, bucketCount)
	for i := 0; i < bucketCount; i++ {
		buckets[i] = make([]int, 0)
	}

	// 将元素放入桶中
	for _, num := range nums {
		index := (num - minValue) / bucketSize
		buckets[index] = append(buckets[index], num)
	}

	// 对每个桶进行排序
	var sortedNums []int
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			sort.Ints(bucket)
			sortedNums = append(sortedNums, bucket...)
		}
	}
	return sortedNums
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
QuickSort
快速排序算法的效率与 pivot 的选择有关。
如果每次选择的 pivot 恰好是数组的中位数，那么算法的时间复杂度为 O(n log n)。
如果每次选择的 pivot 恰好是数组的最小值或最大值，那么算法的时间复杂度为 O(n^2)，效率较低。
*/
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
方法: ClimbStairs3
语言: Go
描述: 爬楼梯时，每次可以走一个或两个或三个台阶，总共有多少种爬法
参数: 这个方法接收以下参数:
x: 台阶总数
输出: 爬台阶的方法数量
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
