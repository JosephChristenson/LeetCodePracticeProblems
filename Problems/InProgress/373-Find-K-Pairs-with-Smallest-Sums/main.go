package main

// You are given two integer arrays nums1 and nums2 sorted in non-decreasing order and an integer k.

// Define a pair (u, v) which consists of one element from the first array and one element from the second array.

// Return the k pairs (u1, v1), (u2, v2), ..., (uk, vk) with the smallest sums.

import (
	"container/heap"
	"fmt"
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
	h := &IntHeap{}
	i := &IntHeap{}

	for _, val := range nums1 {
		heap.Push(h, val)
	}
	for _, val := range nums2 {
		heap.Push(i, val)
	}
	heap.Init(h)
	heap.Init(i)

	result := [][]int{}
	for k-1 >= 0 {
		if storeI != 0 {
			storeH = heap.Pop(h)
		}

		val := heap.Pop(h).([]int)
		result = append(result, val)
		k--
	}
	return result

}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // with the greater than sign it sorts from highest to lowest
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}
