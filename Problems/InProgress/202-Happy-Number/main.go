package main

import (
	"fmt"
	"math"
)

// Write an algorithm to determine if a number n is happy.

// A happy number is a number defined by the following process:

// Starting with any positive integer, replace the number by the sum of the squares of its digits.
// Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
// Those numbers for which this process ends in 1 are happy.
// Return true if n is a happy number, and false if not.

func main() {
	n := 19
	fmt.Println(isHappy(n))
	n = 2
	fmt.Println(isHappy(n))
}

func isHappy(n int) bool {
	var m map[int]bool
	m = make(map[int]bool)
	digits := n
	for {
		fmt.Println(digits)
		total := 0
		for _,{
			total = total + int(math.Pow(float64(digit(digits))))
		}
		
		if total == 1 {
			return true
		}
		if m[total] { // we have reached this total before
			return false
		}
		m[total] = true
		digits = total
	}

}

func digit(num, place int) int {
    r := num % int(math.Pow(10, float64(place)))
    return r / int(math.Pow(10, float64(place-1)))
}
