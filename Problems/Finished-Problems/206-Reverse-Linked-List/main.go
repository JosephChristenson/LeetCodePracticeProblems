package main

import "fmt"

func main() {
	head := []int{1, 2, 3, 4, 5}
	fmt.Println(reverseList(run(head)))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func run(nums []int) *ListNode {
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
	return root
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	prev := head
	curr := head.Next
	prev.Next = nil

	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}
