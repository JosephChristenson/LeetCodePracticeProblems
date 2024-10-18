package main

// Given an integer n, return true if it is a power of three. Otherwise, return false.

// An integer n is a power of three, if there exists an integer x such that n == 3x.

func main() {

}

func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	for n%3 == 0 {
		if n == 3 {
			return true
		}
		n = n / 3
	}
	return false
}
