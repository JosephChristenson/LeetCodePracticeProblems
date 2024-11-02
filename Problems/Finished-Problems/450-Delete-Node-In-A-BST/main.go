package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		return fixTree(root.Left, root.Right)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	} else {
		root.Left = deleteNode(root.Left, key)
		return root
	}
}

func fixTree(left *TreeNode, right *TreeNode) *TreeNode {
	if right == nil {
		return left
	}
	head := right
	for right.Left != nil {
		right = right.Left
	}
	right.Left = left
	return head
}
