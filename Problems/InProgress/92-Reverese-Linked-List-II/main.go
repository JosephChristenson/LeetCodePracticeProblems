package main

import "fmt"

func main() {
	head := []int{1, 2, 3, 4, 5}
	left, right := 2, 4
	fmt.Println(reverseBetween(run(head), left, right))
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
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	start := head
	curr := head

	index := 0

	var leftNode, rightNode *ListNode

	for index <= right {
		if index == left {
			leftNode = curr
		}
		curr = curr.Next
		index++
	}
	rightNode = curr
	rightNode.Next = nil
	leftNode.Next = reverseList(leftNode.Next)
	for leftNode.Next != nil {
		leftNode = leftNode.Next
	}
	leftNode.Next = rightNode
	return start
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
