package main

func main() {

}

type TreeNode struct { //Binary tree
	data  int
	left  *TreeNode
	right *TreeNode
}

func createTree(array []int) *TreeNode {
	if array == nil || len(array) == 0 {
		return nil
	}

}
