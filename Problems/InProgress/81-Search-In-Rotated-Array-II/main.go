package main

import "fmt"

// There is an integer array nums sorted in non-decreasing order (not necessarily with distinct values).

// Before being passed to your function, nums is rotated at an unknown pivot index k (0 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become [4,5,6,6,7,0,1,2,4,4].

// Given the array nums after the rotation and an integer target, return true if target is in nums, or false if it is not in nums.

// You must decrease the overall operation steps as much as possible.

func main() {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	target := 0
	fmt.Println(search(nums, target))
	nums = []int{2, 5, 6, 0, 0, 1, 2}
	target = 3
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) bool {
	front, back := 0, len(nums)-1
	interval := len(nums)/10 + 1
	for back >= front && back > 0 && front < len(nums)-1 {
		if nums[front] < target {
			front += interval
		} else if nums[front] == target {
			return true
		}
		if nums[back] > target {
			back = back - interval
		} else if nums[back] == target {
			return true
		}
	}
	for back < front {
		if nums[back] != target {
			back++
		} else {
			return true
		}
	}
	return false
}

func findFastestStep(length int) int {
	result := 0
	count := 1
	for result < length {
		result = result + count
		count++
	}
	return count
}
