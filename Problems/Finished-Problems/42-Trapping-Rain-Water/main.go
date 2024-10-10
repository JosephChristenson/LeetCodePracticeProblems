package main

import "fmt"

// Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
	height = []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap(height))
}

func trap(height []int) int {
	result := make([]int, len(height))
	blockheight := 0
	for index, val := range height {
		if val <= blockheight {
			result[index] = blockheight - val
		} else {
			blockheight = val
			result[index] = 0
		}
	}
	reverseCount := len(height) - 1
	blockheight = 0
	total := 0

	for reverseCount >= 0 {
		if height[reverseCount] <= blockheight {
			if result[reverseCount] >= blockheight-height[reverseCount] {
				total = total + blockheight - height[reverseCount]
			} else {
				total = total + result[reverseCount]
			}
			result[reverseCount] = blockheight - height[reverseCount]
		} else {
			blockheight = height[reverseCount]
			result[reverseCount] = 0
		}
		reverseCount--
	}

	return total
}
