package main

import "fmt"

// You are given an n x n grid representing a field of cherries, each cell is one of three possible integers.

// 0 means the cell is empty, so you can pass through,
// 1 means the cell contains a cherry that you can pick up and pass through, or
// -1 means the cell contains a thorn that blocks your way.
// Return the maximum number of cherries you can collect by following the rules below:

// Starting at the position (0, 0) and reaching (n - 1, n - 1) by moving right or down through valid path cells (cells with value 0 or 1).
// After reaching (n - 1, n - 1), returning to (0, 0) by moving left or up through valid path cells.
// When passing through a path cell containing a cherry, you pick it up, and the cell becomes an empty cell 0.
// If there is no valid path between (0, 0) and (n - 1, n - 1), then no cherries can be collected.

func main() {
	grid := [][]int{{0, 1, -1}, {1, 0, -1}, {1, 1, 1}}
	fmt.Println(cherryPickup(grid))
	grid = [][]int{{0, 1, -1}, {1, 0, -1}, {1, -1, 1}}
	fmt.Println(cherryPickup(grid))
	grid = [][]int{{1}}
	fmt.Println(cherryPickup(grid))
	grid = [][]int{{0, 0}, {-1, 1}}
	fmt.Println(cherryPickup(grid))
	grid = [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}
	fmt.Println(cherryPickup(grid))
	grid = [][]int{{1, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 1, 0, 0, 0}, {0, 0, 0, 1, 0, 0, 1}, {1, 0, 0, 1, 0, 0, 0}, {0, 0, 0, 1, 0, 0, 0}, {0, 0, 0, 1, 0, 0, 0}, {0, 0, 0, 1, 1, 1, 1}}
	fmt.Println(cherryPickup(grid))
}

func cherryPickup(grid [][]int) int {
	if len(grid) == 1 {
		return grid[0][0]
	}

	start := &chart{}
	reach := false
	start, reach = createGrid(start, grid, 0, 0)
	if reach {
		return pickupCherries(start, true, []*chart{})
	} else {
		return 0
	}
}

func pickupCherries(root *chart, goingDown bool, visitedList []*chart) int {
	if goingDown {
		right, down := 0, 0
		if root.Right == nil && root.Down == nil {
			return root.val + pickupCherries(root, !goingDown, visitedList) // switch Direction
		}
		if root.Right != nil {
			right = root.val + pickupCherries(root.Right, goingDown, append(visitedList, root.Right))
		}
		if root.Down != nil {
			down = root.val + pickupCherries(root.Down, goingDown, append(visitedList, root.Down))
		}
		if right > down {
			return right
		} else {
			return down
		}
	} else {
		left, up := 0, 0
		if root.Left == nil && root.Up == nil { // return when it reaches top left corner again
			return root.val
		}
		if root.Left != nil {
			left = root.val + pickupCherries(root.Left, goingDown, visitedList)
			for _, val := range visitedList {
				if val == root.Left {
					left = left - root.val
					break
				}
			}
		}
		if root.Up != nil {
			up = root.val + pickupCherries(root.Up, goingDown, visitedList)
			for _, val := range visitedList {
				if val == root.Up {
					up = up - root.val
					break
				}
			}
		}
		if left > up {
			return left
		} else {
			return up
		}
	}
}

type chart struct {
	val                   int
	Up, Right, Down, Left *chart
}

func createGrid(root *chart, grid [][]int, X int, Y int) (*chart, bool) { // bool represents if it could reach the bottom right corner
	if grid[Y][X] == -1 {
		return nil, false
	}
	pass := false
	root.val = grid[Y][X]
	val1 := false
	val2 := false
	if X+1 == len(grid) && Y+1 == len(grid) {
		pass = true
	}
	if X < len(grid[0])-1 {
		newChild := &chart{}
		newChild, val1 = createGrid(newChild, grid, X+1, Y)
		if newChild != nil {
			newChild.Left = root
			root.Right = newChild
		}
	}
	if Y < len(grid)-1 {
		newChild := &chart{}
		newChild, val2 = createGrid(newChild, grid, X, Y+1)
		if newChild != nil {
			newChild.Up = root
			root.Down = newChild
		}
	}
	return root, pass || val1 || val2
}
