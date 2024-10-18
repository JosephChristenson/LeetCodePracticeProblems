package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 != nil && root2 != nil {
		var newNode *TreeNode
		leftNode := mergeTrees(root1.Left, root2.Left)
		rightNode := mergeTrees(root1.Right, root2.Right)
		newNode = &TreeNode{}
		newNode.Left = leftNode
		newNode.Right = rightNode
		newNode.Val = root1.Val + root2.Val
		return newNode
	}
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	return nil
}
