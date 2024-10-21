package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	fmt.Println(matrix)

	matrix = [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(matrix)
	fmt.Println(matrix)

	matrix = [][]int{{2, 29, 20, 26, 16, 28}, {12, 27, 9, 25, 13, 21}, {32, 33, 32, 2, 28, 14}, {13, 14, 32, 27, 22, 26}, {33, 1, 20, 7, 21, 7}, {4, 24, 1, 6, 32, 34}}
	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int) {
	dim := len(matrix) - 1
	row := len(matrix) / 2

	for row >= 0 {
		fmt.Println("---------")
		for index := row; index < dim-row; index++ {
			//			matrix[index+row][row], matrix[dim-row][index+row], matrix[dim-index-row][dim-row], matrix[row][dim-index-row] = matrix[dim-row][index+row], matrix[dim-index-row][dim-row], matrix[row][dim-index-row], matrix[index+row][row]

			matrix[index+row][row], matrix[dim-row][index+row], matrix[dim-index-row][dim-row], matrix[row][dim-index-row] = 8, 8, 8, 8

			fmt.Println(matrix)
		}
		row--
	}

	fmt.Println("+++++++++++")
}
