package main

import "fmt"

func main() {
	head := []int{1, 2, 3, 4}
	fmt.Println(swapPairs(run(head)))
	head = []int{}
	fmt.Println(swapPairs(run(head)))
	head = []int{1}
	fmt.Println(swapPairs(run(head)))
	head = []int{1, 2, 3}
	fmt.Println(swapPairs(run(head)))

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

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	} else {
		temp := head.Next
		head.Next = head
		head = swapPairs(temp)
		return head.Next
	}
}
