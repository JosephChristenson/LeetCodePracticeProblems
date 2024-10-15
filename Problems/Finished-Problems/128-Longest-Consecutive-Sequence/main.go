package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
	nums = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(nums))

}

func longestConsecutive(nums []int) int {

	var m map[int]int
	m = make(map[int]int)

	result := 0
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for _, val := range nums {
		if m[val-1] != 0 {
			m[val] = m[val-1] + 1
		} else {
			m[val] = 1
		}
		if m[val] > result {
			result = m[val]
		}
	}
	return result
}
