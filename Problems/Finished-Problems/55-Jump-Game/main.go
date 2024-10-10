package main

import "fmt"

// You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

// Return true if you can reach the last index, or false otherwise.

func main() {
	nums := []int{2, 3, 1, 1, 4}

	fmt.Println(canJump(nums))
	nums = []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
	nums = []int{1, 1, 1, 0}
	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	goal := len(nums) - 1

	result := false
	curr := 0

	furthestIndex := nums[0]

	for furthestIndex >= curr {
		if curr >= goal {
			return true
		}
		if furthestIndex < curr+nums[curr] {
			furthestIndex = curr + nums[curr]
		}

		curr++
	}

	return result
}
