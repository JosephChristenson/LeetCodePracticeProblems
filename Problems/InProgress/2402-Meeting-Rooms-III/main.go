package main

import (
	"container/heap"
	"fmt"
)

// You are given an integer n. There are n rooms numbered from 0 to n - 1.

// You are given a 2D integer array meetings where meetings[i] = [starti, endi] means that a meeting will be held during the half-closed time interval [starti, endi). All the values of starti are unique.

// Meetings are allocated to rooms in the following manner:

// Each meeting will take place in the unused room with the lowest number.
// If there are no available rooms, the meeting will be delayed until a room becomes free. The delayed meeting should have the same duration as the original meeting.
// When a room becomes unused, meetings that have an earlier original start time should be given the room.
// Return the number of the room that held the most meetings. If there are multiple rooms, return the room with the lowest number.

// A half-closed interval [a, b) is the interval between a and b including a and not including b.

func main() {
	meetings := [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}}
	n := 2
	fmt.Println(mostBooked(n, meetings))
	meetings = [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}}
	n = 3
	fmt.Println(mostBooked(n, meetings))
	meetings = [][]int{{10, 11}, {2, 10}, {1, 17}, {9, 13}, {18, 20}}
	n = 2
	fmt.Println(mostBooked(n, meetings))
}

type schedule struct {
	start, end, delay, room int
}

func mostBooked(n int, meetings [][]int) int {
	h := &ScheduleHeap{}

	for _, val := range meetings {
		meet := schedule{
			start: val[0],
			end:   val[1],
			room:  101,
			delay: 0,
		}
		heap.Push(h, meet)
	}
	heap.Init(h)
	time := 0

	rooms := make([]bool, n)
	roomCount := make([]int, n)

	roomFree := &IntHeap{}

	heap.Init(roomFree)
	for n > 0 {
		heap.Push(roomFree, []int{0, n - 1})
		n--
	}

	for h.Len() > 0 {

		meet := heap.Pop(h).(schedule)

		if meet.start == time {
			if meet.end == time {
				rooms[meet.room] = false
			} else {
				room := heap.Pop(roomFree).([]int)

				if room[0] <= time {
					rooms[room[1]] = true
					meet.room = room[1]
					roomCount[room[1]]++
					meet.start = meet.end
					room[0] = meet.end
					meet.delay = 0

					heap.Push(roomFree, room)
					heap.Push(h, meet)
				} else {
					meet.start = meet.start + (time - room[0])
					meet.delay = meet.delay - (time - room[0])
					meet.end = meet.end + (time - room[0])
					heap.Push(h, meet)
				}
			}
		} else {
			heap.Push(h, meet)
			time++
		}
	}
	highest := 0
	roomNum := 0
	for index, val := range roomCount {
		if val > highest {
			highest = val
			roomNum = index
		}
	}

	return roomNum
}

type IntHeap [][]int

func (h IntHeap) Len() int      { return len(h) }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool {
	if h[i][0] < h[j][0] {
		return true
	} else if h[i][0] == h[j][0] && h[i][1] < h[j][1] {
		return true
	} else {
		return false
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
	*h = append(*h, x.([]int))
}

type ScheduleHeap []schedule

func (h ScheduleHeap) Len() int      { return len(h) }
func (h ScheduleHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h ScheduleHeap) Less(i, j int) bool {
	if h[i].start < h[j].start {
		return true // if start time happens first
	} else if h[i].start == h[j].start { // if they both happen at the same time
		if h[i].delay == 0 { // if this is a
			return true
		} else if h[i].delay < 0 && h[j].delay < 0 {
			if h[i].delay < h[j].delay {
				return true
			} else if h[i].delay == h[j].delay {
				if h[i].room < h[j].room {
					return true
				}
			}
		}
	}
	return false
} // with the greater than sign it sorts from highest to lowest
func (h *ScheduleHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *ScheduleHeap) Push(x any) {
	*h = append(*h, x.(schedule))
}
