package main

import "fmt"

// Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(increasingTriplet(nums))
	nums = []int{5, 4, 3, 2, 1}
	fmt.Println(increasingTriplet(nums))
	nums = []int{2, 1, 5, 0, 4, 6}
	fmt.Println(increasingTriplet(nums))
}

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	lowestI := 1 << 31
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] < lowestI {
			lowestI = nums[i]
			lowestJ := 1 << 31
			for j := i + 1; j < len(nums)-1; j++ {
				if nums[i] < nums[j] {
					if nums[j] < lowestJ {
						lowestJ = nums[j]
						highestK := lowestJ
						for k := j + 1; k < len(nums); k++ {
							if nums[k] > highestK {
								highestK = nums[k]
								if nums[j] < nums[k] {
									return true
								}
							}
						}
					}
				}
			}
		}
	}
	return false
}
