package main

import "fmt"

// Given the head of a linked list, rotate the list to the right by k places.

func main() {
	head := []int{1, 2, 3, 4, 5}
	k := 2
	fmt.Println(rotateRight(run(head), k))
	head = []int{0, 1, 2}
	k = 4
	fmt.Println(rotateRight(run(head), k))
	head = []int{1, 2}
	k = 1
	fmt.Println(rotateRight(run(head), k))

}

type ListNode struct {
	Val  int
	Next *ListNode
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
	return root
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil || k == 0 {
		return head
	}

	start := head
	length := 1

	for head.Next != nil {
		head = head.Next
		length++
	}
	head.Next = start //circle
	head = start
	k = k % length
	prev := head

	for ; length-k > 0; length-- {
		prev = head
		head = head.Next
	}
	result := head
	prev.Next = nil

	return result
}
