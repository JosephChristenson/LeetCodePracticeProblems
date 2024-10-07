package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Print(twoSums(nums, target))
}

// Brute Force answer
func twoSums(nums []int, target int) []int {

	for firstindex, element := range nums {
		for secondindex, secondelement := range nums {
			if firstindex != secondindex {
				if (element + secondelement) == target {
					result := []int{firstindex, secondindex}
					return result
				}
			}
		}
	}

	result := []int{0}
	return result
}

// package main

// type IntHeap []int

// func (h IntHeap) Len() int           { return len(h) }
// func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
// func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
// func (h *IntHeap) Pop() any {
//     old := *h
//     n := len(old)
//     x := old[n-1]
//     *h = old[0 : n-1]
//     return x
// }
// func (h *IntHeap) Push(x any) {
//     *h = append(*h, x.(int))
// }
