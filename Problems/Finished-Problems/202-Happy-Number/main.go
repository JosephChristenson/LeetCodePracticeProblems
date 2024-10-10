package main

import (
	"fmt"
	"math"
	"strconv"
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

	digits := []byte(strconv.Itoa(n))
	for {
		total := 0
		for _, digit := range digits {
			number, _ := strconv.Atoi(string(digit))
			total = total + int(math.Pow(float64(number), 2))
		}
		fmt.Println(total)
		if total == 1 {
			return true
		}
		if m[total] { // we have reached this total before
			return false
		}
		m[total] = true
		digits = []byte(strconv.Itoa(total))
	}
}
