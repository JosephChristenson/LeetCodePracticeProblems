package main

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
	results := findConnect(root)
	for i := len(results) - 1; i > 0; i-- {
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
				if results[i][j].Left != nil {
					if len(results[i])-1 == j {
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
	}
	root.Next = nil
	if root.Right != nil {
		root.Right.Next = nil
		if root.Left != nil {
			root.Left.Next = root.Right
		}
	} else {
		if root.Left != nil {
			root.Left.Next = nil
		}
	}
	return root
}

func findConnect(root *Node) [][]*Node {
	if root == nil {
		return nil
	}
	results := [][]*Node{}
	if root.Left != nil && root.Right != nil {
		tempL := findConnect(root.Left)
		tempR := findConnect(root.Right)

		if len(tempL) >= len(tempR) {
			results = make([][]*Node, len(tempL))

			for index := 0; index < len(tempL); index++ {
				if len(tempR) > index+1 {
					results[index] = append(tempL[index], tempR[index]...)
				} else {
					results[index] = tempL[index]
				}
			}
		} else {
			results = make([][]*Node, len(tempR))

			for index := 0; index < len(tempR); index++ {
				if len(tempL) > index+1 {
					results[index] = append(tempL[index], tempR[index]...)
				} else {
					results[index] = tempR[index]
				}
			}
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
