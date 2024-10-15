package main

import "fmt"

// Given the root of a binary tree, return its maximum depth.

// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

const (
	NIL_REPLACEMENT = -101 // what value in the int array represents a nil value
)

func main() {
	root := []int{39, 20, -101, -101, 15, 7}
	fmt.Println(maxDepth(createTree(root)))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {

}

func createTree(array []int) *TreeNode {
	root := &TreeNode{0, nil, nil}
	for _, val := range array {
		root.Insert(val)
	}
	return root
}

func (branch *TreeNode) Insert(val int) {
	branch.InsertRec(branch, val)
}

func (branch *TreeNode) InsertRec(node *TreeNode, val int) *TreeNode {
	if branch == nil {
		branch = &TreeNode{val, nil, nil}
		return node
	}
	if node == nil {
		return &TreeNode{val, nil, nil}
	}
	if val <= node.Val {
		node.Left = branch.InsertRec(node.Left, val)
	}
	if val > node.Val {
		node.Right = branch.InsertRec(node.Right, val)
	}
	return node
}
