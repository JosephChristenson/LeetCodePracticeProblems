package main

import "fmt"

func main() {
	x := 4
	fmt.Println(mySqrt(x))
	x = 8
	fmt.Println(mySqrt(x))
}

func mySqrt(x int) int {
	result := 0

	for result*result <= x {
		result++
	}
	return result - 1
}
