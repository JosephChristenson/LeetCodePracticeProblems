package main

import (
	"container/heap"
	"fmt"
)

// There are n rooms labeled from 0 to n - 1 and all the rooms are locked except for room 0. Your goal is to visit all the rooms. However, you cannot enter a locked room without having its key.

// When you visit a room, you may find a set of distinct keys in it. Each key has a number on it, denoting which room it unlocks, and you can take all of them with you to unlock the other rooms.

// Given an array rooms where rooms[i] is the set of keys that you can obtain if you visited room i, return true if you can visit all the rooms, or false otherwise.

func main() {
	rooms := [][]int{{1}, {2}, {3}, {}}
	fmt.Println(canVisitAllRooms(rooms))
	rooms = [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}
	fmt.Println(canVisitAllRooms(rooms))
}

func canVisitAllRooms(rooms [][]int) bool {
	var m map[int][]int
	m = make(map[int][]int)
	m[0] = rooms[0]
	h := &IntHeap{}
	heap.Push(h, 0)
	for h.Len() > 0 {
		if len(m) == len(rooms) {
			return true
		}
		for _, val := range m[heap.Pop(h).(int)] {
			if _, hit := m[val]; !hit { // only adds new rooms to search when the key is found the first time
				m[val] = rooms[val]
				heap.Push(h, val)
			}
		}
	}
	return false
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
