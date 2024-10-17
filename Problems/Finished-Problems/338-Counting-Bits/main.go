package main

import (
	"fmt"
)

// Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.

func main() {
	n := 2
	fmt.Println(countBits(n))
	n = 5
	fmt.Println(countBits(n))
}

func countBits(n int) []int {
	result := []int{}
	index := 0
	for index <= n {
		currValue := index
		divisor := 2
		oneCount := 0
		for currValue > 0 {
			bit := (currValue % divisor)
			if bit != 0 {
				oneCount++
			}
			currValue = currValue - bit
			divisor = divisor * 2
		}
		result = append(result, oneCount)
		index++
	}
	return result
}
