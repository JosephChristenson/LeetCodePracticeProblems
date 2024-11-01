package main

import "fmt"

func main() {
	fmt.Println(isValidBST(nil))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Left != nil && root.Right != nil {
		_, leftMax, LeftF := findRange(root.Left)
		rightMin, _, rightF := findRange(root.Right)
		if LeftF || rightF {
			return false
		}
		if root.Val >= rightMin {
			return false
		}
		if root.Val <= leftMax {
			return false
		}
	} else if root.Right != nil {
		rightMin, _, rightF := findRange(root.Right)
		if rightF {
			return false
		}
		if root.Val >= rightMin {
			return false
		}
	} else if root.Left != nil {
		_, leftMax, LeftF := findRange(root.Left)
		if LeftF {
			return false
		}
		if root.Val <= leftMax {
			return false
		}
	}
	return true
}
func findRange(root *TreeNode) (int, int, bool) { //min, max
	if root.Left != nil && root.Right != nil {
		leftMin, leftMax, LeftF := findRange(root.Left)
		rightMin, rightMax, rightF := findRange(root.Right)
		if LeftF || rightF {
			return root.Val, root.Val, true
		}
		if root.Val >= rightMin {
			return leftMin, rightMax, true
		}
		if root.Val <= leftMax {
			return leftMin, rightMax, true
		}
		return leftMin, rightMax, false
	} else if root.Right != nil {
		rightMin, rightMax, rightF := findRange(root.Right)
		if rightF {
			return root.Val, root.Val, true
		}
		if root.Val >= rightMin {
			return rightMin, rightMax, true
		}
		return root.Val, rightMax, false
	} else if root.Left != nil {
		leftMin, leftMax, LeftF := findRange(root.Left)
		if LeftF {
			return root.Val, root.Val, true
		}
		if root.Val <= leftMax {
			return leftMin, leftMax, true
		}
		return leftMin, root.Val, false
	} else {
		return root.Val, root.Val, false
	}
}
