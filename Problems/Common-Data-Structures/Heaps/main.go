package main

import (
	"container/heap"
	"fmt"
)

// Problem 215 implementation. given an array returns the k value.
func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(findKthLargest(nums, k))
	nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k = 4
	fmt.Println(findKthLargest(nums, k))
}

func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	for _, val := range nums {
		heap.Push(h, val)
	}
	heap.Init(h)

	for k-1 > 0 {
		heap.Pop(h)
		k--
	}
	return heap.Pop(h).(int)
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] } // with the greater than sign it sorts from highest to lowest
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

// Provided by Jake
