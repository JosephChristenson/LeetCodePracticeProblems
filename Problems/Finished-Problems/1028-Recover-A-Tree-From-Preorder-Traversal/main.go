package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	traversal := "1-2--3--4-5--6--7"
	fmt.Println(recoverFromPreorder(traversal))
	traversal = "1-2--3---4-5--6---7"
	fmt.Println(recoverFromPreorder(traversal))
	traversal = "1-401--349---90--88"
	fmt.Println(recoverFromPreorder(traversal))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type ConstructNode struct {
	Val   int
	Depth int
}

func recoverFromPreorder(traversal string) *TreeNode {
	byteArray := strings.Split(traversal, "-")

	allNodes := make([]ConstructNode, 0, len(byteArray))
	newNode := ConstructNode{}

	leftStart := 0
	rightStart := 0
	for _, val := range byteArray {
		if val == "" {
			newNode.Depth++
		} else {
			if newNode.Depth == 1 {
				if leftStart == 0 {
					leftStart = len(allNodes)
				} else {
					rightStart = len(allNodes)
				}
			}
			newNode.Val, _ = strconv.Atoi(val)
			allNodes = append(allNodes, newNode)
			newNode = ConstructNode{}
			newNode.Depth = 1
		}
	}
	head := &TreeNode{Val: allNodes[0].Val, Left: nil, Right: nil}
	if rightStart != 0 {
		head.Left = itterateThroughBranches(allNodes[leftStart:rightStart], 1)
		head.Right = itterateThroughBranches(allNodes[rightStart:], 1)
	} else if leftStart != 0 {
		head.Left = itterateThroughBranches(allNodes[1:], 1)
	}
	return head
}

func itterateThroughBranches(values []ConstructNode, depth int) *TreeNode {
	if len(values) == 0 {
		return nil
	}
	head := &TreeNode{Val: values[0].Val, Left: nil, Right: nil}

	leftIndex, rightIndex := 0, 0
	for index, val := range values {
		if val.Depth == depth {
			// do nothing
		} else if val.Depth == depth+1 {
			if leftIndex == 0 {
				leftIndex = index
			} else {
				rightIndex = index
			}
		}
	}
	if rightIndex != 0 {
		head.Left = itterateThroughBranches(values[leftIndex:rightIndex], depth+1)
		head.Right = itterateThroughBranches(values[rightIndex:], depth+1)
	} else if leftIndex != 0 {
		head.Left = itterateThroughBranches(values[1:], depth+1)
	}
	return head
}
