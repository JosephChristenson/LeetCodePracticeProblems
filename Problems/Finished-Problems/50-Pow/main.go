package main

import "fmt"

// Implement pow(x, n), which calculates x raised to the power n (i.e., xn).

func main() {
	x := 2.00000
	n := 10
	fmt.Println(myPow(x, n))
	x = 2.10000
	n = 3
	fmt.Println(myPow(x, n))
	x = 2.00000
	n = -2
	fmt.Println(myPow(x, n))
}

func myPow(x float64, n int) float64 {

	if x == float64(1.000000) {
		return x
	}
	if x == float64(-1.00000) {
		if float64(-1*(n+1)%2) == 1 { // determine if it is negative or not
			return x * float64(-1)
		}
		return x
	}
	result := float64(0.00000)

	if n > 0 {
		result = x
		for n > 1 {
			result = result * x
			n--
		}
	} else {
		result = 1
		if n < -100000 {
			return 0
		}
		for n < 0 {
			result = result / x
			n++
		}
	}
	return result
}
