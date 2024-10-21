package main

// Given the root of a binary tree and an integer targetSum, return all root-to-leaf paths where the sum of the node values in the path equals targetSum. Each path should be returned as a list of the node values, not node references.

// A root-to-leaf path is a path starting from the root and ending at any leaf node. A leaf is a node with no children.

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	allNodes := [][]int{}
	if root.Left != nil && root.Right != nil {
		allNodes = append(allNodes, pathSum(root.Left, targetSum-root.Val)...)
		allNodes = append(allNodes, pathSum(root.Right, targetSum-root.Val)...)
	} else if root.Left != nil {
		allNodes = append(allNodes, pathSum(root.Left, targetSum-root.Val)...)
	} else if root.Right != nil {
		allNodes = append(allNodes, pathSum(root.Right, targetSum-root.Val)...)
	} else {
		if targetSum-root.Val == 0 {
			return [][]int{{root.Val}}
		}
	}
	for index := range allNodes {
		allNodes[index] = append([]int{root.Val}, allNodes[index]...)
	}

	return allNodes
}
