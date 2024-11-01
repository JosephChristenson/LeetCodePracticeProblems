package main

// Given a binary tree, find its minimum depth.

// The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

// Note: A leaf is a node with no children.

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left != nil && root.Right != nil {
		left := minDepth(root.Left)
		right := minDepth(root.Right)
		if left < right {
			return left + 1
		} else {
			return right + 1
		}
	}
	if root.Left != nil {
		return 1 + minDepth(root.Left)
	}
	if root.Right != nil {
		return 1 + minDepth(root.Right)
	}
	return -1
}
