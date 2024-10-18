package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left != nil && root.Right != nil {
		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
		return root
	} else if root.Left != nil {
		root.Left = invertTree(root.Right)
		return root
	} else if root.Right != nil {
		root.Right = invertTree(root.Left)
		return root
	} else {
		return root
	}
}
