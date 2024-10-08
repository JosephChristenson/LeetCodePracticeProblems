package main

import "fmt"

// Given the head of a linked list, remove the nth node from the end of the list and return its head.

func main() {
	//list := LinkedList{nil, 0}

	// Example provided by Jake
	head := []int{1, 2, 3, 4, 5}
	n := 2
	fmt.Println(run(head, n))
	fmt.Println(head)
	head = []int{1}
	n = 1
	fmt.Println(run(head, n))
	fmt.Println(head)
	head = []int{1, 2}
	n = 1
	fmt.Println(run(head, n))
	fmt.Println(head)
}

func run(nums []int, n int) *ListNode { //Converts Array to List
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
	return removeNthFromEnd(root, n)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	Curr, Storage := head, head
	if Curr.Next == nil {
		return nil
	}
	Count := 1
	for Curr.Next != nil {
		Count++
		Curr = Curr.Next
	}
	Prev := Storage
	Curr = Prev.Next

	if Count == n+1 {
		Prev.Next = nil
		return Curr
	}

	for Count-n-1 > 0 {
		if Curr.Next != nil {
			Curr = Curr.Next
		} else {
			break
		}
		Prev = Prev.Next
		Count--
	}
	Prev.Next = Curr.Next

	return Storage
}

type ListNode struct {
	Val  int
	Next *ListNode
}
