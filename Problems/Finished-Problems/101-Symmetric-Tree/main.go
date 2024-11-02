package main

import "fmt"

func main() {
	fmt.Println(isSymmetric(nil))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return compareNodes(root.Left, root.Right)
}

func compareNodes(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left != nil && right != nil {
		if left.Val != right.Val {
			return false
		} else {
			return compareNodes(left.Left, right.Right) && compareNodes(left.Right, right.Left)
		}
	} else {
		return false
	}
}
