package main

import "fmt"

func main() {
	input := 3
	fmt.Println(input)
	input = 0
	fmt.Println(input)
	input = 1
	fmt.Println(input)
}

func getRow(rowIndex int) []int {

	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}

	result := make([]int, rowIndex*2-1)

	triangle := [][]int{}
	row := 2
	triangle[0] = []int{1}
	triangle[1] = []int{1, 1}

	for rowIndex > row {
		index := 0
		for index < row*2-1 {
			if index == 0 || index == len(triangle[row-1])-1 {

			}
		}
		row++
	}

	return result
}
