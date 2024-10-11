package main

import (
	"fmt"
)

// You are given a rows x cols matrix grid representing a field of cherries where grid[i][j] represents the number of cherries that you can collect from the (i, j) cell.

// You have two robots that can collect cherries for you:

// Robot #1 is located at the top-left corner (0, 0), and
// Robot #2 is located at the top-right corner (0, cols - 1).
// Return the maximum number of cherries collection using both robots by following the rules below:

// From a cell (i, j), robots can move to cell (i + 1, j - 1), (i + 1, j), or (i + 1, j + 1).
// When any robot passes through a cell, It picks up all cherries, and the cell becomes an empty cell.
// When both robots stay in the same cell, only one takes the cherries.
// Both robots cannot move outside of the grid at any moment.
// Both robots should reach the bottom row in grid.

func main() {
	grid := [][]int{{3, 1, 1}, {2, 5, 1}, {1, 5, 5}, {2, 1, 1}}
	fmt.Println(cherryPickup(grid))
	grid = [][]int{{1, 0, 0, 0, 0, 0, 1}, {2, 0, 0, 0, 0, 3, 0}, {2, 0, 9, 0, 0, 0, 0}, {0, 3, 0, 5, 4, 0, 0}, {1, 0, 2, 3, 0, 0, 6}}
	fmt.Println(cherryPickup(grid))
	grid = [][]int{
		{8, 8, 10, 9, 1, 7},
		{8, 8, 1, 8, 4, 7},
		{8, 6, 10, 3, 7, 7},
		{3, 0, 9, 3, 2, 7},
		{6, 8, 9, 4, 2, 5},
		{1, 1, 5, 8, 8, 1},
		{5, 6, 5, 2, 9, 9},
		{4, 4, 6, 2, 5, 4},
		{4, 4, 5, 3, 1, 6},
		{9, 2, 2, 1, 9, 3}}
	fmt.Println(cherryPickup(grid))
}

type position struct {
	X, Y int
}
type valuePath struct {
	value, bestCol, secondBestCol int
}

// func cherryPickup(grid [][]int) int {

// 	var maxValue map[position]valuePath
// 	maxValue = make(map[position]valuePath)

// 	var preferRightValue map[position]valuePath
// 	preferRightValue = make(map[position]valuePath)

// 	row, col := len(grid)-1, len(grid[0])-1

// 	gridLength := len(grid[0]) - 1

// 	for row >= 0 {
// 		if row == len(grid)-1 {
// 			for col >= 0 {

// 				maxValue[position{row, col}] = valuePath{grid[row][col], 0, 0}
// 				preferRightValue[position{row, col}] = valuePath{grid[row][col], 0, 0}
// 				col--
// 			}
// 		}

// 		for col >= 0 {
// 			left, mid, right := valuePath{0, col - 1, 0}, valuePath{0, col, 0}, valuePath{0, col + 1, 0}
// 			curr := grid[row][col]
// 			if col-1 >= 0 {
// 				left.value = curr + maxValue[position{row + 1, col - 1}].value
// 			}
// 			if col+1 <= gridLength {
// 				right.value = curr + maxValue[position{row + 1, col + 1}].value
// 			}
// 			mid.value = curr + maxValue[position{row + 1, col}].value

// 			if left.value >= right.value && left.value >= mid.value {
// 				if right.value >= mid.value {
// 					left.secondBestCol = col + 1
// 				} else {
// 					left.secondBestCol = col
// 				}
// 				maxValue[position{row, col}] = left

// 			} else if mid.value >= left.value && mid.value >= right.value {
// 				if left.value >= right.value {
// 					mid.secondBestCol = col - 1
// 				} else {
// 					mid.secondBestCol = col + 1
// 				}
// 				maxValue[position{row, col}] = mid

// 			} else {
// 				if left.value >= mid.value {
// 					right.secondBestCol = col - 1
// 				} else {
// 					right.secondBestCol = col
// 				}
// 				maxValue[position{row, col}] = right
// 			}

// 			if col-1 >= 0 {
// 				left.value = curr + preferRightValue[position{row + 1, col - 1}].value
// 			}
// 			if col+1 <= gridLength {
// 				right.value = curr + preferRightValue[position{row + 1, col + 1}].value
// 			}
// 			mid.value = curr + preferRightValue[position{row + 1, col}].value

// 			if right.value >= left.value && right.value >= mid.value {
// 				if mid.value >= left.value {
// 					right.secondBestCol = col
// 				} else {
// 					right.secondBestCol = col - 1
// 				}
// 				preferRightValue[position{row, col}] = right
// 			} else if mid.value >= right.value && mid.value >= left.value {
// 				if right.value >= left.value {
// 					mid.secondBestCol = col + 1
// 				} else {
// 					mid.secondBestCol = col - 1
// 				}
// 				preferRightValue[position{row, col}] = mid
// 			} else {
// 				if right.value >= mid.value {
// 					left.secondBestCol = col + 1
// 				} else {
// 					left.secondBestCol = col
// 				}
// 				preferRightValue[position{row, col}] = left
// 			}

