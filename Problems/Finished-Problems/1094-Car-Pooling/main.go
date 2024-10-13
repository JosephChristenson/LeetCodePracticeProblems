package main

import (
	"container/heap"
	"fmt"
)

// There is a car with capacity empty seats. The vehicle only drives east (i.e., it cannot turn around and drive west).

// You are given the integer capacity and an array trips where trips[i] = [numPassengersi, fromi, toi] indicates that the ith trip has numPassengersi passengers and the locations to pick them up and drop them off are fromi and toi respectively. The locations are given as the number of kilometers due east from the car's initial location.

// Return true if it is possible to pick up and drop off all passengers for all the given trips, or false otherwise.

func main() {
	input := [][]int{{2, 1, 5}, {3, 3, 7}}
	capacity := 4
	fmt.Println(carPooling(input, capacity))
	input = [][]int{{2, 1, 5}, {3, 3, 7}}
	capacity = 5
	fmt.Println(carPooling(input, capacity))
}

func carPooling(trips [][]int, capacity int) bool {
	h := &IntHeap{}
	for _, val := range trips {
		car := toFrom{
			count: val[0],
			to:    val[2],
			from:  val[1],
		}
		heap.Push(h, car)
	}
	heap.Init(h)

	count := 0
	for h.Len() > 0 {
		car := heap.Pop(h).(toFrom)

		if car.to <= car.from {
			count = count - car.count
			if count < 0 {
				count = 0
			}
		} else {
			count = count + car.count
			car.from = car.to
			if count > capacity {
				return false
			}
			heap.Push(h, car)
		}
	}
	return true

}

type toFrom struct {
	count, to, from int
}

type IntHeap []toFrom

func (h IntHeap) Len() int      { return len(h) }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool {
	if h[i].from < h[j].from {
		return true
	} else {
		if h[i].to == h[i].from && h[i].from == h[j].from { // if it has already been visited before
			return true
		} else {
			return false
		}
	}
}
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(toFrom))
}
