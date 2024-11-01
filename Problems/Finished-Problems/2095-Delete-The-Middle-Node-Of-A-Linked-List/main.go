package main

import (
	"fmt"
)

func main() {
	head := []int{1, 3, 4, 7, 1, 2, 6}
	fmt.Println(run(head))
	head = []int{1, 2, 3, 4}
	fmt.Println(run(head))
	head = []int{2, 1}
	fmt.Println(run(head))
}

func run(nums []int) *ListNode { //Converts Array to List
	var prev, root *ListNode
	for i, n := range nums {
		node := &ListNode{Val: n}
		if i == 0 {
			root = node
			prev = root
		} else {
			prev.Next = node
			prev = node
		}
	}
	return deleteMiddle(root)
}

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return nil
	}
	curr := head
	count := 0
	for curr != nil {
		count++
		curr = curr.Next
	}
	middle := count / 2
	previous := head
	curr = head.Next
	count = 1

	for count < middle {
		previous = curr
		curr = curr.Next
		count++
	}
	previous.Next = curr.Next
	fmt.Println(previous)
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
