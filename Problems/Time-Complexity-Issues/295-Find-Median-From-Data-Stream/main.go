package main

import (
	"fmt"
)

// The median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value, and the median is the mean of the two middle values.

// For example, for arr = [2,3,4], the median is 3.
// For example, for arr = [2,3], the median is (2 + 3) / 2 = 2.5.
// Implement the MedianFinder class:

// MedianFinder() initializes the MedianFinder object.
// void addNum(int num) adds the integer num from the data stream to the data structure.
// double findMedian() returns the median of all elements so far. Answers within 10-5 of the actual answer will be accepted.

func main() {

	commands := []string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"}
	num := [][]int{{}, {1}, {2}, {}, {3}, {}}
	// runFunc(commands, num)
	// commands = []string{"MedianFinder", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian"}
	// num = [][]int{{}, {-1}, {}, {-2}, {}, {-3}, {}, {-4}, {}, {-5}, {}}
	// runFunc(commands, num)
	commands = []string{"MedianFinder", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian"}
	num = [][]int{{}, {1}, {}, {2}, {}, {1}, {}, {4}, {}, {5}, {}, {6}, {}, {7}, {}, {8}, {}, {9}, {}, {10}, {}}
	runFunc(commands, num)

}

func runFunc(input []string, values [][]int) {
	var median MedianFinder

	for i, s := range input {
		if s == "MedianFinder" {
			median = Constructor()
		} else if s == "addNum" {
			median.AddNum(values[i][0])
		} else if s == "findMedian" {
			fmt.Println(median.FindMedian())
		}
	}
}

//Linked list attempt. Problem with sorting function
// type MedianFinder struct {
// 	head *LinkedList
// }

// func Constructor() MedianFinder {
// 	var result MedianFinder
// 	return result
// }

// func (this *MedianFinder) AddNum(num int) {
// 	if this.head == nil {
// 		this.head = &LinkedList{}
// 		this.head.insertAtHead(num)
// 	} else {
// 		if this.head.head.data >= num {
// 			this.head.insertAtHead(num)
// 		} else {
// 			tempHead := this.head
// 			for index := 0; index < this.head.length; index++ {
// 				if tempHead.head.data >= num {
// 					this.head.insert(index, num)
// 					break
// 				}
// 				if tempHead.length == index+1 {
// 					this.head.insertAtTail(num)
// 					break
// 				}
// 			}
// 		}
// 	}
// }

// func (this *MedianFinder) FindMedian() float64 {
// 	tempHead := this.head.head
// 	if this.head.length == 1 {
// 		return float64(tempHead.data)
// 	}
// 	for index := 0; index < this.head.length/2-1; index++ {
// 		tempHead = tempHead.next
// 	}
// 	if this.head.length%2 == 1 {
// 		return float64(tempHead.next.data)
// 	} else {
// 		return (float64(tempHead.data) + float64(tempHead.next.data)) / 2
// 	}
// }

// type Node struct {
// 	data int
// 	next *Node
// }

// type LinkedList struct {
// 	head   *Node
// 	length int
// }

// func (l *LinkedList) insertAtHead(data int) {
// 	temp1 := &Node{data, nil}

// 	if l.head == nil {
// 		l.head = temp1
// 	} else {
// 		temp2 := l.head
// 		l.head = temp1
// 		temp1.next = temp2
// 	}
// 	l.length += 1
// }

// func (l *LinkedList) insertAtTail(data int) {
// 	temp1 := &Node{data, nil}

// 	if l.head == nil {
// 		l.head = temp1
// 	} else {
// 		temp2 := l.head
// 		for temp2.next != nil {
// 			temp2 = temp2.next
// 		}
// 		temp2.next = temp1
// 	}
// 	l.length += 1
// }

// func (l *LinkedList) insert(n, data int) {
// 	if n == 0 {
// 		l.insertAtHead(data)
// 	} else if n == l.length-1 {
// 		l.insertAtTail(data)
// 	} else {
// 		temp1 := &Node{data, nil}
// 		temp2 := l.head

// 		for i := 0; i < n-1; i++ {
// 			temp2 = temp2.next
// 		}
// 		temp1.next = temp2.next
// 		temp2.next = temp1
// 		l.length += 1
// 	}
// }

// array solution too slow
// type MedianFinder struct {
// 	list []int
// }

// func Constructor() MedianFinder {
// 	result := MedianFinder{}
// 	return result
// }

// func (this *MedianFinder) AddNum(num int) {
// 	if len(this.list)> 0{
// 		if this.list[0]>num{
// 			this.list = append([]int{num}, this.list...)
// 		}else if this.list[len(this.list)-1] < num{
// 			this.list = append(this.list, num)
// 		}else{
// 			this.list = append(this.list, num)
// 			sort.Slice(this.list, func(i, j int) bool {
// 				return this.list[i] < this.list[j]
// 				})
// 		}
// 	}else{
// 		this.list = append(this.list, num)
// 	}
// }

// func (this *MedianFinder) FindMedian() float64 {

// 	if len(this.list)%2 == 1 { // is it an odd length array
// 		return float64(this.list[len(this.list)/2])
// 	} else {
// 		return (float64(this.list[len(this.list)/2-1]) + float64(this.list[len(this.list)/2])) / 2
// 	}
// }
