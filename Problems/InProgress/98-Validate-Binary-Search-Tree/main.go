package main

func main() {

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
	test := true
	if root.Left != nil {
		if root.Left.Val >= root.Val {
			return false
		}
		test = test && testValues(root.Left, root.Val, root.Val)
	}
	if root.Right != nil {
		if root.Right.Val <= root.Val {
			return false
		}
		test = test && testValues(root.Right, root.Val, root.Val)
	}
	return true
}

func testValues(root *TreeNode, min int, max int) bool {
	test := true

	if root.Val < min {
		min = root.Val
	}
	if root.Val > max {
		max = root.Val
	}
	if root.Left != nil {
		if root.Left.Val >= min {
			return false
		}
		test = test && testValues(root.Left, min, max)
	}
	if root.Right != nil {
		if root.Right.Val <= max {
			return false
		}
		test = test && testValues(root.Left, min, max)
	}
	return true
}
