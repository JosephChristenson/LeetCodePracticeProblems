package main

import "container/heap"

// The median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value, and the median is the mean of the two middle values.

// For example, for arr = [2,3,4], the median is 3.
// For example, for arr = [2,3], the median is (2 + 3) / 2 = 2.5.
// Implement the MedianFinder class:

// MedianFinder() initializes the MedianFinder object.
// void addNum(int num) adds the integer num from the data stream to the data structure.
// double findMedian() returns the median of all elements so far. Answers within 10-5 of the actual answer will be accepted.

func main() {

	commands := []string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"}
	num := []int{{}, {1}, {2}, {}, {3}, {}}

}

type MedianFinder struct {
}

func Constructor() MedianFinder {
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

func (this *MedianFinder) AddNum(num int) {

}

func (this *MedianFinder) FindMedian() float64 {

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

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
