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

	firstHalf := nums[:k]
	secondHalf := nums[k:]

	nums = append(secondHalf, firstHalf...)

}
