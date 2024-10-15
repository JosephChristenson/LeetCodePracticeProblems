package main

import "fmt"

// Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

// You must write an algorithm with O(log n) runtime complexity.

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(nums, target))
	nums = []int{1, 3, 5, 6}
	target = 2
	fmt.Println(searchInsert(nums, target))
	nums = []int{1, 3, 5, 6}
	target = 7
	fmt.Println(searchInsert(nums, target))
	nums = []int{1, 3}
	target = 3
	fmt.Println(searchInsert(nums, target))
	nums = []int{3, 5, 7, 9, 10}
	target = 8
	fmt.Println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	if target <= nums[0] {
		return 0
	} else if target > nums[len(nums)-1] {
		return len(nums)
	}

	mid := low + (high-low)/2
	for low <= high {
		if low == high {
			if nums[low] < target { // could be equal if its a 1 index array
				return low
			} else {
				return low + 1
			}
		}
		mid = low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if target < nums[mid] {
		return mid
	} else {
		return mid + 1
	}

}
