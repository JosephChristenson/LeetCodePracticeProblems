package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {

}

func searchChildren(root *TreeNode) (*TreeNode, *TreeNode) { //min max
	if root.Right != nil && root.Left != nil {
		minR, maxR := searchChildren(root.Right)
		minL, maxL := searchChildren(root.Left)
	} else if root.Right != nil {
		minR, maxR := searchChildren(root.Right)
	} else if root.Right != nil {
		minL, maxL := searchChildren(root.Left)
	} else {

	}
}
