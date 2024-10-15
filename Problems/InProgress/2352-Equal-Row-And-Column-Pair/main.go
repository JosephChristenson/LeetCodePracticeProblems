package main

import "fmt"

// Given a 0-indexed n x n integer matrix grid, return the number of pairs (ri, cj) such that row ri and column cj are equal.

// A row and column pair is considered equal if they contain the same elements in the same order (i.e., an equal array).

func main() {
	grid := [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}
	fmt.Println(equalPairs(grid))
	grid = [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}
	fmt.Println(equalPairs(grid))
}
func equalPairs(grid [][]int) int {
	var m map[string]bool
	m = make(map[string]bool)

}
