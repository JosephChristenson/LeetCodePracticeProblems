package main

import (
	"fmt"
)

// Given the head of a linked list, return the list after sorting it in ascending order

func main() {
	head := []int{4, 2, 1, 3}
	fmt.Println(sortList(run(head)))
	head = []int{-1, 5, 3, 4, 0}
	fmt.Println(sortList(run(head)))
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

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {

	curr := head
	for curr != nil {
		curr = curr.Next
	}

	return head
}
