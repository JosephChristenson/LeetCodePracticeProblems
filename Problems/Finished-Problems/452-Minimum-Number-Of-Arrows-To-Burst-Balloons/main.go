package main

import (
	"fmt"
	"sort"
)

func main() {
	points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	fmt.Println(findMinArrowShots(points))
}

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	bounds := points[0]
	count := 1
	for _, val := range points {
		if val[0] <= bounds[1] {
			if val[1] < bounds[1] {
				bounds[1] = val[1]
			}
		} else {
			bounds = val
			count++
		}
	}
	return count
}
