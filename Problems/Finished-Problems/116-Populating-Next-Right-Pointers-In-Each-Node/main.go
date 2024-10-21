package main

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
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
	for i := len(results) - 1; i > 0; i-- {
		for j := 0; j < len(results[i]); j++ {
			if j == len(results[i])-1 {
				results[i][j].Right.Next = nil
			} else {
				results[i][j].Right.Next = results[i][j+1].Left
			}
			results[i][j].Left.Next = results[i][j].Right
		}
	}
	root.Next = nil
	if root.Left != nil {
		root.Left.Next = root.Right
		root.Right.Next = nil
	}
	return root
}

func findConnect(root *Node) [][]*Node {
	if root == nil {
		return nil
	}
	results := [][]*Node{}
	if root.Left != nil {
		results = findConnect(root.Left)
		temp := findConnect(root.Right)
		for index := range results {
			results[index] = append(results[index], temp[index]...)
		}
		results = append(results, []*Node{root.Left, root.Right})
	}
	return results
}
