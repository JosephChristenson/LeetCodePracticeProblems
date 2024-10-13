package main

import (
	"container/heap"
	"fmt"
)

// You are given an array of CPU tasks, each labeled with a letter from A to Z, and a number n. Each CPU interval can be idle or allow the completion of one task. Tasks can be completed in any order, but there's a constraint: there has to be a gap of at least n intervals between two tasks with the same label.

// Return the minimum number of CPU intervals required to complete all tasks.

func main() {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n := 2
	fmt.Println(leastInterval(tasks, n))
	tasks = []byte{'B', 'C', 'D', 'A', 'A', 'A', 'A', 'G'}
	n = 1
	fmt.Println(leastInterval(tasks, n))
}

func leastInterval(tasks []byte, n int) int {

	var m map[byte]int
	m = make(map[byte]int)

	for _, val := range tasks {
		m[val] = m[val] + 1
	}

	h := &PriorityHeap{}
	i := &PriorityHeap{}

	for index, val := range m {
		task := priority{
			task:     index,
			count:    val,
			cooldown: 0,
		}
		heap.Push(h, task)
	}
	heap.Init(h)
	heap.Init(i)

	result := 0
	for h.Len() > 0 {
		count := h.Len()
		for count > 0 {
			test := priority{}

			test = heap.Pop(h).(priority)
			if test.cooldown < result {

				test.count = test.count - 1

				fmt.Println(test.task)
				if test.count != 0 {
					test.cooldown = n + result
					heap.Push(i, test)
				}
				break
			} else {
				heap.Push(i, test)
			}
			count--
		}
		for i.Len() > 0 {
			heap.Push(h, heap.Pop(i))
		}
		result++
	}
	return result - 1
}

type priority struct {
	task     byte
	count    int
	cooldown int
}

type PriorityHeap []priority

func (h PriorityHeap) Len() int           { return len(h) }
func (h PriorityHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h PriorityHeap) Less(i, j int) bool { return h[i].count > h[j].count }
func (h *PriorityHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *PriorityHeap) Push(x any) {
	*h = append(*h, x.(priority))
}

// Provided by Jake
