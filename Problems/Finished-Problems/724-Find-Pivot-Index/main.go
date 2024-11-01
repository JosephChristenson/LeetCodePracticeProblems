package main

import "fmt"

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	// fmt.Println(pivotIndex(nums))
	nums = []int{-1, -1, -1, -1, -1, 0}

	fmt.Println(pivotIndex(nums))
}

func pivotIndex(nums []int) int {
	total := 0
	for index := 0; index < len(nums); index++ {
		total += nums[index]
	}

	left := 0
	for index := 0; index < len(nums); index++ {
		total = total - nums[index]
		if left == total {
			return index
		}
		left = left + nums[index]
	}
	return -1
}
