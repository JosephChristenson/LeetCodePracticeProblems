package main

import "fmt"

func main() {
	num := 6
	fmt.Println(guessNumber(num))
	//
}

func guess(num int) int {
	if num < 6 {
		return 1
	} else if num > 6 {
		return -1
	}
	return 0
}

func guessNumber(n int) int {
	low, high := 1, n

	result := guess(0)
	for result != 0 {
		mid := low + (high-low)/2
		if guess(mid) == 0 {
			return mid
		} else if guess(mid) == -1 {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return result
}
