package main

import "fmt"

// Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane, return the maximum number of points that lie on the same straight line.

func main() {
	points := [][]int{{1, 1}, {2, 2}, {3, 3}}
	fmt.Println(maxPoints(points))

	points = [][]int{{0, 5}, {1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}
	fmt.Println(maxPoints(points))
}

type line struct {
	slope, Y float32 // records y = mx + b
}

func maxPoints(points [][]int) int {
	var m map[line][]int
	m = make(map[line][]int)
	var straightUp map[int][]int
	straightUp = make(map[int][]int)

	if len(points) <= 2 {
		return len(points)
	}

	pointA := points[len(points)-1]

	points = points[:len(points)-1]
	largestCount := 0

	for len(points) > 1 {
		count := len(points) - 1

		for count >= 0 {

			pointB := points[count]
			val := line{}
			if pointA[0]-pointB[0] == 0 {
				if len(straightUp[pointA[0]]) == 0 {
					straightUp[pointA[0]] = append(straightUp[pointA[0]], len(points))
					straightUp[pointA[0]] = append(straightUp[pointA[0]], count)
				} else if !doesContain(straightUp[pointA[0]], count) {
					straightUp[pointA[0]] = append(straightUp[pointA[0]], count)
				}
			} else {
				val.slope = float32(pointA[1]-pointB[1]) / float32(pointA[0]-pointB[0])
				val.Y = float32(pointA[1]) - val.slope*float32(pointA[0])

				if len(m[val]) == 0 {
					m[val] = append(m[val], len(points))
					m[val] = append(m[val], count)
				} else if !doesContain(m[val], count) {
					m[val] = append(m[val], count)
				}
			}
			count--
		}

		pointA = points[len(points)-1]

		points = points[:len(points)-1]
	}
	for _, val := range m {
		count := len(val)
		if largestCount < count {
			largestCount = count
		}
	}
	for _, val := range straightUp {
		count := len(val)
		if largestCount < count {
			largestCount = count
		}
	}
	return largestCount
}

func doesContain(points []int, val int) bool {
	for _, value := range points {
		if value == val {
			return true
		}
	}
	return false
}
