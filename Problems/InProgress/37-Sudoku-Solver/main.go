package main

import "fmt"

// Write a program to solve a Sudoku puzzle by filling the empty cells.

// A sudoku solution must satisfy all of the following rules:

// Each of the digits 1-9 must occur exactly once in each row.
// Each of the digits 1-9 must occur exactly once in each column.
// Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
// The '.' character indicates empty cells.

func main() {
	board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku(board)
	fmt.Println(board)
}

func solveSudoku(board [][]byte) {
	zones := [][][]int{}
	group := [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {0, 6}, {0, 7}, {0, 8}}
	zones = append(zones, group)
	group = [][]int{{1, 0}, {1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6}, {1, 7}, {1, 8}}
	zones = append(zones, group)
	group = [][]int{{2, 0}, {2, 1}, {2, 2}, {2, 3}, {2, 4}, {2, 5}, {2, 6}, {2, 7}, {2, 8}}
	zones = append(zones, group)
	group = [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 3}, {3, 4}, {3, 5}, {3, 6}, {3, 7}, {3, 8}}
	zones = append(zones, group)
	group = [][]int{{4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}, {4, 5}, {4, 6}, {4, 7}, {4, 8}}
	zones = append(zones, group)
	group = [][]int{{5, 0}, {5, 1}, {5, 2}, {5, 3}, {5, 4}, {5, 5}, {5, 6}, {5, 7}, {5, 8}}
	zones = append(zones, group)
	group = [][]int{{6, 0}, {6, 1}, {6, 2}, {6, 3}, {6, 4}, {6, 5}, {6, 6}, {6, 7}, {6, 8}}
	zones = append(zones, group)
	group = [][]int{{7, 0}, {7, 1}, {7, 2}, {7, 3}, {7, 4}, {7, 5}, {7, 6}, {7, 7}, {7, 8}}
	zones = append(zones, group)
	group = [][]int{{8, 0}, {8, 1}, {8, 2}, {8, 3}, {8, 4}, {8, 5}, {8, 6}, {8, 7}, {8, 8}}
	zones = append(zones, group)
	group = [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {7, 0}, {8, 0}}
	zones = append(zones, group)
	group = [][]int{{0, 1}, {1, 1}, {2, 1}, {3, 1}, {4, 1}, {5, 1}, {6, 1}, {7, 1}, {8, 1}}
	zones = append(zones, group)
	group = [][]int{{0, 2}, {1, 2}, {2, 2}, {3, 2}, {4, 2}, {5, 2}, {6, 2}, {7, 2}, {8, 2}}
	zones = append(zones, group)
	group = [][]int{{0, 3}, {1, 3}, {2, 3}, {3, 3}, {4, 3}, {5, 3}, {6, 3}, {7, 3}, {8, 3}}
	zones = append(zones, group)
	group = [][]int{{0, 4}, {1, 4}, {2, 4}, {3, 4}, {4, 4}, {5, 4}, {6, 4}, {7, 4}, {8, 4}}
	zones = append(zones, group)
	group = [][]int{{0, 5}, {1, 5}, {2, 5}, {3, 5}, {4, 5}, {5, 5}, {6, 5}, {7, 5}, {8, 5}}
	zones = append(zones, group)
	group = [][]int{{0, 6}, {1, 6}, {2, 6}, {3, 6}, {4, 6}, {5, 6}, {6, 6}, {7, 6}, {8, 6}}
	zones = append(zones, group)
	group = [][]int{{0, 7}, {1, 7}, {2, 7}, {3, 7}, {4, 7}, {5, 7}, {6, 7}, {7, 7}, {8, 7}}
	zones = append(zones, group)
	group = [][]int{{0, 8}, {1, 8}, {2, 8}, {3, 8}, {4, 8}, {5, 8}, {6, 8}, {7, 8}, {8, 8}}
	zones = append(zones, group)
	group = [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}
	zones = append(zones, group)
	group = [][]int{{3, 0}, {3, 1}, {3, 2}, {4, 0}, {4, 1}, {4, 2}, {5, 0}, {5, 1}, {5, 2}}
	zones = append(zones, group)
	group = [][]int{{6, 0}, {6, 1}, {6, 2}, {7, 0}, {7, 1}, {7, 2}, {8, 0}, {8, 1}, {8, 2}}
	zones = append(zones, group)
	group = [][]int{{0, 3}, {0, 4}, {0, 5}, {1, 3}, {1, 4}, {1, 5}, {2, 3}, {2, 4}, {2, 5}}
	zones = append(zones, group)
	group = [][]int{{3, 3}, {3, 4}, {3, 5}, {4, 3}, {4, 4}, {4, 5}, {5, 3}, {5, 4}, {5, 5}}
	zones = append(zones, group)
	group = [][]int{{6, 3}, {6, 4}, {6, 5}, {7, 3}, {7, 4}, {7, 5}, {8, 3}, {8, 4}, {8, 5}}
	zones = append(zones, group)
	group = [][]int{{0, 6}, {0, 7}, {0, 8}, {1, 6}, {1, 7}, {1, 8}, {2, 6}, {2, 7}, {2, 8}}
	zones = append(zones, group)
	group = [][]int{{3, 6}, {3, 7}, {3, 8}, {4, 6}, {4, 7}, {4, 8}, {5, 6}, {5, 7}, {5, 8}}
	zones = append(zones, group)
	group = [][]int{{6, 6}, {6, 7}, {6, 8}, {7, 6}, {7, 7}, {7, 8}, {8, 6}, {8, 7}, {8, 8}}
	zones = append(zones, group)

	var possibleValues map[Coordinate][]byte
	possibleValues = make(map[Coordinate][]byte)

	var zoneChecker map[Coordinate][]int
	zoneChecker = make(map[Coordinate][]int)

	for zoneIndex, zone := range zones {
		for _, cell := range zone {
			if board[cell[0]][cell[1]] == '.' {
				possibleValues[Coordinate{X: cell[0], Y: cell[1]}] = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
			} else {
				possibleValues[Coordinate{X: cell[0], Y: cell[1]}] = []byte{board[cell[0]][cell[1]]}
			}
			zoneChecker[Coordinate{X: cell[0], Y: cell[1]}] = append(zoneChecker[Coordinate{X: cell[0], Y: cell[1]}], zoneIndex)
		}
	}

	countFilled := 0
	for countFilled < 81 {
		countFilled = 0
		for cell, values := range possibleValues {
			if len(values) > 1 { // check if there are currently more than one values it could be
				list := values
				for index := 0; index < 3; index++ {
					for _, checkCell := range zones[zoneChecker[cell][index]] {
						if len(possibleValues[Coordinate{X: checkCell[0], Y: checkCell[1]}]) == 1 {
							for index := range list {
								if list[index] == possibleValues[Coordinate{X: checkCell[0], Y: checkCell[1]}][0] {
									list[index] = '.'
									break
								}
							}
						}
					}
				}
				newList := []byte{}
				for index := range list {
					if list[index] != '.' {
						newList = append(newList, list[index])
					}
				}
				if len(newList) == 1 {
					board[cell.X][cell.Y] = newList[0]
				}
				possibleValues[cell] = newList
			} else {
				countFilled += 1
			}
		}
		fmt.Println(countFilled)
	}

}

type Coordinate struct {
	X, Y int
}
