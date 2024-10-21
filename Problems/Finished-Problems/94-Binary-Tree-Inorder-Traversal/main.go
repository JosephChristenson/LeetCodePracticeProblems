package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	} else if root.Left != nil && root.Right != nil {
		left := inorderTraversal(root.Left)
		left = append(left, root.Val)
		right := inorderTraversal(root.Right)
		return append(left, right...)
	} else if root.Left != nil {
		left := inorderTraversal(root.Left)
		return append(left, root.Val)
	} else if root.Right != nil {
		right := inorderTraversal(root.Right)
		return append([]int{root.Val}, right...)
	} else {
		return []int{root.Val}
	}
}
