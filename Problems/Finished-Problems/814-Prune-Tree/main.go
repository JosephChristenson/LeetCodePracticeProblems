package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left != nil && root.Right != nil {
		root.Left = pruneTree(root.Left)
		root.Right = pruneTree(root.Right)
	} else if root.Left != nil {
		root.Left = pruneTree(root.Left)
	} else if root.Right != nil {
		root.Right = pruneTree(root.Right)
	} else {
		if root.Val == 1 {
			return root
		} else {
			return nil
		}
	}
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}
	return root
}
