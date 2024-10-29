package main

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

// Note that you must do this in-place without making a copy of the array.\

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	nums = []int{0}
	moveZeroes(nums)
}

func moveZeroes(nums []int) {
	index := len(nums) - 1
	for index > 0 {
		if nums[index] != 0 {
			if nums[index-1] == 0 {
				nums[index], nums[index-1] = nums[index-1], nums[index]
				index++
				if index == len(nums) {
					index = len(nums) - 1
				}
			} else {
				index--
			}
		} else {
			index--
		}
	}
}