// 			col--
// 		}
// 		col = gridLength
// 		row--
// 	}

// 	robot1Pos, robot2Pos := 0, gridLength
// 	row = 0

// 	path1, path2 := []int{}, []int{}

// 	for row < len(grid)-1 {
// 		valueRob1 := maxValue[position{row, robot1Pos}]
// 		valueRob2 := preferRightValue[position{row, robot2Pos}]

// 		if valueRob1.bestCol == valueRob2.bestCol { // If the next step would have them collide
// 			fmt.Println("collision!")
// 			if maxValue[position{row + 1, valueRob1.secondBestCol}].value >= preferRightValue[position{row + 1, valueRob2.secondBestCol}].value {
// 				robot1Pos = valueRob1.secondBestCol
// 				robot2Pos = valueRob2.bestCol
// 			} else {
// 				robot2Pos = valueRob2.secondBestCol
// 				robot1Pos = valueRob1.bestCol
// 			}
// 		} else { // if they aren't going to collide set the next collumn to the best path
// 			// robot 1 should prefer left, robot 2 should prefer right
// 			robot2Pos = valueRob2.bestCol
// 			robot1Pos = valueRob1.bestCol
// 		}

// 		row++

// 		fmt.Print("Robot 1 Pos:")
// 		fmt.Print(row)
// 		fmt.Print(",")
// 		fmt.Print(robot1Pos)
// 		fmt.Print("Robot 2 Pos:")
// 		fmt.Print(row)
// 		fmt.Print(",")
// 		fmt.Println(robot2Pos)
// 		fmt.Println(totalLostCherries)

// 	}
// 	return maxValue[position{0, 0}].value + preferRightValue[position{0, gridLength}].value - totalLostCherries

// }

type value struct {
	value, bestCol int
}

func cherryPickup(grid [][]int) int {
	storage := grid
	robot1 := simulateRun(storage, true)
	storage = grid
	robot2 := simulateRun(storage, false)

	if robot1 > robot2 {
		return robot1
	} else {
		return robot2
	}
}

func simulateRun(grid [][]int, first bool) int {

	var maxValue map[position]value
	maxValue = make(map[position]value)

	row := 0
	col := 0
	total := 0

	if first {

		maxValue = createMatrix(grid)
		for row < len(grid) {
			total = total + grid[row][col]
			if col > 0 {
				grid[row][col] = grid[row][col-1]
			} else {
				grid[row][col] = 0
			}
			col = maxValue[position{row, col}].bestCol
			row++
		}
		maxValue = createMatrix(grid)

		row = 0
		col = len(grid[0]) - 1

		for row < len(grid) {
			total = total + grid[row][col]
			col = maxValue[position{row, col}].bestCol
			row++
		}
		return total
	} else {

		maxValue = createMatrix(grid)
		row := 0
		col := len(grid[0]) - 1

		for row < len(grid) {
			total = total + grid[row][col]
			if col < len(grid[0])-1 {
				grid[row][col] = grid[row][col+1]
			} else {
				grid[row][col] = 0
			}
			col = maxValue[position{row, col}].bestCol
			row++
		}
		maxValue = createMatrix(grid)
		row = 0
		col = 0

		for row < len(grid) {
			total = total + grid[row][col]
			col = maxValue[position{row, col}].bestCol
			row++
		}
		return total

	}

}

func createMatrix(grid [][]int) map[position]value {
	row, col := len(grid)-1, len(grid[0])-1

	var maxValue map[position]value
	maxValue = make(map[position]value)

	gridLength := len(grid[0]) - 1

	for row >= 0 {
		if row == len(grid)-1 {
			for col >= 0 {
				maxValue[position{row, col}] = value{0, col}
				col--
			}
		} else {
			for col >= 0 {
				left, mid, right := value{0, col - 1}, value{0, col}, value{0, col + 1}

				curr := grid[row][col]

				if col-1 >= 0 {
					left.value = curr + maxValue[position{row + 1, col - 1}].value
				}
				if col+1 <= gridLength {
					right.value = curr + maxValue[position{row + 1, col + 1}].value
				}
				mid.value = curr + maxValue[position{row + 1, col}].value

				if left.value >= mid.value && left.value >= right.value {
					maxValue[position{row, col}] = left
				} else if mid.value >= left.value && mid.value >= right.value {
					maxValue[position{row, col}] = mid
				} else {
					maxValue[position{row, col}] = right
				}
				col--
			}
		}
		col = len(grid[0]) - 1
		row--
	}
	return maxValue
}
