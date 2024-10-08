package main

import (
	"fmt"
	"sort"
)

// You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

// Merge all the linked-lists into one sorted linked-list and return it.

type ListNode struct {
	Val  int
	Next *ListNode
}
type LinkedList struct {
	head   *ListNode
	length int
}

func main() {
	lists := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}
	listheads := []*ListNode{}
	for _, Value := range lists {
		listheads = append(listheads, run(Value))
	}
	fmt.Println(mergeKLists(listheads))
	fmt.Println(lists)
}

func mergeKLists(lists []*ListNode) *ListNode {

	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}

	var m map[int][]*ListNode
	m = make(map[int][]*ListNode)

	for _, List := range lists {
		for List != nil {
			m[List.Val] = append(m[List.Val], List)
			List = List.Next
		}
	}
	keyList := make([]int, 0, len(m))

	for k := range m {
		keyList = append(keyList, k)
	}
	sort.Ints(keyList)

	if len(keyList) == 0 {
		//indicates no values in list
		return lists[0]
	}

	Head := (m[keyList[0]][0]) // Return this value. Is the beginning of the list
	Previous := Head           //stores pointer from last loop to continue list

	for _, Key := range keyList {
		for _, Point := range m[Key] {
			Previous.Next = Point
			Previous = Point
		}
	}

	if Head.Next == Head { // In case there is only one element period in the list. Make sure it doesn't point to itself
		Head.Next = nil
	}

	return Head
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
