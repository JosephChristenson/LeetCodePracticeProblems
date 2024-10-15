package main

func main() {
}

func hasCycle(head *ListNode) bool {
	var m map[*ListNode]int
	m = make(map[*ListNode]int)

	if head == nil {
		return false
	}

	for head.Next != nil {
		if m[head] == 1 {
			return true
		} else {
			m[head] = 1
		}
		head = head.Next
	}
	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}
