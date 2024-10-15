package main

import "fmt"

// You are given the heads of two sorted linked lists list1 and list2.

// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

// Return the head of the merged linked list.

func main() {
	list1, list2 := []int{1, 2, 4}, []int{1, 3, 4}
	fmt.Println(mergeTwoLists(run(list1), run(list2)))
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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	var head, curr *ListNode
	if list1 == nil || list2 == nil {
		if list1 == nil {
			return list2
		} else {
			return list1
		}
	}

	if list1.Val < list2.Val {
		head, curr = list1, list1
		list1 = list1.Next
	} else {
		head, curr = list2, list2
		list2 = list2.Next
	}
	for list1 != nil || list2 != nil {
		if list1 == nil {
			curr.Next = list2
			curr = curr.Next
			list2 = list2.Next
		} else if list2 == nil {
			curr.Next = list1
			curr = curr.Next
			list1 = list1.Next
		} else if list1.Val <= list2.Val {
			curr.Next = list1
			curr = curr.Next
			list1 = list1.Next
		} else {
			curr.Next = list2
			curr = curr.Next
			list2 = list2.Next
		}
	}
	return head
}
