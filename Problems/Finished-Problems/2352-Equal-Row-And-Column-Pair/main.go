package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Given a 0-indexed n x n integer matrix grid, return the number of pairs (ri, cj) such that row ri and column cj are equal.

// A row and column pair is considered equal if they contain the same elements in the same order (i.e., an equal array).

func main() {
	grid := [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}
	fmt.Println(equalPairs(grid))
	grid = [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}
	fmt.Println(equalPairs(grid))
}
func equalPairs(grid [][]int) int {
	var m map[string]int
	m = make(map[string]int)
	row, col := 0, 0
	result := 0

	for col < len(grid[0]) {
		row = 0
		key := make([]string, 0, len(grid[row]))
		for row < len(grid) {
			key = append(key, strconv.Itoa(grid[row][col]))
			row++
		}
		m[strings.Join(key, ",")]++
		col++
	}
	row, col = 0, 0

	for row < len(grid) {
		col = 0
		key := make([]string, 0, len(grid[row]))
		for col < len(grid[0]) {
			key = append(key, strconv.Itoa(grid[row][col]))
			col++
		}
		result = result + m[strings.Join(key, ",")]
		row++
	}
	return result
}
