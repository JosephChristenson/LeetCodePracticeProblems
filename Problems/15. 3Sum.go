package main

import "fmt"

// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

// Notice that the solution set must not contain duplicate triplets.

func main() {
	num := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(num))
	num = []int{0, 1, 1}
	fmt.Println(threeSum(num))
	num = []int{0, 0, 0}
	fmt.Println(threeSum(num))
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

}
