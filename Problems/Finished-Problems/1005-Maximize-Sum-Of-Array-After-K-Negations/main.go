package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums := []int{4, 2, 3}
	k := 1
	fmt.Println(largestSumAfterKNegations(nums, k))
	nums = []int{3, -1, 0, 2}
	k = 3
	fmt.Println(largestSumAfterKNegations(nums, k))
	nums = []int{2, -3, -1, 5, -4}
	k = 2
	fmt.Println(largestSumAfterKNegations(nums, k))
}

func largestSumAfterKNegations(nums []int, k int) int {
	list := &IntHeap{}

	for _, val := range nums {
		heap.Push(list, val)
	}
	lowestValue := 0

	for k > 0 {
		value := heap.Pop(list).(int)
		if value < 0 {
			value = value * -1
			heap.Push(list, value)
		} else {
			lowestValue = value
			heap.Push(list, value)
			break
		}
		k--
	}
	sum := 0

	for list.Len() > 0 {
		sum += heap.Pop(list).(int)
	}
	if k%2 == 1 {
		sum = sum - (2 * lowestValue)
	}
	return sum
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
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
