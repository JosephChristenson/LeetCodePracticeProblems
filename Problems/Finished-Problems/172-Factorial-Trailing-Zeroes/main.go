package main

import "fmt"

// Given an integer n, return the number of trailing zeroes in n!.

// Note that n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1.

func main() {
	n := 3
	fmt.Println(trailingZeroes(n))
	n = 5
	fmt.Println(trailingZeroes(n))
	n = 10000
	fmt.Println(trailingZeroes(n))
}

func trailingZeroes(n int) int {
	// should be written that every power of 5 adds more trailing zeroes
	return (n/5 + n/25 + n/125 + n/625 + n/3125)
}
