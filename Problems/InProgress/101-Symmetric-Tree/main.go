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
	}
	if root.Left != nil && root.Right != nil {
		left := inorderTraversal(root.Left)
		right := inorderTraversal(root.Right)
		if len(left) == len(right) {
			for index, val := range left {
				if val != right[len(right)-index-1] {
					return false
				}
			}
			return true
		} else {
			return false
		}
	} else if root.Left == nil && root.Right == nil {
		return true
	}
	return false
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
		return append(left, append([]int{root.Val}, -1)...)
	} else if root.Right != nil {
		right := inorderTraversal(root.Right)
		return append([]int{-1}, append([]int{root.Val}, right...)...)
	} else {
		return []int{root.Val}
	}
}
