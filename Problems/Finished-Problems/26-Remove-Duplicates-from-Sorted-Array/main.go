package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	index := len(nums) - 2

	for index >= 0 {
		if nums[index] == nums[index+1] {
			nums = remove(nums, index)
		}
		index--
	}
	return len(nums)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
