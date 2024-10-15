package main

import "fmt"

// Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:

// answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
// answer[1] is a list of all distinct integers in nums2 which are not present in nums1.
// Note that the integers in the lists may be returned in any order.

func main() {
	nums1, nums2 := []int{1, 2, 3}, []int{2, 4, 6}
	fmt.Println(findDifference(nums1, nums2))
	nums1, nums2 = []int{1, 2, 3, 3}, []int{1, 1, 2, 2}
	fmt.Println(findDifference(nums1, nums2))
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	var map1 map[int]bool
	map1 = make(map[int]bool)
	var map2 map[int]bool
	map2 = make(map[int]bool)

	for _, val := range nums1 {
		map1[val] = true
	}
	for _, val := range nums2 {
		map2[val] = true
	}
	diff1, diff2 := []int{}, []int{}

	for key := range map1 {
		if map1[key] && !map2[key] {
			diff1 = append(diff1, key)
		} else if map1[key] == false && map2[key] {
			diff2 = append(diff2, key)
		}
	}
	for key := range map2 {
		if map1[key] && !map2[key] {
			diff1 = append(diff1, key)
		} else if map1[key] == false && map2[key] {
			diff2 = append(diff2, key)
		}
	}
	return [][]int{diff1, diff2}
}
