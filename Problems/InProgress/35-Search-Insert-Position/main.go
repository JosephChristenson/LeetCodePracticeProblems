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
}

func searchInsert(nums []int, target int) int {

	if nums[0] >= target {
		return 0
	} else if nums[len(nums)-1] < target {
		return len(nums)
	}

	half := len(nums) / 2
	index := len(nums) - half + half%2

	for half > 1 {
		if nums[index] == target {
			return index
		}
		if nums[index] > target {
			index = index - half/2 + half%2
			half = half / 2
		} else {
			index = index + half/2 + half%2
			half = half / 2
		}
	}
	if index == len(nums) {
		return len(nums) - 1
	} else {
		return index
	}
}
