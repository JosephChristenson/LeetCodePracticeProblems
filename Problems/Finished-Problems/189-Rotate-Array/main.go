package main

import "fmt"

// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)

	nums = []int{-1, -100, 3, 99}
	k = 2
	rotate(nums, k)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	// if rotation is greater than the size of array, reduce size of rotation
	for k > len(nums) {
		k = k - len(nums)
	}

	var m map[int]int
	m = make(map[int]int)
	for index, value := range nums {
		if (index + k) >= len(nums) {
			m[index+k-len(nums)] = value
		} else {
			m[index+k] = value
		}
	}

	for index, _ := range nums {
		nums[index] = m[index]
	}
}
