package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	num := 2736
	fmt.Println(maximumSwap(num))
	num = 9973
	fmt.Println(maximumSwap(num))
}
func maximumSwap(num int) int {
	sortedNumber := []byte(strconv.Itoa(num))
	sort.Slice(sortedNumber, func(i, j int) bool {
		return sortedNumber[i] > sortedNumber[j]
	})

	fullNumber := []byte(strconv.Itoa(num))

	largeIndex, smallIndex := 0, -1
	testIndex := 0

	for index, val := range fullNumber {
		if val == sortedNumber[testIndex] { // if it is equal to the largest digit note it
			// it will keep updating so it finds it at the furthest point back
			largeIndex = index

			if smallIndex == -1 {
				testIndex++
			}
		} else if val < sortedNumber[testIndex] && smallIndex == -1 {
			smallIndex = index
		}
	}
	if largeIndex > smallIndex && smallIndex != -1 {
		fullNumber[smallIndex], fullNumber[largeIndex] = fullNumber[largeIndex], fullNumber[smallIndex]

		intResult, _ := strconv.Atoi(string(fullNumber))

		if intResult > num {
			return intResult
		}
	}
	return num
}
