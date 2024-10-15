package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	k := 3
	fmt.Println(containsNearbyDuplicate(nums, k))
	nums = []int{1, 0, 1, 1}
	k = 1
	fmt.Println(containsNearbyDuplicate(nums, k))
	nums = []int{1, 2, 3, 1, 2, 3}
	k = 2
	fmt.Println(containsNearbyDuplicate(nums, k))
}

func containsNearbyDuplicate(nums []int, k int) bool {

	var m map[int]int
	m = make(map[int]int)

	for index, val := range nums {

		if index+1-k > 0 {
			if m[val] >= index+1-k {
				return true
			} else {
				m[val] = index + 1 // keep track of index position offset by one
			}
		} else if m[val] > 0 {
			return true
		} else {
			m[val] = index + 1 // keep track of index position offset by one
		}
	}
	return false

}
