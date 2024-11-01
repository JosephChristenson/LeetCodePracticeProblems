package main

import "fmt"

func main() {
	gain := []int{-5, 1, 5, 0, -7}
	fmt.Println(largestAltitude(gain))
}

func largestAltitude(gain []int) int {
	result := 0
	altitude := 0
	for index := 0; index < len(gain); index++ {
		altitude = altitude + gain[index]
		if altitude > result {
			result = altitude
		}
	}
	return result
}
