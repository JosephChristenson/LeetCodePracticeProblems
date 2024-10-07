package main

// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

// Find two lines that together with the x-axis form a container, such that the container contains the most water.

// Return the maximum amount of water a container can store.

// Notice that you may not slant the container.
import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
	height = []int{1, 1}
	fmt.Println(maxArea(height))
	height = []int{1, 2, 4, 3}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	result, test, leftindex := 0, 0, 0
	for X, Y := range height {
		if X != leftindex {
			if height[leftindex] <= Y {
				test = height[leftindex] * (X - leftindex)
			} else {
				test = Y * (X - leftindex)
			}
			if test > result {
				result = test
			}
			if leftindex/height[leftindex] <= X/Y {
				leftindex = X
			}
		}
	}
	return result
}
