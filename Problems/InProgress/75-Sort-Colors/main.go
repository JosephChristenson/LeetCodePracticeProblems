package main

func main() {
	nums := []int{2, 0, 1, 1, 0}
	sortColors(nums)
	nums = []int{2, 0, 1}
	sortColors(nums)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortColors(nums []int) {
	red, white, blue := 0, 0, 0
	for _, val := range nums {
		if val == 0 {
			red++
		} else if val == 1 {
			white++
		} else {
			blue++
		}
	}
	order := run(nums)
	for red+white+blue > 0 {
		if red > 0 {
			red--
			order.Val = 0
		} else if white > 0 {
			white--
			order.Val = 1
		} else {
			blue--
			order.Val = 2
		}
		order = order.Next
	}
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
