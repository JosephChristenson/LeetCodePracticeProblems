package main

// You are given two integer arrays nums1 and nums2 sorted in non-decreasing order and an integer k.

// Define a pair (u, v) which consists of one element from the first array and one element from the second array.

// Return the k pairs (u1, v1), (u2, v2), ..., (uk, vk) with the smallest sums.

import (
	"fmt"
	"sort"
)

func main() {
	nums1, nums2 := []int{1, 7, 11}, []int{2, 4, 6}
	k := 3
	fmt.Println(kSmallestPairs(nums1, nums2, k))
	nums1, nums2 = []int{1, 1, 2}, []int{1, 2, 3}
	k = 2
	fmt.Println(kSmallestPairs(nums1, nums2, k))
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	allPairs := [][]int{}
	for _, val1 := range nums1 {
		for _, val2 := range nums2 {
			allPairs = append(allPairs, []int{val1, val2})
		}
	}
	sort.Slice(allPairs, func(i, j int) bool {
		return allPairs[i][0]+allPairs[i][1] < allPairs[j][0]+allPairs[j][1]
	})
	return allPairs[:k]
}
