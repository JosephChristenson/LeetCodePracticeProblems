package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left != nil && root.Right != nil {
		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
		return root
	} else if root.Left != nil {
		root.Left = invertTree(root.Right)
		return root
	} else if root.Right != nil {
		root.Right = invertTree(root.Left)
		return root
	} else {
		return root
	}
}
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	if root.Left != nil && root.Right != nil {
		right := averageOfLevels(root.Right)
		left := averageOfLevels(root.Left)
		index := 0
		result := []float64{float64(root.Val)}
		for index < len(left) && index < len(right) {
			result = append(result, (left[index]+right[index])/2)
			index++
		}
		if index < len(right) {
			result = append(result, right[index:]...)
		} else if index < len(left) {
			result = append(result, left[index:]...)
		}
		return result
	} else if root.Left != nil {
		return append([]float64{float64(root.Val)}, averageOfLevels(root.Left)...)
	} else if root.Right != nil {
		return append([]float64{float64(root.Val)}, averageOfLevels(root.Right)...)
	} else {
		return []float64{float64(root.Val)}
	}
}
