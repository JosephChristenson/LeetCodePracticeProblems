package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Right != nil && root.Left != nil {
		right := rightSideView(root.Right)
		left := rightSideView(root.Left)

		if len(right) >= len(left) {
			return append([]int{root.Val}, right...)
		} else {
			return append([]int{root.Val}, append(right, left[len(right):]...)...)
		}
	} else if root.Right != nil {
		return append([]int{root.Val}, rightSideView(root.Right)...)
	} else if root.Left != nil {
		return append([]int{root.Val}, rightSideView(root.Left)...)
	} else {
		return []int{root.Val}
	}
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
