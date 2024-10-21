package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	} else if root.Left != nil && root.Right != nil {
		return reflectSide(root.Left) == root.Right
	}
	return false
}

func reflectSide(root *TreeNode) *TreeNode {
	if root.Left != nil && root.Right != nil {
		root.Left, root.Right = reflectSide(root.Right), reflectSide(root.Left)
	} else if root.Left != nil {
		root.Right = reflectSide(root.Left)
	} else if root.Right != nil {
		root.Left = reflectSide(root.Right)
	} else {
		return nil
	}
	return root
}
