package main

import (
	"fmt"
	"sort"
)

// You are given an integer array nums and an integer k.

// In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.

// Return the maximum number of operations you can perform on the array.

func main() {
	nums := []int{1, 2, 3, 4}
	k := 5
	fmt.Println(maxOperations(nums, k))
	nums = []int{3, 1, 3, 4, 3}
	k = 6
	fmt.Println(maxOperations(nums, k))
}

func maxOperations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var m map[int]int
	m = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	first, last := 0, len(nums)-1
	result := 0
	for first < last {
		if nums[first]+nums[last] == k {
			result++
			first++
			last--
		} else if nums[first]+nums[last-1] == k {
			last--
		} else if nums[first+1]+nums[last] == k {
			first++
		} else {
			first++
			last--
		}
	}
	return result
}
