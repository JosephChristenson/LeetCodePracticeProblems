package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Right != nil {
		left := maxDepth(root.Left)
		right := maxDepth(root.Right)
		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	} else if root.Left != nil {
		return maxDepth(root.Left) + 1
	} else if root.Right != nil {
		return maxDepth(root.Right) + 1
	} else {
		return 1
	}
}
