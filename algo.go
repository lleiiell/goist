package goist

import (
	"sort"
	"sync"
)

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
The time complexity of bubble sorting is O(n^2), which is less efficient
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

	// Calculate min and max values
	minValue, maxValue := nums[0], nums[0]
	for _, num := range nums {
		if num < minValue {
			minValue = num
		} else if num > maxValue {
			maxValue = num
		}
	}

	// Calculate the number of buckets
	bucketCount := (maxValue-minValue)/bucketSize + 1

	// Initialize bucket
	buckets := make([][]int, bucketCount)
	for i := 0; i < bucketCount; i++ {
		buckets[i] = make([]int, 0)
	}

	// put element into bucket
	for _, num := range nums {
		index := (num - minValue) / bucketSize
		buckets[index] = append(buckets[index], num)
	}

	// Sort each bucket
	var sortedNums []int
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			sort.Ints(bucket)
			sortedNums = append(sortedNums, bucket...)
		}
	}
	return sortedNums
}

func InsertionSort(arr []int) {
	for i := range arr {
		preIndex := i - 1
		current := arr[i]
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex -= 1
		}
		arr[preIndex+1] = current
	}
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
The efficiency of the quicksort algorithm is related to the choice of pivot.
If the pivot selected each time is exactly the median of the array, then the time complexity of the algorithm is O(n log n).
If the pivot selected each time happens to be the minimum or maximum value of the array, then the time complexity of the algorithm is O(n^2), and the efficiency is low.
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
ClimbStairs3 chatgpt generated code
Method: ClimbStairs3
Language: Go
Description: When climbing stairs, you can walk one, two or three steps at a time. How many ways are there in total?
Parameters: This method receives the following parameters:
x: total number of steps
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

type RrItem struct {
	Id int64
}

type RoundRobin struct {
	items []RrItem
	next  int
	lock  sync.Mutex
}

func (r *RoundRobin) Add(item RrItem) {
	r.lock.Lock()
	defer r.lock.Unlock()

	for _, v := range r.items {
		if v.Id == item.Id {
			return
		}
	}

	if r.next == 0 {
		r.items = append(r.items, item)
		return
	}
	r.items = append(r.items[:r.next+1], r.items[r.next:]...)
	r.items[r.next] = item
	r.next = r.next + 1

}

func (r *RoundRobin) Remove(item RrItem) {
	r.lock.Lock()
	defer r.lock.Unlock()

	for k, v := range r.items {
		if v.Id == item.Id {
			r.items = append(r.items[:k], r.items[k+1:]...)
			// 删除的元素位于下一个元素之前，下一个元素-1
			// 删除的元素为下一个元素，判断调头
			// 删除的元素位于下一个元素之后，不用处理
			if r.next > k && r.next > 0 {
				r.next = r.next - 1
			} else if r.next == k && r.next == len(r.items) {
				r.next = 0
			}
			return
		}
	}

}

func (r *RoundRobin) Next() RrItem {
	r.lock.Lock()
	defer r.lock.Unlock()

	if len(r.items) == 0 {
		return RrItem{}
	}

	item := r.items[r.next]

	if r.next == len(r.items)-1 {
		r.next = 0
	} else {
		r.next = r.next + 1
	}

	return item
}
