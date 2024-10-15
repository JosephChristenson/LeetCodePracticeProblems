package main

import "fmt"

// Given an integer numRows, return the first numRows of Pascal's triangle.

// In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:

func main() {
	numRows := 5
	fmt.Println(generate(numRows))
	numRows = 1
	fmt.Println(generate(numRows))
}

func generate(numRows int) [][]int {
	result := [][]int{}

	row := 1
	for row <= numRows {
		rowArray := make([]int, row)
		for index := range rowArray {
			if index == 0 || index == len(rowArray)-1 {
				rowArray[index] = 1
			} else {
				rowArray[index] = result[row-2][index-1] + result[row-2][index]
			}
		}
		result = append(result, rowArray)
		row++
	}

	return result
}
