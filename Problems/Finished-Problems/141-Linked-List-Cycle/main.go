package main

// Given head, the head of a linked list, determine if the linked list has a cycle in it.

// There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

// Return true if there is a cycle in the linked list. Otherwise, return false.

func main() {
	//No given input method
}

func hasCycle(head *ListNode) bool {
	var m map[*ListNode]bool
	m = make(map[*ListNode]bool)

	if head == nil {
		return false
	}

	for head.Next != nil {
		if m[head] {
			return true
		} else {
			m[head] = true
		}
		head = head.Next
	}
	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}
