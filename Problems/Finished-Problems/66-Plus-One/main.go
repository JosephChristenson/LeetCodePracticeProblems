package main

// You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.

// Increment the large integer by one and return the resulting array of digits.

func main() {

}

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}
	if digits[len(digits)-1] < 9 {
		digits[len(digits)-1]++
		return digits
	}

	carry := len(digits) - 1
	for carry >= 0 {
		if digits[carry] < 9 {
			digits[carry]++
			return digits
		} else {
			digits[carry] = 0
		}
		carry--
	}
	digits = append([]int{1}, digits...)

	return digits
}
