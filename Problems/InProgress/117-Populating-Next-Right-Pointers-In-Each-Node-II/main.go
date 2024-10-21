package main

import "slices"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func main() {

}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	results := [][]*Node{}
	results = append(results, []*Node{root})

	if root.Left != nil {
		results = findConnect(root.Left)
		temp := findConnect(root.Right)
		for index := range results {
			results[index] = append(results[index], temp[index]...)
		}
		results = append(results, []*Node{root.Left, root.Right})
	}
	slices.Reverse(results)
	for i := 0; i < len(results)-1; i++ {
		for j := 0; j < len(results[i]); j++ {

			if results[i][j].Right != nil {
				if j == len(results[i])-1 {
					results[i][j].Right.Next = nil
				} else {
					if results[i][j+1].Left != nil {
						results[i][j].Right.Next = results[i][j+1].Left
					} else {
						results[i][j].Right.Next = results[i][j+1].Right
					}
				}
				if results[i][j].Left != nil {
					results[i][j].Left.Next = results[i][j].Right
				}
			} else {
				if j == len(results[i])-1 {
					results[i][j].Left.Next = nil
				} else {
					if results[i][j+1].Left != nil {
						results[i][j].Left.Next = results[i][j+1].Left
					} else {
						results[i][j].Left.Next = results[i][j+1].Right
					}
				}
			}
		}
	}
	root.Next = nil
	if root.Left != nil && root.Right != nil {
		root.Left.Next = root.Right
	}
	if root.Right != nil {
		root.Right.Next = nil
	}

	return root
}

func findConnect(root *Node) [][]*Node {
	if root == nil {
		return nil
	}
	results := [][]*Node{}
	if root.Left != nil && root.Right != nil {
		results = findConnect(root.Left)
		temp := findConnect(root.Right)
		for index := range results {
			results[index] = append(results[index], temp[index]...)
		}
		results = append(results, []*Node{root.Left, root.Right})
	} else if root.Left != nil {
		results = findConnect(root.Left)
		results = append(results, []*Node{root.Left})
	} else if root.Right != nil {
		results = findConnect(root.Right)
		results = append(results, []*Node{root.Right})
	}
	return results
}
