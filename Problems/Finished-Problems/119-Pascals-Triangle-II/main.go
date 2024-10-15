package main

import "fmt"

func main() {
	input := 3
	fmt.Println(getRow(input))
	input = 0
	fmt.Println(getRow(input))
	input = 1
	fmt.Println(getRow(input))
}

func getRow(rowIndex int) []int {
	result := [][]int{}
	row := 1
	for row <= rowIndex+1 {
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

	return result[len(result)-1]

}
