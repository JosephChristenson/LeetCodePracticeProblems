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
}

func cherryPickup(grid [][]int) int {

}

// type position struct {
// 	X, Y int
// }

// func cherryPickup(grid [][]int) int {
// 	rowCount, colCount := len(grid), len(grid[0])

// 	totalCherries := 0

// 	var directions map[position]int
// 	directions = make(map[position]int)

// 	for y := rowCount - 1; y >= 0; y-- {
// 		for x := colCount - 1; x >= 0; x-- {
// 			if grid[y][x] == -1 {
// 				directions[position{X: x, Y: y}] = -1
// 			} else {
// 				valueX, valueY := 0, 0

// 				if x+1 == colCount && y+1 == rowCount {
// 					directions[position{X: x, Y: y}] = grid[y][x]
// 				} else {
// 					if x+1 == colCount {
// 						valueX = -1
// 					} else {
// 						if directions[position{X: x + 1, Y: y}] == -1 {
// 							valueX = -1
// 						} else {
// 							valueX = grid[y][x] + directions[position{X: x + 1, Y: y}]
// 						}
// 					}
// 					if y+1 == rowCount {
// 						valueY = -1
// 					} else {
// 						if directions[position{X: x, Y: y + 1}] == -1 {
// 							valueY = -1
// 						} else {
// 							valueY = grid[y][x] + directions[position{X: x, Y: y + 1}]
// 						}

// 					}

// 					if valueX == -1 && valueY == -1 {
// 						directions[position{X: x, Y: y}] = -1 // if they are both dead ends then going to this node is a dead end
// 					} else if valueX == -1 {
// 						directions[position{X: x, Y: y}] = valueY
// 					} else if valueY == -1 {
// 						directions[position{X: x, Y: y}] = valueX
// 					} else if valueX < valueY {
// 						directions[position{X: x, Y: y}] = valueY
// 					} else {
// 						directions[position{X: x, Y: y}] = valueX
// 					}
// 				}
// 			}
// 		}
// 	}
// 	if directions[position{X: 0, Y: 0}] == -1 { // if there is no valid path or no cherries in the maze return now
// 		return 0
// 	}
// 	totalCherries = directions[position{X: 0, Y: 0}]

// 	x, y := 0, 0
// 	for x < colCount && y < rowCount {
// 		if directions[position{X: x + 1, Y: y}] > directions[position{X: x, Y: y + 1}] {
// 			grid[y][x] = 0
// 			x++
// 		} else {
// 			grid[y][x] = 0
// 			y++
// 		}
// 	}

// 	for y := 0; y < rowCount; y++ {
// 		for x := 0; x < colCount; x++ {
// 			if grid[y][x] == -1 {
// 				directions[position{X: x, Y: y}] = -1
// 			} else {
// 				valueX, valueY := 0, 0

// 				if x == 0 && y == 0 {
// 					directions[position{X: x, Y: y}] = grid[y][x]
// 				} else {
// 					if x == 0 {
// 						valueX = -1
// 					} else {
// 						if directions[position{X: x - 1, Y: y}] == -1 {
// 							valueX = -1
// 						} else {
// 							valueX = grid[y][x] + directions[position{X: x - 1, Y: y}]
// 						}
// 					}
// 					if y == 0 {
// 						valueY = -1
// 					} else {
// 						if directions[position{X: x, Y: y - 1}] == -1 {
// 							valueY = -1
// 						} else {
// 							valueY = grid[y][x] + directions[position{X: x, Y: y - 1}]
// 						}

// 					}

// 					if valueX == -1 && valueY == -1 {
// 						directions[position{X: x, Y: y}] = -1 // if they are both dead ends then going to this node is a dead end
// 					} else if valueX == -1 {
// 						directions[position{X: x, Y: y}] = valueY
// 					} else if valueY == -1 {
// 						directions[position{X: x, Y: y}] = valueX
// 					} else if valueX < valueY {
// 						directions[position{X: x, Y: y}] = valueY
// 					} else {
// 						directions[position{X: x, Y: y}] = valueX
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return totalCherries + directions[position{X: colCount - 1, Y: rowCount - 1}]
// }
