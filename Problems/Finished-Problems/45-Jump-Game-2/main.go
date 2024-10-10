package main

import (
	"container/heap"
	"fmt"
)

// You are given a 0-indexed array of integers nums of length n. You are initially positioned at nums[0].

// Each element nums[i] represents the maximum length of a forward jump from index i. In other words, if you are at nums[i], you can jump to any nums[i + j] where:

// 0 <= j <= nums[i] and
// i + j < n
// Return the minimum number of jumps to reach nums[n - 1]. The test cases are generated such that you can reach nums[n - 1].

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums))
	nums = []int{2, 3, 0, 1, 4}
	fmt.Println(jump(nums))
}

type step struct {
	index, jump int
}

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	h := &stepHeap{}
	heap.Init(h)
	heap.Push(h, step{0, 0}) // start at index 0, we have jumped 0 times

	var m map[int]int
	m = make(map[int]int)

	for h.Len() > 0 {
		curr := heap.Pop(h).(step)
		distance := nums[curr.index] // find how far it can jump

		if curr.index+nums[curr.index] >= len(nums)-1 { // did we reach the goal
			return curr.jump + 1 // 1 last jump
		}
		for distance > 0 {
			if m[curr.index+distance] == 0 { // if we haven't been to this index add it to the list
				heap.Push(h, step{curr.index + distance, curr.jump + 1})
				m[curr.index+distance] = curr.jump + 1
			}
			distance--
		}
	}

	return -1 // should not reach this point

}

type stepHeap []step

func (h stepHeap) Len() int      { return len(h) }
func (h stepHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h stepHeap) Less(i, j int) bool {
	if h[i].jump < h[j].jump {
		return true
	} else if h[i].jump == h[j].jump && h[i].index > h[j].index {
		return true
	} else {
		return false
	}
}

func (h *stepHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *stepHeap) Push(x any) {
	*h = append(*h, x.(step))
}
