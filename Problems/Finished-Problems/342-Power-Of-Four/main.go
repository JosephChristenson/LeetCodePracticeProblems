package main

// Given an integer n, return true if it is a power of four. Otherwise, return false.

// An integer n is a power of four, if there exists an integer x such that n == 4x.
func main() {

}

func isPowerOfFour(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	for n%4 == 0 {
		if n == 4 {
			return true
		}
		n = n / 4
	}
	return false
}
