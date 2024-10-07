package main

import "fmt"

func main() {
	nums := []int{3, 2, 3}
	fmt.Print(majorityElement(nums))
	nums = []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Print(majorityElement(nums))
}
func majorityElement(nums []int) int {
	var m map[int]int
	m = make(map[int]int)
	goal := len(nums)/2 + len(nums)%2

	for _, val := range nums {
		m[val] = m[val] + 1
		if m[val] >= goal {
			return val
		}
	}
	return 0
}
