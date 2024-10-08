package main

import "fmt"

// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1, list2 := []int{2, 4, 3}, []int{5, 6, 4}
	l1, l2 := run(list1), run(list2)
	fmt.Println(addTwoNumbers(l1, l2))

	list1, list2 = []int{0}, []int{0}
	l1, l2 = run(list1), run(list2)
	fmt.Println(addTwoNumbers(l1, l2))

	list1, list2 = []int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}
	l1, l2 = run(list1), run(list2)
	fmt.Println(addTwoNumbers(l1, l2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	Index1, Index2 := l1, l2
	carry := 0
	result := []int{}
	add := 0

	for Index1 != nil && Index2 != nil {
		add = Index1.Val + Index2.Val + carry
		carry = 0
		if add > 9 {
			carry = 1
			add = add - 10
		}

		result = append(result, add)

		Index1 = Index1.Next
		Index2 = Index2.Next
	}
	for Index1 != nil { // if l1 is longer than l2 get rest of l1
		add = Index1.Val + carry
		carry = 0
		if add > 9 {
			carry = 1
			add = add - 10
		}
		result = append(result, add)

		Index1 = Index1.Next
	}
	for Index2 != nil { // if l2 is longer than l1 get rest of l2
		add = Index2.Val + carry
		carry = 0
		if add > 9 {
			carry = 1
			add = add - 10
		}
		result = append(result, add)

		Index2 = Index2.Next
	}
	if carry == 1 {
		result = append(result, 1)
	}
	fmt.Println(result)
	return run(result)
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
