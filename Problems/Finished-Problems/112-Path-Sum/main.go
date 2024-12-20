package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left != nil && root.Right != nil {
		return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
	} else if root.Left != nil {
		return hasPathSum(root.Left, targetSum-root.Val)
	} else if root.Right != nil {
		return hasPathSum(root.Right, targetSum-root.Val)
	} else {
		return targetSum-root.Val == 0
	}
}
