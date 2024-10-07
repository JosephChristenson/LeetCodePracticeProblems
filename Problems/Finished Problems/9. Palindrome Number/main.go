package main

import (
	"fmt"
	"strconv"
)

// //Given an integer x, return true if x is a
// //palindrome
// //, and false otherwise.

func main() {
	x := 121
	fmt.Print(isPalindrome(x))
	x = -121
	fmt.Print(isPalindrome(x))
	x = 10
	fmt.Print(isPalindrome(x))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	stringX := strconv.Itoa(x)
	byteArray := []byte(stringX)

	for i, _ := range byteArray {
		if byteArray[i] != byteArray[len(byteArray)-(i+1)] {
			return false
		}
	}
	return true
}
