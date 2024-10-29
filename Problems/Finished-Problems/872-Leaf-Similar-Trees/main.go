package main

// Consider all the leaves of a binary tree, from left to right order, the values of those leaves form a leaf value sequence.

// For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).

// Two binary trees are considered leaf-similar if their leaf value sequence is the same.

// Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

	result1, result2 := findAllLeaves(root1), findAllLeaves(root2)

	if len(result1) != len(result2) {
		return false
	} else {
		for index := range result1 {
			if result1[index] != result2[index] {
				return false
			}
		}
	}
	return true
}

func findAllLeaves(root *TreeNode) []int {
	results := []int{}

	if root.Left != nil && root.Right != nil {
		results = append(results, findAllLeaves(root.Left)...)
		results = append(results, findAllLeaves(root.Right)...)
	} else if root.Left != nil {
		results = append(results, findAllLeaves(root.Left)...)
	} else if root.Right != nil {
		results = append(results, findAllLeaves(root.Right)...)
	} else {
		return []int{root.Val}
	}
	return results
}
