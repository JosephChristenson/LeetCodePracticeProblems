package main

import "fmt"

// Given the head of a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list. Return the linked list sorted as well.

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := []int{1, 2, 3, 3, 4, 4, 5}
	fmt.Println(deleteDuplicates(run(head)))
	head = []int{1, 1, 1, 2, 3}
	fmt.Println(deleteDuplicates(run(head)))
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

func deleteDuplicates(head *ListNode) *ListNode {

	curr := head

	var root *ListNode
	var empty *ListNode
	var prev *ListNode

	for curr != nil {
		if curr.Next != empty {
			if curr.Next.Val == curr.Val {
				curr = curr.Next
				prev.Next = nil
			} else {
				if root == empty {
					root = curr
				}
				prev = curr
				curr = curr.Next
			}
		} else {
			if prev != nil {
				if prev.Val != curr.Val {
					prev.Next = curr
				}
			}
			curr = empty
		}
	}
	return root
}

// exceeds memory limit
// func deleteDuplicates(head *ListNode) *ListNode {
// 	curr := head

// 	var pointers map[int]*ListNode
// 	pointers = make(map[int]*ListNode)

// 	var includePoint map[int]bool
// 	includePoint = make(map[int]bool)

// 	var empty *ListNode
// 	for curr != nil {
// 		if pointers[curr.Val] != empty {
// 			includePoint[curr.Val] = false
// 		} else {
// 			pointers[curr.Val] = curr
// 			includePoint[curr.Val] = true
// 		}
// 		curr = curr.Next
// 	}

// 	var root *ListNode
// 	var prev *ListNode
// 	for index := range pointers {
// 		if includePoint[index] {
// 			if root == empty {
// 				root = pointers[index]
// 				prev = root
// 			} else {
// 				prev.Next = pointers[index]
// 				prev = prev.Next
// 			}
// 		}
// 	}
// 	return root
// }
