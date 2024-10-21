package main

import "fmt"

// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
// Note:

// A Sudoku board (partially filled) could be valid but is not necessarily solvable.
// Only the filled cells need to be validated according to the mentioned rules.

func main() {
	board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	fmt.Println(isValidSudoku(board))
}

func isValidSudoku(board [][]byte) bool {
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

	for _, pos := range zones {
		var list []byte
		for _, cell := range pos {
			list = append(list, board[cell[0]][cell[1]])
		}
		if !check(list) {
			return false
		}
	}
	return true
}

func check(set []byte) bool {
	var m map[byte]bool
	m = make(map[byte]bool)

	for _, val := range set {
		if val != '.' {
			if m[val] {
				return false
			} else {
				m[val] = true
			}
		}
	}
	return true
}
