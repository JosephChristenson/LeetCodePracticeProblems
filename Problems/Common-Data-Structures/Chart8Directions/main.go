package main

import "fmt"

// 4 direction grid

func main() {
	grid := [][]int{{0, 1, -1}, {1, 0, -1}, {1, 1, 1}}
	fmt.Println(cherryPickup(grid))
}

func cherryPickup(grid [][]int) int {
	start := &chart{}
	start = createGrid(start, grid, 0, 0)

	return 0
}

type chart struct {
	val                          int
	Up, Right, Down, Left        *chart
	URight, ULeft, DRight, DLeft *chart
}

func createGrid(root *chart, grid [][]int, X int, Y int) *chart {
	root.val = grid[Y][X]
	if X < len(grid[0])-1 {
		newChild := &chart{}
		newChild = createGrid(newChild, grid, X+1, Y)
		newChild.Left = root
		root.Right = newChild
	}
	if Y < len(grid)-1 {
		newChild := &chart{}
		newChild = createGrid(newChild, grid, X, Y+1)
		newChild.Up = root
		root.Down = newChild
	}

	if root.Up != nil {
		root.ULeft = root.Up.Left
		root.URight = root.Up.Right
	}
	if root.Down != nil {
		root.DLeft = root.Down.Left
		root.DRight = root.Down.Right
	}
	return root
}
