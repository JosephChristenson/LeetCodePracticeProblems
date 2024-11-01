package main

import "fmt"

func main() {
	maze := [][]byte{{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'}}
	entrance := []int{1, 2}
	fmt.Println(nearestExit(maze, entrance))
}

func nearestExit(maze [][]byte, entrance []int) int {
	start := &chart{}
	start = createGrid(start, maze, entrance[1], entrance[0])
	return start.val
}
func createGrid(root *chart, grid [][]byte, X int, Y int) *chart {
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
	return root
}

type chart struct {
	val                   byte
	Up, Right, Down, Left *chart
}
