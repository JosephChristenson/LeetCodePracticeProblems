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
	index1, index2 := 0, len(height)-1

	result := 0
	for index1 < index2 {
		Y1, Y2 := height[index1], height[index2]
		area := 0
		if Y1 >= Y2 {
			area = (index2 - index1) * Y2
			index2--
		} else {
			area = (index2 - index1) * Y1
			index1++
		}
		if result < area {
			result = area
		}
	}
	return result
}
