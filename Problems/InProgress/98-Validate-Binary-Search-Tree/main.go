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
}
func findRange(root *TreeNode) (int, int,bool){//min, max
	min, max := root.Val, root.Val
	fail = false
	if root.Right != nil{
		temp1, temp2, tempF = findLowest(root.Right)
		if temp1 <
	}
	if root.Left != nil{
		max, max, fail = findHighest(root.Left)
	}
	if root.Left == nil && root.Right == nil{
		return root.Val,root.Val false
	}
}